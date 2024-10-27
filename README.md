# IPv6 flags
Go library to determine if an interface's IPv6 address is temporary or deprecated. This library is meant to be a temporary fix until [go/issues/42694](https://github.com/golang/go/issues/42694) is implemented.

Use the function `func GetAddrs() ([]IPv6, error)` to get a list of all the IPv6 addresses on the system along with their flags.

```go
type IPv6 struct {
	address net.IP   // IPv6 address
	netlink int      // Netlink device number
	prefix  int      // Prefix length
	scope   string   // Scope of the address
	flags   []string // Interface flags
	dev     string   // Device name
}
```

## Example
```go
package main

import (
	"fmt"

	ipv6flags "github.com/RLado/ipv6flags"
)

func main() {
	addrs, err := ipv6flags.GetAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		fmt.Println(addr.String())
	}
}
```

**Result:**
```
-----
Address: fe80::beed:3b05:9bf3:9975
Netlink: 3
Prefix: 64
Scope: LinkLocal
Flags: [Permanent]
Device: wlp15s0

-----
Address: ::1
Netlink: 1
Prefix: 128
Scope: NodeLocal
Flags: [Permanent]
Device: lo
```

## Sources:
#### Linux
- https://kernel.org/
- https://mirrors.deepspace6.net/Linux+IPv6-HOWTO/proc-net.html

#### Windows
> Under development

#### MacOS
> Under development
