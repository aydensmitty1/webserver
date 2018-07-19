package methodID

import (
		"net/http"
	"path/filepath"
	"fmt"
	"io/ioutil"
	"bytes"
	"net"
		)

//func MethodID(w http.ResponseWriter, r *http.Request) {
func MethodID(conn net.Conn, r *http.Request) {
	recker:=r.URL.Path
	fmt.Println("recker:",recker)
		switch r.Method {
		case http.MethodGet:
			switch recker {
			case "/index.html":
				dir, err := filepath.Abs(filepath.Dir("index.html"))
				direct := fmt.Sprintf("%s"+"/configs/index.html", dir)
				dat, err := ioutil.ReadFile(direct)
				if err != nil {
					fmt.Println(err)
				}
				data := fmt.Sprintf("data: %s \r\n", dat)
				res := http.Response{
					Status:     "OK",
					StatusCode: 200,
					Proto:      "HTTP/1.1",
					ProtoMajor: 1,
					ProtoMinor: 1,
					Body:       ioutil.NopCloser(bytes.NewBufferString(data)),
				}
				resBuff := bytes.NewBuffer(nil)
				res.Write(resBuff)
				resbuffer := *resBuff
				fmt.Print("\r\n resbuf: \r\n", resBuff)
				conn.Write(resbuffer.Bytes())
				defer conn.Close()
			default:
				failure := http.Response{
					Status:"Not Found",
					StatusCode:404,
					Proto:"HTTP/1.1",
					ProtoMajor:1,
					ProtoMinor:1,
				}
				fmt.Println("\r\n fail: \r\n", failure)
				fbuffs := bytes.NewBuffer(nil)
				failure.Write(fbuffs)
				failbuffs:=*fbuffs
				conn.Write(failbuffs.Bytes())
				defer conn.Close()

			}
		default:
			fail := http.Response{
				Status:"Method Not Allowed",
				StatusCode:405,
				Proto:"HTTP/1.1",
				ProtoMajor:1,
				ProtoMinor:1,
			}
			fmt.Println("\r\n fail: \r\n", fail)
			fbuff := bytes.NewBuffer(nil)
			fail.Write(fbuff)
			failbuff:=*fbuff
			conn.Write(failbuff.Bytes())
			defer conn.Close()


		}
	}
