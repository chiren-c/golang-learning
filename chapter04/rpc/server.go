package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/chiren-c/golang-learning/chapter04/rpc/utils"
)

func main() {
	// 注册 RPC 服务
	math := new(MathService)
	rpc.Register(math)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("启动服务监听失败:", err)
	}
	// 使用 HTTP 协议处理 RPC 调用
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("启动 HTTP 服务失败:", err)
	}
}

type MathService struct{}

func (m *MathService) Multiply(args *utils.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (m *MathService) Divide(args *utils.Args, reply *int) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*reply = args.A / args.B
	return nil
}
