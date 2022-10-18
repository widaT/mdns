package main

import (
	"fmt"
	"time"

	"github.com/widaT/mdns"
)

func main() {
	//这边的数量影响一次查询出的解析数量
	entries := make(chan *mdns.ServiceEntry, 30)
	err := mdns.Lookup("_foobar._tcp", entries)

	/*	默认超时是1s 可以使用如下修改
		params := mdns.DefaultParams("_foobar._tcp")
		params.Timeout = 20 * time.Millisecond
		params.Entries = entries
		err := Query(params)
	*/
	if err != nil {
		panic(err)
	}

	// for item := range entries {
	// 	fmt.Println(item.Host)
	// 	fmt.Println(item.Name)
	// 	fmt.Println(item.AddrV4)
	// 	fmt.Println(item.Port)
	// }
	s := time.Now()
out:
	for {
		select {
		case item := <-entries:
			fmt.Println(item.Host)
			fmt.Println(item.Name)
			fmt.Println(item.AddrV4)
			fmt.Println(item.Port)
		default:
			break out
			// case <-time.After(20 * time.Millisecond):
			// 	break out
		}
	}

	fmt.Println(time.Since(s))
}
