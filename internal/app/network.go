package app

import (
	"log"
	"net"
)

func SearchIP(host string) (ips []string) {
	ipList, err := net.LookupIP(host)
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}

	for _, ip := range ipList {
		ips = append(ips, ip.String())
	}
	return
}

func SearchServer(host string) (servers []string) {
	serverList, err := net.LookupNS(host)
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}

	for _, server := range serverList {
		servers = append(servers, server.Host)
	}
	return
}
