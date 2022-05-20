package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := new(cli.App)
	usage(app)
	app.Action = func(c *cli.Context) error {
		fmt.Printf("h: %v, u: %v, p: %v, c: %v",
			c.String("h"),
			c.String("u"), c.Int("P"), c.String("c"))
		return nil
	}
	err := app.Run(os.Args) //app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

//构建命令使用说明
func usage(app *cli.App) {
	app.Name = "项目名称" //项目名称
	app.Authors = []*cli.Author{{
		Name:  "RicoLy",
		Email: "Rico",
	}}
	app.Version = "0.0.1"         //版本号
	app.Copyright = "@Copyright " //版权保护
	app.Usage = "Generator"       //说明
	app.Commands = []*cli.Command{
		{
			Name:    "help",
			Aliases: []string{"?"},
			Usage:   "show help",
			Action: func(c *cli.Context) error {
				_ = cli.ShowAppHelp(c)
				return nil
			},
		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "print the version",
			Action: func(c *cli.Context) error {
				cli.ShowVersion(c)
				return nil
			},
		},
	}
	app.HideVersion = true
	app.HideHelp = true
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "h", Value: "127.0.0.1", Usage: "Database address"},
		&cli.IntFlag{Name: "P", Value: 3306, Usage: "port number"},
		&cli.StringFlag{Name: "u", Value: "root", Usage: "database username", Required: true},
		&cli.StringFlag{Name: "p", Value: "", Usage: "database password", Required: true, DefaultText: ""},
		&cli.StringFlag{Name: "c", Value: "utf8mb4", Usage: "database format"},
		&cli.StringFlag{Name: "d", Usage: "database name"},
		&cli.BoolFlag{Name: "debug", Usage: "debug", Value: false},
	}
}
