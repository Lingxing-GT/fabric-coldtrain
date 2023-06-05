package main

import (
	"fmt"
	"github.com/Lingxing-GT/fabric-coldtrain/chaincode/api"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

func main() {
	timeLocal, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	time.Local = timeLocal
	dt := time.Now()
	fmt.Println(dt)
	err = shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting SmartContract: %s", err)
	}
}

// Init 初始化时会执行该方法
func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("*****SmartContract Init*****")
	// 初始化默认数据
	// 初始化养殖场
	var FarmOwner = "producer.com"
	var FarmIDs = [5]string{"FM001", "FM002", "FM003", "FM004", "FM005"}
	var Addrs = [5]string{"Changan, XiAn", "Donghu, Nanchang", "Zhanggong, Ganzhou", "Rongxian, Yulin", "Shunde, Foshan"}
	for i, val := range FarmIDs {
		var args = []string{val, Addrs[i], FarmOwner}
		fmt.Println(i)
		api.AddFarm(stub, args)
	}
	fmt.Println("养殖场初始化结束")
	// 初始化饲养员
	var BreederIDs = [5]string{"FM1B1", "FM2B1", "FM3B1", "FM4B1", "FM5B1"}
	var BNames = [5]string{"Zhang San", "Li Si", "Wang Wu", "Liu Er", "Ye Liu"}
	var Sex = [5]string{"male", "male", "female", "male", "female"}
	var Age = [5]string{"33", "25", "47", "38", "30"}
	for i, val := range BreederIDs {
		var args = []string{val, FarmIDs[i], BNames[i], Sex[i], Age[i]}
		api.AddBreeder(stub, args)
	}
	fmt.Println("饲养员初始化结束")
	// 初始化加工厂
	var FactoryIDs = [2]string{"FT001", "FT002"}
	var FactoryOwner = "processor.com"
	var FTAddrs = [2]string{"Haizhu, Guangzhou", "Jiangxia, Wuhan"}
	for i, val := range FactoryIDs {
		var args = []string{val, FTAddrs[i], FactoryOwner}
		api.AddFactory(stub, args)
	}
	fmt.Println("加工厂初始化结束")
	// 加入操作员
	var OperatorIDs = [2]string{"FT1OP1", "FT2OP1"}
	var ONames = [2]string{"LiLi", "YeYe"}
	var OSex = [2]string{"female", "female"}
	var OAge = [2]string{"29", "31"}
	for i, val := range OperatorIDs {
		var args = []string{val, FactoryIDs[i], ONames[i], OSex[i], OAge[i]}
		api.AddOperator(stub, args)
	}
	fmt.Println("操作员初始化结束")
	// 加入卡车
	var PlateNos = [3]string{"TK1", "TK2", "TK3"}
	var DriverIDs = [3]string{"D1", "D2", "D3"}
	var TruckOwner = "logtistics.com"
	for i, val := range PlateNos {
		var args = []string{val, DriverIDs[i], TruckOwner}
		api.AddTruck(stub, args)
	}
	fmt.Println("卡车初始化结束")
	//加入驾驶员
	var DNames = [3]string{"XiaoKa", "XiaoYe", "XiaoZhang"}
	var DAge = [3]string{"44", "39", "35"}
	for i, val := range DriverIDs {
		fmt.Println(i)
		var args = []string{val, DNames[i], DAge[i], TruckOwner}
		api.AddDriver(stub, args)
	}
	fmt.Println("驾驶员初始化结束")
	//加入市场
	var MarketIDs = [4]string{"MK001", "MK002", "MK003", "MK004"}
	var MKAddr = [4]string{"Xinfeng, Ganzhou", "Changan, XiAn", "Rongxian, Yulin", "Shunde, Foshan"}
	var MarketOwner = "retailer.com"
	for i, val := range MarketIDs {
		var args = []string{val, MKAddr[i], MarketOwner}
		api.AddMarket(stub, args)
	}
	fmt.Println("市场初始化结束")
	return shim.Success(nil)
}

// Invoke 智能合约的功能函数定义
func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "hello":
		return api.Hello(stub, args)
	case "addFarm":
		return api.AddFarm(stub, args)
	case "addBreeder":
		return api.AddBreeder(stub, args)
	case "addFactory":
		return api.AddFactory(stub, args)
	case "addCattle":
		return api.AddCattle(stub, args)
	case "addBeff":
		return api.AddBeff(stub, args)
	case "addCattleGrowInfo":
		return api.AddCattleGrowInfo(stub, args)
	case "deleteCattle":
		return api.DeleteCattle(stub, args)
	case "frozenProcess":
		return api.FrozenProcess(stub, args)
	case "addOperator":
		return api.AddOperator(stub, args)
	case "addTruck":
		return api.AddTruck(stub, args)
	case "addDriver":
		return api.AddDriver(stub, args)
	case "createWaybill":
		return api.CreateWaybill(stub, args)
	case "addWaybillInfo":
		return api.AddWaybillInfo(stub, args)
	case "addMarket":
		return api.AddMarket(stub, args)
	case "addRetailBeff":
		return api.AddRetailBeff(stub, args)
	case "queryByBeffID":
		return api.QueryByBeffID(stub, args)
	case "queryByWaybillNo":
		return api.QueryByWaybillNo(stub, args)
	case "queryByCattleID":
		return api.QueryByCattleID(stub, args)
	/*case "queryByFarmID":
		return api.QueryByFarmID(stub, args)
	case "queryByCattleID":
		return api.QueryByCattleID(stub, args)
	case "queryByWaybillNo":
		return api.QueryBywaybillNo(stub, args)
	case "queryByPlateNo":
		return api.QueryByPlateNo(stub, args)*/
	case "addSaleBill":
		return api.AddSaleBill(stub, args)
	default:
		return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))
	}
}
