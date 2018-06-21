package main

import (
	"os"
	"fmt"
	"net"
)

func DealError(err error,tips string) bool{
	if err !=nil {
		fmt.Println(tips)
		return false
	}
	return true
}

func StartServer(port string){
	tcpAddr,err:=net.ResolveTCPAddr("tcp4",":"+port)
	DealError(err,"ResolveTCPAddr")

	tcpListener,err:=net.ListenTCP("tcp",tcpAddr)

	conns:=make(map[string]net.Conn)
	messages:=make(chan string,10)
	go BroadCast(&conns,messages) //广播

	for{
		con,err:=tcpListener.Accept()
		DealError(err,"Accept")
		conns[con.RemoteAddr().String()]=con
		go Handler(con,messages)
	}
}

func Handler(conn net.Conn,message chan string){
	buf:=make([]byte,1024)
	for {
		n,err:=conn.Read(buf)
		if err!=nil {
			conn.Close()
			break
		}
		msg:=buf[0:n]
		message<- string(msg)
	}
}

func  BroadCast(conns *map[string]net.Conn,messages chan string){
	for {
		msg:= <- messages
		fmt.Println(msg)

		for name,con:=range  *conns{
			fmt.Println("connection is from ",name)
			_,err:=con.Write([]byte(msg))
			if err!=nil {
				delete(*conns,name)
			}
		}
	}
}

func StartClient(port string){
	dialAddr,_:=net.ResolveTCPAddr("tcp4",":"+port)
	conn,_:=net.DialTCP("tcp",nil,dialAddr)
	go sendChat(conn)

	buf:=make([]byte,1024)
	for {
		n,err:=conn.Read(buf)
		if DealError(err,"conn fail")==false{
			fmt.Println("the world stopped")
			os.Exit(1)
		}

		str:=string(buf[0:n])
		fmt.Println(str)
	}
}
func sendChat(con *net.TCPConn){
	var input string

	for  {
		fmt.Scanln(&input)
		if input=="exit" {
			con.Close()
			os.Exit(1)
		}
		con.Write([]byte("from"+con.RemoteAddr().String()+":::"+input))
	}
}
func main(){
	//使用：
	//客户端 chat server 9999
	//服务端 chat client 9999


	if len(os.Args)!=3{
		fmt.Println("wrong param")
		os.Exit(1)
	}

	if len(os.Args)==3 && os.Args[1]=="server" {
		//启动server
		StartServer(os.Args[2])
	}

	if len(os.Args)==3 && os.Args[1]=="client"{
		//处理client
		StartClient(os.Args[2])
	}
}

