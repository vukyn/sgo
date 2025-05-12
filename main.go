package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vukyn/sgo/pkg/analyzer"
)

var (
	Version = "1.0.0"
)

func main() {
	app := &cli.App{
		Name:    "sgo",
		Usage:   "Analyze and visualize Go project structure",
		Version: Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Value:   "text",
				Usage:   "Output format (json or text)",
				EnvVars: []string{"SGO_OUTPUT"},
			},
			&cli.StringFlag{
				Name:    "path",
				Aliases: []string{"p"},
				Value:   ".",
				Usage:   "Path to the Go project",
				EnvVars: []string{"SGO_PATH"},
			},
		},
		Action: func(c *cli.Context) error {
			path := c.String("path")
			outputFormat := c.String("output")

			analyzer := analyzer.NewAnalyzer(path)
			result, err := analyzer.Analyze()
			if err != nil {
				return fmt.Errorf("analysis failed: %w", err)
			}

			fmt.Println()
			if outputFormat == "json" {
				return result.ToJSON(os.Stdout)
			}
			return result.ToText(os.Stdout)
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
