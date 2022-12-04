package ipcalc

import (
	"errors"
	"ip-calculator/regex"
)

const (

	IPv4 = 0
	IPv6 = 1
	IPv4NetMask = 2
	IPv6NetMask = 3
)

func ValidateIP(ip string) (int, error) {
	//TODO check ipv6
	if regex.CheckIpv4(ip) {
		return IPv4, nil
	}

	return -1, errors.New("Unrecognized IP address")
}

func ValidateNetMask(mask string) (int, error){
	//TODO check ipv6
	if regex.CheckIpv4Mask(mask){
		return IPv4NetMask, nil
	}

	return -1, errors.New("Unrecognized netmask")
}
