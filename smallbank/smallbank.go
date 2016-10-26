package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SmallBank struct {
}

var MAX_ACCOUNTS int = 100000
var BALANCE int = 1000000
var accountTab string = "accounts"
var savingTab string = "saving"
var checkingTab string = "checking"

func main() {
	err := shim.Start(new(SmallBank))
	if err != nil {
		fmt.Printf("Error starting smallbank: %s", err)
	}
}

func (t *SmallBank) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	for i := 0; i < MAX_ACCOUNTS; i++ {
		stub.PutState(accountTab+"_"+strconv.Itoa(i), []byte("name_"+strconv.Itoa(i)))
		stub.PutState(savingTab+"_"+strconv.Itoa(i), []byte(strconv.Itoa(BALANCE)))
		stub.PutState(checkingTab+"_"+strconv.Itoa(i), []byte(strconv.Itoa(BALANCE)))
	}

	return nil, nil
}

func (t *SmallBank) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	/*
		if function == "updateBalance" {
			return t.updateBalance(stub, args)
		} else {
			return nil, nil
		}

	*/
	return nil, nil

}
func (t *SmallBank) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	valAsbytes, err := stub.GetState(args[0])
	return valAsbytes, err
}

func (t *SmallBank) updateBalance(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	stub.GetState(accountTab + "_" + args[0])
	bal_str, err := stub.GetState(checkingTab + "_" + args[0])
	var bal1, bal2 int
	bal1, err = strconv.Atoi(string(bal_str))
	bal2, err = strconv.Atoi(args[1])
	bal1 += bal2

	err = stub.PutState(checkingTab+"_"+args[0], []byte(strconv.Itoa(bal1)))

	if err != nil {
		return nil, err
	}
	return nil, nil
}
