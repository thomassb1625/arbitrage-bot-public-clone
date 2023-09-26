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