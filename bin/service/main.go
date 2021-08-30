package main

import (
	"net"
	"os"

	"github.com/widaT/mdns"
)

func main() {

	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	serv, err := mdns.NewServer(&mdns.Config{Zone: makeServiceWithServiceName([]net.IP{mdns.GetOutboundIP()}, hostName, "_foobar._tcp", 80), LocalhostChecking: true})
	if err != nil {
		panic(err)
	}
	defer serv.Shutdown()

	serv1, err := mdns.NewServer(&mdns.Config{Zone: makeServiceWithServiceName([]net.IP{mdns.GetOutboundIP()}, hostName+"1", "_foobar._tcp", 8008), LocalhostChecking: true})
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
