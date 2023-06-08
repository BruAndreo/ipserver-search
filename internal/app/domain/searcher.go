package domain

import (
	"errors"
	"ipserver-search/internal/app"
)

func SearchIpByHost(host string) (err error, ips []string) {
	if isInvalidHost(host) {
		err = errors.New("Invalid host")
		return
	}

	ips = app.SearchIP(host)
	return
}

func SearchServerByHost(host string) (err error, servers []string) {
	if isInvalidHost(host) {
		err = errors.New("Invalid host")
		return
	}

	servers = app.SearchServer(host)
	return
}

func isInvalidHost(host string) bool {
	return len(host) < 4
}
