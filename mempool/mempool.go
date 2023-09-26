package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func DecodeTx(contractABI *abi.ABI, data []byte) error {
	methodSig := data[:4]
	method, err := contractABI.MethodById(methodSig)
	if err != nil {
		return err
	}

	inputData := data[4:]
	v, err := method.Inputs.UnpackValues(inputData)
	if err != nil {
		return err
	}

	fmt.Println(v)
	return nil
}
