package main

import (
	"awesomeProject/logger"
	"fmt"
	"net"
)

func ListLocalHostAddrs() (*[]net.Addr, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		logger.DefaultLogger.Warnf("net.Interfaces failed, err:", err.Error())
		return nil, err
	}
	var allAddrs []net.Addr
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) == 0 {
			continue
		}
		addrs, err := netInterfaces[i].Addrs()
		if err != nil {
			logger.DefaultLogger.Warnf("failed to get Addrs, %s", err.Error())
		}
		for j := 0; j < len(addrs); j++ {
			allAddrs = append(allAddrs, addrs[j])
		}
	}
	return &allAddrs, nil
}

func IsLocalIP(ip string, addrs *[]net.Addr) bool {
	if defaultIP, _, err := net.SplitHostPort(ip); err == nil {
		ip = defaultIP
	}
	for _, address := range *addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil && ipnet.IP.String() == ip {
			return true
		}
	}
	return false
}
func LocalIP(addrs *[]net.Addr) string {
	for _, address := range *addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}
	return ""
}
func main() {
	addr, err := ListLocalHostAddrs()
	if err != nil {
		logger.DefaultLogger.Errorf("addr is %s", err)
	}
	ip := IsLocalIP("127.0.0.1", addr)
	fmt.Println(ip)
}
