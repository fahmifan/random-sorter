package utils

import (
	"net"
	"os"
)

// PrivateIP get private ip of current network
// if not connected to network it will return the hostname
func PrivateIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}

			if ip = ip.To4(); ip == nil {
				continue // not an ipv4 address
			}

			return ip.String(), nil
		}
	}

	host, err := os.Hostname()
	if err != nil {
		return "", err
	}

	return host, nil
}
