# @version 0.3.6
# (c) Curve.Fi, 2020

@pure
@external
def get_D(xp: uint256[8], amp: uint256, N_COINS: uint256) -> uint256:
    S: uint256 = 0
    for _x in xp:
        S += _x
    if S == 0:
        return 0

    Dprev: uint256 = 0
    D: uint256 = S
    N_: uint256 = N_COINS
    Ann: uint256 = amp * N_
    for _i in range(255):
        D_P: uint256 = D
        for _x in xp:
            if _x != 0:
                D_P = D_P * D / (_x * N_)  
            else:
                continue
        Dprev = D
        D = (Ann * S + D_P * N_) * D / ((Ann - 1) * D + (N_ + 1) * D_P)
        
        if D > Dprev:
            if D - Dprev <= 1:
                break
        else:
            if Dprev - D <= 1:
                break
    return D