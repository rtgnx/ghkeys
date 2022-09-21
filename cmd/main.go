package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	cli "github.com/jawher/mow.cli"
	"github.com/rtgnx/ghkeys"
)

func main() {
	app := cli.App("ghkeys", "github keys")

	var (
		allowedUsers   = app.StringsOpt("users", []string{}, "list of allowed users")
		allowedSources = app.StringsOpt("sources", []string{"github"}, "allowed sources")
		userName       = app.StringArg("USER", "", "user name")
	)

	app.Action = func() {
		keys, err := ghkeys.Keys(*userName, ghkeys.Config{AllowedSources: *allowedSources, AllowedUsers: *allowedUsers})
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Print(strings.Join(keys, "\n"))
	}

	app.Run(os.Args)
}
