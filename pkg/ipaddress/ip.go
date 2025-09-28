package ipaddress

import (
	"net"
)

func LocalIPAddr() (localAddr string) {
	if addrs, err := net.InterfaceAddrs(); err == nil {
		for _, v := range addrs {
			ip, ok := v.(*net.IPNet)
			if ok && !ip.IP.IsLoopback() && ip.IP.IsPrivate() && ip.IP.To4() != nil {
				localAddr = ip.IP.String()
				break
			}
		}
	}

	if localAddr == "" {
		conn, err := net.Dial("udp", "8.8.8.8:80")
		if err != nil {
			return
		}
		defer conn.Close()
		localAddr = conn.LocalAddr().String()
	}
	return
}

func IPv6() (ipv6Addr string) {
	if addrs, err := net.InterfaceAddrs(); err == nil {
		for _, v := range addrs {
			IPNet, ok := v.(*net.IPNet)
			ip := IPNet.IP
			if ok && !ip.IsLoopback() && ip.To4() == nil && ip.IsGlobalUnicast() &&
				!ip.IsInterfaceLocalMulticast() && !ip.IsLinkLocalUnicast() && !ip.IsUnspecified() {
				ipv6Addr = ip.String()
				break
			}
		}
	}

	if ipv6Addr == "" {
		conn, err := net.Dial("udp", "[2001:db8::1]:domain")
		if err != nil {
			return
		}
		defer conn.Close()
		localAddr := conn.LocalAddr().(*net.UDPAddr)
		ipv6Addr = localAddr.IP.String()
	}
	return
}
