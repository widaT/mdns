# mdns

纯go实现基于mdns服务发现


## example

service endpoint

```go
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

```


client endpoint

```go
package main

import (
	"fmt"

	"github.com/widaT/mdns"
)

func main() {
	//这边的数量影响一次查询出的解析数量
	entries := make(chan *mdns.ServiceEntry, 3)
	err := mdns.Lookup("_foobar._tcp", entries)
	if err != nil {
		panic(err)
	}

		s := time.Now()
out:
	for {
		select {
		case item := <-entries:
			fmt.Println(item.Host)
			fmt.Println(item.Name)
			fmt.Println(item.AddrV4)
			fmt.Println(item.Port)

		case <-time.After(20 * time.Millisecond):
			break out
		}
	}

	fmt.Println(time.Since(s))
}

```
