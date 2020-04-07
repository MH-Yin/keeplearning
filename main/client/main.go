package main

import "net"

func main() {
	c ,  _ := net.Dial("tcp", "0.0.0.0:8888")
	c.Write([]byte("hello world"))
	defer c.Close()
}
