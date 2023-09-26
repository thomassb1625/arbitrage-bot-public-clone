from web3 import Web3, HTTPProvider
import json

infura = 'https://mainnet.infura.io/v3/95a42ee2051d4c759fe380b03d72c98d'

w3 = Web3(HTTPProvider(infura))

pool_addrs = [
    '0x453D92C7d4263201C69aACfaf589Ed14202d83a4', # yCRV/CRV contract
]
ycrv_addr = '0xFCc5c47bE19d06BF83eB04298b026F81069ff65b'

crv_pool_abi = json.load(open('abis/crv_pool.json'))
ycrv_abi = json.load(open('abis/ycrv.json'))

pools = [
    w3.eth.contract(addr,abi=crv_pool_abi)
    for addr in pool_addrs
]

ycrv = w3.eth.contract(ycrv_addr,abi=ycrv_abi)

block_header_filter = w3.eth.filter('latest')

def get_pegs(pools):
    global requests
    pegs = []
    for pool in pools:
        # for yCRV/CRV pool, r1 is CRV
        # not sure about other pools
        r1, r2 = pool.functions.get_balances().call()
        pegs.append(r1/r2)
        requests += 1
    return pegs

def arb():
    # calculate optimal size

    # take flashloan

    # swap

    # repay flashloan

    # calculate bribe

    # submit transaction
    return 

def main():
    new_entries = block_header_filter.get_new_entries()
    global requests
    requests += 1

    if new_entries:
        # get block number
        block_number = w3.eth.get_block_number()
        print(f'\nblock_number: {block_number}')
        # call reserves and updates peg(s)
        pegs = get_pegs(pools)
        for i, peg in enumerate(pegs):
            print(f'pool: {pool_addrs[i]}')
            if peg > 1.05:
                arb()
                print('\tsubmitted')
            else:
                print(f'\tpeg: {peg}')
                print(f'\trequests: {requests}')
            print('')
        requests = 0

if __name__ == '__main__':
    requests = 0
    while True:
        main()