package main

import (
	"flag"
	"fmt"
	"os"

	"jlowell000.github.io/budgeting/internal/service"
)

const (
	CMD_CREATE         = "create"
	CMD_READ           = "read"
	FLG_ALL            = "all"
	VAR_CONTENT        = "content"
	ENTRYLIST_FILENAME = "./entrylist.json"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("expected '" + CMD_CREATE + ", or " + CMD_READ + "' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case CMD_CREATE:
		create()
	case CMD_READ:
		read()
	default:
		fmt.Println("expected '" + CMD_CREATE + ", or " + CMD_READ + "' subcommands")
		os.Exit(1)
	}
}

func create() {
	cmd := flag.NewFlagSet(CMD_CREATE, flag.ExitOnError)
	content := cmd.String(VAR_CONTENT, "", VAR_CONTENT)
	cmd.Parse(os.Args[2:])
	entry := service.CreateEntry(*content, ENTRYLIST_FILENAME)
	fmt.Println("Added entry:", entry)
}

func read() {
	cmd := flag.NewFlagSet(CMD_READ, flag.ExitOnError)
	all := cmd.Bool(FLG_ALL, false, FLG_ALL)
	cmd.Parse(os.Args[2:])

	if *all {
		fmt.Println(service.GetEntryList(ENTRYLIST_FILENAME))
	} else {
		fmt.Println(service.GetLatestEntry(ENTRYLIST_FILENAME))
	}
}
