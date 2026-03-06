package collectors

import (
	"syscall"
)

type Disk struct{
	path string
}

func NewDisk(path string) *Disk {
	return &Disk{path: path}
}
 
func (d *Disk)Name() string {
    return "disk"
}

func (d *Disk)Collect() (float64, error) {
	// Call Statfs
	var stat syscall.Statfs_t
	err := syscall.Statfs(d.path, &stat)
	if err != nil {
		return 0.0, err
	}

	usagePercent := (1.0 - float64(stat.Bavail) / float64(stat.Blocks)) * 100.0
	return usagePercent, nil
}