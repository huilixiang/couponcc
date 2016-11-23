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
	"strings"
)

const (
	SENT = iota
	APPLY
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
	EffectiveDate int64 `json:"effectiveDate,omitempty"`
	//失效时间
	ExpiringDate int64 `json:"expiringDate,omitempty"`
	//可申请开始时间
	ApplyStartDate int64 `json:"applyStartDate,omitempty"`
	//可申请结束时间
	ApplyEndDate int64 `json:"applyEndDate,omitempty"`
	//发布时间
	PublishDate int64  `json:"publishDate,omitempty"`
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
	CreateDate   int64  `json:"createDate,omitempty"`
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
	Id      uint64 `json:"id,omitempty"`
	BatchSn string `json:"batchSn,omitempty"`
	Sn      string `json:"sn,omitempty"`
	//1: 发放 2：领取
	ReceiveType int    `json:"receiveType,omitempty"`
	CreateDate  int64  `json:"createDate,omitempty"`
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
	EffectiveStartTime int64                  `json:"effectiveStartTime,omitempty"`
	EffectiveEndTime   int64                  `json:"effectiveEndTime,omitempty"`
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
	PublishTime        int64                  `json:"publishTime,omitempty"`
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
		return re, err
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
		BatchType: cbDto.BatchType, Denomination: cbDto.Money,
		EffectiveDate: cbDto.EffectiveStartTime, ExpiringDate: cbDto.EffectiveEndTime,
		PublishDate: cbDto.PublishTime, Desc: cbDto.Description, ShopId: cbDto.ShopId,
		Picture: cbDto.Picture, Status: cbDto.Status, CouponType: cbDto.UsageRuleType}
	if cbDto.UsageRuleType == 1 {
		cb.OrderPriceLimit = cbDto.UsageRule["minAmount"].(int)
	}
	return cb

}

//couponbatch cb --> dto
func (cc *CouponChaincode) cb2Dto(cb *CouponBatch) *CouponBatchDto {
	dto := &CouponBatchDto{
		Id: cb.Id, BatchSn: cb.Sn, BatchType: cb.BatchType, UsageRuleType: cb.CouponType,
		EffectiveStartTime: cb.EffectiveDate, EffectiveEndTime: cb.ExpiringDate,
		Money: cb.Denomination, ShopId: cb.ShopId, BatchName: cb.Name, Status: cb.Status,
		Description: cb.Desc, Picture: cb.Picture, PublishTime: cb.PublishDate}
	dto.UsageRule["minAmount"] = cb.OrderPriceLimit
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

/**
* 申请优惠券，优惠券合法性验证&用户合法性验证&生成优惠券
 */
func (cc *CouponChaincode) applyCoupon(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Invalid number of arguments.")
	}
	cpDto := &CouponInfoDto{}
	err := json.Unmarshal([]byte(args[0]), cpDto)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal args[0] error: %v", err)
	}
	cp := cc.cpDto2Cp(cpDto)
	couponBatch, err := cc.getCouponBatch(stub, cp.Sn)
	if err != nil {
		return nil, fmt.Errorf("getCouponBatch error: %v", err)
	}
	ok, err := cc.checkCouponBatch(couponBatch)
	if err != nil {
		return nil, fmt.Errorf("check CB error: %v", err)
	}
	if !ok {
		return nil, fmt.Errorf("invalid CB")
	}
	ok, err = cc.checkUserLegality(stub, cp, couponBatch)
	if err != nil {
		return nil, fmt.Errorf("check user error:%v", err)
	}
	if !ok {
		return nil, fmt.Errorf("illegal user")
	}
	err = cc.saveCoupon(stub, cp)
	if err != nil {
		return nil, fmt.Errorf("save coupon error: %v", err)
	}
	err = cc.saveCouponOfBatch(stub, cp)
	if err != nil {
		return nil, fmt.Errorf("save coupon batch relation error: %v", err)
	}
	err = cc.saveCouponsOfUserBatch(stub, cp)
	if err != nil {
		return nil, fmt.Errorf("save user coupon relation error: %v", err)
	}
	err = cc.saveCouponBatchOfUser(stub, cp.Owner, cp.BatchSn)
	if err != nil {
		return nil, fmt.Errorf("save user couponbatch relation error: %v", err)
	}
	return []byte("success"), nil
}

func (cc *CouponChaincode) saveCoupon(stub *shim.ChaincodeStub, cp *Coupon) error {
	jsonBytes, err := json.Marshal(cp)
	if err != nil {
		return fmt.Errorf("json.Marshal cb error, %v", err)
	}
	err = stub.PutState(COUPON_KEY+cp.BatchSn+cp.Sn, jsonBytes)
	return err
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
func (cc *CouponChaincode) checkCouponBatch(cb *CouponBatch) (bool, error) {
	//status 验证
	if cb.Status != 1 && cb.Status != 2 {
		return false, nil
	}
	now := getCurMilliSeconds()
	if now < cb.PublishDate || now < cb.ApplyStartDate || now > cb.ExpiringDate {
		return false, nil
	}
	return true, nil
}

func (cc *CouponChaincode) checkUserLegality(stub *shim.ChaincodeStub, coupon *Coupon, cb *CouponBatch) (bool, error) {
	//获取用户在当前批次下已经领取的优惠券
	coupons, err := cc.getCouponsOfUserBatch(stub, coupon.Owner, cb.Sn)
	if err != nil {
		return false, err
	}
	//如果没有领取过，则可以合法领取
	if coupons == nil || len(coupons) == 0 {
		return true, nil
	}
	ac, _ := cc.applyAndSendAmount(coupons)
	//已领取数量不能大于批次限制
	if ac >= cb.ApplyLimit {
		return false, nil
	}
	return true, nil
}

func (cc *CouponChaincode) getCouponsOfUserBatch(stub *shim.ChaincodeStub, uid uint64, sn string) ([]string, error) {
	key := COUPONS_OF_USER_BATCH_KEY + string(uid) + sn
	return cc.getSplitableValues(stub, key)
}

func (cc *CouponChaincode) saveCouponsOfUserBatch(stub *shim.ChaincodeStub, cp *Coupon) error {
	newSn := ""
	if cp.ReceiveType == APPLY {
		newSn = cp.Sn + "a"
	} else {
		newSn = cp.Sn + "s"
	}
	key := COUPONS_OF_USER_BATCH_KEY + string(cp.Owner) + cp.BatchSn
	return cc.appendSplitableValue(stub, key, newSn)
}

func (cc *CouponChaincode) getCouponBatchsOfUser(stub *shim.ChaincodeStub, uid uint64) ([]string, error) {
	key := COUPONBATCH_OF_USER_KEY + string(uid)
	return cc.getSplitableValues(stub, key)
}

func (cc *CouponChaincode) saveCouponBatchOfUser(stub *shim.ChaincodeStub, uid uint64, sn string) error {
	key := COUPONBATCH_OF_USER_KEY + string(uid)
	return cc.appendSplitableValue(stub, key, sn)
}

func (cc *CouponChaincode) applyAndSendAmount(coupons []string) (int, int) {
	send := 0
	apply := 0
	for _, cp := range coupons {
		if strings.HasSuffix(cp, "a") {
			apply++
		} else {
			send++
		}
	}
	return apply, send
}

func (cc *CouponChaincode) getCouponOfBatch(stub *shim.ChaincodeStub, sn string) ([]string, error) {
	key := COUPONS_OF_BATCH_KEY + sn
	return cc.getSplitableValues(stub, key)
}

func (cc *CouponChaincode) saveCouponOfBatch(stub *shim.ChaincodeStub, cp *Coupon) error {
	newSn := ""
	if cp.ReceiveType == APPLY {
		newSn = cp.Sn + "a"
	} else {
		newSn = cp.Sn + "s"
	}
	key := COUPONS_OF_BATCH_KEY + cp.BatchSn
	return cc.appendSplitableValue(stub, key, newSn)

}

func (cc *CouponChaincode) getSplitableValues(stub *shim.ChaincodeStub, key string) ([]string, error) {
	bytes, err := stub.GetState(key)
	if err != nil {
		return nil, err
	}

	if bytes == nil || len(bytes) < 1 {
		return []string{}, nil
	}
	return strings.Split(string(bytes), ","), nil
}

func (cc *CouponChaincode) appendSplitableValue(stub *shim.ChaincodeStub, key string, value string) error {
	values, err := cc.getSplitableValues(stub, key)
	if err != nil {
		return err
	}
	for _, v := range values {
		//重复的数据不保存
		if v == value {
			return nil
		}

	}
	values = append(values, value)
	s := strings.Join(values, ",")
	return stub.PutState(key, []byte(s))
}

func getCurMilliSeconds() int64 {
	return time.Now().UnixNano() / 1000000
}

func main() {
	err := shim.Start(new(CouponChaincode))
	if err != nil {
		log.Fatalf("bootstrap failed,%v", err)
	}
}
