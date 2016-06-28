package conntrack

import (
	"fmt"
	"net"
	"strconv"
)

type ConntrackInfo struct {
	MsgType        NfConntrackEventType
	Proto          int
	Src            net.IP
	SrcPort        uint16
	Dst            net.IP
	DstPort        uint16
	Packets        uint64
	Bytes          uint64
	StartTimestamp uint64
	DeltaTime      uint64
	TCPState       string
}

func (c ConntrackInfo) String() string {
	return fmt.Sprintf("%s:%d->%s:%d, packets=%d, bytes=%d, start_time=%d, delta_time=%d",
		c.Src, c.SrcPort, c.Dst, c.DstPort, c.Packets, c.Bytes, c.StartTimestamp, c.DeltaTime)
}

// ConnTCP decides which way this connection is going and makes a ConnTCP.
func (c ConntrackInfo) BuildTCPConn(addressSet map[string]struct{}) []*ConnTCP {
	var res []*ConnTCP
	// conntrack gives us all connections, even things passing through. But here we only
	// care connection those are sourced from or destinated to address defined in addressSet
	src := c.Src.String()
	dst := c.Dst.String()
	_, srcLocal := addressSet[src]
	_, dstLocal := addressSet[dst]
	if srcLocal {
		srcConn := &ConnTCP{
			Local:      src,
			LocalPort:  strconv.Itoa(int(c.SrcPort)),
			Remote:     dst,
			RemotePort: strconv.Itoa(int(c.DstPort)),
		}
		res = append(res, srcConn)
	}
	if dstLocal {
		dstConn := &ConnTCP{
			Local:      dst,
			LocalPort:  strconv.Itoa(int(c.DstPort)),
			Remote:     src,
			RemotePort: strconv.Itoa(int(c.SrcPort)),
		}
		res = append(res, dstConn)
	}
	// Neither is in addressSet. conntrack also reports NAT connections.
	return res
}
