package devicestate

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getSocketByNUMANode(numaNode int64) (int64, error) {
	cpulistPath := fmt.Sprintf("/sys/devices/system/node/node%d/cpulist", numaNode)
	data, err := os.ReadFile(cpulistPath)
	if err != nil {
		return -1, fmt.Errorf("reading cpulist for NUMA %d: %w", numaNode, err)
	}
	cpulist := strings.TrimSpace(string(data))
	if cpulist == "" {
		return -1, fmt.Errorf("empty cpulist for NUMA %d", numaNode)
	}
	firstCPU := cpulist
	if idx := strings.IndexAny(cpulist, ",-"); idx > 0 {
		firstCPU = cpulist[:idx]
	}
	socketPath := fmt.Sprintf("/sys/devices/system/cpu/cpu%s/topology/physical_package_id", firstCPU)
	socketData, err := os.ReadFile(socketPath)
	if err != nil {
		return -1, fmt.Errorf("reading socket for CPU %s: %w", firstCPU, err)
	}
	socket, err := strconv.ParseInt(strings.TrimSpace(string(socketData)), 10, 64)
	if err != nil {
		return -1, fmt.Errorf("parsing socket ID: %w", err)
	}
	return socket, nil
}
