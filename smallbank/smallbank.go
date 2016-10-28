package main

import (
	"fmt"
	//"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SmallBank struct {
}

var MAX_ACCOUNTS int = 100000
var BALANCE int = 100000
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

	return nil, nil
}

func (t *SmallBank) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "amalgate" {
		return t.almagate(stub, args)
	} else if function == "getBalance" {
		return t.getBalance(stub, args)
	} else if function == "updateBalance" {
		return t.updateBalance(stub, args)
	} else if function == "updateSaving" {
		return t.updateSaving(stub, args)
	} else if function == "sendPayment" {
		return t.sendPayment(stub, args)
	} else if function == "writeCheck" {
		return t.writeCheck(stub, args)
	}

	return nil, nil

}
func (t *SmallBank) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	valAsbytes, err := stub.GetState(args[0])
	return valAsbytes, err
}

func (t *SmallBank) almagate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return nil, nil
	/*
		bal_str1, err1 := stub.GetState(accountTab + "_" + args[0])
		if err1 != nil {
			stub.PutState(accountTab+"_"+args[0], []byte([]byte(strconv.Itoa(BALANCE))))
		}
		bal_str2, err2 := stub.GetState(checkingTab + "_" + args[0])
		if err2 != nil {
			stub.PutState(checkingTab+"_"+args[0], []byte(strconv.Itoa(BALANCE)))
		}
		_, err3 := stub.GetState(savingTab + "_" + args[0])
		if err3 != nil {
			stub.PutState(savingTab+"_"+args[0], []byte(strconv.Itoa(BALANCE)))
		}
		bal_str1, err1 = stub.GetState(accountTab + "_" + args[1])
		if err1 != nil {
			stub.PutState(accountTab+"_"+args[1], []byte(strconv.Itoa(BALANCE)))
		}
		bal_str2, err2 = stub.GetState(checkingTab + "_" + args[1])
		if err2 != nil {
			stub.PutState(checkingTab+"_"+args[1], []byte(strconv.Itoa(BALANCE)))
		}
		_, err3 = stub.GetState(savingTab + "_" + args[0])
		if err3 != nil {
			stub.PutState(savingTab+"_"+args[1], []byte(strconv.Itoa(BALANCE)))
		}

		var bal1, bal2 int
		var err error
		bal_str1, err = stub.GetState(savingTab + "_" + args[0])
		if err != nil {
			bal_str1 = []byte(strconv.Itoa(BALANCE))
		}
		bal_str2, err = stub.GetState(checkingTab + "_" + args[1])
		if err != nil {
			bal_str2 = []byte(strconv.Itoa(BALANCE))
		}

		bal1, err = strconv.Atoi(string(bal_str1))
		bal2, err = strconv.Atoi(string(bal_str2))
		bal1 += bal2

		err = stub.PutState(checkingTab+"_"+args[0], []byte(strconv.Itoa(0)))

		if err != nil {
			return nil, err
		}

		err = stub.PutState(savingTab+"_"+args[1], []byte(strconv.Itoa(bal1)))

		if err != nil {
			fmt.Printf("Error almagate: %s", err)
			return nil, err
		}
		return nil, nil
	*/
}

func (t *SmallBank) getBalance(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return nil, nil
	/*
		bal_str1, err1 := stub.GetState(accountTab + "_" + args[0])
		if err1 != nil {
			stub.PutState(accountTab+"_"+args[0], []byte(strconv.Itoa(BALANCE)))
		}
		bal_str2, err2 := stub.GetState(checkingTab + "_" + args[0])
		if err2 != nil {
			stub.PutState(checkingTab+"_"+args[0], []byte(strconv.Itoa(BALANCE)))
		}
		_, err3 := stub.GetState(savingTab + "_" + args[0])
		if err3 != nil {
			stub.PutState(savingTab+"_"+args[0], []byte(strconv.Itoa(BALANCE)))
		}

		var bal1, bal2 int
		var err error
		bal_str1, err = stub.GetState(savingTab + "_" + args[0])
		if err != nil {
			bal_str1 = []byte(strconv.Itoa(BALANCE))
		}
		bal_str2, err = stub.GetState(checkingTab + "_" + args[0])
		if err != nil {
			bal_str2 = []byte(strconv.Itoa(BALANCE))
		}

		bal1, err = strconv.Atoi(string(bal_str1))
		bal2, err = strconv.Atoi(string(bal_str2))
		bal1 += bal2

		if err != nil {
			fmt.Printf("Error getBalance: %s", err)
			return nil, err
		}
		return nil, nil
	*/
}

func (t *SmallBank) updateBalance(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return nil, nil
	/*
		_, err1 := stub.GetState(accountTab + "_" + args[0])
		if err1 != nil {
			stub.PutState(accountTab+"_"+args[0], []byte(strconv.Itoa(BALANCE)))
		}
		bal_str, err2 := stub.GetState(checkingTab + "_" + args[0])
		if err2 != nil {
			stub.PutState(checkingTab+"_"+args[0], []byte(strconv.Itoa(BALANCE)))
		}
		_, err3 := stub.GetState(savingTab + "_" + args[0])
		if err3 != nil {
			stub.PutState(savingTab+"_"+args[0], []byte(strconv.Itoa(BALANCE)))
		}
		var bal1, bal2 int
		var err error
		bal1, err = strconv.Atoi(string(bal_str))
		bal2, err = strconv.Atoi(args[1])
		bal1 += bal2

		err = stub.PutState(checkingTab+"_"+args[0], []byte(strconv.Itoa(bal1)))

		if err != nil {
			fmt.Printf("Error updateBalance: %s", err)
			return nil, err
		}
		return nil, nil
	*/
}

func (t *SmallBank) updateSaving(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return nil, nil
	/*
		_, err1 := stub.GetState(accountTab + "_" + args[0])
		if err1 != nil {
			stub.PutState(accountTab+"_"+args[0], []byte(strconv.Itoa(BALANCE)))
		}
		_, err2 := stub.GetState(checkingTab + "_" + args[0])
		if err2 != nil {
			stub.PutState(checkingTab+"_"+args[0], []byte(strconv.Itoa(BALANCE)))
		}
		bal_str3, err3 := stub.GetState(savingTab + "_" + args[0])
		if err3 != nil {
			stub.PutState(savingTab+"_"+args[0], []byte(strconv.Itoa(BALANCE)))
		}
		var bal1, bal2 int
		var err error

		bal1, err = strconv.Atoi(string(bal_str3))
		bal2, err = strconv.Atoi(args[1])
		bal1 += bal2

		err = stub.PutState(savingTab+"_"+args[0], []byte(strconv.Itoa(bal1)))

		if err != nil {
			fmt.Printf("Error updateSaving: %s", err)
			return nil, err
		}
		return nil, nil
	*/
}

func (t *SmallBank) sendPayment(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return nil, nil
	/*
		bal_str1, err1 := stub.GetState(accountTab + "_" + args[0])
		if err1 != nil {
			stub.PutState(accountTab+"_"+args[0], []byte(strconv.Itoa(BALANCE)))
		}
		bal_str2, err2 := stub.GetState(checkingTab + "_" + args[0])
		if err2 != nil {
			stub.PutState(checkingTab+"_"+args[0], []byte(strconv.Itoa(BALANCE)))
		}
		_, err3 := stub.GetState(savingTab + "_" + args[0])
		if err3 != nil {
			stub.PutState(savingTab+"_"+args[0], []byte(strconv.Itoa(BALANCE)))
		}
		bal_str1, err1 = stub.GetState(accountTab + "_" + args[1])
		if err1 != nil {
			stub.PutState(accountTab+"_"+args[1], []byte(strconv.Itoa(BALANCE)))
		}
		bal_str2, err2 = stub.GetState(checkingTab + "_" + args[1])
		if err2 != nil {
			stub.PutState(checkingTab+"_"+args[1], []byte(strconv.Itoa(BALANCE)))
		}
		_, err3 = stub.GetState(savingTab + "_" + args[0])
		if err3 != nil {
			stub.PutState(savingTab+"_"+args[1], []byte(strconv.Itoa(BALANCE)))
		}

		var bal1, bal2, amount int
		var err error

		bal_str1, err = stub.GetState(checkingTab + "_" + args[0])
		if err != nil {
			bal_str1 = []byte(strconv.Itoa(BALANCE))
		}
		bal_str2, err = stub.GetState(checkingTab + "_" + args[1])
		if err != nil {
			bal_str2 = []byte(strconv.Itoa(BALANCE))
		}
		amount, err = strconv.Atoi(args[2])

		bal1, err = strconv.Atoi(string(bal_str1))
		bal2, err = strconv.Atoi(string(bal_str2))
		bal1 -= amount
		bal2 += amount

		err = stub.PutState(checkingTab+"_"+args[0], []byte(strconv.Itoa(bal1)))

		if err != nil {
			fmt.Printf("Error sendPayment: %s", err)
			return nil, err
		}

		err = stub.PutState(checkingTab+"_"+args[1], []byte(strconv.Itoa(bal2)))

		if err != nil {
			fmt.Printf("Error sendPayment: %s", err)
			return nil, err
		}
		return nil, nil
	*/
}

func (t *SmallBank) writeCheck(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return nil, nil
	/*
		_, err1 := stub.GetState(accountTab + "_" + args[0])
		if err1 != nil {
			stub.PutState(accountTab+"_"+args[0], []byte(strconv.Itoa(BALANCE)))
		}
		bal_str2, err2 := stub.GetState(checkingTab + "_" + args[0])
		if err2 != nil {
			stub.PutState(checkingTab+"_"+args[0], []byte(strconv.Itoa(BALANCE)))
		}
		bal_str3, err3 := stub.GetState(savingTab + "_" + args[0])
		if err3 != nil {
			stub.PutState(savingTab+"_"+args[0], []byte(strconv.Itoa(BALANCE)))
		}

		var bal1, bal2 int
		var err error
		var amount int
		bal1, err = strconv.Atoi(string(bal_str2))
		bal2, err = strconv.Atoi(string(bal_str3))
		amount, err = strconv.Atoi(args[1])
		if amount < bal1+bal2 {
			err = stub.PutState(checkingTab+"_"+args[0], []byte(strconv.Itoa(bal1-amount-1)))
		} else {
			err = stub.PutState(checkingTab+"_"+args[0], []byte(strconv.Itoa(bal1-amount)))
		}

		if err != nil {
			fmt.Printf("Error writeCheck: %s", err)
			return nil, err
		}
		return nil, nil
	*/
}
