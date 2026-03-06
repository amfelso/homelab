package collectors

import "testing"

func TestDiskName(t *testing.T) {
	d := &Disk{}
	if d.Name() != "disk" {
		t.Errorf("expected name 'disk', got %s", d.Name())
	}
}

func TestDiskCollect_ReturnsValidPercent(t *testing.T) {
	d := &Disk{path: "/tmp"}
	val, err := d.Collect()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val < 0 || val > 100 {
		t.Errorf("expected value between 0 and 100, got %.2f", val)
	}
}

func TestDiskCollect_InvalidPath(t *testing.T) {
	d := &Disk{path: "/nonexistent/path"}
	_, err := d.Collect()
	if err == nil {
		t.Error("expected error for invalid path, got nil")
	}
}
