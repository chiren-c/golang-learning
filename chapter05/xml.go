package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// 抽取单个server对象
type Server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

type Servers struct {
	Name    xml.Name `xml:"servers"`
	Version int      `xml:"version"`
	Servers []Server `xml:"server"`
}

func main() {
	fmt.Println("----------使用 os.ReadFile 读取 XML 文件----------")
	data, err := os.ReadFile("chapter05/my.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	var servers Servers
	err = xml.Unmarshal(data, &servers)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("xml: %#v\n", servers)
	for _, server := range servers.Servers {
		fmt.Printf("server: %#v\n", server)
	}
	fmt.Println("----------使用 os.Open 流式读取 XML 文件----------")
	streamParseXML()
}

func streamParseXML() {
	file, err := os.Open("chapter05/my.xml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := xml.NewDecoder(file)
	for {
		token, err := decoder.Token()
		if err != nil {
			break // 文件结束或出错
		}
		switch se := token.(type) {
		case xml.StartElement:
			if se.Name.Local == "servers" {
				var servers Servers
				decoder.DecodeElement(&servers, &se)
				fmt.Printf("流式解析: %+v\n", servers)
			}
		}
	}
}
