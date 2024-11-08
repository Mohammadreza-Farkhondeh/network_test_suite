package tests

import (
	"fmt"
	"github.com/showwin/speedtest-go/speedtest"
	"github.com/showwin/speedtest-go/speedtest/transport"
	"log"
	"time"
)

type SpeedTestResult struct {
	Latency        time.Duration
	DownloadSpeed  float64 // in MB/s
	UploadSpeed    float64 // in MB/s
	PacketLoss     float64 // percentage
	TestTime       time.Time
	ServerLocation string
}

// RunSpeedTest performs a speed test using the speedtest-go library and returns a SpeedTestResult.
func RunSpeedTest() (SpeedTestResult, error) {
	speedtestClient := speedtest.New()

	serverList, err := speedtestClient.FetchServers()
	if err != nil {
		return SpeedTestResult{}, fmt.Errorf("failed to fetch servers: %w", err)
	}

	// Select the nearest server automatically
	targets, err := serverList.FindServer([]int{})
	if err != nil {
		return SpeedTestResult{}, fmt.Errorf("failed to find server: %w", err)
	}

	analyzer := speedtest.NewPacketLossAnalyzer(nil)

	for _, server := range targets {
		server.PingTest(nil)
		server.DownloadTest()
		server.UploadTest()

		packetLossResult := 0.0
		err := analyzer.Run(server.Host, func(pl *transport.PLoss) {
			if pl != nil {
				packetLoss := (1.0 - float64(pl.Sent-pl.Dup)/float64(pl.Max+1)) * 100
				packetLossResult = packetLoss
			}
		})
		if err != nil {
			log.Printf("packet loss test failed for server %s: %v", server.Host, err)
		}

		return SpeedTestResult{
			Latency:        server.Latency,
			DownloadSpeed:  float64(server.DLSpeed) / (1024 * 1024),
			UploadSpeed:    float64(server.ULSpeed) / (1024 * 1024),
			PacketLoss:     packetLossResult,
			TestTime:       time.Now(),
			ServerLocation: server.Name,
		}, nil
	}

	return SpeedTestResult{}, fmt.Errorf("no servers available for speed test")
}

