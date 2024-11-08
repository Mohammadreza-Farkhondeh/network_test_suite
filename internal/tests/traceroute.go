package tests

import (
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"net"
	"time"
)

type TracerouteResult struct {
	Hops     []string
	TestTime time.Time
}

// RunTraceroute performs a basic traceroute to the host.
func RunTraceroute(host string) (TracerouteResult, error) {
	const maxHops = 10
	ipAddr, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		return TracerouteResult{}, fmt.Errorf("failed to resolve host: %w", err)
	}

	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return TracerouteResult{}, fmt.Errorf("failed to listen for ICMP: %w", err)
	}
	defer c.Close()

	hops := []string{}
	for ttl := 1; ttl <= maxHops; ttl++ {
		p := ipv4.NewPacketConn(c)
		p.SetTTL(ttl)

		msg := icmp.Message{
			Type: ipv4.ICMPTypeEcho, Code: 0,
			Body: &icmp.Echo{ID: ttl, Seq: ttl, Data: []byte("traceroute")},
		}

		start := time.Now()
		msgBytes, _ := msg.Marshal(nil)
		c.WriteTo(msgBytes, ipAddr)

		reply := make([]byte, 1500)
		c.SetReadDeadline(time.Now().Add(time.Second))
		n, peer, err := c.ReadFrom(reply)
		if err != nil {
			hops = append(hops, fmt.Sprintf("%d * * *", ttl))
			continue
		}

		rtt := time.Since(start)
		hops = append(hops, fmt.Sprintf("%d %v %v", ttl, peer, rtt))

		parsedMsg, _ := icmp.ParseMessage(1, reply[:n])
		if parsedMsg.Type == ipv4.ICMPTypeEchoReply {
			break
		}
	}

	return TracerouteResult{
		Hops:     hops,
		TestTime: time.Now(),
	}, nil
}

