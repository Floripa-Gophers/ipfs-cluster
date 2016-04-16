package main

import (
	"os"
	"log"
)

func main() {
	var (
		subcmd string
		err error
	)

	if len(os.Args) < 2 {
		log.Fatal("Requires a sub command...\n")
	}
		
	subcmd = os.Args[1]

	switch subcmd {
	case "init":
		err = initCluster(os.Args[1:])
	case "join":
		err = joinCluster(os.Args[1:])
	case "add":
		err = addNode(os.Args[1:])
	case "sync":
		err = syncCluster(os.Args[1:])
	}

	if err != nil {
		log.Fatal("cmd failed: %s\n", err.Error())
	}
}
