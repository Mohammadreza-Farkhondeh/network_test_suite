package tests

import (
	"time"
)

type NetworkTestResult struct {
	Timestamp     time.Time
	Latency       float64
	PacketLoss    float64
	DownloadSpeed float64
	UploadSpeed   float64
}
