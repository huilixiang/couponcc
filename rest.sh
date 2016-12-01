#curl -X POST -d '{"jsonrpc": "2.0","method":"deploy", "params": {"type": 1,"chaincodeID":{"name":"SimpleSample"},"ctorMsg": {"args":["init", "a", "10", "b", "20"]}},"id": 3}' http://127.0.0.1:7051 
#curl -X POST http://0.0.0.0:7050/registrar -d '{"enrollId":"jim","enrollSecret":"6avZQLwcUe9b"}'

#curl -X POST -d '{"jsonrpc": "2.0","method":"deploy", "params": {"type": 1,"chaincodeID":{"name":"couponcc"},"ctorMsg": {"args":["init"]}, "secureContext":"jim"},"id": 3}' http://127.0.0.1:7050/chaincode 
curl -X POST -d '{"jsonrpc": "2.0","method":"deploy", "params": {"type": 1,"chaincodeID":{"name":"couponcc"},	"ctorMsg": {"args":["init","a","10"]}},"id": 3}' http://127.0.0.1:7050/chaincode 
sleep 1
echo -e "\nafter deploy......\n"
curl -X POST -d '{"jsonrpc": "2.0","method":"invoke", "params": {"type": 1,"chaincodeID":{"name":"couponcc"},"ctorMsg": {"args":["createCouponBatch", "{\"id\":1938,\"usageRuleType\":2,\"batchTypeDesc\":\"商家劵\",\"batchSn\":\"2016111035306\",\"status\":0,\"publishedTime\":0,\"shopId\":1045,\"batchType\":1,\"money\":200,\"effectiveEndTime\":1480521599000,\"usageRule\":{\"minAmount\":250},\"effectiveStartTime\":1478707200000}"]}},"id": 5}' http://127.0.0.1:7050/chaincode 
echo -e "\nafter invoke create CB\n"
sleep 1
curl -X POST -d '{"jsonrpc": "2.0","method":"query", "params": {"type": 1,"chaincodeID":{"name":"couponcc"},"ctorMsg": {"args":["queryCouponBatch", "2016111035306"]}},"id": 7}' http://127.0.0.1:7050/chaincode 
echo -e "\nafter queryCB\n"
curl -X POST -d '{"jsonrpc": "2.0","method":"query", "params": {"type": 1,"chaincodeID":{"name":"couponcc"},"ctorMsg": {"args":["queryCouponBatch", "basdf"]}, "secureContext":"jim"},"id": 8}' http://127.0.0.1:7050/chaincode 
