package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func GetImageData() {
	// runs the ../scripts/image_info.sh every 10minutes
	for {
		cmd := exec.Command("../scripts/image_info.sh")
		err := cmd.Run()
		if err != nil {
			log.Printf("Error executing script: %v", err)
		} else {
			fmt.Println("Script executed successfully.")
		}

		time.Sleep(10 * time.Minute)
	}
}
