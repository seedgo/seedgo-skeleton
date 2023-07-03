package util

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"net"
)

var snowNode *snowflake.Node

func init() {
	localIp, err := getClientIp()
	if err != nil {
		fmt.Println(err.Error())
	}
	nodeId, err := Ipv4ToLong(localIp)
	if err != nil {
		fmt.Println(err.Error())
	}

	id := nodeId % 1024
	fmt.Printf("localIp: %s, nodeId: %d \n", localIp, id)
	snowNode, err = snowflake.NewNode(int64(id))
	if err != nil {
		fmt.Printf("init snowflake failed: %s \n", localIp, err)
	}
}

func Ipv4ToLong(ip string) (uint, error) {
	p := net.ParseIP(ip).To4()
	if p == nil {
		return 0, errors.New("invalid ipv4 format")
	}

	return uint(p[0])<<24 | uint(p[1])<<16 | uint(p[2])<<8 | uint(p[3]), nil
}

func getClientIp() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}

		}
	}

	return "", errors.New("Can not find the client ip address!")

}

func NextUid() string {
	// Generate a snowflake ID.
	return snowNode.Generate().String()
}
