package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Image struct {
	repo string
	id   string
}

// TODO: Memory usage per container.
// TODO: CPU Usage per container.
// TODO: Delete;
// TODO: Create; ~MAYBE~
func ReadInfoFile() []Image {
	var rValue []Image
	f, err := os.Open("image_info.txt")
	if err != nil {
		log.Fatalf("Erorr opening file %v", err)
		return nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		if len(fields) >= 3 {
			repository := fields[0]
			imageID := fields[2]
			rValue = append(rValue, Image{repo: repository, id: imageID})
			//	fmt.Printf("Reposiory: %s, Image ID: %s\n", repository, imageID)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return rValue
}

func ConvertImagesToStringArr(images []Image) []string {
	var result []string
	for _, img := range images {
		result = append(result, fmt.Sprintf("Repository: %s, Image ID: %s", img.repo, img.id))
	}

	return result
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize termui: %v", err)
	}
	defer ui.Close()

	go GetImageData()

	images := ReadInfoFile()
	p := widgets.NewParagraph()
	l := widgets.NewList()

	l.Title = "Images"
	l.Rows = ConvertImagesToStringArr(images)
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 60, 20)

	p.Text = "Hello World!"
	p.SetRect(10, 20, 25, 5)

	ui.Render(p, l)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
