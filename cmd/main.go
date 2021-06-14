package main

import (
	"github.com/evgeny-klyopov/bashColor"
	"github.com/evgeny-klyopov/blockchain-node-export/internal/params"
	"github.com/evgeny-klyopov/blockchain-node-export/internal/report"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	mode := os.Getenv("BNE_MODE")

	appHelp, commandHelp := helpTemplate()

	cli.AppHelpTemplate = appHelp
	cli.CommandHelpTemplate = commandHelp

	if mode == "dev" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
		}
	}

	p := params.NewParams()

	app := &cli.App{
		Name: "BNE" +
			"",
		Version:  "v0.0.5",
		Flags:    p.GetFlags(),
		HelpName: "bne",
		Usage:    "Blockchain node export transactions",
		Action: func(c *cli.Context) error {
			return report.NewReport(p).Run()
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func helpTemplate() (appHelp string, commandHelp string) {
	c := bashColor.NewColor()
	info := c.White(`{{.Name}} - {{.Usage}}`)
	info += c.Green(`{{if .Version}} {{.Version}}{{end}}`)

	appHelp = info + `

` + c.Yellow("Usage:") + `
	{{.HelpName}} {{if .VisibleFlags}}[options]{{end}}
 {{if .Commands}}
` + c.Yellow("Commands:") + `
{{range .Commands}}{{if not .HideHelp}}` + "	" + c.GetColor(bashColor.Green) + `{{join .Names ", "}}` + c.GetColor(bashColor.Default) + `{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
` + c.Yellow("Options:") + `
{{range .VisibleFlags}}  {{.}}
{{end}}{{end}}`

	commandHelp = c.Yellow("Description:") + ` 
   {{.Usage}}
` + c.Yellow("Usage:") + `
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}}{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}
{{if .VisibleFlags}}
` + c.Yellow("Arguments:") + `
	` + c.GetColor(bashColor.Green) + `stage` + c.GetColor(bashColor.Default) + `{{ "\t"}}{{ "\t"}}{{ "\t"}}{{ "\t"}} Stage or hostname
` + c.Yellow("Options:") + `
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}
`
	return appHelp, commandHelp
}
