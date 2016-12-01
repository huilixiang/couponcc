package main

import (
	//"bytes"
	"encoding/json"
	//"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"log"
	"strconv"
	"strings"
	"time"
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
	// 1->未使用，2->已使用，3-->冻结
	Status int `json:"status,omitempty"`
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
		re, err := cc.sendCoupon(ccStub, args)
		return re, err
	} else if "consumeCoupon" == function {
		re, err := cc.consumeCoupon(ccStub, args)
		return re, err
	} else if "disableCouponBatch" == function {
		re, err := cc.disableCouponBatch(ccStub, args)
		return re, err
	} else if "publishCouponBatch" == function {
		re, err := cc.publishCouponBatch(ccStub, args)
		return re, err
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
	fmt.Printf("in queryCouponBatch\n")
	if len(args) != 1 {
		return nil, fmt.Errorf("Invalid number of arguments")
	}
	sn := args[0]
	cb, err := cc.getCouponBatch(stub, sn)
	if err != nil {
		return nil, fmt.Errorf("Get Coupon Batch error: %v", err)
	}
	fmt.Print("new cc\n")
	if cb == nil {
		return []byte("not found"), nil
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
	dto.UsageRule = make(map[string]interface{})
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
		return nil, fmt.Errorf("Request arguments format error: %v", err)
	}
	log.Printf("CouponBatchDto:%v", cbDto)
	existed_cb, err := cc.getCouponBatch(stub, cbDto.BatchSn)
	if err != nil {
		return nil, fmt.Errorf("Check CouponBatch error: %v", err)
	}
	if existed_cb != nil {
		return nil, fmt.Errorf("CouponBatch existed")
	}
	cb := cc.cbDto2Cb(cbDto)
	err = cc.saveCouponBatch(stub, cb)
	if err != nil {
		return nil, fmt.Errorf("Save CouponBatch error: %v", err)
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
		return nil, fmt.Errorf("Request CouponDto format error: %v", err)
	}
	cp := cc.cpDto2Cp(cpDto)
	couponBatch, err := cc.getCouponBatch(stub, cp.Sn)
	if err != nil {
		return nil, fmt.Errorf("Get CouponBatch error: %v", err)
	}
	if couponBatch == nil {
		return []byte("CouponBatch not found"), nil
	}
	ok, err := cc.checkCouponBatch(stub, cp, couponBatch)
	if err != nil {
		return nil, fmt.Errorf("Error occured while checking CouponBatch: %v", err)
	}
	if !ok {
		return []byte("Invalid CouponBatch"), nil
	}
	ok, err = cc.checkUserLegality(stub, cp, couponBatch)
	if err != nil {
		return nil, fmt.Errorf("Error occured while checking UserLegality: %v", err)
	}
	if !ok {
		return []byte("illegal applyment"), nil
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
	err = cc.saveCouponBatchOfUser(stub, strconv.FormatUint(cp.Owner, 10), cp.BatchSn)
	if err != nil {
		return nil, fmt.Errorf("save user couponbatch relation error: %v", err)
	}
	return []byte("success"), nil
}

/**
*发送优惠券处理,现使用与领取优惠券相同的处理流程，由coupon.ReceiveType做逻辑区分
*
 */
func (cc *CouponChaincode) sendCoupon(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return cc.applyCoupon(stub, args)
}

func (cc *CouponChaincode) consumeCoupon(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	if len(args) != 5 {
		return nil, fmt.Errorf("Invalid number of arguments")
	}
	uid, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("uid format error: %v", err)
	}
	couponSn := args[1]
	//couponBatchSn := args[2]
	orderId := args[2]
	fmt.Printf("orderId:%s", orderId)
	price, err := strconv.Atoi(args[3])
	if err != nil {
		return nil, fmt.Errorf("price format error: %v", err)
	}
	sku := args[4]

	cp, err := cc.getCoupon(stub, couponSn)
	if err != nil {
		return nil, fmt.Errorf("GetCoupon error: %v", err)
	}
	//优惠券状态验证
	if cp.Status != 1 {
		return nil, fmt.Errorf("illegal coupon status")
	}
	cb, err := cc.getCouponBatch(stub, cp.BatchSn)
	if err != nil {
		return nil, fmt.Errorf("GetCouponBatch error: %v", err)
	}
	//优惠券拥有者身份验证
	if uid != cp.Owner {
		return nil, fmt.Errorf("illegal user of consuming coupon")
	}
	//优惠券批次状态验证
	if cb.Status != 1 && cb.Status != 2 {
		return nil, fmt.Errorf("illegal couponbatch status")
	}
	//优惠券有效期验证
	now := getCurMilliSeconds()
	if now > cb.ExpiringDate {
		return nil, fmt.Errorf("expired coupon")
	}
	//满减订单金额验证
	if cb.CouponType == 0 && price < cb.OrderPriceLimit {
		return nil, fmt.Errorf("price limit not satisfied")
	}
	//下单立减sku验证
	if cb.CouponType == 1 {
		if !strings.Contains(cb.Sku, sku) {
			return nil, fmt.Errorf("sku limit not satisfied")
		}
	}
	cp.Status = 2
	err = cc.saveCoupon(stub, cp)
	if err != nil {
		return nil, fmt.Errorf("update coupon status error: %v", err)
	}
	return []byte("success"), nil
}

func (cc *CouponChaincode) disableCouponBatch(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Invalid number of auguments")
	}
	sn := args[0]
	cb, err := cc.getCouponBatch(stub, sn)
	if err != nil {
		return nil, fmt.Errorf("Get CouponBatch error: %v", err)
	}
	if cb == nil {
		return []byte("CouponBatch not found"), nil
	}
	cb.Status = -2
	err = cc.saveCouponBatch(stub, cb)
	if err != nil {
		return nil, fmt.Errorf("Save CouponBatch error: %v", err)
	}
	return []byte("success"), nil
}

func (cc *CouponChaincode) publishCouponBatch(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Invalid number of arguments")
	}
	sn := args[0]
	cb, err := cc.getCouponBatch(stub, sn)
	if err != nil {
		return nil, fmt.Errorf("Check CouponBatch error: %v", err)
	}
	if cb == nil {
		return []byte("CouponBatch not found"), nil
	}
	now := getCurMilliSeconds()
	cb.PublishDate = now
	cb.Status = 1
	err = cc.saveCouponBatch(stub, cb)
	if err != nil {
		return nil, fmt.Errorf("Save CouponBatch error: %v", err)
	}
	return []byte("success"), nil

}

func (cc *CouponChaincode) saveCoupon(stub *shim.ChaincodeStub, cp *Coupon) error {
	jsonBytes, err := json.Marshal(cp)
	if err != nil {
		return fmt.Errorf("json.Marshal cb error, %v", err)
	}
	err = stub.PutState(COUPON_KEY+cp.Sn, jsonBytes)
	return err
}

func (cc *CouponChaincode) getCoupon(stub *shim.ChaincodeStub, sn string) (*Coupon, error) {
	key := COUPON_KEY + sn
	bytes, err := stub.GetState(key)
	if err != nil {
		return nil, err
	}
	cp := &Coupon{}
	err = json.Unmarshal(bytes, cp)
	if err != nil {
		return nil, err
	}
	return cp, nil
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
		fmt.Printf("get cb from state error, %v", err)
		return nil, err
	}
	if bytes == nil || len(bytes) == 0 {
		return nil, nil
	}
	cb := &CouponBatch{}
	fmt.Printf("couponbatch:%s", string(bytes))
	err = json.Unmarshal(bytes, cb)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error:%v", err)
	}
	return cb, nil
}

//检验cb的有效性
func (cc *CouponChaincode) checkCouponBatch(stub *shim.ChaincodeStub, cp *Coupon, cb *CouponBatch) (bool, error) {
	//status 验证
	if cb.Status != 1 && cb.Status != 2 {
		return false, nil
	}
	now := getCurMilliSeconds()
	if now < cb.PublishDate || now < cb.ApplyStartDate || now > cb.ExpiringDate {
		return false, nil
	}
	cps, err := cc.getCouponOfBatch(stub, cb.Sn)
	if err != nil {
		return false, nil
	}
	ac, _ := cc.applyAndSendAmount(cps)
	if cp.ReceiveType == APPLY {
		//当前批次可领取总量限制
		if ac >= cb.CouponAmount {
			return false, nil
		}
	} else {
		//发放现在没有限制

	}
	return true, nil
}

func (cc *CouponChaincode) checkUserLegality(stub *shim.ChaincodeStub, coupon *Coupon, cb *CouponBatch) (bool, error) {
	//获取用户在当前批次下已经领取的优惠券
	coupons, err := cc.getCouponsOfUserBatch(stub, strconv.FormatUint(coupon.Owner, 10), cb.Sn)
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

func (cc *CouponChaincode) getCouponsOfUserBatch(stub *shim.ChaincodeStub, uid string, sn string) ([]string, error) {
	key := COUPONS_OF_USER_BATCH_KEY + uid + sn
	return cc.getSplitableValues(stub, key)
}

func (cc *CouponChaincode) saveCouponsOfUserBatch(stub *shim.ChaincodeStub, cp *Coupon) error {
	newSn := ""
	if cp.ReceiveType == APPLY {
		newSn = cp.Sn + "a"
	} else {
		newSn = cp.Sn + "s"
	}
	key := COUPONS_OF_USER_BATCH_KEY + strconv.FormatUint(cp.Owner, 10) + cp.BatchSn
	return cc.appendSplitableValue(stub, key, newSn)
}

func (cc *CouponChaincode) getCouponBatchsOfUser(stub *shim.ChaincodeStub, uid string) ([]string, error) {
	key := COUPONBATCH_OF_USER_KEY + uid
	return cc.getSplitableValues(stub, key)
}

func (cc *CouponChaincode) saveCouponBatchOfUser(stub *shim.ChaincodeStub, uid string, sn string) error {
	key := COUPONBATCH_OF_USER_KEY + uid
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
