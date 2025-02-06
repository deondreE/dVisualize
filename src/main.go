package main

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func returnImageArray(images []image.Summary) []string {
	var result []string

	for _, image := range images {
		result = append(result, image.ID)
	}

	return result
}

func returnContainerArray(containers []types.Container) []string {
	var result []string

	for _, container := range containers {
		result = append(result, container.ID)
	}

	return result
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize termui: %v", err)
	}
	defer ui.Close()

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Docker daemon is likely not runing: %v", err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, containertypes.ListOptions{})
	if err != nil {
		log.Fatalf("Docker daemon is likely not runing: %v", err)
	}

	l := widgets.NewList()
	l2 := widgets.NewList()

	images, err := cli.ImageList(ctx, image.ListOptions{})
	if err != nil {
		log.Fatalf("Docker daemon is likely not runing: %v", err)
	}

	l.Title = "Current Images"
	l.Rows = returnImageArray(images)
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 10, 50, 20)

	l2.Title = "Containers"
	l2.Rows = returnContainerArray(containers)
	l.TextStyle = ui.NewStyle(ui.ColorCyan)
	l.WrapText = true
	l2.SetRect(51, 10, 101, 20)

	ui.Render(l, l2)

	// TODO: Selection based rendering for the cpu,mem usage inside of tabs.
	for e := range ui.PollEvents() {
		switch e.ID {
		case "q", "<C-c>":
			return
		}

		switch e.Type {
		case ui.KeyboardEvent:
			continue
		}
	}
}
