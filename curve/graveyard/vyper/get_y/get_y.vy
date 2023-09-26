# @version 0.3.6

@view
@external
def get_y(i: int128, j: int128, x: uint256, xp_: uint256[8], A: uint256, N_COINS: uint256, D: uint256) -> uint256:
    amp: uint256 = A
    c: uint256 = D
    S_: uint256 = 0
    N_: uint256 = N_COINS
    Ann: uint256 = amp * N_

    _x: uint256 = 0
    for _i in range(8):
        check: int128 = convert(N_, int128)
        if _i == i:
            _x = x
        elif _i != j:
            _x = xp_[_i]
        else:
            continue

        if _i < check:
            c = c * D / (_x * N_)
        else:
            continue
        S_ += _x        
            
    c = c * D / (Ann * N_)
    b: uint256 = S_ + D / Ann  
    y_prev: uint256 = 0
    y: uint256 = D
    for _i in range(255):
        y_prev = y
        y = (y*y + c) / (2 * y + b - D)
        
        if y > y_prev:
            if y - y_prev <= 1:
                break
        else:
            if y_prev - y <= 1:
                break
    return y