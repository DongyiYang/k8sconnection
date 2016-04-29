package conntrack

const (
	// #defined in libnfnetlink/include/libnfnetlink/linux_nfnetlink.h
	NFNL_SUBSYS_CTNETLINK = 1
	NFNETLINK_V0          = 0

	// #defined in libnfnetlink/include/libnfnetlink/linux_nfnetlink_compat.h
	NF_NETLINK_CONNTRACK_NEW     = 0x00000001
	NF_NETLINK_CONNTRACK_UPDATE  = 0x00000002
	NF_NETLINK_CONNTRACK_DESTROY = 0x00000004

	// #defined in libnfnetlink/include/libnfnetlink/libnfnetlink.h
	NLA_F_NESTED        = uint16(1 << 15)
	NLA_F_NET_BYTEORDER = uint16(1 << 14)
	NLA_TYPE_MASK       = ^(NLA_F_NESTED | NLA_F_NET_BYTEORDER)
)

type NfConntrackMsg int

const (
	NfctMsgUnknown NfConntrackMsg = 0
	NfctMsgNew     NfConntrackMsg = 1 << 0
	NfctMsgUpdate  NfConntrackMsg = 1 << 1
	NfctMsgDestroy NfConntrackMsg = 1 << 2
)

// taken from libnetfilter_conntrack/src/conntrack/snprintf.c
var tcpState = []string{
	"NONE",
	"SYN_SENT",
	"SYN_RECV",
	"ESTABLISHED",
	"FIN_WAIT",
	"CLOSE_WAIT",
	"LAST_ACK",
	"TIME_WAIT",
	"CLOSE",
	"LISTEN",
	"MAX",
	"IGNORE",
}
