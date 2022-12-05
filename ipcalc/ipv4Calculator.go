package ipcalc

import (
	"errors"
	"fmt"
	"ip-calculator/structs"
	"net"
	"strconv"
	"strings"
)

func FindIpv4Results(addr string, netmask string) (structs.Ipv4Result,error) {
	var result structs.Ipv4Result

	//Create an actual ip to work with

	ip := net.ParseIP(addr)
	var mask net.IPMask
	broadcast := net.IP(make([]byte, 4))
	netAddr := net.IP(make([]byte, 4))
	var startHost net.IP
	var endHost net.IP
	var totalAddresses int 
	var usableAdresses int
	var maskOnes int 
	var err error

	if ip == nil {
		return result, errors.New("Failed to get results")
	}
	
	//Build the netmask
	mask, err = ipv4MaskFromNonCIDR(netmask)

	if err != nil {
		return result, err
	}

	//Find the broadcast and net address
	for i := range ip {
		broadcast[i] = ip[i] | ^ mask[i]
		netAddr[i] = ip[i] & mask[i]
	}

	//Find the start and end Host
	startHost = nextIP(netAddr,1)
	endHost = previousIP(broadcast, 1)

	//Find the total and usable ip addresses
	maskOnes, _ = mask.Size()
	totalAddresses = 2^(32-maskOnes)
	usableAdresses = totalAddresses-2

	//TODO get IP type

	//Build the result struct
	result = structs.Ipv4Result{
		Addr: addr,
		Netmask: mask.String(),
		NetworkAddress: netAddr.String(),
		BroadcastAddress: broadcast.String(),
		UsableRange: fmt.Sprint(startHost.String()," - ",endHost.String()),
		TotalIPAddresses: totalAddresses,
		UsableIPAddresses: usableAdresses,
		IPType: "TODO",
	}

	return result, nil
}

func ipv4MaskFromNonCIDR (netmask string) (net.IPMask, error) {
	var mask net.IPMask
	//Split for each period
	maskParts := strings.Split(netmask,".")
	var decimalMask []int

	//Convert to integer
	for i := range maskParts {
		value, err := strconv.Atoi(maskParts[i])

		if err != nil {
			return nil, errors.New("Could not parse the netmask")
		}
		decimalMask = append(decimalMask, value)
	}
	
	//Build a mask 
	mask = net.IPv4Mask(byte(decimalMask[0]),
						byte(decimalMask[1]),
						byte(decimalMask[2]),
						byte(decimalMask[3]))

	return mask, nil
}

//https://gist.github.com/udhos/b468fbfd376aa0b655b6b0c539a88c03
func nextIP(ip net.IP, inc uint) net.IP {
	i := ip.To4()
	v := uint(i[0])<<24 + uint(i[1])<<16 + uint(i[2])<<8 + uint(i[3])
	v += inc
	v3 := byte(v & 0xFF)
	v2 := byte((v >> 8) & 0xFF)
	v1 := byte((v >> 16) & 0xFF)
	v0 := byte((v >> 24) & 0xFF)
	return net.IPv4(v0, v1, v2, v3)
}

func previousIP(ip net.IP, dec uint) net.IP {
	i := ip.To4()
	v := uint(i[0])<<24 + uint(i[1])<<16 + uint(i[2])<<8 + uint(i[3])
	v -= dec
	v3 := byte(v & 0xFF)
	v2 := byte((v >> 8) & 0xFF)
	v1 := byte((v >> 16) & 0xFF)
	v0 := byte((v >> 24) & 0xFF)
	return net.IPv4(v0, v1, v2, v3)
}
