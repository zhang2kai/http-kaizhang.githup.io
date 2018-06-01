package main

import (
	"net/http"
	"net/rpc"
	"net"
	"fmt"
	"io"
)

type Watcher int

func (w *Watcher)Getinfo(arg int, result *int)error{
	*result = 1
	return nil
}

func main(){
	http.HandleFunc("/kaizhang", Kz2018Test)

	watcher := new(Watcher)
	rpc.Register(watcher)
	rpc.HandleHTTP()

	l,err := net.Listen("tcp", ":1234")
	if err != nil{
		fmt.Println("监听失败")
	}
	fmt.Println("正在监听1234端口")
	http.Serve(l,nil)
}

func Kz2018Test(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"<html><body>贺淑萍 我爱你 一辈子 </body></html>")
}


