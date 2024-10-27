package ipv6flags

import (
	"fmt"
	"net"
)

// IPv6 struct
type IPv6 struct {
	address net.IP   // IPv6 address
	netlink int      // Netlink device number
	prefix  int      // Prefix length
	scope   string   // Scope of the address
	flags   []string // Interface flags
	dev     string   // Device name
}

// Print IPv6 struct
func (addr *IPv6) String() string { // Improve
	return fmt.Sprintf("-----\nAddress: %s\nNetlink: %d\nPrefix: %d\nScope: %s\nFlags: %v\nDevice: %s\n", addr.address, addr.netlink, addr.prefix, addr.scope, addr.flags, addr.dev)
}
