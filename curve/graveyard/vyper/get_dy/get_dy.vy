# @version 0.3.6

@view
@external
def get_dy(i: int128, j: int128, dx: uint256, xp: uint256[8], y: uint256, RATES: uint256[8], N_COINS: uint256, PRECISION: uint256, FEE_DENOMINATOR: uint256, fee: uint256) -> uint256:    
    rates: uint256[8] = RATES

    x: uint256 = xp[i] + (dx * rates[i] / PRECISION)
    dy: uint256 = (xp[j] - y - 1) * PRECISION / rates[j]
    _fee: uint256 = fee * dy / FEE_DENOMINATOR
    return dy - _fee