package telnetUtil

import (
	"fmt"
	"net"
)

func StartTCPServer(addres string) {
	listener, err := net.Listen("tcp", addres)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server is listening...")
	go func() {
		defer listener.Close()
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		for {
			buf := make([]byte,1024)
			n,err := conn.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			conn.Write([]byte(string(buf[0:n])))
		}
	}()
}