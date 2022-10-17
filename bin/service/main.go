package main

import (
	"net"

	"github.com/widaT/mdns"
)

func main() {

	hostName1, hostName2 := "host1", "host2"
	serv, err := mdns.NewServer(&mdns.Config{Zone: makeServiceWithServiceName([]net.IP{mdns.GetOutboundIP()}, hostName1, "_foobar._tcp", 80), LocalhostChecking: true})
	if err != nil {
		panic(err)
	}
	defer serv.Shutdown()

	serv1, err := mdns.NewServer(&mdns.Config{Zone: makeServiceWithServiceName([]net.IP{mdns.GetOutboundIP()}, hostName2, "_foobar._tcp", 8008), LocalhostChecking: true})
	if err != nil {
		panic(err)
	}
	defer serv1.Shutdown()

	select {}
}

func makeServiceWithServiceName(ip []net.IP, hostName, service string, port int) *mdns.MDNSService {
	m, err := mdns.NewMDNSService(
		hostName,
		service,
		"local.",
		hostName+".",
		port, // port
		ip,
		[]string{"Local web server"}) // TXT

	if err != nil {
		panic(err)
	}

	return m
}
