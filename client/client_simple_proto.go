package main

import (
	"github/leiwb/simple_proto/utils"
	"net"
	"time"
)

func Start() {
	conn, err := net.Dial("tcp4", "127.0.0.1:8030")
	check(err)
	defer conn.Close()
	for range time.NewTicker(time.Second * 2).C {
		write_byte := make([]byte, 0)
		write_byte = append(write_byte, 0)
		send_str := "i am send server msg"
		write_byte = append(write_byte, utils.Int32_2_byte(int32(len(send_str)))...)
		write_byte = append(write_byte, []byte(send_str)...)
		_, err := conn.Write(write_byte)
		check(err)
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
