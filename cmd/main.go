package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	cli "github.com/jawher/mow.cli"

	"github.com/rtgnx/ghkey"
)

func main() {
	app := cli.App("ghkey", "github keys")

	var (
		allowedUsers   = app.StringsOpt("users", []string{}, "list of allowed users")
		allowedSources = app.StringsOpt("sources", []string{"github"}, "allowed sources")
		userName       = app.StringArg("USER", "", "user name")
	)

	app.Action = func() {
		keys, err := ghkey.Keys(*userName, ghkey.Config{AllowedSources: *allowedSources, AllowedUsers: *allowedUsers})
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Print(strings.Join(keys, "\n"))
	}

	app.Run(os.Args)
}
