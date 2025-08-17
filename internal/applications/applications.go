package applications

import (
	"context"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/filters"
	"github.com/moby/moby/client"
)

type Application struct {
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

	var apps []Application

	for _, c := range containers {
		labels := c.Labels

		app := Application{
			Name:   labels["x-homelab-name"],
			WebURL: labels["x-homelab-web-url"],
			Icon:   labels["x-homelab-icon"],
		}

		// Only add if Name and WebURL are present (optional safeguard)
		if app.Name != "" && app.WebURL != "" {
			apps = append(apps, app)
		}
	}

	return apps, nil
}

