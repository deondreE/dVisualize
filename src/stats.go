package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Stats struct {
	ID   string
	Name string
	CPU  float64
	MP   float64 // Memory Usage / Limit
}

func parseMemory(memStr string) (float64, error) {
	units := map[string]float64{
		"B":   1,
		"KiB": 1024, "MiB": 1024 * 1024, "GiB": 1024 * 1024 * 1024,
		"KB": 1000, "MB": 1000 * 1000, "GB": 1000 * 1000 * 1000,
	}

	re := regexp.MustCompile(`([\d.]+)([KMGT]i?B)`)
	matches := re.FindStringSubmatch(memStr)
	if len(matches) != 3 {
		return 0, fmt.Errorf("invalid memory format: %s", memStr)
	}

	value, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0, err
	}

	unit, exists := units[matches[2]]
	if !exists {
		return 0, fmt.Errorf("unknown unit: %s", matches[2])
	}

	return value * unit, nil
}

func GetImageStats() ([]Stats, error) {
	var statsA []Stats

	file, err := os.Open("image_stats.txt")
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`^(\S+)\s+(\S+)\s+([\d.]+)%\s+([\d.]+[KMGiB]+) / ([\d.]+[KMGiB]+)`)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "CONTAINER ID") { // Skip header
			continue
		}

		matches := re.FindStringSubmatch(line)
		if len(matches) < 6 {
			continue // Skip invalid lines
		}

		// Parse CPU usage
		cpuUsage, err := strconv.ParseFloat(matches[3], 64)
		if err != nil {
			return nil, fmt.Errorf("error converting CPU usage for container %s (ID: %s): %v", matches[2], matches[1], err)
		}

		// Parse memory usage and limit
		memUsage, err := parseMemory(matches[4])
		if err != nil {
			return nil, fmt.Errorf("error parsing memory usage for container %s: %v", matches[2], err)
		}

		memLimit, err := parseMemory(matches[5])
		if err != nil {
			return nil, fmt.Errorf("error parsing memory limit for container %s: %v", matches[2], err)
		}

		// Compute memory usage ratio
		memRatio := 0.0
		if memLimit > 0 {
			memRatio = memUsage / memLimit
		}

		stats := Stats{
			ID:   matches[1],
			Name: matches[2],
			CPU:  cpuUsage,
			MP:   memRatio,
		}
		statsA = append(statsA, stats)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return statsA, nil
}

func GetCpuValues(stats []Stats) []float64 {
	var cpuValues []float64
	for _, stat := range stats {
		cpuValues = append(cpuValues, stat.CPU)
	}
	return cpuValues
}

func GetMemVals(stats []Stats) []float64 {
	var memValues []float64
	for _, stat := range stats {
		memValues = append(memValues, stat.MP)
	}
	return memValues
}
