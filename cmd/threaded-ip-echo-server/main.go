// /* 
// ThreadedIPEchoServer
// */

// package main

// import (

// )

// func main() {

// 	service := ":1200"
// 	listener, err := net.Listen("tcp", service)
// 	checkError(err)

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			continue
// 		}

// 		go handleClient(conn)
// 	}

// }

// func handleClient(conn net.Conn) {
// 	defer conn.Close

// 	var buf [512]byte

// 	for {
// 		n, err := conn.Read(buf[0:])
// 		if err != nil {
// 			return
// 		}

// 		_, err2 := conn.Write(buf[0:n])
// 		if err2 != nil {
// 			return
// 		}
// 	}
// }

// func checkError(err error) {
	
// 	Log.fatal(err)
// 	os.Exit(1)
// }