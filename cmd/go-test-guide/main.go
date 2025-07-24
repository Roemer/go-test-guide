package main

import (
	"context"
	"fmt"
	"log"
	"os"

	gotestguide "github.com/roemer/go-test-guide"
	"github.com/roemer/go-test-guide/internal"
	gotestguideapp "github.com/roemer/go-test-guide/internal/app/go-test-guide"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:    "go-test-guide",
		Usage:   "A CLI tool for managing Test.Guide",
		Version: internal.Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "base-url",
				Usage:   "Base URL of the Test.Guide server",
				Sources: cli.EnvVars("TEST_GUIDE_BASE_URL"),
			},
			&cli.StringFlag{
				Name:    "token",
				Usage:   "API token for authenticating with the Test.Guide server",
				Sources: cli.EnvVars("TEST_GUIDE_TOKEN"),
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "report-management",
				Aliases: []string{"rm"},
				Usage:   "Manage reports",
				Commands: []*cli.Command{
					{
						Name:  "upload-report",
						Usage: "Upload a new report",
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:     "project",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "converter",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "report",
								Required: true,
							},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							return runAction(ctx, cmd, func(client *gotestguide.Client) error {
								projectID := cmd.Int("project")
								converter := cmd.String("converter")
								report := cmd.String("report")
								return gotestguideapp.UploadReport(client, projectID, converter, report)
							})
						},
					},
					{
						Name:  "add-artifact",
						Usage: "Add a new artifact",
						Flags: []cli.Flag{
							&cli.Int64Flag{
								Name:     "tce",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "artifact",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "comment",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "category",
								Required: true,
							},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							return runAction(ctx, cmd, func(client *gotestguide.Client) error {
								tceId := cmd.Int64("tce")
								filePath := cmd.String("artifact")
								comment := cmd.String("comment")
								category := cmd.String("category")
								return gotestguideapp.AddArtifact(client, tceId, filePath, comment, category)
							})
						},
					},
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func runAction(ctx context.Context, cmd *cli.Command, clientFunc func(client *gotestguide.Client) error) error {
	client, err := createClient(cmd)
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}
	return clientFunc(client)
}

func createClient(cmd *cli.Command) (*gotestguide.Client, error) {
	baseURL := cmd.String("base-url")
	token := cmd.String("token")

	if baseURL == "" || token == "" {
		return nil, fmt.Errorf("base-url and token are required")
	}

	return gotestguide.NewClient(baseURL, token)
}
