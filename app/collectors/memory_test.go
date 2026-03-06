package collectors

import (
	"os"
	"testing"
)

func writeMemInfo(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp("", "proc-meminfo-*")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Remove(f.Name()) })
	if _, err := f.WriteString(content); err != nil {
		t.Fatal(err)
	}
	f.Close()
	return f.Name()
}

func TestMemoryName(t *testing.T) {
	m := &Memory{}
	if m.Name() != "memory" {
		t.Errorf("expected name 'memory', got %s", m.Name())
	}
}

func TestMemoryCollect_ReturnsUsagePercent(t *testing.T) {
	// MemTotal=1000, MemAvailable=600 → usage = (1000-600)/1000 * 100 = 40.0
	content := "MemTotal:       1000 kB\nMemFree:         400 kB\nMemAvailable:    600 kB\n"
	path := writeMemInfo(t, content)
	m := &Memory{fileName: path}

	val, err := m.Collect()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 40.0 {
		t.Errorf("expected 40.0, got %.2f", val)
	}
}

func TestMemoryCollect_FullMemoryUsed(t *testing.T) {
	// MemAvailable=0 → usage = 100.0
	content := "MemTotal:       1000 kB\nMemFree:           0 kB\nMemAvailable:      0 kB\n"
	path := writeMemInfo(t, content)
	m := &Memory{fileName: path}

	val, err := m.Collect()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 100.0 {
		t.Errorf("expected 100.0, got %.2f", val)
	}
}

func TestMemoryCollect_EmptyFile(t *testing.T) {
	path := writeMemInfo(t, "")
	m := &Memory{fileName: path}

	_, err := m.Collect()
	if err == nil {
		t.Error("expected error for empty file, got nil")
	}
}
