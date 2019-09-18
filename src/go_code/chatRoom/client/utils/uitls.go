package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatRoom/common/message"
	"io"
	"net"
)

// 将这里的方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte //传输时使用的缓冲
}

func (t *Transfer) ReadPkg() (mes1 message.Message, err error) {

	fmt.Println("读取服务端发送的数据, 我是客户端......")
	n3, err := t.Conn.Read(t.Buf[:4]) // 读取四个,这里会阻塞
	fmt.Println("n3是什么", n3)
	if err != nil {
		if err == io.EOF {
			fmt.Println("客户端退出触发EOF,我是客户端的EOF")
			return
		} else {
			fmt.Println("读取客户端返回的数据出错:", err)
			return
		}

	}

	fmt.Println("buf读到的内容是:", t.Buf[:4])
	// 把 buf[:4] 转成unit32的类型,作为读取字节的长度
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(t.Buf[:4]) // Uint32方法:传入一个byte类型的切片,返回一个uint32的数据
	fmt.Println("pkgLen---->", pkgLen)

	// 根据pkgLen的长度来读取数据,从conn中读取pkgLen那么多字节,放入buf中
	n, err := t.Conn.Read(t.Buf[:pkgLen]) // 这里出错了,n返回的是0
	fmt.Println("n------->", n)
	if n != int(pkgLen) {
		// 说明丢包了
		fmt.Println("丢包了!!!!!!!")
		return
	}
	fmt.Println("buf[:pkgLen]---->", t.Buf[:pkgLen])
	// 把读取的数据反序列化Message,再根据Type,把Data反序列化成相应的结构体
	err = json.Unmarshal(t.Buf[:pkgLen], &mes1) // 一定要记得加&,如果不加,mes是空的
	fmt.Println("客户端的mes是什么---------->", mes1)
	if err != nil {
		fmt.Println("服务器反序列化客户端的数据出错:", err)
		return
	}

	return
}

// 将数据写入包中,并发送
func (t *Transfer) WriterPkg(data []byte) (err error) {

	// 先发送一个长度给对方
	var pkgLen uint32                             // 无符号的,能表示比较大的数字
	pkgLen = uint32(len(data))                    // 将int类型的数字强转成uint32
	binary.BigEndian.PutUint32(t.Buf[:4], pkgLen) // PutUint32将一个uint32的数据类型转换成一个byte切片

	// 发送长度
	n, err := t.Conn.Write(t.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("发送长度出错", err)
		return
	}

	// 发送内容
	n1, err := t.Conn.Write(data)
	if n1 != int(pkgLen) || err != nil {
		fmt.Println("发送内容出错", err)
		return
	}
	return
}
