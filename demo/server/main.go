package main

import (
	"fmt"
	"net"
)

func main(){

	tcpAddr,_:=net.ResolveTCPAddr("tcp",":8888")
	listener,_:=net.ListenTCP("tcp",tcpAddr)


	for{
		coon,err:=listener.AcceptTCP()
		if err!=nil{
			fmt.Println(err)

		}

	go handleConnection(coon)

	}


}

func handleConnection(coon *net.TCPConn){

	for{
		buf:=make([]byte,1024)
		n,err:=coon.Read(buf)
		if err!=nil{
			fmt.Println(err)
			break

		}
		fmt.Println(string(buf[0:n]))
	}

}