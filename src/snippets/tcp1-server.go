package main

import (
	"net"
	"os"
    "fmt"
)

/*
A server registers itself on a port and listens on that port
It blocks on an accept op waiting for clients to connect
When a client connects, the accept() returns a connection object
*/

const (
  SERVICE_INTERFACE = "" // all interfaces
  SERVICE_PORT = "43444" // default service port
  MAX_BYTES = 1 << 10 /* max bytes a server can read from amp */
)

func main() {

	service := SERVICE_INTERFACE + ":" + SERVICE_PORT //default
    if len(os.Args) == 2 {
      service = os.Args[1]
    }

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	CheckAndShowErrorExit(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
    fmt.Printf("Got listener of type %T\n", listener)
	CheckAndShowErrorExit(err)

	for {
		conn, err := listener.Accept()

        fmt.Println("Accepted!")
		if err != nil {
          continue //logErr
        }
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {

    // cleanup when process() function ends
	defer conn.Close()

    // allocate a zeroed buffer of bytes
    fmt.Println("allocating", MAX_BYTES, "bytes")
    var buf [MAX_BYTES]byte

    // read from client
    for {

      // read bytes from client into a buffer
      nbytes, err := conn.Read(buf[:])

      if err != nil {
          fmt.Println(err.Error())
        return
      }

      // log input to stdout
      fmt.Println(string(buf[:]))

      // echo n bytes read back
      _, err2 := conn.Write(buf[:nbytes])
      if err2 != nil {
          fmt.Println(err.Error())
        return
      }
    }
}

func CheckAndShowErrorExit(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Got an error: %v\n", err.Error())
		os.Exit(1)
	}
}

func CheckAndShowError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Got an error: %v\n", err.Error())
	}
}
