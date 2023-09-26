package dex

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

type Token struct {
	Addr        common.Address
	IsTaxed     bool
	IsTradable  bool //False if estimated amount after tax is 0
	AfterTaxNum *uint256.Int
	AfterTaxDen *uint256.Int
}

func NewTaxed(addr common.Address, numerator, denominator *uint256.Int, tradable bool) Token {
	return Token{Addr: addr, IsTaxed: true, AfterTaxNum: new(uint256.Int).Set(numerator), AfterTaxDen: new(uint256.Int).Set(denominator), IsTradable: tradable}
}

func NewNotTaxed(addr common.Address) Token {
	return Token{Addr: addr, IsTaxed: false, AfterTaxNum: nil, AfterTaxDen: nil, IsTradable: true}
}

func MakeTokenMap(tokens []Token) map[common.Address]Token {
	tokensMap := make(map[common.Address]Token)
	for _, t := range tokens {
		x, ok := tokensMap[t.Addr]
		//In case of repated token(likely from different market) keep the one with lower output
		if ok {
			if x.AfterTaxNum.Cmp(t.AfterTaxNum) < 0 {
				tokensMap[t.Addr] = x
			}
		} else {
			tokensMap[t.Addr] = t
		}
	}
	return tokensMap
}

func SaveTokens(list []Token, filename string) {
	data, _ := json.Marshal(list)
	ioutil.WriteFile(filename, data, 0666)
}

func LoadTokens(filename string) []Token {
	tokens := make([]Token, 0)
	jsonFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	bytes, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bytes, &tokens)
	return tokens
}
