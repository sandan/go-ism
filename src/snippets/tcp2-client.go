package main

import "net"
import "fmt"
import "bufio"

func main() {

  // connect to this socket
  conn, _ := net.Dial("tcp", "127.0.0.1:31826")
  text := "hi"
  for {
    // send to socket
    fmt.Fprintf(conn, text + "\n")
    // listen for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message from server: "+message)
  }
}
