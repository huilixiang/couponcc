package main

import (
	//"bytes"
	"encoding/json"
	//"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"log"
	//"strconv"
	//"strings"
)

//优惠券批次信息
type CouponBatch struct {
	Id              uint64 `json:"id,omitempty"`
	Sn              string `json:"sn,omitempty"`
	Name            string `json:"name,omitempty"`
	Type            uint8  `json:"type,omitempty"`
	Denomination    uint32 `json:"denomination,omitempty"`
	EffectiveDate   uint64 `json:"effectiveDate,omitempty"`
	ExpiringDate    uint64 `json:"expiringDate,omitempty"`
	ApplyStartDate  uint64 `json:"applyStartDate,omitempty"`
	ApplyEndDate    uint64 `json:"applyEndDate,omitempty"`
	PublishDate     uint64 `json:"publishDate,omitempty"`
	BudgetNo        string `json:"budgetNo,omitempty"`
	Channel         string `json:"channel,omitempty"`
	Sku             string `json:"sku,omitempty"`
	Desc            string `json:"desc,omitempty"`
	CouponAmount    int    `json:"couponAmount,omitempty"`
	AppliedAmount   int    `json:"appliedAmount,omitempty"`
	SentAmount      int    `json:"sentAmount,omitempty"`
	OrderPriceLimit int    `json:"orderPriceLimit,omitempty"`
	ApplyLimit      int    `json:"applyLimit,omitempty"`
	Creator         uint64 `json:"creator,omitempty"`
	CreateDate      uint64 `json:"createDate,omitempty"`
	Disabled        bool   `json:"disabled,omitempty"`
	ShopId          uint64 `json:"shopId,omitempty"`
	Picture         string `json:"picture,omitempty"`
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
	Status      uint8  `json:"status,omitempty"`
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
	return nil, nil
}

func (cc *CouponChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	ccStub := stub.(*shim.ChaincodeStub)
	if "createCouponBatch" == function {
		re, err := cc.createCouponBatch(ccStub, args)
		return re, err
	} else if "applyCoupon" == function {

	} else if "sendCoupon" == function {

	} else if "consumeCoupon" == function {

	} else if "disableCouponBatch" == function {

	} else if "publishCouponBatch" == function {

	} else {

	}
	return nil, nil
}

func (cc *CouponChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil

}

func (cc *CouponChaincode) createCouponBatch(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Invalid number of arguments.")
	}
	cbDto := &CouponBatchDto{}
	err := json.Unmarshal([]byte(args[0]), cbDto)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error,%v", err)
	}
	return []byte("success"), nil
}

func main() {
	err := shim.Start(new(CouponChaincode))
	if err != nil {
		log.Fatalf("bootstrap failed,%v", err)
	}
}
