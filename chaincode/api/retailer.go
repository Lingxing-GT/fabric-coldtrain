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

//Three args:MarketID,Addr,Owner
func AddMarket(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("The number of args is unqualified! It should be 3.")
	}
	marketID := args[0]
	addr := args[1]
	owner := args[2]
	market := &model.Market{
		MarketID: marketID,
		Addr:     addr,
		Owner:    owner,
	}
	// 写入账本
	if err := utils.WriteLedger(market, stub, model.Marketkey, []string{marketID, owner}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	return shim.Success(nil)
}

//to add a retail beff, the order of args is "BeffID, MarketID, WaybillNo, Price"
func AddRetailBeff(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("The number of args is unqualified! It should be 4.")
	}
	beffID := args[0]
	marketID := args[1]
	waybillNo := args[2]
	price, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		return shim.Error(fmt.Sprintf("the price is illegal: %s", err))
	}
	//Find the postbeff by beffID
	resultPostBeff, err := utils.GetStateByPartialCompositeKeys2(stub, model.PostBeffkey, []string{beffID})
	if err != nil || len(resultPostBeff) != 1 {
		return shim.Error(fmt.Sprintf("Error finding post processing beff %s: %s", beffID, err))
	}
	//Find the beff by beffID
	resultBeff, err := utils.GetStateByPartialCompositeKeys2(stub, model.Beffkey, []string{beffID})
	if err != nil || len(resultBeff) != 1 {
		return shim.Error(fmt.Sprintf("Error finding beff %s: %s", beffID, err))
	}

	var beff model.Beff
	var postBeff model.PostBeff
	if err = json.Unmarshal(resultPostBeff[0], &postBeff); err != nil {
		return shim.Error(fmt.Sprintf("AddRetailBeff-Error Unmarshaling: %s", err))
	}
	if err = json.Unmarshal(resultBeff[0], &beff); err != nil {
		return shim.Error(fmt.Sprintf("AddRetailBeff-Error Unmarshaling: %s", err))
	}
	retailBeff := &model.RetailBeff{
		BeffID:    beffID,
		CattleID:  beff.CattleID,
		FactoryID: postBeff.FactoryID,
		MarketID:  marketID,
		WaybillNo: waybillNo,
		Quality:   beff.Quality,
		Weight:    postBeff.Weight,
		Price:     price,
		Time:      postBeff.Time,
		Life:      postBeff.Life,
		Status:    false,
	}
	if err := utils.WriteLedger(retailBeff, stub, model.RetailBeffkey, []string{beffID, marketID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	return shim.Success(nil)
}

//to add a sale bill, the order of args is "BeffID, BillID"
func AddSaleBill(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("The number of args is unqualified! It should be 2.")
	}
	beffID := args[0]
	billNo := args[1]
	createTime, _ := stub.GetTxTimestamp()
	time := time.Unix(int64(createTime.GetSeconds()), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")

	//Find the retail beff by beffID
	resultRetailBeff, err := utils.GetStateByPartialCompositeKeys2(stub, model.RetailBeffkey, []string{beffID})
	if err != nil || len(resultRetailBeff) != 1 {
		return shim.Error(fmt.Sprintf("Error finding retail beff %s: %s", beffID, err))
	}

	var retailBeff model.RetailBeff
	if err = json.Unmarshal(resultRetailBeff[0], &retailBeff); err != nil {
		return shim.Error(fmt.Sprintf("AddSaleBill-Error Unmarshaling: %s", err))
	}
	retailBeff.Status = true

	salebill := model.SaleBill{
		BeffID:   beffID,
		BillNo:   billNo,
		Price:    retailBeff.Price,
		MarketID: retailBeff.MarketID,
		Time:     time,
	}
	if err := utils.WriteLedger(retailBeff, stub, model.RetailBeffkey, []string{beffID, retailBeff.MarketID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if err := utils.WriteLedger(salebill, stub, model.SaleBillkey, []string{beffID, billNo}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	return shim.Success(nil)
}

//to query beff info and trans Info by beffID
func QueryByBeffID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		shim.Error("please enter a BeffID")
	}
	beffID := args[0]
	//Find the retail beff by beffID
	resultRetailBeff, err := utils.GetStateByPartialCompositeKeys2(stub, model.RetailBeffkey, []string{beffID})
	if err != nil || len(resultRetailBeff) != 1 {
		return shim.Error(fmt.Sprintf("Error finding retail beff %s: %s", beffID, err))
	}
	var retailBeff model.RetailBeff
	if err = json.Unmarshal(resultRetailBeff[0], &retailBeff); err != nil {
		return shim.Error(fmt.Sprintf("QueryByBeffID-Error Unmarshaling: %s", err))
	}
	//Find the sale bill by beffID
	resultSaleBill, err := utils.GetStateByPartialCompositeKeys2(stub, model.SaleBillkey, []string{beffID})
	if err != nil || len(resultRetailBeff) != 1 {
		return shim.Error(fmt.Sprintf("Error finding sale bill by beffID %s: %s", beffID, err))
	}
	var saleBill model.SaleBill
	if err = json.Unmarshal(resultSaleBill[0], &saleBill); err != nil {
		return shim.Error(fmt.Sprintf("QueryByBeffID-Error Unmarshaling: %s", err))
	}
	//Find the beff by beffID
	resultBeff, err := utils.GetStateByPartialCompositeKeys2(stub, model.Beffkey, []string{beffID})
	if err != nil || len(resultRetailBeff) != 1 {
		return shim.Error(fmt.Sprintf("Error finding beff %s: %s", beffID, err))
	}
	var beff model.Beff
	if err = json.Unmarshal(resultBeff[0], &beff); err != nil {
		return shim.Error(fmt.Sprintf("QueryByBeffID-Error Unmarshaling: %s", err))
	}
	//Find the cattle by CattleID
	resultCattle, err := utils.GetStateByPartialCompositeKeys2(stub, model.Cattlekey, []string{beff.CattleID})
	if err != nil || len(resultRetailBeff) != 1 {
		return shim.Error(fmt.Sprintf("Error finding cattle %s: %s", beff.CattleID, err))
	}
	var cattle model.Cattle
	if err = json.Unmarshal(resultCattle[0], &cattle); err != nil {
		return shim.Error(fmt.Sprintf("QueryByBeffID-Error Unmarshaling: %s", err))
	}

	returnInfo := model.QueryBeffID{
		BeffID:      beffID,
		Quality:     retailBeff.Quality,
		Weight:      retailBeff.Weight,
		Price:       saleBill.Price,
		TimeProcess: retailBeff.Time,
		Life:        retailBeff.Life,
		FarmID:      cattle.FarmID,
		CattleID:    cattle.CattleID,
		FactoryID:   retailBeff.FactoryID,
		WaybillNo:   retailBeff.WaybillNo,
		BillNo:      saleBill.BillNo,
		MarketID:    saleBill.MarketID,
		TimeSale:    saleBill.Time,
	}
	returnInfoByte, err := json.Marshal(returnInfo)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryByBeffID-Error Marshaling: %s", err))
	}
	return shim.Success(returnInfoByte)
}
