package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	ig "github.com/felipeinf/instago"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	switch os.Args[1] {
	case "login":
		loginCmd(os.Args[2:])
	case "me":
		meCmd(os.Args[2:])
	case "-h", "--help", "help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n\n", os.Args[1])
		usage()
		os.Exit(2)
	}
}

func usage() {
	fmt.Println("cli - minimal instago CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  cli <command> [flags]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  login   Login using INSTAGO_USERNAME/INSTAGO_PASSWORD and write session")
	fmt.Println("  me      Print logged-in username using an existing session")
	fmt.Println()
	fmt.Println("Run:")
	fmt.Println("  cli <command> -h")
}

func loginCmd(args []string) {
	fs := flag.NewFlagSet("login", flag.ExitOnError)
	sessionPath := fs.String("session", "session.json", "Path to session settings JSON file")
	_ = fs.Parse(args)

	username := os.Getenv("INSTAGO_USERNAME")
	password := os.Getenv("INSTAGO_PASSWORD")
	if username == "" || password == "" {
		log.Fatal("missing env vars: INSTAGO_USERNAME and/or INSTAGO_PASSWORD")
	}

	c := ig.NewClient()
	if err := c.Login(username, password, ""); err != nil {
		log.Fatal(err)
	}
	if err := c.DumpSettings(*sessionPath); err != nil {
		log.Fatal(err)
	}
	fmt.Println("ok:", *sessionPath)
}

func meCmd(args []string) {
	fs := flag.NewFlagSet("me", flag.ExitOnError)
	sessionPath := fs.String("session", "session.json", "Path to session settings JSON file")
	_ = fs.Parse(args)

	c := ig.NewClient()
	if err := c.LoadSettings(*sessionPath, false); err != nil {
		log.Fatal(err)
	}

	me, err := c.AccountInfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(me.Username)
}

