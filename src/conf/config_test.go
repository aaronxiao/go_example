package conf

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	InitConfig("D:\\JunXiao\\bin\\config.ini")
	fmt.Println( ReadConf("host", "tcpAddr") )
}