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

func GetContainerStats(cli *client.Client, ctx context.Context, containerID string) containertypes.StatsResponseReader {
	// TODO: make this a goroutine so that the main thread is not effected by this stream.
	stats, err := cli.ContainerStats(ctx, containerID, false)
	if err != nil {
		log.Fatalf("Container is not running: %v", err)
	}

	return stats
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
	tabpane := widgets.NewTabPane("cpu", "memory", "net", "pd")

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

	tabpane.SetRect(0, 1, 50, 4)
	tabpane.Border = true

	// TODO: add ui to render inside tabs
	renderTab := func() {
		switch tabpane.ActiveTabIndex {
		case 0:
			return
		case 1:
			return
		}
	}

	ui.Render(l, l2, tabpane)

	uiEvents := ui.PollEvents()

	// TODO: Selection based rendering for the cpu,mem usage inside of tabs.
	// TODO: Format so that container info is on left
	//
	//	Container Info -> Container Stats
	//	Image Info -> No Render
	// 	Kuberneties Cluster -> Tree View
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "h":
			tabpane.FocusLeft()
			ui.Clear()
			ui.Render(tabpane)
			renderTab()
		case "l":
			tabpane.FocusRight()
			ui.Clear()
			ui.Render(tabpane)
			renderTab()
		}
	}
}
