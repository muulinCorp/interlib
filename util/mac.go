package util

import (
	"errors"
	"fmt"
	"net"
)

func GetMacAddrs() (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("fail to get net interfaces: %v", err)
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		return macAddr, nil
	}
	return "", errors.New("not found")
}
