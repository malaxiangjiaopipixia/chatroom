package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("等待读取数据...")
	time.Sleep(time.Second * 2)
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}

	pkgLen := binary.BigEndian.Uint32(buf[:4])

	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read(buf[:pkgLen] err=", err)
		return
	}

	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal faild err=", err)
		return
	}

	return

}

func writePkg(conn net.Conn, data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte

	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	n, err := conn.Write(buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.Write(bytes) faild!", err)
		return
	}

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err=", err)
		return
	}

	return

}
