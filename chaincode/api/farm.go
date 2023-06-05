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
func AddFarm(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("The number of args is unqualified! It should be 3.")
	}
	farmID := args[0]
	addr := args[1]
	owner := args[2]
	farm := &model.Farm{
		FarmID: farmID,
		Addr:   addr,
		Owner:  owner,
	}
	// 写入账本
	if err := utils.WriteLedger(farm, stub, model.Farmkey, []string{farmID, owner}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	return shim.Success(nil)
}

func AddBreeder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 5 {
		return shim.Error("The number of args is unqualified! It should be 5.")
	}
	breederID := args[0]
	farmID := args[1]
	name := args[2]
	sex := args[3]
	age, err := strconv.Atoi(args[4])
	if err != nil {
		return shim.Error("The age is illegal!")
	}
	breeder := &model.Breeder{
		BreederID: breederID,
		FarmID:    farmID,
		Name:      name,
		Sex:       sex,
		Age:       age,
	}
	if err := utils.WriteLedger(breeder, stub, model.Breederkey, []string{breederID, farmID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	return shim.Success(nil)
}

//add cattle, the order of args is "CattleID, FarmID, BreederID, Remarks(optional)"
func AddCattle(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 3 || len(args) > 4 {
		return shim.Error("The number of args is unqualified! It should be 3 or 4.")
	}

	cattleID := args[0]
	farmID := args[1]
	breederID := args[2]
	//dayin := args[3]
	createTime, _ := stub.GetTxTimestamp()
	dayin := time.Unix(int64(createTime.GetSeconds()), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	remarks := "NONE"
	if len(args) == 4 {
		remarks = args[3]
	}

	cattle := &model.Cattle{
		CattleID:  cattleID,
		FarmID:    farmID,
		BreederID: breederID,
		DayIn:     dayin,
		DayOut:    "NONE",
		Remarks:   remarks,
	}
	if err := utils.WriteLedger(cattle, stub, model.Cattlekey, []string{cattleID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	cattleBytes, err := json.Marshal(cattle)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	return shim.Success(cattleBytes)
}

//delete cattle， the order of args is (CattleID, DayOut, Remarks)
func DeleteCattle(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("The number of args is unqualified! It should be 2.")
	}
	cattleID := args[0]
	createTime, _ := stub.GetTxTimestamp()
	dayout := time.Unix(int64(createTime.GetSeconds()), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	remarks := args[1]

	//Find the cattle by cattleID
	resultsCattle, err := utils.GetStateByPartialCompositeKeys2(stub, model.Cattlekey, []string{cattleID})
	fmt.Println("DeleteCattle, resultsCattle =", resultsCattle)
	if err != nil || len(resultsCattle) != 1 {
		return shim.Error(fmt.Sprintf("Error finding cattle %s: %s", cattleID, err))
	}
	var realCattle model.Cattle
	if err = json.Unmarshal(resultsCattle[0], &realCattle); err != nil {
		return shim.Error(fmt.Sprintf("DeleteCattle-Error Unmarshaling: %s", err))
	}

	//Determine if the cattle has been deleted
	if realCattle.DayOut != "NONE" {
		return shim.Error("This cattle has been deleted! Can't delete twice.")
	}
	//delete the cattle
	realCattle.DayOut = dayout
	realCattle.Remarks = remarks
	if err := utils.WriteLedger(realCattle, stub, model.Cattlekey, []string{cattleID, realCattle.FarmID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	return shim.Success(nil)
}

//add cattle growth record, the order of args is "CattleID, TEMP, Health, Weight, Remarks(optional)"
func AddCattleGrowInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 4 || len(args) > 5 {
		return shim.Error("The number of args is unqualified! It should be 4 or 5.")
	}
	cattleID := args[0]
	createTime, _ := stub.GetTxTimestamp()
	var recordTimes []string
	var temps []float64
	var healths []string
	var remarks []string
	var weights []float64
	recordTime := time.Unix(int64(createTime.GetSeconds()), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	temp, err := strconv.ParseFloat(args[1], 64)
	recordTimes = append(recordTimes, recordTime)
	temps = append(temps, temp)
	if err != nil {
		return shim.Error(fmt.Sprintf("the temperature is illegal: %s", err))
	}
	health := args[2]
	healths = append(healths, health)
	weight, err := strconv.ParseFloat(args[3], 64)
	weights = append(weights, weight)
	if err != nil {
		return shim.Error(fmt.Sprintf("the weight is illegal: %s", err))
	}
	remark := "NONE"
	if len(args) == 5 {
		remark = args[4]
	}
	remarks = append(remarks, remark)
	cattleGrowInfo := &model.CattleGrowInfo{
		CattleID:   cattleID,
		RecordTime: recordTimes,
		TEMP:       temps,
		Health:     healths,
		Weight:     weights,
		Remarks:    remarks,
	}
	if err := utils.WriteLedger(cattleGrowInfo, stub, model.CattleGrowkey, []string{cattleID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	return shim.Success(nil)
}

//Add Beff, the order of args is "BeffID, CattleID, Quality, Weight"
func AddBeff(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 4 {
		return shim.Error("The number of args is unqualified! It should be 4.")
	}
	beffID := args[0]
	cattleID := args[1]
	quality, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error(fmt.Sprintf("the quality of beff is illegal: %s", err))
	}
	weight, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		return shim.Error(fmt.Sprintf("the weight of beef is illegal: %s", err))
	}
	createTime, _ := stub.GetTxTimestamp()
	time := time.Unix(int64(createTime.GetSeconds()), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	status := false
	beff := &model.Beff{
		BeffID:   beffID,
		CattleID: cattleID,
		Quality:  quality,
		Weight:   weight,
		Time:     time,
		Status:   status,
	}
	if err := utils.WriteLedger(beff, stub, model.Beffkey, []string{beffID, cattleID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	return shim.Success(nil)
}

//to query cattle grow info by cattleID
func QueryByCattleID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		shim.Error("please enter a CattleID")
	}
	cattleID := args[0]
	//Find the waybill by number
	resultGrowInfo, err := utils.GetStateByPartialCompositeKeys2(stub, model.CattleGrowkey, []string{cattleID})
	if err != nil || len(resultGrowInfo) != 1 {
		return shim.Error(fmt.Sprintf("Error finding CattleGrowInfo %s: %s", cattleID, err))
	}
	return shim.Success(resultGrowInfo[0])
}
