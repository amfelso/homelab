package collectors

import (
	"os"
	"testing"
)

// writeProcStat creates a temp file with the given content and returns its path.
// t.Cleanup ensures it's deleted when the test finishes.
func writeProcStat(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp("", "proc-stat-*")
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

func TestCpuName(t *testing.T) {
	c := &Cpu{}
	if c.Name() != "cpu" {
		t.Errorf("expected name 'cpu', got %s", c.Name())
	}
}

func TestCpuCollect_FirstCallReturnsZero(t *testing.T) {
	path := writeProcStat(t, "cpu  180 50 90 500 90 40 50 0 0 0\n")
	c := &Cpu{fileName: path}

	val, err := c.Collect()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 0.0 {
		t.Errorf("expected 0.0 on first call, got %.2f", val)
	}
}

func TestCpuCollect_ReturnsUsagePercent(t *testing.T) {
	// Previous snapshot: total=1000, idle=500
	// Current snapshot:  total=1100, idle=540
	// idleDelta=40, totalDelta=100 → usage = (1 - 40/100) * 100 = 60.0
	path := writeProcStat(t, "cpu  200 50 100 540 100 50 60 0 0 0\n")
	c := &Cpu{
		fileName:  path,
		prevIdle:  500,
		prevTotal: 1000,
	}

	val, err := c.Collect()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 60.0 {
		t.Errorf("expected 60.0, got %.2f", val)
	}
}

func TestCpuCollect_UpdatesState(t *testing.T) {
	path := writeProcStat(t, "cpu  200 50 100 540 100 50 60 0 0 0\n")
	c := &Cpu{
		fileName:  path,
		prevIdle:  500,
		prevTotal: 1000,
	}

	c.Collect()

	if c.prevIdle != 540 {
		t.Errorf("expected prevIdle=540, got %d", c.prevIdle)
	}
	if c.prevTotal != 1100 {
		t.Errorf("expected prevTotal=1100, got %d", c.prevTotal)
	}
}
