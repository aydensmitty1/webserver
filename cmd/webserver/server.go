package main

import (

	"flag"
	"github.houston.softwaregrp.net/CSB/webserver/pkg/wserver"
)
func main() {
	port:=flag.String("port", "9992", "A port for hosting server")
	flag.Parse()
	srv :=wserver.Wserver{}
	srv.Listener(port)

}
