package main

import (
	"github.com/floholz/imgut/internal/imgut"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var app *cli.App

func main() {
	app = &cli.App{
		Name:        "imgut",
		Description: "Perform url pattern shenanigans",
		Usage:       "Image Utility Tool",
		UsageText:   "imgut [command] [arguments]",
		Commands: []*cli.Command{
			docCommand(),
			downloadCommand(),
			fuzzCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func downloadCommand() *cli.Command {
	return &cli.Command{
		Name:      "download",
		Aliases:   []string{"d"},
		UsageText: "imgut download [arguments]",
		Args:      true,
		ArgsUsage: "Url pattern",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "outDir",
				Aliases: []string{"o"},
				Value:   "out",
				Usage:   "output directory for downloaded images",
			},
		},
		Action: func(c *cli.Context) error {
			baseUrl := c.Args().First()
			outDir := c.String("outDir")
			return imgut.DownloadImages(baseUrl, outDir)
		},
	}
}

func fuzzCommand() *cli.Command {
	return &cli.Command{
		Name:      "fuzz",
		Aliases:   []string{"f"},
		UsageText: "imgut fuzz [arguments]",
		Args:      true,
		ArgsUsage: "Url pattern to fuzz",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "outPath",
				Aliases: []string{"o"},
				Value:   "out/fuzz.json",
				Usage:   "output path for fuzz file",
			},
		},
		Action: func(c *cli.Context) error {
			url := c.Args().First()
			outPath := c.String("outPath")
			return imgut.FuzzUrl(url, outPath)
		},
	}
}

func docCommand() *cli.Command {
	return &cli.Command{
		Name:   "doc",
		Usage:  "Generate documentation",
		Hidden: true,
		Action: func(*cli.Context) error {
			md, err := app.ToMarkdown()
			if err != nil {
				return err
			}

			fileName := "doc.md"
			err = os.WriteFile(fileName, []byte(md), 0644)
			if err != nil {
				return err
			}

			log.Printf("Documentation created at %s\n", fileName)
			return nil
		},
	}
}
