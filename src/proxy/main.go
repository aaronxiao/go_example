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
	"net"
	"net/http"
	"sync"
	"time"
)


//设置集群地址，最好内网IP
var hostArray= []string{"127.0.0.1:30000","127.0.0.1:30001", "127.0.0.1:30002"}
var hashConsistent *common.Consistent
var hostMap sync.Map					//tcp服信息 [string]1


func GetTcpServer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	rsp := new(proto.GetTcpServerRsp)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read request.Body failed, err: \n", err)
		w.Write([]byte(`{"code":4, "msg":"http读取数据失败"}`))
		return
	}
	fmt.Println(string(b))

	var body proto.GetTcpServerBody
	addrPort, err := hashConsistent.Get(string(b))
	if err != nil{
		fmt.Printf("获取一致性hash节点失败: %v \n", err)
	}

	if len(addrPort) == 0 {
		rsp.Code = proto.StatusCode_ConsistentEmptyError
		rsp.Msg = fmt.Sprintf("获取一致性hash空的节点")
	}else{
		rsp.Code = 0
		rsp.Msg = fmt.Sprintf("获取到一致性hash节点: %s", addrPort)
		body.AddrPort = addrPort
		rsp.Body = body
	}

	jsonStr,err := json.Marshal(rsp)
    if err != nil{
		w.Write([]byte(`{"code":1, "msg":"json序列化数据失败"}`))
		return
	}
	w.Write([]byte(jsonStr))

}



func RegisterTcpServer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read request.Body failed, err: \n", err)
		w.Write([]byte(`{"code":4, "msg":"http读取数据失败"}`))
		return
	}
	fmt.Println("添加一致性hash节点:", string(b))
	hashConsistent.Add(string(b))
	hostMap.Store(string(b), 1)

	w.Write([]byte(`{"code":0, "msg":"添加一致性hash节点成功"}`))
	return
}

func main()  {
	path := common.GetCurPwd()
	if len(path) == 0{
		panic("GetCurPwd path is empty")
		return
	}

	conf.InitConfig(path + "\\config3.ini" )

	hashConsistent = common.NewConsistent()
	//采用一致性hash算法，添加节点
	//for _,v :=range hostArray {
	//	hashConsistent.Add(v)
	//}

	http.HandleFunc("/getTcpServer", GetTcpServer)
	http.HandleFunc("/registerServer", RegisterTcpServer)
	
	
	go HealthyCheck()
	go PrintInfo()

	listen := conf.ReadConf("host", "proxyHttp")
	fmt.Printf("listen: %s \n", listen)

	//启动服务
	http.ListenAndServe(listen, nil)
}


//打印服务器信息
func PrintInfo()  {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("PrintInfo捕获到的错误：%s\n", r)
		}
	}()

	//lastNum := UserCache.GetCacheNum()
	ticker := time.Tick(time.Second * 2) //定义一个1秒间隔的定时器
	for {
		select {
		case <-ticker:
			var printInfo string
			hostMap.Range(func(k, _ interface{}) bool {
				tmp := fmt.Sprintf("%s : ", k)
				printInfo = printInfo + tmp
				return true
			})
			fmt.Println("当前节点信息：", printInfo)
		}
	}
}

//HealthyCheck
func HealthyCheck()  {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("PrintInfo捕获到的错误：%s\n", r)
		}
	}()

	//lastNum := UserCache.GetCacheNum()
	ticker := time.Tick(time.Second * 2) //定义一个2秒间隔的定时器
	for {
		select {
		case <-ticker:
			hostMap.Range(func(k, _ interface{}) bool {
				go HandleTcp( k.(string) )
				return true
			})
		}
	}
}



func HandleTcp(addr string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到的错误：%s\n", r)
		}
	}()
	var client proto.ClientInfo

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		hostMap.Delete(addr)
		hashConsistent.Remove(addr)
		fmt.Println("检查tcp服务节点连接失败 删除该节点信息", err)
		return
	}
	defer conn.Close()
	sendHealthCheckMsg(conn, &client)
	if recvMsg(conn, &client) < 0 {
		hostMap.Delete(addr)
		hashConsistent.Remove(addr)
		return
	}
}



func sendHealthCheckMsg(conn net.Conn,client *proto.ClientInfo) {

	var healthCheck proto.HealthCheck
	healthCheck.Msg = "停服了么"

	message1, _ := json.Marshal(&healthCheck)

	data, err := proto.Encode(proto.HealthCheckType, string(message1))
	if err != nil {
		fmt.Println("encode msg failed, err:", err)
		return
	}
	conn.Write(data)

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
	case proto.HealthCheckTypeRsp:
		fmt.Println("收到Server发来健康检查数据：", msgBody)
		var resp proto.HealthCheckRsp

		err := json.Unmarshal([]byte(msgBody), &resp)
		if err == nil {
			if resp.Code != proto.StatusCode_Success{
				return -1
			}else {
				return 0
			}
		}else{
			fmt.Printf("json.Unmarshal err: %v \n", err )
			return -1
		}
	default:
		fmt.Println("无法识别Server发来的数据：", msgBody)
	}
	return 0
}