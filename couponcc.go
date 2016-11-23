package main

import (
	//"bytes"
	"encoding/json"
	//"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"log"
	"time"
	//"strconv"
	//"strings"
)

const (
	COUPON_BATCH_KEY          = "CPB_"
	COUPON_KEY                = "CP_"
	COUPONS_OF_BATCH_KEY      = "COB_"
	COUPONBATCH_OF_USER_KEY   = "CBOU_"
	COUPONS_OF_USER_BATCH_KEY = "COUB_"
)

//优惠券批次信息
type CouponBatch struct {
	Id   uint64 `json:"id,omitempty"`
	Sn   string `json:"sn,omitempty"`
	Name string `json:"name,omitempty"`
	//1->商家优惠券，2->平台优惠券
	BatchType int `json:"batchType,omitempty"`
	//0普通优惠券 ，1下单立减券，2折扣券，3一元换购
	CouponType int `json:"couponType,omitempty"`
	//面额
	Denomination int `json:"denomination,omitempty"`
	//生效时间
	EffectiveDate uint64 `json:"effectiveDate,omitempty"`
	//失效时间
	ExpiringDate uint64 `json:"expiringDate,omitempty"`
	//可申请开始时间
	ApplyStartDate uint64 `json:"applyStartDate,omitempty"`
	//可申请结束时间
	ApplyEndDate uint64 `json:"applyEndDate,omitempty"`
	//发布时间
	PublishDate uint64 `json:"publishDate,omitempty"`
	BudgetNo    string `json:"budgetNo,omitempty"`
	//渠道号
	Channel uint64 `json:"channel,omitempty"`
	//渠道名称
	ChannelName    string `json:"channelName,omitempty"`
	Sku            string `json:"sku,omitempty"`
	Desc           string `json:"desc,omitempty"`
	CouponAmount   int    `json:"couponAmount,omitempty"`
	AppliedAmount  int    `json:"appliedAmount,omitempty"`
	SentAmount     int    `json:"sentAmount,omitempty"`
	ConsumedAmount int    `json:"consumedAmount,omitempty"`
	//优惠券起用订单价格
	OrderPriceLimit int `json:"orderPriceLimit,omitempty"`
	//优惠券起用条件：订单满多少件商品可用
	OrderQuantityLimit int `json:"orderQuantityLimit,omitempty"`
	//每个订单最多可使用张数
	MaximumUsage int    `json:"maximumUsage,omitempty"`
	ApplyLimit   int    `json:"applyLimit,omitempty"`
	Creator      uint64 `json:"creator,omitempty"`
	CreateDate   uint64 `json:"createDate,omitempty"`
	Disabled     bool   `json:"disabled,omitempty"`
	ShopId       uint64 `json:"shopId,omitempty"`
	ShopName     string `json:"shopName,omitempty"`
	Picture      string `json:"picture,omitempty"`
	//状态，0->未发布,1->已发布，未生效,2->已生效,-1->失效,-2->停发（只能用券，不能送券）
	Status       int `json:"status,omitempty"`
	DiscountRate int `json:"discountRate,omitempty"`
}

//优惠券信息
type Coupon struct {
	Id          uint64 `json:"id,omitempty"`
	BatchSn     string `json:"batchSn,omitempty"`
	Sn          string `json:"sn,omitempty"`
	ReceiveType int    `json:"receiveType,omitempty"`
	CreateDate  uint64 `json:"createDate,omitempty"`
	Creator     uint64 `json:"creator,omitempty"`
	Owner       uint64 `json:"owner,omitempty"`
	Status      int    `json:"status,omitempty"`
}

//优惠券批次dto
type CouponBatchDto struct {
	Id                 uint64                 `json:"id,omitempty"`
	BatchSn            string                 `json:"batchSn,omitempty"`
	BatchType          int                    `json:"omitempty,omitempty"`
	BatchTypeDesc      string                 `json:"batchTypeDesc,omitempty"`
	EffectiveStartTime uint64                 `json:"effectiveStartTime,omitempty"`
	EffectiveEndTime   uint64                 `json:"effectiveEndTime,omitempty"`
	Money              int                    `json:"money,omitempty"`
	ShopId             uint64                 `json:"shopId,omitempty"`
	UsageRuleType      int                    `json:"usageRuleType,omitempty"`
	UsageRule          map[string]interface{} `json:"usageRule,omitempty"`
	Quantity           map[string]interface{} `json:"quantity,omitempty"`
	BatchName          string                 `json:"batchName,omitempty"`
	Status             int                    `json:"status,omitempty"`
	StatusDesc         string                 `json:"statusDesc,omitempty"`
	Description        string                 `json:"description,omitempty"`
	Picture            string                 `json:"picture,omitempty"`
	PublishTime        uint64                 `json:"publishTime,omitempty"`
}

// 优惠券dto
type CouponInfoDto struct {
	Id          uint64 `json:"id,omitempty"`
	BatchSn     string `json:"batchSn,omitempty"`
	CouponSn    string `json:"couponSn,omitempty"`
	UserId      uint64 `json:"userId,omitempty"`
	Status      int    `json:"status,omitempty"`
	StatusDesc  string `json:"statusDesc,omitempty"`
	ReceiveType int    `json:"receiveType,omitempty"`
}

type CouponChaincode struct {
}

/**
 *
 */
func (cc *CouponChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("couponcc init")
	return nil, nil
}

func (cc *CouponChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("couponcc Invoke")
	ccStub := stub.(*shim.ChaincodeStub)
	if "createCouponBatch" == function {
		re, err := cc.createCouponBatch(ccStub, args)
		return re, err
	} else if "applyCoupon" == function {
		re, err := cc.applyCoupon(ccStub, args)
	} else if "sendCoupon" == function {

	} else if "consumeCoupon" == function {

	} else if "disableCouponBatch" == function {

	} else if "publishCouponBatch" == function {

	} else {

	}
	return nil, nil
}

func (cc *CouponChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	ccStub := stub.(*shim.ChaincodeStub)
	if "queryCouponBatch" == function {
		return cc.queryCouponBatch(ccStub, args)
	}
	return nil, nil

}

func (cc *CouponChaincode) queryCouponBatch(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	fmt.Printf("in queryCouponBatch")
	if len(args) != 1 {
		return nil, fmt.Errorf("Invalid number of arguments")
	}
	sn := args[0]
	cb, err := cc.getCouponBatch(stub, sn)
	if err != nil {
		return nil, fmt.Errorf("Get Coupon Batch error: %v", err)
	}
	fmt.Printf("coupon batch: %v", cb)
	dto := cc.cb2Dto(cb)
	bytes, err := json.Marshal(dto)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal error: %v", err)
	}
	return bytes, nil

}

//couponbatch dto --> cb
func (cc *CouponChaincode) cbDto2Cb(cbDto *CouponBatchDto) *CouponBatch {
	cb := &CouponBatch{Id: cbDto.Id, Sn: cbDto.BatchSn, Name: cbDto.BatchName,
		Type: cbDto.BatchType, Denomination: cbDto.Money,
		EffectiveDate: cbDto.EffectiveStartTime, ExpiringDate: cbDto.EffectiveEndTime,
		//ApplyStartDate:,ApplyEndDate:,BudgetNo:,
		PublishDate: cbDto.PublishTime, Desc: cbDto.Description, ShopId: cbDto.ShopId,
		Picture: cbDto.Picture, Status: cbDto.Status}
	return cb

}

//couponbatch cb --> dto
func (cc *CouponChaincode) cb2Dto(cb *CouponBatch) *CouponBatchDto {
	dto := &CouponBatchDto{
		Id: cb.Id, BatchSn: cb.Sn, BatchType: cb.Type,
		EffectiveStartTime: cb.EffectiveDate, EffectiveEndTime: cb.ExpiringDate,
		Money: cb.Denomination, ShopId: cb.ShopId, BatchName: cb.Name, Status: cb.Status,
		Description: cb.Desc, Picture: cb.Picture, PublishTime: cb.PublishDate}
	return dto
}

//coupon dto --> cp
func (cc *CouponChaincode) cpDto2Cp(dto *CouponInfoDto) *Coupon {
	cp := &Coupon{Id: dto.Id, BatchSn: dto.BatchSn, Sn: "",
		ReceiveType: dto.ReceiveType, Owner: dto.UserId, Status: dto.Status}
	return cp
}

//coupon cp --> dto
func (cc *CouponChaincode) cp2Dto(cp *Coupon) *CouponInfoDto {
	dto := &CouponInfoDto{Id: cp.Id, BatchSn: cp.BatchSn, UserId: cp.Owner,
		Status: cp.Status, StatusDesc: "", ReceiveType: cp.ReceiveType}
	return dto
}

//创建优惠券批次， 将dto-->couponbatch, 涉及字段的逻辑转换
func (cc *CouponChaincode) createCouponBatch(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Invalid number of arguments.")
	}
	cbDto := &CouponBatchDto{}
	err := json.Unmarshal([]byte(args[0]), cbDto)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal args[0] error: %v", err)
	}
	fmt.Printf("CouponBatchDto:%v", cbDto)
	cb := cc.cbDto2Cb(cbDto)
	err = cc.saveCouponBatch(stub, cb)
	if err != nil {
		return nil, fmt.Errorf("save CB error: %v", err)
	}
	return []byte("success"), nil
}

func (cc *CouponChaincode) applyCoupon(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Invalid number of arguments.")
	}
	cpDto := &CouponDto{}
	err := json.Unmarshal([]byte(args[0]), cpDto)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal args[0] error: %v", err)
	}
	cp := cc.cpDto2Cp(cpDto)
	couponBatch, err := cc.getCouponBatch(stub, cp.Sn)
	if err != nil {
		return nil, fmt.Errorf("getCouponBatch error: %v", err)
	}

}

//将cb json后，存入worldstate
func (cc *CouponChaincode) saveCouponBatch(stub *shim.ChaincodeStub, cb *CouponBatch) error {
	jsonBytes, err := json.Marshal(cb)
	if err != nil {
		return fmt.Errorf("json.Marshal cb error, %v", err)
	}
	err = stub.PutState(COUPON_BATCH_KEY+cb.Sn, jsonBytes)
	return err
}

//从world state中查询cb的json bytes, 然后unmarshal返回对象
func (cc *CouponChaincode) getCouponBatch(stub *shim.ChaincodeStub, sn string) (*CouponBatch, error) {
	bytes, err := stub.GetState(COUPON_BATCH_KEY + sn)
	if err != nil {
		return nil, err
	}
	cb := &CouponBatch{}
	err = json.Unmarshal(bytes, cb)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error:%v", err)
	}
	return cb, nil
}

//检验cb的有效性
func (cc *CouponChaincode) checkCouponBatch(cb *CouponBatch) (bool, err) {
	//status 验证
	if cb.status != 1 {
		return false, nil
	}

}

func getCurMilliseconds() uint64 {
	return time.Now().UnixNano() / 1000000
}

func main() {
	err := shim.Start(new(CouponChaincode))
	if err != nil {
		log.Fatalf("bootstrap failed,%v", err)
	}
}
