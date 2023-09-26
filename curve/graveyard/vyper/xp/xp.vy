# @version 0.3.6

@view
@external
def _xp(balances: uint256[8], N_COINS: uint256, RATES: uint256[8], LENDING_PRECISION: uint256) -> uint256[8]:
    result: uint256[8] = RATES
    for i in range(8):
        result[i] = result[i] * balances[i] / LENDING_PRECISION
    return result