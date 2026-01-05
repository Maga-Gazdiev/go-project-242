package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"code"

	"github.com/urfave/cli/v3"
)

func main() {
	// Создаём команду
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
			},
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			// Считываем значения флагов
			recursive := cmd.Bool("recursive")
			human := cmd.Bool("human")
			all := cmd.Bool("all")

			// Проверяем аргументы
			if cmd.Args().Len() == 0 {
				return fmt.Errorf("path is required")
			}

			path := cmd.Args().First()

			// Получаем размер
			size, err := code.GetPathSize(path, recursive, human, all)
			if err != nil {
				return err
			}

			fmt.Printf("%s\t%s\n", size, path)
			return nil
		},
	}

	// Запускаем команду напрямую с аргументами и контекстом
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
