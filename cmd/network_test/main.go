package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Mohammadreza-Farkhondeh/network_test_suite/internal/report"
	"github.com/Mohammadreza-Farkhondeh/network_test_suite/internal/tests"
)

func main() {
	dnsServers := []string{"1.1.1.1", "8.8.8.8", "4.2.2.4", "172.26.146.34", "172.26.146.35"}
	websites := []string{"http://net2.sharif.edu", "http://edu.sharif.edu", "http://cw.sharif.edu", "http://sharif.edu"}
	ips := []string{"172.17.1.214", "172.26.146.34", "172.26.146.35", "4.2.2.4", "31.13.64.1", "185.112.250.1"}

	reportFileName := fmt.Sprintf("network_report_%s.txt", time.Now().Format("20060102_150405"))
	reportFile, err := os.Create(reportFileName)
	if err != nil {
		log.Fatalf("failed to create report file: %v", err)
	}
	defer reportFile.Close()

	fmt.Fprintln(reportFile, "NETWORK DIAGNOSTIC REPORT")
	fmt.Fprintf(reportFile, "Report generated on: %s\n\n", time.Now().Format("2006-01-02 15:04:05"))

	// Ping Test
	fmt.Fprintln(reportFile, "PING TESTS")
	for _, ip := range ips {
		fmt.Fprintf(reportFile, "\nPinging IP: %s\n", ip)
		pingResult, err := tests.RunPing(ip)
		if err != nil {
			fmt.Fprintf(reportFile, "Ping to %s failed: %v\n", ip, err)
		} else {
			report.WritePingResult(reportFile, pingResult)
		}
	}

	// DNS Lookup Test
	fmt.Fprintln(reportFile, "\nDNS TESTS")
	for _, dns := range dnsServers {
		fmt.Fprintf(reportFile, "\nTesting DNS Server: %s\n", dns)
		dnsResult, err := tests.RunDNSLookupWithServer("google.com", dns)
		if err != nil {
			fmt.Fprintf(reportFile, "Failed to resolve using DNS %s: %v\n", dns, err)
		} else {
			report.WriteDNSResult(reportFile, dnsResult)
		}
	}

	// Website Accessibility Test
	fmt.Fprintln(reportFile, "\nWEBSITE ACCESSIBILITY TESTS")
	for _, website := range websites {
		fmt.Fprintf(reportFile, "\nChecking website: %s\n", website)
		accessibilityResult, err := tests.CheckWebsiteAccessibility(website)
		if err != nil {
			fmt.Fprintf(reportFile, "Accessibility check for %s failed: %v\n", website, err)
		} else {
			report.WriteWebsiteAccessibilityResult(reportFile, accessibilityResult)
		}
	}

	// Traceroute Test
	fmt.Fprintln(reportFile, "\nTRACEROUTE TEST")
	tracerouteResult, err := tests.RunTraceroute("8.8.8.8")
	if err != nil {
		fmt.Fprintf(reportFile, "Traceroute failed: %v\n", err)
	} else {
		report.WriteTracerouteResult(reportFile, tracerouteResult)
	}

	// Speed Test
	fmt.Fprintln(reportFile, "\nSPEED TEST")
	speedResult, err := tests.RunSpeedTest()
	if err != nil {
		fmt.Fprintf(reportFile, "Speed test failed: %v\n", err)
	} else {
		report.WriteSpeedResult(reportFile, speedResult)
	}

	fmt.Printf("Network report saved to %s\n", reportFileName)
}

