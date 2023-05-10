package model

type Breeder struct {
	BreederID string `json:"BreederID"`
	FarmID    string `json:"FarmID"`
	Name      string `json:"UserName"`
	Sex       string `json:"Sex"`
	Age       int    `json:"Age"`
}

//————————————————————————————————————————————————————————————————————————————————————————————//
//食品原产地生产商节点producer
//肉牛信息
type Cattle struct {
	CattleID  string `json:"CattleID"`
	FarmID    string `json:"FarmID"`
	BreederID string `json:"BreederID"`
	DayIn     string `json:"EntryTime"`
	DayOut    string `json:"MarketingTime"`
	Remarks   string `json:"Remakrs"` //CSS meaning "cattle slaughter and sectioning"; DEAD meaning "dead and not on the market""
}

//肉牛成长记录
type CattleGrowInfo struct {
	CattleID   string  `json:"CattleID"`
	RecordTime string  `json:"RecordTime"`
	TEMP       float64 `json:"Temperature"`
	Health     string  `json:"Health"` //3 Rank: Good, Ill, Emergency
	Weight     float64 `json:"Weight"`
	Remarks    string  `json:"Remarks"`
}

//牛肉信息（销售单位）
type Beff struct {
	BeffID   string  `json:"BeffID"`
	CattleID string  `json:"FromCattleID"`
	Quality  int     `json:"Quality"` //分为5级，0为特级，1-4级数字越大品质越低
	Weight   float64 `json:"Weight"`
	//Price    float64 `json:"Price"`
	Time   string `json:"ProducingTime"`
	Status bool   `json:"Status"` //交易状态，true表示已运出，false表示存在仓库
}

//生产商信息
type Farm struct {
	FarmID string `json:"FarmID"`
	Addr   string `json:"Address"`
	Owner  string `json:"Owner"` //所属组织
}

//————————————————————————————————————————————————————————————————————————————————————————————//
//加工产商Processor
type Factory struct {
	FactoryID string `json:"FactoryID"`
	Addr      string `json:"Address"`
	Owner     string `json:"Owner"` //所属组织
}

type Operator struct {
	OperatorID string `json:"OperatorID"`
	FactoryID  string `json:"FactoryID"`
	Name       string `json:"UserName"`
	Sex        string `json:"Sex"`
	Age        int    `json:"Age"`
}

type PostBeff struct {
	BeffID     string  `json:"BeffID"`
	OperatorID string  `json:"OperatorID"`
	FactoryID  string  `json:"FactoryID"`
	Weight     float64 `json:"Weight"`
	//Price    float64 `json:"Price"`
	Time     string `json:"ProcessingTime"`
	Life     int    `json:"Shelf-Life"`
	Quality  int    `json:"Quality"`
	CattleID string `json:"FromCattleID"`
	FarmID   string `json:"FarmID"`
	Status   bool   `json:"Status"`
}

//————————————————————————————————————————————————————————————————————————————————————————————//
//冷链运输Logitics
type Truck struct {
	PlateNo  string `json:"PlateNumber"`
	DriverID string `json:"DriverID"`
	Owner    string `json:"Owner"`
}

type Driver struct {
	ID    string `json:"DriverID"`
	Name  string `json:"DriverName"`
	Age   int    `json:"Age"`
	Owner string `json:"Owner"`
}

type Waybill struct {
	WaybillNo string    `json:"WaybillNumber"`
	PlateNo   string    `json:"PlateNumber"`
	Nodes     []string  `json:"Nodes"`
	Timeline  []string  `json:"Timeline"`
	CTEMP     []float64 `json:"Carriage Temperature"`
	BeffID    string    `json:"ProductID"`
	Status    int       `json:"Status"` //0正在运输，1已签收
}

//———————————————————————————————————————————————————————————————————————————————————————————//
//零售商Retailer
type Market struct {
	MarketID string `json:"MarketID"`
	Addr     string `json:"Address"`
	Owner    string `json:"Owner"`
}

type RetailBeff struct {
	BeffID    string  `json:"BeffID"`
	CattleID  string  `json:"FromCattleID"`
	FarmID    string  `json:"FarmID"`
	FactoryID string  `json:"FactoryID"`
	MarketID  string  `json:"MarketID"`
	WaybillNo string  `json:"WaybillNumber"`
	Quality   int     `json:"Quality"`
	Weight    float64 `json:"Weight"`
	Price     float64 `json:"Price"`
	Time      string  `json:"ProcessingTime"`
	Life      int     `json:"Shelf-Life"`
	Status    bool    `json:"Status"`
}

type SaleBill struct {
	BeffID   string  `json:"RetailBeffID"`
	BillNo   string  `json:"SaleBillNumber"`
	MarketID string  `json:"MarketID"`
	Price    float64 `json:"Price"`
	Time     string  `json:"SalingTime"`
}

type QueryBeffID struct {
	BeffID      string  `json:"RetailBeffID"`
	Quality     int     `json:"Quality"`
	Weight      float64 `json:"Weight"`
	Price       float64 `json:"Price"`
	TimeProcess string  `json:"ProcessingTime"`
	Life        int     `json:"Shelf-Life(days)"`
	FarmID      string  `json:"FarmID"`
	CattleID    string  `json:"FromCattleID"`
	FactoryID   string  `json:"FactoryID"`
	WaybillNo   string  `json:"WaybillNumber"`
	BillNo      string  `json:"SaleBillNumber"`
	MarketID    string  `json:"MarketID"`
	TimeSale    string  `json:"SalingTime"`
}

const (
	Farmkey       = "Farm-key"
	Breederkey    = "Breeder-key"
	Cattlekey     = "Cattle-key"
	CattleGrowkey = "CatttleGrowInfo-key"
	Beffkey       = "Beff-Key"
	PostBeffkey   = "PostBeff-key"
	Factorykey    = "Factory-key"
	Operatorkey   = "Operator-key"
	Marketkey     = "Market-key"
	RetailBeffkey = "RetailBeff-key"
	Waybillkey    = "Waybill-key"
	Truckkey      = "Truck-key"
	Driverkey     = "Driver-key"
	SaleBillkey   = "SaleBill-key"
)
