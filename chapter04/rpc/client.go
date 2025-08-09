package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/chiren-c/golang-learning/chapter04/rpc/utils"
)

func main() {
	var serverAddress = "localhost:8080"
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		// 处理错误 ...
		return
	}
	defer client.Close()
	args := &utils.Args{7, 8}
	var reply int
	err = client.Call("MathService.Multiply", args, &reply)
	if err != nil {
		log.Fatal("调用远程方法 MathService.Multiply 失败:", err)
		return
	}
	fmt.Printf("%d*%d=%d\n", args.A, args.B, reply)
	// 异步方式进行 RPC 调用
	divideCall := client.Go("MathService.Divide", args, &reply, nil)
	for {
		select {
		case <-divideCall.Done:
			fmt.Printf("%d/%d=%d\n", args.A, args.B, reply)
			return
		}
	}
}
