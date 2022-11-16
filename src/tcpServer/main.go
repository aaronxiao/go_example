package main

import (
	"Practice/userCache"
	"Practice/common"
	"Practice/conf"
	"Practice/proto"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type GlobalInfo struct {
	StartUid 	int32
	lock  		sync.Mutex
}


var global GlobalInfo
var stopServer int			//停服状态 0未停服 1停服
var once sync.Once

func init()  {
	global.StartUid = 1000000
}

func PickOneUid() int32  {
	//defer global5.lock.Unlock()
	//global5.lock.Lock()
	//global5.StartUid++
	//return global5.StartUid

	new := atomic.AddInt32(&global.StartUid, 3) //加操作
	return new
}

func delClientInfo(client *proto.ClientInfo) {
	if client.Uid > 0 {
		UserCache.DelCache(client.Uid)
	}
}


func process(conn *net.Conn) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到的错误：%s\n", r)
		}
	}()

	defer (*conn).Close()
	reader := bufio.NewReader(*conn)
	clientInfo := new(proto.ClientInfo)

	defer delClientInfo(clientInfo)

	var resetTime time.Duration = 6			//心跳重置时间
	d := time.Duration(time.Second * resetTime)
	t := time.NewTimer(d)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			return //没收到客户端消息   返回
		default:
			//fmt.Println("Not timeout...")
			msgType, msgBody, err := proto.Decode(reader)
			if err == io.EOF {
				return
			}
			if err != nil {
				fmt.Println("decode msg failed, err:", err)
				return
			}
			if msgType == proto.RegisterType {
				//fmt.Printf("收到client发来的注册数据：%s \n", msgBody)
				if err = dealRegister(conn, msgBody, clientInfo); err != nil {
					return
				}
			} else if msgType == proto.HeartBeatType {
				//fmt.Printf("收到client:发来的心跳数据: %s \n", msgBody)
				if err = dealHeartBeat(conn, msgBody); err != nil {
					return
				}
			}else if msgType == proto.HealthCheckType {
				//HealthCheck
				if err = dealHealthCheck(conn, msgBody); err != nil {
					return
				}
			} else if msgType == proto.ExitType {
				return
			} else {
				fmt.Printf("没有匹配client发来%d类型的数据 即将退出 \n", msgType)
			}
			t.Reset(time.Second * resetTime)  //收到消息  重置定时器
		}
	}
}

func dealRegister(conn *net.Conn, msgBody string,clientInfo *proto.ClientInfo) error {
	//fmt.Println("收到来自客户端的地址:", conn.RemoteAddr().String() )

	var resp proto.RegisterRsp
	register := new(proto.Register)

	err := json.Unmarshal([]byte(msgBody), register)
	if err != nil{
		resp.Code = proto.StatusCode_JsonUnmarshalError
		resp.Msg  = "请求数据反序列化失败"
		resp.SendToClient(conn)
		return errors.New(fmt.Sprintf("json.Unmarshal err :%v", err) )
	}


	//clientInfo := new(proto.ClientInfo)
	clientInfo.Uid = int(PickOneUid() )
	splitSlice := strings.Split( (*conn).RemoteAddr().String(),":")
	if len(splitSlice) == 2{
		clientInfo.Ip 		= splitSlice[0]
		clientInfo.Port, _  = strconv.Atoi(splitSlice[1] )
	}

	if !UserCache.AddCache(clientInfo) {
		resp.Code = proto.StatusCode_MaxConnectError
		resp.Msg  = "超过最大连接数"
		resp.SendToClient(conn)
		return errors.New(fmt.Sprintf("超过最大连接数") )
	}

	resp.Code = proto.StatusCode_Success
	resp.Msg  = "请求注册成功"
	resp.Body.Uid = clientInfo.Uid
	resp.Body.NickName = register.NickName
	resp.Body.PassWord = register.PassWord

	resp.SendToClient(conn)

	return nil
}


func dealHeartBeat(conn *net.Conn, msgBody string) error {
	var resp proto.HeartBeatRsp
	heartBeat :=new( proto.HeartBeat)

	err := json.Unmarshal([]byte(msgBody), heartBeat)
	if err != nil {
		resp.Code = proto.StatusCode_JsonUnmarshalError
		resp.Msg  = "请求数据反序列化失败"
		resp.SendToClient(conn)
		return errors.New(fmt.Sprintf("json.Unmarshal err :%v", err) )
	}


	resp.Msg = "PONG"
	resp.Code = proto.StatusCode_Success

	resp.SendToClient(conn)
	return nil
}



func dealHealthCheck(conn *net.Conn, msgBody string) error {
	var resp proto.HealthCheckRsp
	healthCheck :=new( proto.HealthCheck)

	err := json.Unmarshal([]byte(msgBody), healthCheck)
	if err != nil {
		resp.Code = proto.StatusCode_JsonUnmarshalError
		resp.Msg  = "请求数据反序列化失败"
		resp.SendToClient(conn)
		return errors.New(fmt.Sprintf("json.Unmarshal err :%v", err) )
	}

	resp.Msg = "在的"
	resp.Code = proto.StatusCode_Success
	if stopServer == 1 {		//停服了
		resp.Msg = "抱歉我已经停服"
		resp.Code = proto.StatusCode_StopServerError
	}

	resp.SendToClient(conn)
	return nil
}

var connectNum int

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("main捕获到的错误：%s\n", r)
		}
	}()

	path := common.GetCurPwd()
	if len(path) == 0{
		panic("GetCurPwd path is empty")
		return
	}

	if runtime.GOOS != "windows"  {
		conf.InitConfig(path + "//config3.ini" )
	}else {
		conf.InitConfig(path + "\\config3.ini" )
	}
	tcpAddr := conf.ReadConf("host", "tcpAddr")
	tcpPort := conf.ReadConf("host", "tcpPort")
	if len(tcpAddr) == 0 || len(tcpPort) == 0{
		panic("tcpAddr or tcpPort is empty")
		return
	}

	fmt.Printf("NumCPU: %d \n", runtime.NumCPU() )
	//runtime.GOMAXPROCS(4)

	addr := tcpAddr + ":" + tcpPort
	listen, err := net.Listen("tcp", addr )
	if err != nil {
		fmt.Println("listen failed, err:", err )
		return
	}
	defer listen.Close()

	go PrintInfo()
	go RegisterTcpServer()
	go HttpServer()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v \n", err )
			continue
		}
		//fmt.Printf("收tcp连接 \n")
		dealStopServer(&conn)				//停服判断
		connectNum++
		go process(&conn)
	}
}

func dealStopServer(conn *net.Conn)  {
	once.Do(func() {
		stop := conf.ReadConf("host", "stopServer")
		stopServer,_ = strconv.Atoi(stop)
	})

	//fmt.Printf("打印服务端停服状态值：%d \n", stopServer)
	if stopServer == 1 {		//停服状态  通知客户端
		exit := new(proto.ExitRsp)
		exit.SendToClient(conn)
	}
}

//打印服务器信息
func PrintInfo()  {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("PrintInfo捕获到的错误：%s\n", r)
		}
	}()

	lastConnectNum := connectNum
	//lastNum := UserCache.GetCacheNum()
	ticker := time.Tick(time.Second * 1) //定义一个1秒间隔的定时器
	for {
		select {
		case <-ticker:
			currNum := UserCache.GetCacheNum()
			//diff := currNum - lastNum
			//lastNum = currNum

			curConnectNum := connectNum
			diffConnect := curConnectNum - lastConnectNum
			lastConnectNum = curConnectNum

			//fmt.Printf("当前客户端连接数：%d 总连接数%d, 每秒新建连数%d, goroutine数: %d \n",
			//				currNum, curConnectNum, diffConnect, runtime.NumGoroutine() )

			fmt.Printf("当前客户端连接数：%d 总连接数%d, 每秒新建连数%d \n",
				currNum, curConnectNum, diffConnect )

			NewMonitor()
			//var stats runtime.MemStats
			//runtime.ReadMemStats(&stats)
			//fmt.Printf("MemStats: %+v\n", stats)
		}
	}
}




func RegisterTcpServer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("RegisterTcpServer捕获到的错误：%s\n", r)
		}
	}()

	regMark :=  conf.ReadConf("host", "register")
	if strings.Compare(regMark, "1") == 0 {
		url := "http://127.0.0.1:8081/registerServer"

		contentType := "application/json"

		tcpAddr := conf.ReadConf("host", "tcpAddr")
		tcpPort := conf.ReadConf("host", "tcpPort")

		sendData := fmt.Sprintf("%s:%s", tcpAddr, tcpPort)
		resp, err := http.Post(url, contentType, strings.NewReader(sendData))
		if err != nil {
			fmt.Println("post failed, err:", err)
			return
		}
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("get resp failed,err:", err)
			return
		}

		fmt.Printf("注册tcp服务返回：%s", string(b))
	}
	return
}

func HttpServer()  {

	listen := conf.ReadConf("host", "tcpHttp")
	http.HandleFunc("/server/status", doStatus)				//修改服务器状态
	//启动服务
	fmt.Printf("listen: %s \n", listen)
	http.ListenAndServe(listen, nil)
}


func doStatus(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	//data, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	fmt.Println("request failed, err:%v\n", err)
	//	return
	//}

	if stopServer == 1{		//停服
		stopServer = 0		//开启
	}else {
		stopServer = 1		//停服
	}
	w.Write( []byte(`{"code": 0, "msg":"修改成功"}`) )
}



//增大客户端连接数  在服务器上查看实际连接数
//在客户端心跳包上做统计    统计每次发送一个心跳来回需要的时间和发送总数   总时间/总发送数 = 响应延迟   由于是在同一个局域网  不考虑数据传输延迟