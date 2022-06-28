package main


import (
	"encoding/json"
	"fmt"
	"sync"
)

var (
	cfgPoint = `{"login":{"enabled":false,"conds":{"1":[{"jackpot":100,"usable":20000}],"2":[{"jackpot":200,"usable":300000}],"3":[{"jackpot":300,"usable":5000000}]},"y":{"1":120,"2":120,"3":200,"4":200},"default_y":10,"fishes":[3,4],"up_limit":5},"newer":{"enabled":true,"moneys":{"A":80000,"B":20000},"y":{"2":200,"1":200,"3":1000,"4":1000},"default_y":10,"fishes":[3,4,1,2],"up_limit":5},"bankruptcy":{"enabled":false,"conds":{"1":[{"money":1000,"jackpot":20000,"usable":20000}],"2":[{"money":1000,"jackpot":500000,"usable":300000}],"3":[{"money":5000,"jackpot":10000000,"usable":1000000}]},"y":{"1":120,"2":120,"3":200,"4":200},"default_y":10,"fishes":[4,3],"up_limit":5},"recharge":{"enabled":false,"usables":{"A":14},"y":{"1":234,"2":567},"default_y":10},"low_chips":{"enabled":false,"conds":{"1":[{"money":2000,"jackpot":20000,"usable":20000}],"2":[{"money":50000,"jackpot":500000,"usable":300000}],"3":[{"money":2000000,"jackpot":10000000,"usable":5000000}]},"y":{"1":120,"2":120,"3":200,"4":200},"default_y":10,"fishes":[4,3],"up_limit":5},"texas_consume":{"enabled":false,"up_limit":{"1":1233},"moneys":{"A":1444},"default_y":10}}`
)

var cfgPoint1 = `{"login":{"enable":true,"usables":null,"conds":{"1":[{"money":0,"jackpot":100,"usable":20000}],"2":[{"money":0,"jackpot":200,"usable":300000}],"3":[{"money":0,"jackpot":300,"usable":5000000}]},"fishes":[3,4],"up_limit":5,"y":{"1":120,"2":120,"3":200,"4":200},"default_y":10},"newer":{"enable":false,"moneys":{"A":80000,"B":20000},"fishes":[3,4,1,2],"up_limit":5,"y":{"1":200,"2":200,"3":1000,"4":1000},"default_y":10},"low_chips":{"enable":false,"conds":{"1":[{"money":2000,"jackpot":20000,"usable":20000}],"2":[{"money":50000,"jackpot":500000,"usable":300000}],"3":[{"money":2000000,"jackpot":10000000,"usable":5000000}]},"fishes":[4,3],"up_limit":5,"y":{"1":120,"2":120,"3":200,"4":200},"default_y":10},"bankruptcy":{"enable":false,"conds":{"1":[{"money":1000,"jackpot":20000,"usable":20000}],"2":[{"money":1000,"jackpot":500000,"usable":300000}],"3":[{"money":5000,"jackpot":10000000,"usable":1000000}]},"fishes":[4,3],"up_limit":5,"y":{"1":120,"2":120,"3":200,"4":200},"default_y":10},"recharge":{"enable":false,"usables":{"A":14},"y":{"1":234,"2":567},"default_y":10},"texas_consume":{"enable":false,"moneys":{"A":1444},"up_limit":{"1":1233}}}`


type PointTargetCondCfg struct {
	Money   int64 `json:"money"`   //筹码游戏币
	Jackpot int64 `json:"jackpot"` //个人财经池
	Usable  int64 `json:"usable"`  //取出多少可用值
}

type Point struct {
	Cfg   *PointCfg `json:"cfg"`

	sync.RWMutex
}

type FishCategory int
type TableLevel int

type NewerLabelType string
type ConsumeLabelType string
type RechargeLabelType string


type PointTargetCfg struct {
	Conds map[TableLevel][]*PointTargetCondCfg `json:"conds"` //场次下面多个触发条件
}


type PointStopCfg struct {
	Fishes  []FishCategory `json:"fishes"`   //终止条件 累计计数category
	UpLimit int              `json:"up_limit"` //终止条件 累计总数
}

type PointYCfg struct {
	Y        map[FishCategory]int `json:"y"`         //Y值
	DefaultY int                    `json:"default_y"` //默认Y值，在Y值中没有找到合适的就是用当前值
}

type PointCfg struct {
	Login struct {
		Enable  bool                   `json:"enable"`
		Usables map[TableLevel]int64 `json:"usables"` //可用值
		PointTargetCfg
		*PointStopCfg
		*PointYCfg
	} `json:"login"`

	Newer struct {
		Enable bool                       `json:"enable"`
		Moneys map[NewerLabelType]int64 `json:"moneys"` //赠送筹码
		PointStopCfg
		PointYCfg
	} `json:"newer"`

	LowerMoney struct {
		Enable bool `json:"enable"`
		PointTargetCfg
		PointStopCfg
		PointYCfg
	} `json:"low_chips"`

	Bankruptcy struct {
		Enable bool `json:"enable"`
		PointTargetCfg
		PointStopCfg
		PointYCfg
	} `json:"bankruptcy"`

	Recharge struct {
		Enable  bool                          `json:"enable"`
		Usables map[RechargeLabelType]int64 `json:"usables"` //可用值
		PointYCfg
	} `json:"recharge"`

	Consume struct {
		Enable  bool                         `json:"enable"`
		Moneys  map[ConsumeLabelType]int64 `json:"moneys"`   //赠送筹码
		UpLimit map[TableLevel]int64       `json:"up_limit"` //投放上限
	} `json:"texas_consume"`
}

func NewPoint() *Point {
	return &Point{
	}
}

func main() {
	c := NewPoint()
	err := json.Unmarshal([]byte(cfgPoint1), &c.Cfg)
	//t.Fatalf(fmt.Sprintf("val %#v", c.Cfg.Login))
	fmt.Printf("%+v, %v \n" , c.Cfg.Login, err)
	//assert.Nil(t, err, "load cfg error")

	c.Cfg.Login.Enable = true

	jsonByte,_ :=  json.Marshal(&c.Cfg)
	fmt.Println( string(jsonByte))

	c1 := NewPoint()
	json.Unmarshal([]byte(jsonByte), &c1.Cfg)

	fmt.Printf("%+v \n" , c1.Cfg)

}