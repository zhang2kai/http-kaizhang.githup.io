package main

import (

	"net/rpc"
	"fmt"
)
func main(){
	client, err := rpc.DialHTTP("tcp", "192.168.1.112:1234")
	if err != nil {
		fmt.Println("链接rpc服务器失败: ",err)
	}

	var reply int
	err = client.Call("Watcher.Getinfo",1,&reply)
	if err != nil{
		fmt.Println("调用远程服务器失败")
	}
	fmt.Println("远程服务返回结果: ",reply)

}

