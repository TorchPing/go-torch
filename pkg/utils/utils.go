package utils

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func TimeIt(f func() interface{}) (int64, interface{}) {
	startAt := time.Now()
	res := f()
	endAt := time.Now()
	return endAt.UnixNano() - startAt.UnixNano(), res
}

func TimeItWithResult(f func() (interface{}, interface{})) (int64, interface{}, interface{}) {
	startAt := time.Now()
	res, err := f()
	endAt := time.Now()
	return endAt.UnixNano() - startAt.UnixNano(), res, err
}

// FormatIP - trim spaces and format IP
//
// IP - the provided IP
//
// string - return "" if the input is neither valid IPv4 nor valid IPv6
//          return IPv4 in format like "192.168.9.1"
//          return IPv6 in format like "[2002:ac1f:91c5:1::bd59]"
func FormatIP(IP string) string {

	host := strings.Trim(IP, "[ ]")
	if parseIP := net.ParseIP(host); parseIP != nil && parseIP.To4() == nil {
		host = fmt.Sprintf("[%s]", host)
	}

	return host
}
