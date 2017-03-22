package main

import (
	"net"
	"log"
	"fmt"

	"time"
	"os"
)


func main() {

	interval := "1s"
	if len(os.Args) > 1 {
		interval = os.Args[1]
	}

	d, err := time.ParseDuration(interval)
	if err != nil {
		log.Fatal(err)
	}
	ticker := time.NewTicker(d)
	for {
		select {
		case <- ticker.C:
			displayIPs()
		}
	}
}

func displayIPs() {
	fmt.Print("\u001B[2J\u001B[0;0f")
	fmt.Println("IPs:", time.Now())
	ifaces, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	for _, iface := range ifaces {
		if ipnet, ok := iface.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Printf("ip:%s - mask:%v \n",ipnet.IP.To4(), ipnet.IP.Mask(ipnet.IP.DefaultMask()) )
			}
		}

	}
}