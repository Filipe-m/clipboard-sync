package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hashicorp/mdns"
)

func main() {
	serviceName := "_clipboard._quic"
	host, _ := os.Hostname()
	info := []string{"My awesome service"}
	service, _ := mdns.NewMDNSService(host, serviceName, "", "", 9000, nil, info)

	server, _ := mdns.NewServer(&mdns.Config{Zone: service})
	defer server.Shutdown()

	for {
		entriesCh := make(chan *mdns.ServiceEntry, 4)

		go func() {
			for entry := range entriesCh {
				fmt.Printf("Service found: %s\n", entry.Host)
				for _, ip := range entry.AddrV4 {
					fmt.Printf("  -> IP: %b\n", ip)
				}
			}
		}()

		mdns.Lookup(serviceName, entriesCh)
		close(entriesCh)

		time.Sleep(5 * time.Second)
	}

}
