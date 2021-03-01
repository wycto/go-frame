package weigo

/*
函数都写在这儿
*/
import (
	"fmt"
	"net"
	"os"
)

/**
获取IP地址
*/
func GetIP() string {
	var ip = ""
	address, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range address {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}

		}
	}

	return ip
}
