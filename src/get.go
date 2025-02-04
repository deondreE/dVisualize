package main

import (
	"log"
	"os/exec"
	"runtime"
	"time"
)

func GetImageData() {
	var script string
	if runtime.GOOS == "windows" {
		script = ".\\scripts\\windows\\image_info.bat"
	} else {
		script = "./scripts/image_info.sh"
	}

	for {
		cmd := exec.Command(script)
		err := cmd.Run()
		if err != nil {
			log.Printf("Error executing script: %v", err)
		}

		time.Sleep(10 * time.Minute)
	}
}

func GetImageStatsFile() ([]Stats, error) {
	var script string
	var stats []Stats

	if runtime.GOOS == "windows" {
		script = ".\\scripts\\windows\\image_stats.bat"
	} else {
		script = "./scripts/image_stats.sh"
	}

	statsCh := make(chan []Stats)
	errCh := make(chan error)

	// Start the process of running the script in a separate goroutine
	go func() {
		for {
			cmd := exec.Command(script)
			err := cmd.Run()
			if err != nil {
				errCh <- err
				return
			}

			stats, err := GetImageStats() // Assuming GetImageStats() works and returns the Stats
			if err != nil {
				errCh <- err
				return
			}

			statsCh <- stats

			time.Sleep(4 * time.Second)
		}
	}()

	// Return first stats or error
	select {
	case stats = <-statsCh:
		return stats, nil
	case err := <-errCh:
		log.Printf("Error executing script or getting stats: %v", err)
		return nil, err
	}
}

func GetContainerInfo() {
	var script string
	if runtime.GOOS == "windows" {
		script = ".\\scripts\\windows\\container_info.bat"
	} else {
		script = ".//scripts//container_info.sh"
	}

	for {
		cmd := exec.Command(script)
		err := cmd.Run()
		if err != nil {
			log.Printf("Error executing script: %v", err)
		}
		time.Sleep(10 * time.Second)
	}
}
