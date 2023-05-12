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

//Three args:PlateNo,DriverID,Owner
func AddTruck(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("The number of args is unqualified! It should be 3.")
	}
	plateNo := args[0]
	driverID := args[1]
	owner := args[2]
	truck := &model.Truck{
		PlateNo:  plateNo,
		DriverID: driverID,
		Owner:    owner,
	}
	// 写入账本
	if err := utils.WriteLedger(truck, stub, model.Truckkey, []string{plateNo, owner}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	return shim.Success(nil)
}

//three args:ID,owner,name,age
func AddDriver(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("The number of args is unqualified! It should be 4.")
	}
	driverID := args[0]
	owner := args[1]
	name := args[2]
	age, err := strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("The age is illegal!")
	}
	driver := &model.Driver{
		ID:    driverID,
		Owner: owner,
		Name:  name,
		Age:   age,
	}
	if err := utils.WriteLedger(driver, stub, model.Driverkey, []string{driverID, owner}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	return shim.Success(nil)
}

//to create a new waybill, the order of args is "WaybillNo, PlateNo, BeginCity, BeffID"
func CreateWaybill(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("The number of args is unqualified! It should be 4.")
	}
	waybillNo := args[0]
	plateNo := args[1]
	nodes := []string{args[2]}
	beffID := args[3]
	var ctemp []float64
	createTime, _ := stub.GetTxTimestamp()
	time := time.Unix(int64(createTime.GetSeconds()), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	timeline := []string{time}
	//Find the post Beef by BeffID
	resultPostBeff, err := utils.GetStateByPartialCompositeKeys2(stub, model.PostBeffkey, []string{beffID})
	if err != nil || len(resultPostBeff) != 1 {
		return shim.Error(fmt.Sprintf("Error finding postBeff %s: %s", beffID, err))
	}
	var postBeff model.PostBeff
	if err = json.Unmarshal(resultPostBeff[0], &postBeff); err != nil {
		return shim.Error(fmt.Sprintf("CreateWaybill-Error Unmarshaling: %s", err))
	}
	postBeff.Status = true
	if err := utils.WriteLedger(postBeff, stub, model.PostBeffkey, []string{beffID, postBeff.FactoryID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	waybill := model.Waybill{
		WaybillNo: waybillNo,
		PlateNo:   plateNo,
		Nodes:     nodes,
		Timeline:  timeline,
		CTEMP:     ctemp,
		BeffID:    beffID,
		Status:    0,
	}
	if err := utils.WriteLedger(waybill, stub, model.Waybillkey, []string{waybillNo, beffID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	return shim.Success(nil)
}

//to add a waybill info, the order of args is "WaybillNo, Node, CTEMP, Status(end)"
func AddWaybillInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 3 || len(args) > 4 {
		return shim.Error("The number of args is unqualified! It should be 3 or 4.")
	}
	waybillNo := args[0]
	node := args[1]
	ctemp, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return shim.Error(fmt.Sprintf("the charrier temperature is illegal!"))
	}
	createTime, _ := stub.GetTxTimestamp()
	time := time.Unix(int64(createTime.GetSeconds()), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	//Find the waybill by number
	resultBill, err := utils.GetStateByPartialCompositeKeys2(stub, model.Waybillkey, []string{waybillNo})
	if err != nil || len(resultBill) != 1 {
		return shim.Error(fmt.Sprintf("Error finding waybill %s: %s", waybillNo, err))
	}
	var realBill model.Waybill
	if err = json.Unmarshal(resultBill[0], &realBill); err != nil {
		return shim.Error(fmt.Sprintf("AddWaybillInfo-Error Unmarshaling: %s", err))
	}
	if realBill.Status == 1 {
		return shim.Error("The waybill has been over!")
	}
	realBill.Nodes = append(realBill.Nodes, node)
	realBill.Timeline = append(realBill.Timeline, time)
	realBill.CTEMP = append(realBill.CTEMP, ctemp)
	if len(args) == 4 {
		realBill.Status, err = strconv.Atoi(args[3])
		if err != nil {
			return shim.Error(fmt.Sprintf("the status of waybill is illegal!"))
		}
	}
	if err := utils.WriteLedger(realBill, stub, model.Waybillkey, []string{waybillNo, realBill.BeffID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	return shim.Success(nil)
}
