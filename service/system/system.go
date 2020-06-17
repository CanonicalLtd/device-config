/*
 * Copyright (C) 2020 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package system

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"log"
	"os"
)

const (
	snapData = "SNAP_DATA"
)

// Service interface for system resources
type Service interface {
	CPU() (float64, error)
	Memory() (float64, error)
	Disk() (float64, error)
}

// System implements a system service
type System struct{}

// NewSystem creates a new system service
func NewSystem() *System {
	return &System{}
}

// CPU returns the current CPU usage
func (sys *System) CPU() (float64, error) {
	vv, err := cpu.Percent(0, false)
	if err != nil {
		log.Printf("Error getting cpu usage: %v\n", err)
		return 0, err
	}

	var total float64
	if len(vv) > 0 {
		total = vv[0]
	}

	return total, nil
}

// Memory returns the current memory usage
func (sys *System) Memory() (float64, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("Error getting memory usage: %v\n", err)
		return 0, err
	}

	return v.UsedPercent, nil
}

// Disk returns the current disk usage
func (sys *System) Disk() (float64, error) {
	// Check the disk space of the host FS not the snap
	v, err := disk.Usage(os.Getenv(snapData))
	if err != nil {
		log.Printf("Error getting disk usage: %v\n", err)
		return 0, err
	}

	return v.UsedPercent, nil
}
