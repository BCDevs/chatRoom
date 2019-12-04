package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type Marks struct {
	Biology string `json:"biology"`
	Physics string `json:"physics"`
	Chemestry  string `json:"chemestry"`
	English  string `json:"english"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()
	if function == "addMarks" {
		return s.addMarks(APIstub, args)
	} else if function == "getMarks" {
		return s.getMarks(APIstub, args)
	} else if function == "deleteMarks" {
		return s.deleteMarks(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) getMarks(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	marksAsBytes, _ := APIstub.GetState(args[0])
	if marksAsBytes == nil {
		return shim.Error("Student with this Id doesnot exist")
	}
	return shim.Success(marksAsBytes)
}
func (s *SmartContract) deleteMarks(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	marksAsBytes, _ := APIstub.DeleteState(args[0])

}


func (s *SmartContract) addMarks(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var studentMarks = Marks{ Biology: args[1], Physics: args[2], Chemestry: args[3], English: args[4] }

	marksAsBytes, _ := json.Marshal(studentMarks)
	err := APIstub.PutState(args[0], marksAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record Student Marks: %s", args[0]))
	}

	return shim.Success(nil)
}

func main() {

	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
