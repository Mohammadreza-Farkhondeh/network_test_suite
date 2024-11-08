package tests

import (
	"context"
	"fmt"
	"net"
	"time"
)

type DNSLookupResult struct {
	Domain     string
	ResolvedIP string
	LookupTime time.Duration
	TestTime   time.Time
}

// RunDNSLookup performs a DNS lookup on the given domain using the system default DNS server.
func RunDNSLookup(domain string) (DNSLookupResult, error) {
	start := time.Now()
	ips, err := net.LookupIP(domain)
	if err != nil {
		return DNSLookupResult{}, fmt.Errorf("DNS lookup failed: %w", err)
	}
	elapsed := time.Since(start)

	return DNSLookupResult{
		Domain:     domain,
		ResolvedIP: ips[0].String(),
		LookupTime: elapsed,
		TestTime:   time.Now(),
	}, nil
}

func RunDNSLookupWithServer(domain, dnsServer string) (DNSLookupResult, error) {
	start := time.Now()

	resolver := net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return net.Dial("udp", dnsServer+":53")
		},
	}

	ips, err := resolver.LookupHost(context.Background(), domain)
	if err != nil {
		return DNSLookupResult{}, fmt.Errorf("DNS lookup with server %s failed: %w", dnsServer, err)
	}
	elapsed := time.Since(start)

	return DNSLookupResult{
		Domain:     domain,
		ResolvedIP: ips[0],
		LookupTime: elapsed,
		TestTime:   time.Now(),
	}, nil
}
