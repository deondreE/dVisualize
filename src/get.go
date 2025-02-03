package main

import (
	"fmt"
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
		} else {
			fmt.Println("Script executed successfully.")
		}

		time.Sleep(10 * time.Minute)
	}
}
