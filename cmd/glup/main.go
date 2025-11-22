package main

import (
	"context"
	"os"

	cli "github.com/urfave/cli/v3"

	clone "github.com/adzpm/glup/internal/app/clone"
	logger "github.com/adzpm/glup/internal/logger"
)

func main() {
	// Create logger instance
	lgr := logger.New()

	app := &cli.Command{
		Name:  "glup",
		Usage: "Clones all available repositories from GitLab",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "gitlab-host",
				Usage:   "GitLab host (e.g., gitlab.com)",
				Sources: cli.EnvVars("GITLAB_HOST"),
			},
			&cli.StringFlag{
				Name:    "gitlab-user",
				Usage:   "GitLab user",
				Sources: cli.EnvVars("GITLAB_USER"),
			},
			&cli.StringFlag{
				Name:    "gitlab-token",
				Usage:   "GitLab access token",
				Sources: cli.EnvVars("GITLAB_TOKEN"),
			},
		},
		Commands: []*cli.Command{
			{
				Name:      "clone",
				Usage:     "Clones all available repositories",
				ArgsUsage: "[directory]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "group",
						Usage: "Clone repositories only from specified group",
					},
				},
				Action: clone.Clone,
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		lgr.Fatal(err)
	}
}
