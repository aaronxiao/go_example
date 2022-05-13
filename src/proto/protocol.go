package proto

import (
	"encoding/json"
	"fmt"
	"net"
)

const (
	HeartBeatType = iota + 1
	RegisterType
	HealthCheckType
	ExitType
)


const (
	HeartBeatTypeRsp = iota + 1
	RegisterTypeRsp
	HealthCheckTypeRsp
	ExitTypeRsp
)


//注册请求
type Register struct {
	NickName  string 	`json:"nick_name"`
	PassWord  string 	`json:"pass_word"`
}

type MsgHead struct {
	Code     int32 		`json:"code"`
	Msg      string   	`json:"msg"`
}

type NormalRsp struct {
	MsgHead
	Body    string `json:"body"`
}


//注册返回
type RegisterRsp struct {
	MsgHead
	Body    RegisterBody `json:"body"`
}

func (r *RegisterRsp)SendToClient(conn *net.Conn)  {
	jsonStr, err := json.Marshal(r)

	data, err := Encode(RegisterTypeRsp, string(jsonStr) )
	if err != nil {
		fmt.Println("RegisterRsp SendToClient", err)
		return
	}
	(*conn).Write(data)
}

type RegisterBody struct {
	Uid    		int 		`json:"uid"`
	NickName    string 		`json:"nick_name"`
	PassWord    string 		`json:"pass_word"`
}


//心跳请求
type HeartBeat struct {
	Uid    		int 		`json:"uid"`
	Msg         string 		`json:"msg"`
}

//心跳返回
type HeartBeatRsp struct {
	MsgHead
}

func (r *HeartBeatRsp)SendToClient(conn *net.Conn)  {
	jsonStr, err := json.Marshal(r)

	data, err := Encode(HeartBeatTypeRsp, string(jsonStr) )
	if err != nil {
		fmt.Println("HeartBeatRsp SendToClient", err)
		return
	}
	(*conn).Write(data)
}



//退出
type Exit struct {
	Msg  string `json:"msg"`
}

//退出返回
type ExitRsp struct {
	MsgHead
}


func (r *ExitRsp)SendToClient(conn *net.Conn)  {
	r .Code = 0
	r.Msg = "服务停服，请退出"
	jsonStr, err := json.Marshal(r)

	data, err := Encode(ExitTypeRsp, string(jsonStr) )
	if err != nil {
		fmt.Println("RegisterRsp SendToClient", err)
		return
	}
	(*conn).Write(data)
}


//获得TcpServer信息返回
type GetTcpServerRsp struct {
	MsgHead
	Body GetTcpServerBody `json:"body"`
}


type GetTcpServerBody struct {
	AddrPort      string `json:"addr_port"`
}


type ClientInfo struct {
	Uid     int
	Ip      string
	Port    int
}



//健康检查请求
type HealthCheck struct {
	Msg         string 		`json:"msg"`
}

//心跳返回
type HealthCheckRsp struct {
	MsgHead
}

func (r *HealthCheckRsp)SendToClient(conn *net.Conn)  {
	jsonStr, err := json.Marshal(r)

	data, err := Encode(HealthCheckTypeRsp, string(jsonStr) )
	if err != nil {
		fmt.Println("HeartBeatRsp SendToClient", err)
		return
	}
	(*conn).Write(data)
}




//token 设置
type CacheTokenSet struct {
	Uuid    string `json:"uuid"`
	Token   string `json:"token"`
}
type CacheTokenSetRsp struct {
	MsgHead
	Body CacheTokenSetBody `json:"body"`
}
type CacheTokenSetBody struct {
	Token   string `json:"token"`
}


//token 获取
type CacheTokenGet struct {
	Uuid    string `json:"uuid"`
}
type CacheTokenGetRsp struct {
	MsgHead
	Body CacheTokenGetBody `json:"body"`
}
type CacheTokenGetBody struct {
	Token   string `json:"token"`
}

