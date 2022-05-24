package main

import (
	"Practice/common"
	"Practice/conf"
	"Practice/proto"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net"
	"sync"
	"time"
)

func sendRegisterMsg(conn net.Conn)  {
	var register proto.Register

	register.NickName = fmt.Sprintf("boyaa_%d", rand.Intn(100) )
	register.PassWord = fmt.Sprintf("by_%d", 1000 + rand.Intn(100) )

	message1, _ := json.Marshal(&register)

	data, err := proto.Encode(proto.RegisterType, string(message1) )
	if err != nil {
		fmt.Println("encode msg failed, err:", err)
		return
	}
	conn.Write(data)
}


func sendHeartBeatMsg(conn net.Conn,client *proto.ClientInfo) {

	var heartBeat proto.HeartBeat
	heartBeat.Msg = "PING"
	heartBeat.Uid = client.Uid

	message1, _ := json.Marshal(&heartBeat)

	data, err := proto.Encode(proto.HeartBeatType, string(message1))
	if err != nil {
		fmt.Println("encode msg failed, err:", err)
		return
	}
	conn.Write(data)

}

func dealRegisterRsp(msgBody string, client *proto.ClientInfo)  {
	//fmt.Println("收到Server发来注册返回的数据：", msgBody)
	var resp proto.RegisterRsp

	err := json.Unmarshal([]byte(msgBody), &resp)
	if err == nil {
		if resp.Code != proto.StatusCode_Success{
			fmt.Printf("收到Server发来注册错误：: %s \n", resp.Msg )
		}else {
			client.Uid = resp.Body.Uid
		}
	}else{
		fmt.Printf("json.Unmarshal err: %v \n", err )
	}
}


func dealHeartBeatRsp(msgBody string)  {
	//fmt.Println("收到Server发来心跳返回的数据：", msgBody)
	var resp proto.RegisterRsp

	err := json.Unmarshal([]byte(msgBody), &resp)
	if err == nil {
		if resp.Code != proto.StatusCode_Success{
			fmt.Printf("收到Server发来心跳返回错误： %s \n", resp.Msg )
		}else {
			//
		}
	}else{
		fmt.Printf("json.Unmarshal err: %v \n", err )
	}
}


//接收数据
func recvMsg(conn net.Conn, client *proto.ClientInfo) int  {

	reader := bufio.NewReader(conn)
	msgType, msgBody, err := proto.Decode(reader)

	if err == io.EOF {
		return -1
	}
	if err != nil {
		fmt.Println("decode msg failed, err:", err)
		return -1
	}

	switch msgType {
	case proto.RegisterTypeRsp:
		dealRegisterRsp(msgBody, client)
	case proto.HeartBeatTypeRsp:
		dealHeartBeatRsp(msgBody)
	case proto.ExitTypeRsp:
		fmt.Println("收到服务端停服消息 准备退出！")
		return -1
	default:
		fmt.Println("无法识别Server发来的数据：", msgBody)
	}

	return 0
}

func HandleTcp(addr string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到的错误：%s\n", r)
		}
	}()
	defer wg1.Done()

	var client proto.ClientInfo

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		sendRegisterMsg(conn)
		for {
			sendHeartBeatMsg(conn, &client)
			select {
			case <-ctx.Done():
				//fmt.Println("监控退出，停止了...")
				return
			default:
				//fmt.Println("goroutine监控中...")
				time.Sleep(time.Duration(time.Second) * time.Duration(1 ) )
			}
		}
	}(ctx)


	for {
		//fmt.Printf("客户端：%d发送心跳数据！！\n", client.Uid)
		//sendHeartBeatMsg(conn, &client)
		if recvMsg(conn, &client) < 0{
			fmt.Printf("客户端：%d退出\n", client.Uid)
			cancel()
			return
		}
		//time.Sleep(time.Duration(time.Second) * time.Duration(rand.Intn(1) + 3 ) )
		//time.Sleep(time.Duration(time.Second) * time.Duration(1 ) )
	}
}

var wg1 sync.WaitGroup



func main() {
	path := common.GetCurPwd()
	if len(path) == 0 {
		panic("GetCurPwd path is empty")
		return
	}
	conf.InitConfig(path + "\\config.ini")

	addr := "192.168.56.101:9002"
	//tcpAddr := conf.ReadConf("host", "tcpAddr")
	//tcpPort := conf.ReadConf("host","tcpPort")
	//
	//if len(tcpAddr) > 0 && len(tcpPort) > 0{
	//	addr = fmt.Sprintf("%s:%s", tcpAddr, tcpPort)
	//}
	for i := 0; i < 100; i++ {
		wg1.Add(1)
		go HandleTcp(addr)
		time.Sleep(time.Microsecond * 1000)
	}
	wg1.Wait()

}
