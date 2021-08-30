package main

import (
	"errors"
	"net"
	"os"
	"strings"

	"github.com/widaT/mdns"
)

func main() {
	ip, err := GetInternalIP()
	if err != nil {
		panic(err)
	}

	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	serv, err := mdns.NewServer(&mdns.Config{Zone: makeServiceWithServiceName(hostName, ip, "_foobar._tcp", 80), LocalhostChecking: true})
	if err != nil {
		panic(err)
	}
	defer serv.Shutdown()

	serv1, err := mdns.NewServer(&mdns.Config{Zone: makeServiceWithServiceName(hostName+"1", ip, "_foobar._tcp", 8008), LocalhostChecking: true})
	if err != nil {
		panic(err)
	}
	defer serv1.Shutdown()

	select {}
}

func makeServiceWithServiceName(hostName, ip, service string, port int) *mdns.MDNSService {
	m, err := mdns.NewMDNSService(
		hostName,
		service,
		"local.",
		hostName+".",
		port, // port
		[]net.IP{net.IP([]byte{172, 30, 20, 106})},
		[]string{"Local web server"}) // TXT

	if err != nil {
		panic(err)
	}

	return m
}

func GetInternalIP() (string, error) {
	// 思路来自于Python版本的内网IP获取， udp 面向无连接
	conn, err := net.Dial("udp", "114.114.114.114:80")
	if err != nil {
		return "", errors.New("internal IP fetch failed, detail:" + err.Error())
	}
	defer conn.Close()

	res := conn.LocalAddr().String()
	//	fmt.Println(res)
	res = strings.Split(res, ":")[0]
	return res, nil
}
