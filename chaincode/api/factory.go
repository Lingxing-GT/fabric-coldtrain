package api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Lingxing-GT/fabric-coldtrain/chaincode/model"
	"github.com/Lingxing-GT/fabric-coldtrain/chaincode/pkg/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//三个参数:FarmID,Addr,Owner
func AddFactory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("The number of args is unqualified! It should be 3.")
	}
	factoryID := args[0]
	addr := args[1]
	owner := args[2]
	factory := &model.Factory{
		FactoryID: factoryID,
		Addr:      addr,
		Owner:     owner,
	}
	// 写入账本
	if err := utils.WriteLedger(factory, stub, model.Factorykey, []string{factoryID, owner}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	return shim.Success(nil)
}

func AddOperator(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 5 {
		return shim.Error("The number of args is unqualified! It should be 5.")
	}
	operatorID := args[0]
	factoryID := args[1]
	name := args[2]
	sex := args[3]
	age, err := strconv.Atoi(args[4])
	if err != nil {
		return shim.Error("The age is illegal!")
	}
	operator := &model.Operator{
		OperatorID: operatorID,
		FactoryID:  factoryID,
		Name:       name,
		Sex:        sex,
		Age:        age,
	}
	if err := utils.WriteLedger(operator, stub, model.Operatorkey, []string{operatorID, factoryID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	return shim.Success(nil)
}

//to process and frozen the beff, the order of args is (BeffID, OperatorID, FactoryID, Weight)
func FrozenProcess(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 5 {
		return shim.Error("The number of args is unqualified! It should be 5.")
	}
	beffID := args[0]
	operatorID := args[1]
	factoryID := args[2]
	weight, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		return shim.Error("the weight of post-processing beff is illegal!")
	}
	createTime, _ := stub.GetTxTimestamp()
	time := time.Unix(int64(createTime.GetSeconds()), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")

	//Find the beff by beffID
	resultBeff, err := utils.GetStateByPartialCompositeKeys2(stub, model.Beffkey, []string{beffID})
	if err != nil || len(resultBeff) != 1 {
		return shim.Error(fmt.Sprintf("Error finding beff %s: %s", beffID, err))
	}
	var realBeff model.Beff
	if err = json.Unmarshal(resultBeff[0], &realBeff); err != nil {
		return shim.Error(fmt.Sprintf("FozenProcess-Error Unmarshaling: %s", err))
	}
	realBeff.Status = true
	postBeff := model.PostBeff{
		BeffID:     beffID,
		OperatorID: operatorID,
		FactoryID:  factoryID,
		Weight:     weight,
		Time:       time,
		Life:       365,
		Quality:    realBeff.Quality,
		CattleID:   realBeff.CattleID,
		Status:     false,
	}
	if err := utils.WriteLedger(realBeff, stub, model.Beffkey, []string{beffID, realBeff.CattleID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err := utils.WriteLedger(postBeff, stub, model.PostBeffkey, []string{beffID, factoryID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	return shim.Success(nil)
}
