package collectors

import (
	"os"
	"testing"
)

func writeTempInfo(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp("", "sys-class-thermal-thermal_zone0-temp-*")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Remove(f.Name()) })
	if _, err := f.WriteString(content); err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	return f.Name()
}

func TestTempName(t *testing.T) {
	tp := &Temperature{}
	if tp.Name() != "temp_f" {
		t.Errorf("expected name 'temp_f', got %s", tp.Name())
	}
}

func TestTempCollect_ZeroVal(t *testing.T) {
	// Celsius=0 → fahrenheight = 32.0
	content := "00000\n"
	path := writeTempInfo(t, content)
	tp := &Temperature{fileName: path}

	val, err := tp.Collect()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 32.0 {
		t.Errorf("expected 32.0, got %.2f", val)
	}
}

func TestTempCollect_PositiveVal(t *testing.T) {
	// Celsius=50.634 → fahrenheight = 123.14
	content := "50634\n"
	path := writeTempInfo(t, content)
	tp := &Temperature{fileName: path}

	val, err := tp.Collect()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val < 123.14 || val > 123.15 {
		t.Errorf("expected ~123.14, got %.4f", val)
	}
}

func TestTempCollect_NegativeVal(t *testing.T) {
	// Celsius=-6.666 → fahrenheight = 20.0
	content := "-6666\n"
	path := writeTempInfo(t, content)
	tp := &Temperature{fileName: path}

	val, err := tp.Collect()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val < 20.0 || val > 20.01 {
		t.Errorf("expected ~20.0, got %.4f", val)
	}
}

func TestTempCollect_EmptyFile(t *testing.T) {
	path := writeTempInfo(t, "")
	tp := &Temperature{fileName: path}

	_, err := tp.Collect()
	if err == nil {
		t.Error("expected error for empty file, got nil")
	}
}