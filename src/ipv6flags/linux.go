//go:build linux

package ipv6flags

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Scope constants (linux/include/net/ipv6.h)
const (
	IPV6_ADDR_SCOPE_NODELOCAL = 0x10
	IPV6_ADDR_SCOPE_LINKLOCAL = 0x20
	IPV6_ADDR_SCOPE_SITELOCAL = 0x40
	IPV6_ADDR_SCOPE_COMPAT_V4 = 0x80
	IPV6_ADDR_SCOPE_GLOBAL    = 0x00
)

// Flags constants (linux/include/uapi/linux/if_addr.h)
const (
	IFA_F_SECONDARY      = 0x01
	IFA_F_TEMPORARY      = IFA_F_SECONDARY
	IFA_F_NODAD          = 0x02
	IFA_F_OPTIMISTIC     = 0x04
	IFA_F_DADFAILED      = 0x08
	IFA_F_HOMEADDRESS    = 0x10
	IFA_F_DEPRECATED     = 0x20
	IFA_F_TENTATIVE      = 0x40
	IFA_F_PERMANENT      = 0x80
	IFA_F_MANAGETEMPADDR = 0x100
	IFA_F_NOPREFIXROUTE  = 0x200
	IFA_F_MCAUTOJOIN     = 0x400
	IFA_F_STABLE_PRIVACY = 0x800
)

var flag_map = map[uint64]string{
	IFA_F_SECONDARY:      "Secondary/Temporary",
	IFA_F_NODAD:          "NoDAD",
	IFA_F_OPTIMISTIC:     "Optimistic",
	IFA_F_DADFAILED:      "DADFailed",
	IFA_F_HOMEADDRESS:    "HomeAddress",
	IFA_F_DEPRECATED:     "Deprecated",
	IFA_F_TENTATIVE:      "Tentative",
	IFA_F_PERMANENT:      "Permanent",
	IFA_F_MANAGETEMPADDR: "ManageTempAddr",
	IFA_F_NOPREFIXROUTE:  "NoPrefixRoute",
	IFA_F_MCAUTOJOIN:     "MCAutoJoin",
	IFA_F_STABLE_PRIVACY: "StablePrivacy",
}

// Read the system's IPv6 addresses including their flags
func GetAddrs() ([]IPv6, error) {
	file, err := os.Open("/proc/net/if_inet6")
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return nil, err
	}
	defer file.Close()

	// Read the file and print line by line translating the contents
	reader := bufio.NewScanner(file)
	var addrs []IPv6
	for reader.Scan() {
		var addr IPv6
		line := reader.Text()
		line_split := strings.Fields(line)

		// Parse IPv6 address
		addr.address, err = hex.DecodeString(line_split[0])
		if err != nil {
			err = fmt.Errorf("error parsing address: %v", err)
			return nil, err
		}

		// Parse Netlink device number hex to int
		fmt.Sscanf(line_split[1], "%x", &addr.netlink)

		// Parse Prefix length hex to int
		fmt.Sscanf(line_split[2], "%x", &addr.prefix)

		// Parse Scope
		scope, err := strconv.ParseUint(line_split[3], 16, 64)
		if err != nil {
			err = fmt.Errorf("error parsing scope: %v", err)
			return nil, err
		}
		switch scope {
		case IPV6_ADDR_SCOPE_NODELOCAL:
			addr.scope = "NodeLocal"
		case IPV6_ADDR_SCOPE_LINKLOCAL:
			addr.scope = "LinkLocal"
		case IPV6_ADDR_SCOPE_SITELOCAL:
			addr.scope = "SiteLocal"
		case IPV6_ADDR_SCOPE_COMPAT_V4:
			addr.scope = "CompatV4"
		case IPV6_ADDR_SCOPE_GLOBAL:
			addr.scope = "Global"
		default:
			addr.scope = "Unknown"
		}

		// Parse Flags
		flags, err := strconv.ParseUint(line_split[4], 16, 64)
		if err != nil {
			err = fmt.Errorf("error parsing flags: %v", err)
			return nil, err
		}
		// Split flags
		flag_l1 := flags & 0xF
		flag_l2 := flags & 0xF0
		flag_l3 := flags & 0xF00

		if flag_l1 != 0 {
			addr.flags = append(addr.flags, flag_map[flag_l1])
		}
		if flag_l2 != 0 {
			addr.flags = append(addr.flags, flag_map[flag_l2])
		}
		if flag_l3 != 0 {
			addr.flags = append(addr.flags, flag_map[flag_l3])
		}

		// Parse device name
		addr.dev = line_split[5]

		addrs = append(addrs, addr)
	}

	return addrs, nil
}
