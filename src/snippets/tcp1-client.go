package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

/*
  tcp client establishes connection to a service
  prints a response from the service
*/
func main() {

	if len(os.Args) != 2 {
      fmt.Fprintln(os.Stderr, "Provide an <interface>:<port>\n")
      os.Exit(1)
	}

	service := os.Args[1]
    fmt.Println("Client given addr:", service)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	CheckAndShowError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	CheckAndShowError(err)

	res1, err := ioutil.ReadAll(conn)
	CheckAndShowError(err)
	fmt.Println(string(res1))

	msg := "hello!\n"
	_, err2 := conn.Write([]byte(msg))
	CheckAndShowError(err2)

	res, err := ioutil.ReadAll(conn)
	CheckAndShowError(err)

	fmt.Println(string(res))
    defer conn.Close()
}

func CheckAndShowError(err error) {
	if err != nil {
		fmt.Printf("Got an error: %v\n", err)
		os.Exit(1)
	}
}
