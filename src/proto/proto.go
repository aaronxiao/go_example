package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)


func Encode(msgType int32, message string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)	//包体长度
	if err != nil {
		return nil, err
	}
	err = binary.Write(pkg, binary.LittleEndian, msgType)	//消息类型
	if err != nil {
		return nil, err
	}

	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode 解码消息
func Decode(reader *bufio.Reader) (int32, string, error) {
	// 读取消息的长度
	lengthByte, _ := reader.Peek(8) // 读取前4个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	//var length int32
	type Head struct {
		length 		int32					//消息长度
		ntype       int32					//消息类型
	}
	head := new(Head)
	err := binary.Read(lengthBuff, binary.LittleEndian, &head.length)
	if err != nil {
		return head.ntype, "", err
	}

	err = binary.Read(lengthBuff, binary.LittleEndian, &head.ntype)
	if err != nil {
		return head.ntype, "", err
	}

	//fmt.Println("head info: ", head)

	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < head.length+8 {
		return head.ntype, "", err
	}
	// 读取真正的消息数据
	pack := make([]byte, int(8+head.length))
	_, err = reader.Read(pack)
	if err != nil {
		return head.ntype, "", err
	}

	return head.ntype, string(pack[8:]), nil
}