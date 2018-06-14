// Log all received data

package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	addr := flag.String("addr", "", "The address to listen to; default is \"\" (all interfaces).")
	port := flag.Int("port", 8000, "The port to listen on.")

	flag.Parse()

	src := *addr + ":" + strconv.Itoa(*port)
	listener, _ := net.Listen("tcp", src)
	fmt.Printf("Listening on %s.\n", src)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		remoteAddr := conn.RemoteAddr().String()
		log.Println("Accepted connection from " + remoteAddr)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		size, err := conn.Read(buf)
		if err != nil {
			if io.EOF == err {
				break
			}
			log.Panicln(err)
		}
		bytes := buf[:size]
		fmt.Printf("%s", hex.Dump(bytes))
	}

}
