package conf

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	InitConfig("D:\\JunXiao\\bin\\config3.ini")
	fmt.Println( ReadConf("host", "tcpAddr") )
}