package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Stats struct {
	ID   string
	Name string
	CPU  string
	MP   string // Memory Usage / Limit
}

func GetImageStats() ([]Stats, error) {
	var statsA []Stats

	file, err := os.Open("image_stats.txt")
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) >= 4 {
			stats := Stats{
				ID:   fields[0],
				Name: fields[1],
				CPU:  fields[2],
				MP:   fields[3],
			}
			statsA = append(statsA, stats)
			fmt.Printf("CPU: %v, MP: %v", stats.CPU, stats.MP)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return statsA, nil
}
