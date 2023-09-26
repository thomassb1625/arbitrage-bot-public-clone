# @version 0.2.11
"""
@title Curve Registry
@license MIT
@author Curve.Fi
"""

MAX_COINS: constant(int128) = 8
CALC_INPUT_SIZE: constant(int128) = 100

struct CoinInfo:
    index: uint256
    register_count: uint256
    swap_count: uint256
    swap_for: address[MAX_INT128]

struct PoolArray:
    location: uint256
    decimals: uint256
    underlying_decimals: uint256
    rate_info: bytes32
    base_pool: address
    coins: address[MAX_COINS]
    ul_coins: address[MAX_COINS]
    n_coins: uint256  # [coins, underlying coins] tightly packed as uint128[2]
    has_initial_A: bool
    is_v1: bool
    name: String[64]
    asset_type: uint256

struct PoolParams:
    A: uint256
    future_A: uint256
    fee: uint256
    admin_fee: uint256
    future_fee: uint256
    future_admin_fee: uint256
    future_owner: address
    initial_A: uint256
    initial_A_time: uint256
    future_A_time: uint256


interface AddressProvider:
    def admin() -> address: view
    def get_address(_id: uint256) -> address: view

interface ERC20:
    def balanceOf(_addr: address) -> uint256: view
    def decimals() -> uint256: view
    def totalSupply() -> uint256: view

interface CurvePool:
    def A() -> uint256: view
    def future_A() -> uint256: view
    def fee() -> uint256: view
    def admin_fee() -> uint256: view
    def future_fee() -> uint256: view
    def future_admin_fee() -> uint256: view
    def future_owner() -> address: view
    def initial_A() -> uint256: view
    def initial_A_time() -> uint256: view
    def future_A_time() -> uint256: view
    def coins(i: uint256) -> address: view
    def underlying_coins(i: uint256) -> address: view
    def balances(i: uint256) -> uint256: view
    def get_virtual_price() -> uint256: view

interface CurvePoolV1:
    def coins(i: int128) -> address: view
    def underlying_coins(i: int128) -> address: view
    def balances(i: int128) -> uint256: view

interface CurveMetapool:
    def base_pool() -> address: view

interface GasEstimator:
    def estimate_gas_used(_pool: address, _from: address, _to: address) -> uint256: view

interface LiquidityGauge:
    def lp_token() -> address: view

interface GaugeController:
    def gauge_types(gauge: address) -> int128: view

interface RateCalc:
    def get_rate(_coin: address) -> uint256: view


event PoolAdded:
    pool: indexed(address)
    rate_method_id: Bytes[4]

event PoolRemoved:
    pool: indexed(address)


address_provider: public(AddressProvider)
gauge_controller: public(address)
pool_list: public(address[65536])   # master list of pools
pool_count: public(uint256)         # actual length of pool_list

pool_data: HashMap[address, PoolArray]

coin_count: public(uint256)  # total unique coins registered
coins: HashMap[address, CoinInfo]
get_coin: public(address[65536])  # unique list of registered coins
# bitwise_xor(coina, coinb) -> (coina_pos, coinb_pos) sorted
# stored as uint128[2]
coin_swap_indexes: HashMap[uint256, uint256]

# lp token -> pool
get_pool_from_lp_token: public(HashMap[address, address])

# pool -> lp token
get_lp_token: public(HashMap[address, address])

# mapping of estimated gas costs for pools and coins
# for a pool the values are [wrapped exchange, underlying exchange]
# for a coin the values are [transfer cost, 0]
gas_estimate_values: HashMap[address, uint256[2]]

# pool -> gas estimation contract
# used when gas costs for a pool are too complex to be handled by summing
# values in `gas_estimate_values`
gas_estimate_contracts: HashMap[address, address]

# mapping of coins -> pools for trading
# a mapping key is generated for each pair of addresses via
# `bitwise_xor(convert(a, uint256), convert(b, uint256))`
markets: HashMap[uint256, address[65536]]
market_counts: HashMap[uint256, uint256]

liquidity_gauges: HashMap[address, address[10]]

last_updated: public(uint256)


@external
def __init__(_address_provider: address, _gauge_controller: address):
    """
    @notice Constructor function
    """
    self.address_provider = AddressProvider(_address_provider)
    self.gauge_controller = _gauge_controller


# internal functionality for getters

@view
@internal
def _unpack_decimals(_packed: uint256, _n_coins: uint256) -> uint256[MAX_COINS]:
    # decimals are tightly packed as a series of uint8 within a little-endian bytes32
    # the packed value is stored as uint256 to simplify unpacking via shift and modulo
    decimals: uint256[MAX_COINS] = empty(uint256[MAX_COINS])
    n_coins: int128 = convert(_n_coins, int128)
    for i in range(MAX_COINS):
        if i == n_coins:
            break
        decimals[i] = shift(_packed, -8 * i) % 256

    return decimals


@view
@internal
def _get_rates(_pool: address) -> uint256[MAX_COINS]:
    rates: uint256[MAX_COINS] = empty(uint256[MAX_COINS])
    base_pool: address = self.pool_data[_pool].base_pool
    if base_pool == ZERO_ADDRESS:
        rate_info: bytes32 = self.pool_data[_pool].rate_info
        rate_calc_addr: uint256 = convert(slice(rate_info, 8, 20), uint256)
        rate_method_id: Bytes[4] = slice(rate_info, 28, 4)

        for i in range(MAX_COINS):
            coin: address = self.pool_data[_pool].coins[i]
            if coin == ZERO_ADDRESS:
                break
            if rate_info == EMPTY_BYTES32 or coin == self.pool_data[_pool].ul_coins[i]:
                rates[i] = 10 ** 18
            elif rate_calc_addr != 0:
                rates[i] = RateCalc(convert(rate_calc_addr, address)).get_rate(coin)
            else:
                rates[i] = convert(
                    raw_call(coin, rate_method_id, max_outsize=32, is_static_call=True),  # dev: bad response
                    uint256
                )
    else:
        base_coin_idx: uint256 = shift(self.pool_data[_pool].n_coins, -128) - 1
        rates[base_coin_idx] = CurvePool(base_pool).get_virtual_price()
        for i in range(MAX_COINS):
            if i == base_coin_idx:
                break
            rates[i] = 10 ** 18

    return rates

@view
@internal
def _get_balances(_pool: address) -> uint256[MAX_COINS]:
    is_v1: bool = self.pool_data[_pool].is_v1

    balances: uint256[MAX_COINS] = empty(uint256[MAX_COINS])
    for i in range(MAX_COINS):
        if self.pool_data[_pool].coins[i] == ZERO_ADDRESS:
            assert i != 0
            break

        if is_v1:
            balances[i] = CurvePoolV1(_pool).balances(i)
        else:
            balances[i] = CurvePool(_pool).balances(convert(i, uint256))

    return balances


@view
@internal
def _get_underlying_balances(_pool: address) -> uint256[MAX_COINS]:
    balances: uint256[MAX_COINS] = self._get_balances(_pool)
    rates: uint256[MAX_COINS] = self._get_rates(_pool)
    decimals: uint256 = self.pool_data[_pool].underlying_decimals
    underlying_balances: uint256[MAX_COINS] = balances
    for i in range(MAX_COINS):
        coin: address = self.pool_data[_pool].coins[i]
        if coin == ZERO_ADDRESS:
            break
        ucoin: address = self.pool_data[_pool].ul_coins[i]
        if ucoin == ZERO_ADDRESS:
            continue
        if ucoin != coin:
            underlying_balances[i] = balances[i] * rates[i] / 10**(shift(decimals, -8 * i) % 256)

    return underlying_balances


@view
@internal
def _get_meta_underlying_balances(_pool: address, _base_pool: address) -> uint256[MAX_COINS]:
    base_coin_idx: uint256 = shift(self.pool_data[_pool].n_coins, -128) - 1
    is_v1: bool = self.pool_data[_base_pool].is_v1
    base_total_supply: uint256 = ERC20(self.get_lp_token[_base_pool]).totalSupply()

    underlying_balances: uint256[MAX_COINS] = empty(uint256[MAX_COINS])
    ul_balance: uint256 = 0
    underlying_pct: uint256 = 0
    if base_total_supply > 0:
        underlying_pct = CurvePool(_pool).balances(base_coin_idx) * 10**36 / base_total_supply

    for i in range(MAX_COINS):
        if self.pool_data[_pool].ul_coins[i] == ZERO_ADDRESS:
            break
        if i < base_coin_idx:
            ul_balance = CurvePool(_pool).balances(i)
        else:
            if is_v1:
                ul_balance = CurvePoolV1(_base_pool).balances(convert(i - base_coin_idx, int128))
            else:
                ul_balance = CurvePool(_base_pool).balances(i-base_coin_idx)
            ul_balance = ul_balance * underlying_pct / 10**36
        underlying_balances[i] = ul_balance

    return underlying_balances


@view
@internal
def _get_coin_indices(
    _pool: address,
    _from: address,
    _to: address
) -> uint256[3]:
    """
    Convert coin addresses to indices for use with pool methods.
    """
    # the return value is stored as `uint256[3]` to reduce gas costs
    # from index, to index, is the market underlying?
    result: uint256[3] = empty(uint256[3])

    found_market: bool = False

    # check coin markets
    for x in range(MAX_COINS):
        coin: address = self.pool_data[_pool].coins[x]
        if coin == ZERO_ADDRESS:
            # if we reach the end of the coins, reset `found_market` and try again
            # with the underlying coins
            found_market = False
            break
        if coin == _from:
            result[0] = x
        elif coin == _to:
            result[1] = x
        else:
            continue

        if found_market:
            # the second time we find a match, break out of the loop
            break
        # the first time we find a match, set `found_market` to True
        found_market = True

    if not found_market:
        # check underlying coin markets
        for x in range(MAX_COINS):
            coin: address = self.pool_data[_pool].ul_coins[x]
            if coin == ZERO_ADDRESS:
                raise "No available market"
            if coin == _from:
                result[0] = x
            elif coin == _to:
                result[1] = x
            else:
                continue

            if found_market:
                result[2] = 1
                break
            found_market = True

    return result


# targetted external getters, optimized for on-chain calls

@view
@external
def find_pool_for_coins(_from: address, _to: address, i: uint256 = 0) -> address:
    """
    @notice Find an available pool for exchanging two coins
    @param _from Address of coin to be sent
    @param _to Address of coin to be received
    @param i Index value. When multiple pools are available
            this value is used to return the n'th address.
    @return Pool address
    """
    key: uint256 = bitwise_xor(convert(_from, uint256), convert(_to, uint256))
    return self.markets[key][i]


@view
@external
def get_n_coins(_pool: address) -> uint256[2]:
    """
    @notice Get the number of coins in a pool
    @dev For non-metapools, both returned values are identical
         even when the pool does not use wrapping/lending
    @param _pool Pool address
    @return Number of wrapped coins, number of underlying coins
    """
    n_coins: uint256 = self.pool_data[_pool].n_coins
    return [shift(n_coins, -128), n_coins % 2**128]


@view
@external
def get_coins(_pool: address) -> address[MAX_COINS]:
    """
    @notice Get the coins within a pool
    @dev For pools using lending, these are the wrapped coin addresses
    @param _pool Pool address
    @return List of coin addresses
    """
    coins: address[MAX_COINS] = empty(address[MAX_COINS])
    n_coins: uint256 = shift(self.pool_data[_pool].n_coins, -128)
    for i in range(MAX_COINS):
        if i == n_coins:
            break
        coins[i] = self.pool_data[_pool].coins[i]

    return coins


@view
@external
def get_underlying_coins(_pool: address) -> address[MAX_COINS]:
    """
    @notice Get the underlying coins within a pool
    @dev For pools that do not lend, returns the same value as `get_coins`
    @param _pool Pool address
    @return List of coin addresses
    """
    coins: address[MAX_COINS] = empty(address[MAX_COINS])
    n_coins: uint256 = self.pool_data[_pool].n_coins % 2**128
    for i in range(MAX_COINS):
        if i == n_coins:
            break
        coins[i] = self.pool_data[_pool].ul_coins[i]

    return coins


@view
@external
def get_decimals(_pool: address) -> uint256[MAX_COINS]:
    """
    @notice Get decimal places for each coin within a pool
    @dev For pools using lending, these are the wrapped coin decimal places
    @param _pool Pool address
    @return uint256 list of decimals
    """
    n_coins: uint256 = shift(self.pool_data[_pool].n_coins, -128)
    return self._unpack_decimals(self.pool_data[_pool].decimals, n_coins)


@view
@external
def get_underlying_decimals(_pool: address) -> uint256[MAX_COINS]:
    """
    @notice Get decimal places for each underlying coin within a pool
    @dev For pools that do not lend, returns the same value as `get_decimals`
    @param _pool Pool address
    @return uint256 list of decimals
    """
    n_coins: uint256 = self.pool_data[_pool].n_coins % 2**128
    return self._unpack_decimals(self.pool_data[_pool].underlying_decimals, n_coins)


@view
@external
def get_rates(_pool: address) -> uint256[MAX_COINS]:
    """
    @notice Get rates between coins and underlying coins
    @dev For coins where there is no underlying coin, or where
         the underlying coin cannot be swapped, the rate is
         given as 1e18
    @param _pool Pool address
    @return Rates between coins and underlying coins
    """
    return self._get_rates(_pool)


@view
@external
def get_gauges(_pool: address) -> (address[10], int128[10]):
    """
    @notice Get a list of LiquidityGauge contracts associated with a pool
    @param _pool Pool address
    @return address[10] of gauge addresses, int128[10] of gauge types
    """
    liquidity_gauges: address[10] = empty(address[10])
    gauge_types: int128[10] = empty(int128[10])
    gauge_controller: address = self.gauge_controller
    for i in range(10):
        gauge: address = self.liquidity_gauges[_pool][i]
        if gauge == ZERO_ADDRESS:
            break
        liquidity_gauges[i] = gauge
        gauge_types[i] = GaugeController(gauge_controller).gauge_types(gauge)

    return liquidity_gauges, gauge_types


@view
@external
def get_balances(_pool: address) -> uint256[MAX_COINS]:
    """
    @notice Get balances for each coin within a pool
    @dev For pools using lending, these are the wrapped coin balances
    @param _pool Pool address
    @return uint256 list of balances
    """
    return self._get_balances(_pool)


@view
@external
def get_underlying_balances(_pool: address) -> uint256[MAX_COINS]:
    """
    @notice Get balances for each underlying coin within a pool
    @dev  For pools that do not lend, returns the same value as `get_balances`
    @param _pool Pool address
    @return uint256 list of underlyingbalances
    """
    base_pool: address = self.pool_data[_pool].base_pool
    if base_pool == ZERO_ADDRESS:
        return self._get_underlying_balances(_pool)
    return self._get_meta_underlying_balances(_pool, base_pool)


@view
@external
def get_virtual_price_from_lp_token(_token: address) -> uint256:
    """
    @notice Get the virtual price of a pool LP token
    @param _token LP token address
    @return uint256 Virtual price
    """
    return CurvePool(self.get_pool_from_lp_token[_token]).get_virtual_price()


@view
@external
def get_A(_pool: address) -> uint256:
    return CurvePool(_pool).A()


@view
@external
def get_parameters(_pool: address) -> PoolParams:
    """
    @notice Get parameters for a pool
    @dev For older pools where `initial_A` is not public, this value is set to 0
    @param _pool Pool address
    @return Pool amp, future amp, fee, admin fee, future fee, future admin fee,
            future owner, initial amp, initial amp time, future amp time
    """
    pool_params: PoolParams = empty(PoolParams)
    pool_params.A = CurvePool(_pool).A()
    pool_params.future_A = CurvePool(_pool).future_A()
    pool_params.fee = CurvePool(_pool).fee()
    pool_params.future_fee = CurvePool(_pool).future_fee()
    pool_params.admin_fee = CurvePool(_pool).admin_fee()
    pool_params.future_admin_fee = CurvePool(_pool).future_admin_fee()
    pool_params.future_owner = CurvePool(_pool).future_owner()

    if self.pool_data[_pool].has_initial_A:
        pool_params.initial_A = CurvePool(_pool).initial_A()
        pool_params.initial_A_time = CurvePool(_pool).initial_A_time()
        pool_params.future_A_time = CurvePool(_pool).future_A_time()

    return pool_params


@view
@external
def get_fees(_pool: address) -> uint256[2]:
    """
    @notice Get the fees for a pool
    @dev Fees are expressed as integers
    @return Pool fee as uint256 with 1e10 precision
            Admin fee as 1e10 percentage of pool fee
    """
    return [CurvePool(_pool).fee(), CurvePool(_pool).admin_fee()]


@view
@external
def get_admin_balances(_pool: address) -> uint256[MAX_COINS]:
    """
    @notice Get the current admin balances (uncollected fees) for a pool
    @param _pool Pool address
    @return List of uint256 admin balances
    """
    balances: uint256[MAX_COINS] = self._get_balances(_pool)
    n_coins: uint256 = shift(self.pool_data[_pool].n_coins, -128)
    for i in range(MAX_COINS):
        coin: address = self.pool_data[_pool].coins[i]
        if i == n_coins:
            break
        if coin == 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE:
            balances[i] = _pool.balance - balances[i]
        else:
            balances[i] = ERC20(coin).balanceOf(_pool) - balances[i]

    return balances


@view
@external
def get_coin_indices(
    _pool: address,
    _from: address,
    _to: address
) -> (int128, int128, bool):
    """
    @notice Convert coin addresses to indices for use with pool methods
    @param _from Coin address to be used as `i` within a pool
    @param _to Coin address to be used as `j` within a pool
    @return int128 `i`, int128 `j`, boolean indicating if `i` and `j` are underlying coins
    """
    result: uint256[3] = self._get_coin_indices(_pool, _from, _to)
    return convert(result[0], int128), convert(result[1], int128), result[2] > 0


@view
@external
def estimate_gas_used(_pool: address, _from: address, _to: address) -> uint256:
    """
    @notice Estimate the gas used in an exchange.
    @param _pool Pool address
    @param _from Address of coin to be sent
    @param _to Address of coin to be received
    @return Upper-bound gas estimate, in wei
    """
    estimator: address = self.gas_estimate_contracts[_pool]
    if estimator != ZERO_ADDRESS:
        return GasEstimator(estimator).estimate_gas_used(_pool, _from, _to)

    # here we call `_get_coin_indices` to find out if the exchange involves wrapped
    # or underlying coins, and use the result as an index in `gas_estimate_values`
    # 0 == wrapped   1 == underlying
    idx_underlying: uint256 = self._get_coin_indices(_pool, _from, _to)[2]

    total: uint256 = self.gas_estimate_values[_pool][idx_underlying]
    assert total != 0  # dev: pool value not set

    for addr in [_from, _to]:
        _gas: uint256 = self.gas_estimate_values[addr][0]
        assert _gas != 0  # dev: coin value not set
        total += _gas

    return total

@view
@external
def is_meta(_pool: address) -> bool:
    """
    @notice Verify `_pool` is a metapool
    @param _pool Pool address
    @return True if `_pool` is a metapool
    """
    return self.pool_data[_pool].base_pool != ZERO_ADDRESS


@view
@external
def get_pool_name(_pool: address) -> String[64]:
    """
    @notice Get the given name for a pool
    @param _pool Pool address
    @return The name of a pool
    """
    return self.pool_data[_pool].name


@view
@external
def get_coin_swap_count(_coin: address) -> uint256:
    """
    @notice Get the number of unique coins available to swap `_coin` against
    @param _coin Coin address
    @return The number of unique coins available to swap for
    """
    return self.coins[_coin].swap_count


@view
@external
def get_coin_swap_complement(_coin: address, _index: uint256) -> address:
    """
    @notice Get the coin available to swap against `_coin` at `_index`
    @param _coin Coin address
    @param _index An index in the `_coin`'s set of available counter
        coin's
    @return Address of a coin available to swap against `_coin`
    """
    return self.coins[_coin].swap_for[_index]


@view
@external
def get_pool_asset_type(_pool: address) -> uint256:
    """
    @notice Query the asset type of `_pool`
    @param _pool Pool Address
    @return The asset type as an unstripped string
    """
    return self.pool_data[_pool].asset_type