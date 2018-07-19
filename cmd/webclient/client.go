package main

import (
	"flag"
	"github.houston.softwaregrp.net/CSB/webserver/pkg/wclient"
)

func main(){
	port:=flag.String("port", "9992", "A port")
	address:=flag.String("ipad","10.12.182.133", "an Ip address")
	flag.Parse()
	cli:= wclient.Client{}
	cli.Call(address,port)
}
