package wserver

import (
	"fmt"
	"net"
	"bufio"
	"net/http"
	"github.houston.softwaregrp.net/CSB/webserver/pkg/methodID"
)

type Wserver struct {
	message string
	rw bufio.ReadWriter
}


func (w *Wserver)Listener(port *string) {
	prt := *port
	fmt.Printf("prt: %s \r\n", prt)
	ln, err := net.Listen("tcp", ":"+prt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("server is listening on %s \r\n", prt)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("yo failed to accept connection. skipping")
		}

		go w.Handle(conn)
	}
}
func (w *Wserver)Handle(conn net.Conn) {
	fmt.Println("handling the connection")
	//scanner := bufio.NewScanner(conn)
	reader:=bufio.NewReader(conn)
	//scanner.Scan()
		//line := scanner.Text()
		//fmt.Println(line)
		req	,err:=http.ReadRequest(reader)
		if err !=nil{
			fmt.Println("somthing went wrong",err)
		}
		fmt.Println("Request:",req)
		if err !=nil {
			fmt.Println(err)
			return
		}
		methodID.MethodID(conn,req)
		//reader := bufio.NewReader(conn)
		//req, err := http.ReadRequest(reader)

	}