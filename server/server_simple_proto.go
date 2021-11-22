package main

import (
	"fmt"
	"github/leiwb/simple_proto/utils"
	"log"
	"net"
	"os"
	"time"
)

func handler(conn net.Conn) {
	defer conn.Close()
	handler_byte := make([]byte, 0)
	for range time.NewTicker(time.Second).C {
		read_byte := make([]byte, 4096)
		n, err := conn.Read(read_byte)
		check(err)
		if n > 0 {
			handler_byte = append(handler_byte, read_byte[0:n]...)
			handler_byte = handler_redundancy(handler_byte)
		}
	}
}

func handler_redundancy(in_byte []byte) []byte {
	if len(in_byte) > 5 {
		data_length_slice := in_byte[1:5]
		data_length := utils.Byte_2_int32(data_length_slice)
		var complete_proto_size int32 = data_length + 5
		if len(in_byte) > int(complete_proto_size) {
			handler_proto(in_byte[0:complete_proto_size])
			rest_part_byte := make([]byte, 0)
			return append(rest_part_byte, in_byte[complete_proto_size:]...)
		}
	}
	return in_byte
}
func handler_proto(in_byte []byte) {
	log.Println("-------------------")
	data_length_slice := in_byte[1:5]
	data_length := utils.Byte_2_int32(data_length_slice)
	data_conten := in_byte[5:]
	f, err := os.OpenFile("./out.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModeAppend|os.ModePerm)
	check(err)
	fmt.Fprintln(f, in_byte[0])
	fmt.Fprintln(f, data_length)
	fmt.Fprintln(f, string(data_conten))
	fmt.Fprintln(f, string("------------------"))
	defer f.Close()

	fmt.Println(in_byte[0])
	fmt.Println(data_length)
	fmt.Println(string(data_conten))
	fmt.Println(string("------------------"))

}

func Start() {
	listen, err := net.Listen("tcp4", ":8030")
	check(err)
	defer listen.Close()
	for range time.NewTicker(time.Second * 2).C {
		remote_tcp, err := listen.Accept()
		check(err)
		go handler(remote_tcp)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	Start()
}
