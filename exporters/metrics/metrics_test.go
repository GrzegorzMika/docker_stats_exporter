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
