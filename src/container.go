package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type Container struct {
	Name   string
	ID     string
	STATUS string
	PORTS  string
}

func GetConInfo() []Container {
	var con []Container

	file, err := os.Open("container_info.txt")
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`^(\S+)\s+(\S+)\s+(.+?)\s+(\S+)$`)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "CONTAINER ID") { // Skip header
			continue
		}

		matches := re.FindStringSubmatch(line)
		if len(matches) < 5 {
			continue
		}

		container := Container{
			ID:     matches[1],
			Name:   matches[2],
			STATUS: matches[3],
			PORTS:  matches[4],
		}
		con = append(con, container)
	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	return con
}
