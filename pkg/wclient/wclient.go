package wclient

import (
	"fmt"
	"net"
	"bufio"
	"os"
	)

type Client struct {

}

func (c *Client) Call(address *string,port *string){
	addr := fmt.Sprintf("%s:%s", *address, *port)
	fmt.Printf("dialing %s\n", addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err !=nil{
			return
		}

		fmt.Fprintf(conn, text+"\n")
		go c.handleServerMessages(conn)
	}
}
func (c *Client) handleServerMessages(conn net.Conn) {
	for {
		scanner,err:=bufio.NewReader(conn).ReadString('!')
		fmt.Println(err)
		fmt.Println("scanner:",scanner)

	}

}