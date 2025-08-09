package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 服务端的rpc处理器
type ServiceHandler struct{}

func (s *ServiceHandler) GetName(id int, i *Item) error {
	// log.Printf("receive GetName call, id: %d", id)
	i.Id = id
	i.Name = "name-1"
	return nil
}

func (s *ServiceHandler) SaveName(i *Item, reply *Response) error {
	// log.Printf("receive SaveName call, item: %+v", i)
	reply.Ok = true
	reply.Id = i.Id
	reply.Message = "success"
	return nil
}

func main() {
	// 创建一个新的 RPC 服务实例
	service := rpc.NewServer()
	// 监听端口 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("监听端口失败：%v", err)
	}
	defer listener.Close()

	// log.Println("Start listen on port localhost:8080")

	serviceHandler := &ServiceHandler{}
	// 注册服务处理器
	err = service.Register(serviceHandler)
	if err != nil {
		log.Fatalf("注册服务失败：%v", err)
	}
	// 等待并处理客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("接收客户端连接请求失败: %v", err)
		}

		// 自定义 RPC 编码器：新建一个 jsonrpc 编码器，并将该编码器绑定给 RPC 服务端处理器
		go service.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
