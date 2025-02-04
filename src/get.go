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

func GetImageStatsFile() {
	var script string
	if runtime.GOOS == "windows" {
		script = ".\\scripts\\windows\\image_stats.bat"
	} else {
		script = "./scripts/image_stats.sh"
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
