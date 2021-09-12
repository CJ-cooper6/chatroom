package main

import (
	"fmt"
	"net"
	"time"
)

//全局变量

var Message=make(chan string,10)		//广播信息
var Usermap=make(map[string]User)	//用户列表

func main(){

	listener,err:=net.Listen("tcp","127.0.0.1:8000")
	if err!=nil{
		fmt.Println("Listen err",err)
		return
	}
	defer listener.Close()


	go Manager(Message)//用于监听message通道，并向每个用户发送消息

	for {
		coon,err:=listener.Accept()
		if err!= nil{
			fmt.Println("Accept err",err)
			return
		}
		//启动go程处理数据请求,这相当于新来了一个用户
		go handleConnect(coon)

	}


}

func handleConnect(conn net.Conn){
	defer conn.Close()
	Addr:=conn.RemoteAddr().String()	//获取客户端端口信息

	go WriteMsgToClient(Addr,conn)		//创建一个go程监听用户自己的通道

	//创建一个User结构
	newUser:=User{
		Addr,		//name可以修改
		Addr,
		make(chan string),
	}
	Usermap[Addr]=newUser		//保存user信息在map中
	Message<-fmt.Sprintf("%s上线了",newUser.Name)		//向其他人说自己上线了

	Activequit:=make(chan string)	//判断用户主动退出状态
	timequit:=make(chan string)		//判断用户是否超时

	go Watch(Activequit,timequit,&newUser,conn)


	res:=make([]byte,1024)
	for  {

		n,err:=conn.Read(res)
		if n==0{	//代表用户ctrl+c推出了，向Activequit通道发送消息
		Activequit<-"quit"
			return
		}
		if err!=nil{
			fmt.Println("read fail,err:",err)
			return
		}


		if n==4&&string(res[0:3])=="who"{		//查看当前在线用户
			for _,v:=range Usermap{
				newUser.Msg<-v.Name
			}	//遍历Usermap，将用户名加到用户的通道中



		}else if n==7&&string(res[0:6])=="rename"{		//修改用户名称
			newUser.Msg<-fmt.Sprintf("请输入要修改的名称")
			rename:=make([]byte,1024)
			re,_:=conn.Read(rename)
			newUser.Name=string(rename[:re-1])
			Usermap[Addr]=newUser
			newUser.Msg<-fmt.Sprintf("修改完毕")
			//TODO	私聊功能
		}else{
			b:=fmt.Sprintf(newUser.Name+" :"+string(res[:n-1]))
			Message<-b
		}

		timequit<-"nouit"

	}
}


func Manager(message chan string){		//监听全局的message通道
	for{
		msg:=<-message
		for _,v:=range Usermap{
			v.Msg <- msg
		}
	}
}


func WriteMsgToClient(Addr string,conn net.Conn){	//将用户通道里的信息写回给客户端，每个用户一个
	for{
		res:=<-Usermap[Addr].Msg
		_,err:=conn.Write([]byte(res+"\n"))
		if err!=nil{
			fmt.Println("%s,qqqqq",err)
		}
	}

}

func Watch(Activequit,timequir <-chan string,user *User,coon net.Conn){		//监控退出信号

for{		//这里一定要有这个for循环，不然就只能执行一次

	select {
	case <-Activequit:
		delete(Usermap,user.Id)

		Message<-fmt.Sprintf("%s已退出",user.Name)
		coon.Close()
		return

	case <-timequir:		//重置时间
		//fmt.Println("重置了")

	case <-time.After(60 * time.Second):	//设置超时时间
		delete(Usermap,user.Id)

		Message<-fmt.Sprintf("%s已超时退出",user.Name)
		coon.Close()
		return

	}
}
}


