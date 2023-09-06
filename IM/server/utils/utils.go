package utils

import (
	"IM/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8192]byte
}

func (tf *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 4096)
	//conn.Read is blocked only when conn(or Client) is not be closed.
	_, err = tf.Conn.Read(tf.Buf[:4])
	if err != nil {
		//fmt.Println("read pkg header error!")
		return
	}

	pkgLen := binary.BigEndian.Uint32(tf.Buf[:4])
	fmt.Println("PkgLen:", pkgLen)

	n, err := tf.Conn.Read(tf.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//fmt.Println("read pkg body error!")
		return
	}

	err = json.Unmarshal(tf.Buf[:pkgLen], &mes) //& is important!
	if err != nil {
		fmt.Println("json.Unmarshal(buf[:pkgLen], &mes) failed! Error:", err)
		return
	}
	return

	//fmt.Println("buf Sever has read =", buf[:4])
}

func (tf *Transfer) WritePkg(data []byte) (err error) {
	pkgLen := uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(tf.Buf[:4], pkgLen)
	//Send length
	_, err = tf.Conn.Write(tf.Buf[:4])
	if err != nil {
		fmt.Println("conn.Write(buf) failed! Error:", err)
		return
	}

	//Send data
	n, err := tf.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) failed! Error:", err)
		return
	}

	return
}
