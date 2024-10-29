//go:build !linux

package ipv6flags

import "errors"

func GetAddrs() ([]IPv6, error) {
	return nil, errors.New("Unsupported platform")
}
