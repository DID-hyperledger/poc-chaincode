package main

import (
	"fmt"
	"strconv"
	"crypto/sha256"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

 
type PerfTestChaincode struct {
}
 
func (t *PerfTestChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
    return shim.Success(nil)
}
 
func (t *PerfTestChaincode) Query(stub shim.ChaincodeStubInterface) peer.Response {
    return shim.Success(nil)
}
 
func (t *PerfTestChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
    if function == "writeBlock"{
		return t.writeBlock(stub, args)
	} else {
		return shim.Error("Invalid invoke function name")
	}
}

func (t *PerfTestChaincode) writeBlock(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 3 {
		fmt.Printf("Invalid number of argument")
		return shim.Error("Incorrect number of arguments")
	}

	var blockId = args[0]
	var blockData = args[1]
	var shaOperations = args[2]

	// Run SHA loop as specify

	operations, shaErr := strconv.Atoi(shaOperations)
	if shaErr != nil {
		fmt.Printf("Expect integer for number of SHA operations")
		return shim.Error("Argument Invalid")
	}

	for i:=0; i < operations; i++ {
		dummy := sha256.New()
		dummy.Write([]byte(blockData))
	}

	// Write transaction
	err := stub.PutState(blockId, []byte(blockData))
	if err != nil {
		fmt.Printf("Could not write block")
		return shim.Error("Cannot write block")
	}

	fmt.Printf("Successfully saved new block")
	return shim.Success(nil)
}
 
func main() {
    err := shim.Start(new(PerfTestChaincode))
    if err != nil {
        fmt.Println("Could not start PerfTestChaincode")
    } else {
        fmt.Println("SampleChaincode successfully started")
    }
 
}