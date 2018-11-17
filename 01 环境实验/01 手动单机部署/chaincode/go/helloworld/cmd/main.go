/*
 * 
 * 文件名称 : chaincode.go 
 * 创建者 : linwf
 * 创建日期: 2018/08/23 
 * 文件描述: 实现存储Helloworld字符串的智能合约
 * 历史记录: 无 
 */
package main

import (
    "fmt"

    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/hyperledger/fabric/examples/chaincode/go/helloworld"
)

func main() {
    err := shim.Start(new(helloworld.SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}