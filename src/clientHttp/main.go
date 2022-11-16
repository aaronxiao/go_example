package main

import (
	"Practice/common"
	"Practice/conf"
	"Practice/proto"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)


//tcp客户端，需要获得tcp连接信息
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


func sendHeartBeatMsg(conn net.Conn,client *proto.ClientInfo)  {
	if client.Uid > 0 {
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
	}else {
		fmt.Printf("用户：%d 心跳数据发送失败 \n", client.Uid)
	}
}

func dealRegisterRsp(msgBody string, client *proto.ClientInfo)  {
	fmt.Println("收到Server发来注册返回的数据：", msgBody)
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
	fmt.Println("收到Server发来心跳返回的数据：", msgBody)
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
	default:
		fmt.Println("无法识别Server发来的数据：", msgBody)
	}

	return 0
}

func HandleTcp(addr string) {
	defer wg1.Done()

	var client proto.ClientInfo

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()

	sendRegisterMsg(conn)
	if recvMsg(conn, &client) < 0{
		return
	}

	for {
		fmt.Printf("客户端：%d发送心跳数据！！\n", client.Uid)
		sendHeartBeatMsg(conn, &client)

		if recvMsg(conn, &client) < 0{
			fmt.Printf("客户端：%d退出\n", client.Uid)
			return
		}

		time.Sleep(time.Duration(time.Second) * time.Duration(rand.Intn(1) + 1 ) )
	}

}

var wg1 sync.WaitGroup


func GetServerAddr() string {
	url := "http://127.0.0.1:8081/getTcpServer"

	contentType := "application/json"

	sendData := fmt.Sprintf("%d", time.Now().Nanosecond() )
	resp, err := http.Post(url, contentType, strings.NewReader(sendData) )
	if err != nil {
		fmt.Println("post failed, err:", err)
		return ""
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:", err)
		return ""
	}

	getTcpServerRsp := new(proto.GetTcpServerRsp)
	json.Unmarshal(b, getTcpServerRsp)

	fmt.Println("获得一致性hash节点: ", getTcpServerRsp.Body.AddrPort)
	return getTcpServerRsp.Body.AddrPort
}


func main() {
	path := common.GetCurPwd()
	if len(path) == 0 {
		panic("GetCurPwd path is empty")
		return
	}
	conf.InitConfig(path + "\\config3.ini")

	addr := "127.0.0.1:30000"
	if remote := conf.ReadConf("client", "remote"); strings.Compare(remote, "1") == 0 {
		if addr = GetServerAddr(); len(addr) == 0 {
			fmt.Printf("获得tcp Server地址为空 客户端即将退出 \n")
			return
		}
	}

	for i := 0; i < 1; i++ {
		wg1.Add(1)
		go HandleTcp(addr)
		time.Sleep(time.Microsecond * 10)
	}
	wg1.Wait()

}
