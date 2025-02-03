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
	go GetImageStatsFile()

	images := ReadInfoFile()
	stats, err := GetImageStats()
	if err != nil {
		log.Fatalf("Failed to worky: %v", err)
	}

	cpuUsage := GetCpuValues(stats)
	memUsage := GetMemVals(stats)

	l := widgets.NewList()

	sl := widgets.NewSparkline()
	sl.Data = cpuUsage
	sl.Title = "Cpu Usage"
	sl.LineColor = ui.ColorGreen

	sl1 := widgets.NewSparkline()
	sl1.Title = "Memory Usage"
	sl1.Data = memUsage
	sl1.LineColor = ui.ColorCyan

	slg := widgets.NewSparklineGroup(sl, sl1)
	slg.Title = "Usages"
	slg.SetRect(0, 0, 50, 10)

	l.Title = "Images View"
	l.Rows = ConvertImagesToStringArr(images)
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 10, 50, 20)

	ui.Render(l, slg)

	for e := range ui.PollEvents() {
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
