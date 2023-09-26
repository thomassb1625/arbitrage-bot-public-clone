# @version 0.3.6

@view
@external
def get_dy_underlying(i: int128, j: int128, dx: uint256) -> uint256:
    # dx and dy in underlying units
    xp: uint256[N_COINS] = self._xp()
    precisions: uint256[N_COINS] = PRECISION_MUL

    x: uint256 = xp[i] + dx * precisions[i]
    y: uint256 = self.get_y(i, j, x, xp)
    dy: uint256 = (xp[j] - y - 1) / precisions[j]
    _fee: uint256 = self.fee * dy / FEE_DENOMINATOR
    return dy - _fee