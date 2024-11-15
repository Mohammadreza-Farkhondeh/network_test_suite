package tests

import (
	"fmt"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"net"
)

type PingResult struct {
	Host       string
	AvgLatency time.Duration
	PacketLoss float64
	TestTime   time.Time
}

// RunPing performs a ping test to the given host.
func RunPing(host string) (PingResult, error) {
	const (
		count       = 4
		timeout     = 2 * time.Second
		packetCount = 4
	)

	ipAddr, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		return PingResult{}, fmt.Errorf("failed to resolve host: %w", err)
	}

	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return PingResult{}, fmt.Errorf("failed to listen for ICMP: %w", err)
	}
	defer c.Close()

	var totalLatency time.Duration
	var received int
	for i := 0; i < count; i++ {
		msg := icmp.Message{
			Type: ipv4.ICMPTypeEcho, Code: 0,
			Body: &icmp.Echo{ID: i + 1, Seq: i, Data: []byte("PING")},
		}
		msgBytes, err := msg.Marshal(nil)
		if err != nil {
			return PingResult{}, fmt.Errorf("failed to marshal ICMP message: %w", err)
		}

		start := time.Now()
		_, err = c.WriteTo(msgBytes, ipAddr)
		if err != nil {
			return PingResult{}, fmt.Errorf("failed to send ICMP request: %w", err)
		}

		reply := make([]byte, 1500)
		c.SetReadDeadline(time.Now().Add(timeout))
		n, _, err := c.ReadFrom(reply)
		if err != nil {
			continue // Considering packet lost
		}

		rtt := time.Since(start)
		totalLatency += rtt
		received++

		// Checking if we received an echo reply
		rm, err := icmp.ParseMessage(1, reply[:n])
		if err == nil && rm.Type == ipv4.ICMPTypeEchoReply {
			fmt.Printf("Received reply from %v: %v\n", ipAddr, rtt)
		}
	}

	var avgLatency time.Duration
	if received == 0 {
		return PingResult{}, fmt.Errorf("no packets received")
	}
	avgLatency = totalLatency / time.Duration(received)

	packetLoss := float64(count-received) / float64(count) * 100

	return PingResult{
		Host:       host,
		AvgLatency: avgLatency,
		PacketLoss: packetLoss,
		TestTime:   time.Now(),
	}, nil
}
