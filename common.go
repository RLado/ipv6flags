package ipv6flags

import (
	"fmt"
	"net"
)

// IPv6 struct
type IPv6 struct {
	Address net.IP   // IPv6 address
	Netlink int      // Netlink device number
	Prefix  int      // Prefix length
	Scope   string   // Scope of the address
	Flags   []string // Interface flags
	Dev     string   // Device name
}

// Print IPv6 struct
func (addr *IPv6) String() string { // Improve
	return fmt.Sprintf("-----\nAddress: %s\nNetlink: %d\nPrefix: %d\nScope: %s\nFlags: %v\nDevice: %s\n", addr.Address, addr.Netlink, addr.Prefix, addr.Scope, addr.Flags, addr.Dev)
}
