package domain

import (
	"errors"
	"ipserver-search/internal/app/pkg"
)

func SearchIpByHost(host string) (ips []string, err error) {
	if isInvalidHost(host) {
		err = errors.New("invalid host")
		return
	}

	ips = pkg.SearchIP(host)
	return
}

func SearchServerByHost(host string) (servers []string, err error) {
	if isInvalidHost(host) {
		err = errors.New("invalid host")
		return
	}

	servers = pkg.SearchServer(host)
	return
}

func isInvalidHost(host string) bool {
	return len(host) < 4
}
