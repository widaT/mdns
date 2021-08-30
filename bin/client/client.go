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

	for item := range entries {
		fmt.Println(item.Host)
		fmt.Println(item.Name)
		fmt.Println(item.AddrV4)
		fmt.Println(item.Port)
	}

}
