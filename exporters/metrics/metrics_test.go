package metrics

import (
	"testing"
)

func TestReadTimeMetric(t *testing.T) {
	m := ReadTimeMetric(example_response)
	expected := 1726671800.129723
	if m != expected {
		t.Errorf("Expected %f, got %f", expected, m)
	}
}

func TestCPUUsageTotalMetric(t *testing.T) {
	m := CPUUsageTotalMetric(example_response)
	expected := 20281474000.
	if m != expected {
		t.Errorf("Expected %f, got %f", expected, m)
	}
}

func TestCPUSystemUsageTotalMetric(t *testing.T) {
	m := CPUSystemUsageTotalMetric(example_response)
	expected := 13020180000000.
	if m != expected {
		t.Errorf("Expected %f, got %f", expected, m)
	}
}

func TestCPUUsageDeltaMetric(t *testing.T) {
	m := CPUUsageDeltaMetric(example_response)
	expected := 20847000.
	if m != expected {
		t.Errorf("Expected %f, got %f", expected, m)
	}
}

func TestCPUSystemUsageDeltaMetric(t *testing.T) {
	m := CPUSystemUsageDeltaMetric(example_response)
	expected := 15970000000.
	if m != expected {
		t.Errorf("Expected %f, got %f", expected, m)
	}
}

func TestCPUNumberMetric(t *testing.T) {
	m := CPUNumberMetric(example_response)
	expected := 16.
	if m != expected {
		t.Errorf("Expected %f, got %f", expected, m)
	}
}

func TestMemoryUsageMetric(t *testing.T) {
	m := MemoryUsageMetric(example_response)
	expected := 1975832576.
	if m != expected {
		t.Errorf("Expected %f, got %f", expected, m)
	}
}

func TestMemoryCachedUsageMetric(t *testing.T) {
	m := MemoryCachedUsageMetric(example_response)
	expected := 0.
	if m != expected {
		t.Errorf("Expected %f, got %f", expected, m)
	}
}

func TestMemoryLimitMetric(t *testing.T) {
	m := MemoryLimitMetric(example_response)
	expected := 33013448704.
	if m != expected {
		t.Errorf("Expected %f, got %f", expected, m)
	}
}

func TestNetworkBytesReceivedMetric(t *testing.T) {
	m := NetworkBytesReceivedMetric(example_response)
	expected := 88452.
	if m != expected {
		t.Errorf("Expected %f, got %f", expected, m)
	}
}

func TestNetworkBytesSentMetric(t *testing.T) {
	m := NetworkBytesSentMetric(example_response)
	expected := 1786.
	if m != expected {
		t.Errorf("Expected %f, got %f", expected, m)
	}
}

func TestNetworkPacketsReceivedMetric(t *testing.T) {
	m := NetworkPacketsReceivedMetric(example_response)
	expected := 1944.
	if m != expected {
		t.Errorf("Expected %f, got %f", expected, m)
	}
}

func TestNetworkPacketsSentMetric(t *testing.T) {
	m := NetworkPacketsSentMetric(example_response)
	expected := 25.
	if m != expected {
		t.Errorf("Expected %f, got %f", expected, m)
	}
}

func TestNetworkErrorsReceivedMetric(t *testing.T) {
	m := NetworkErrorsReceivedMetric(example_response)
	expected := 0.
	if m != expected {
		t.Errorf("Expected %f, got %f", expected, m)
	}
}

func TestNetworkErrorsSentMetric(t *testing.T) {
	m := NetworkErrorsSentMetric(example_response)
	expected := 0.
	if m != expected {
		t.Errorf("Expected %f, got %f", expected, m)
	}
}
