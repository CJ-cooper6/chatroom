package main

import (
	"bufio"
	"net"
	"os"
)

func main(){
	tcpAddrs,_:=net.ResolveTCPAddr("tcp",":8888")


		coon,_:=net.DialTCP("tcp",nil,tcpAddrs)

		for{
			reader:=bufio.NewReader(os.Stdin)
			bytes,_,_:=reader.ReadLine()
			coon.Write(bytes)

		}







}

func serverConnection(coon *net.TCPConn){

}