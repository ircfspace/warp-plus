package statute

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"net/netip"
	"time"
)

type TIPQueueChangeCallback func(ips []IPInfo)

type (
	TDialerFunc     func(ctx context.Context, network, addr string) (net.Conn, error)
	THTTPClientFunc func(rawDialer TDialerFunc, tlsDialer TDialerFunc, targetAddr ...string) *http.Client
)

var (
	HTTPPing = 1 << 1
	TLSPing  = 1 << 2
	TCPPing  = 1 << 3
	WARPPing = 1 << 5
)

type IPInfo struct {
	AddrPort  netip.AddrPort
	RTT       time.Duration
	CreatedAt time.Time
}

type ScannerOptions struct {
	UseIPv4               bool
	UseIPv6               bool
	CidrList              []netip.Prefix // CIDR ranges to scan
	SelectedOps           int
	Logger                *slog.Logger
	InsecureSkipVerify    bool
	RawDialerFunc         TDialerFunc
	TLSDialerFunc         TDialerFunc
	HttpClientFunc        THTTPClientFunc
	UseHTTP2              bool
	DisableCompression    bool
	HTTPPath              string
	Referrer              string
	UserAgent             string
	Hostname              string
	WarpPrivateKey        string
	WarpPeerPublicKey     string
	WarpPresharedKey      string
	Port                  uint16
	IPQueueSize           int
	IPQueueTTL            time.Duration
	MaxDesirableRTT       time.Duration
	IPQueueChangeCallback TIPQueueChangeCallback
	ConnectionTimeout     time.Duration
	HandshakeTimeout      time.Duration
	TlsVersion            uint16
}
