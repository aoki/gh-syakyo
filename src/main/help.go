package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var cmdVersion = &Command{
	Run:   runVersion,
	Usage: "version",
	Short: "show gh version",
	Long:  `Version show the gh client version showstring.`,
}

func runVersion(cmd *Command, args []string) {
	fmt.Println(Version)
}

var cmdHelp = &Command{
	Usage: "help [command]",
	Short: "show help",
	Long:  `Help shows usage for a command.`,
}

func init() {
	cmdHelp.Run = runHelp // break init loop
}

func runHelp(cmd *Command, args []string) {
	if len(args) == 0 {
		printUsage()
		return // not os.Exit(2); success
	}
	if len(args) != 1 {
		log.Fatal("too many arguments")
	}
	for _, cmd := range commands {
		if cmd.Name() == args[0] {
			cmd.printUsage()
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown help topic: %q. Run 'acldt helpp'. \n", args[0])
	os.Exit(2)
}

var usageTplText = `Usage: gh [command] [options] [arguments]

Supported command are:
{{range .Commands}}{{if .Runnnable}}{{if .ShowUsage}}
{{.Name | printf "%-8s"}} {{.Short}}{{end}}{{end}}{{end}}

See 'gh help [command]' for more information about a command.
`

var usageTemplate = template.Must(template.New("usage").Parse(usageTplText))

func printUsage() {
	usageTemplate.Execute(os.Stdout, struct {
		Commands []*Command
	}{
		commands,
	})
}

func usage() {
	printUsage()
	os.Exit(2)
}
