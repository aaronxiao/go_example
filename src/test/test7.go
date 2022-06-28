package main

import (
	"encoding/json"
	"fmt"
)

var cfgTide = `
  [
    {
      "id": 1,
      "name": "\u9c7c\u6f6e1",
      "table_level": [1,2,3],
      "table_type": 1,
      "swim_time": 60000,
      "swim_speed": 20,
      "area_fish": {
        "1": {
          "fish": [1,2,3]
        },
        "2": {
          "fish": [4,5,6]
        },
        "3": {
          "fish": [7,8,9]
        }
      }
    },
    {
      "id": 2,
      "name": "\u9c7c\u6f6e2",
      "table_level": [2,1,3],
      "table_type": 2,
      "swim_time": 65000,
      "swim_speed": 25,
      "area_fish": {
        "1": {
          "fish": [1,2,3]
        },
        "2": {
          "fish": [4,5,6]
        },
        "3": {
          "fish": [7,8,9]
        },
        "4": {
          "fish": [10,11,12]
        },
        "5": {
          "fish": [13,14,15]
        },
        "6": {
          "fish": [16,17,18]
        }
      }
    }
  ]
`

type TideFish struct {
	Fish []int `json:"fish"` //鱼id
}

type Tide struct {
	Id         int                  `json:"id"`
	Name       string               `json:"name"`        //鱼潮名称
	TableLevel []int       			`json:"table_level"` //场次类型 可多选
	TideType   int                  `json:"tide_type"`   //鱼潮阵型 1-n
	SwimTime   int64                `json:"swim_time"`   //游动时常 毫秒
	SwimSpeed  int                  `json:"swim_speed"`  //游动速度  像素每秒
	SwimDirect int                  `json:"swim_direct"` //游动方向  1 上下方向 2 左右方向  服务端自己赋值
	AreaFish   map[string]*TideFish `json:"area_fish"`   //色块鱼配置

}

type Tides struct {
	Cfg   []*Tide `json:"cfg"` //鱼潮配置配置
}


func main() {
	var tides Tides
	err := json.Unmarshal([]byte(cfgTide), &tides.Cfg)

	fmt.Println(err)

	//tides.Cfg = make([]*Tide, 2) //鱼潮配置配置
	//tides.Cfg[0] = new(Tide)
	//tides.Cfg[0].Id = 1
	//tides.Cfg[0].Name = "鱼潮1"
	//tides.Cfg[0].TableLevel = []int{1,2,3}
	//tides.Cfg[0].TideType = 1
	//tides.Cfg[0].SwimTime = 60000
	//tides.Cfg[0].SwimSpeed = 20
	//tides.Cfg[0].SwimDirect = 1
	//tides.Cfg[0].AreaFish = map[string]*TideFish{ "1": {[]int{1,2,3} }, "2":{[]int{4,5,6} } }
	//
	//tides.Cfg[1] = new(Tide)
	//tides.Cfg[1].Id = 2
	//tides.Cfg[1].Name = "鱼潮2"
	//tides.Cfg[1].TableLevel = []int{1,2,3}
	//tides.Cfg[1].TideType = 1
	//tides.Cfg[1].SwimTime = 60000
	//tides.Cfg[1].SwimSpeed = 20
	//tides.Cfg[1].SwimDirect = 1
	//tides.Cfg[1].AreaFish = map[string]*TideFish{ "1": {[]int{7,8,9} }, "2":{[]int{10,11,12} } }
	//
	//res ,err := json.Marshal(&tides.Cfg)
	//
	//fmt.Println(string(res), err)
	//
	//var tides1 Tides
	//err1 := json.Unmarshal(res,&tides1.Cfg)
	//
	//fmt.Println(err1)
}
