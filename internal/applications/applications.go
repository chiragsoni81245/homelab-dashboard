package applications

import (
	"context"
	"sort"
	"strconv"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/filters"
	"github.com/moby/moby/client"
)

type Application struct {
	Index  int    `json:"index"`
	Name   string `json:"name"`
	WebURL string `json:"web_url"`
	Icon   string `json:"icon"`
}

func GetApplications() ([]Application, error) {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	// Filter containers that must have label x-homelab=true
	args := filters.NewArgs()
	args.Add("label", "x-homelab=true")

	containers, err := cli.ContainerList(ctx, container.ListOptions{
		All:     true,
		Filters: args,
	})
	if err != nil {
		return nil, err
	}

	apps := []Application{}

	for _, c := range containers {
		labels := c.Labels
		indexString := labels["x-homelab-index"]
		index, err := strconv.Atoi(indexString)	
		if err != nil {
			return nil, err
		}

		app := Application{
			Index:  index,
			Name:   labels["x-homelab-name"],
			WebURL: labels["x-homelab-web-url"],
			Icon:   labels["x-homelab-icon"],
		}

		// Only add if Name and WebURL are present (optional safeguard)
		if app.Name != "" && app.WebURL != "" {
			apps = append(apps, app)
		}
	}

	sort.Slice(apps, func(i, j int) bool {
		if apps[i].Index != apps[j].Index {
			return apps[i].Index < apps[j].Index
		} else {
			return apps[i].Name < apps[j].Name
		}
	})

	return apps, nil
}

