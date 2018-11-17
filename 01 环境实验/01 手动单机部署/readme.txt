
--实例化Chaincode
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export FABRIC_START_TIMEOUT=300
peer chaincode instantiate -o orderer.example.com:7050 --tls  --cafile $ORDERER_CA -C mychannel -n hello -v 1.0 -c '{"Args":["a","helloworld!!!"]}' -P "OR ('Org1MSP.peer')"


&&&& 遗留问题 合约vm
2018-10-14 10:59:32.298 UTC [shim] userChaincodeStreamGetter -> ERRO 001 context deadline exceeded
error trying to connect to local peer
github.com/hyperledger/fabric/core/chaincode/shim.userChaincodeStreamGetter
	/opt/gopath/src/github.com/hyperledger/fabric/core/chaincode/shim/chaincode.go:111
github.com/hyperledger/fabric/core/chaincode/shim.Start
	/opt/gopath/src/github.com/hyperledger/fabric/core/chaincode/shim/chaincode.go:150
main.main
	/chaincode/input/src/github.com/hyperledger/fabric/examples/chaincode/go/helloworld/cmd/main.go:19
runtime.main
	/opt/go/src/runtime/proc.go:198
runtime.goexit
	/opt/go/src/runtime/asm_amd64.s:2361

