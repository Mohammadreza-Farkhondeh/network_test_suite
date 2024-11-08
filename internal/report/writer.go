package report

import (
	"fmt"
	"os"

	"github.com/Mohammadreza-Farkhondeh/network_test_suite/internal/tests"
)

// WritePingResult writes the ping test results to the report file
func WritePingResult(file *os.File, result tests.PingResult) {
	fmt.Fprintf(file, "PING TEST\n")
	fmt.Fprintf(file, "Host: %s\n", result.Host)
	fmt.Fprintf(file, "Average Latency: %v\n", result.AvgLatency)
	fmt.Fprintf(file, "Packet Loss: %.2f%%\n", result.PacketLoss)
	fmt.Fprintf(file, "Test Time: %v\n\n", result.TestTime)
}

// WriteDNSResult writes the DNS lookup test results to the report file
func WriteDNSResult(file *os.File, result tests.DNSLookupResult) {
	fmt.Fprintf(file, "DNS LOOKUP TEST\n")
	fmt.Fprintf(file, "Domain: %s\n", result.Domain)
	fmt.Fprintf(file, "Resolved IP: %s\n", result.ResolvedIP)
	fmt.Fprintf(file, "Lookup Time: %v\n", result.LookupTime)
	fmt.Fprintf(file, "Test Time: %v\n\n", result.TestTime)
}

// WriteSpeedResult writes the detailed speed test results to the report file.
func WriteSpeedResult(file *os.File, result tests.SpeedTestResult) {
	fmt.Fprintln(file, "SPEED TEST")
	fmt.Fprintf(file, "Latency: %v\n", result.Latency)
	fmt.Fprintf(file, "Download Speed: %.2f MB/s\n", result.DownloadSpeed)
	fmt.Fprintf(file, "Upload Speed: %.2f MB/s\n", result.UploadSpeed)
	fmt.Fprintf(file, "Packet Loss: %.2f%%\n", result.PacketLoss)
	fmt.Fprintf(file, "Server Location: %s\n", result.ServerLocation)
	fmt.Fprintf(file, "Test Time: %v\n\n", result.TestTime)
}

// WriteTracerouteResult writes the traceroute test results to the report file
func WriteTracerouteResult(file *os.File, result tests.TracerouteResult) {
	fmt.Fprintf(file, "TRACEROUTE TEST\n")
	for _, hop := range result.Hops {
		fmt.Fprintf(file, "%s\n", hop)
	}
	fmt.Fprintf(file, "Test Time: %v\n\n", result.TestTime)
}

// WriteWebsiteAccessibilityResult writes the accessibility test result to the report file.
func WriteWebsiteAccessibilityResult(file *os.File, result tests.WebsiteAccessibilityResult) {
	fmt.Fprintf(file, "WEBSITE ACCESSIBILITY TEST\n")
	fmt.Fprintf(file, "Website: %s\n", result.Website)
	if result.Accessible {
		fmt.Fprintf(file, "Status: Accessible\n")
	} else {
		fmt.Fprintf(file, "Status: Not Accessible\n")
	}
	fmt.Fprintf(file, "Status Code: %d\n", result.StatusCode)
	fmt.Fprintf(file, "Test Time: %v\n\n", result.TestTime)
}
