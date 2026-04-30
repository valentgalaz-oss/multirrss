package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	ig "github.com/felipeinf/instago"
)

func main() {
	var (
		sessionPath = flag.String("session", "session.json", "Path to session settings JSON file")
		login       = flag.Bool("login", false, "If true, login using INSTAGO_USERNAME/INSTAGO_PASSWORD and save session")
	)
	flag.Parse()

	c := ig.NewClient()

	if _, err := os.Stat(*sessionPath); err == nil {
		if err := c.LoadSettings(*sessionPath, false); err != nil {
			log.Fatal(err)
		}
	} else if *login {
		username := os.Getenv("INSTAGO_USERNAME")
		password := os.Getenv("INSTAGO_PASSWORD")
		if username == "" || password == "" {
			log.Fatal("missing env vars: INSTAGO_USERNAME and/or INSTAGO_PASSWORD")
		}
		if err := c.Login(username, password, ""); err != nil {
			log.Fatal(err)
		}
		if err := c.DumpSettings(*sessionPath); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("No session file found. Create one by running:")
		fmt.Println(`  $env:INSTAGO_USERNAME="your_username"; $env:INSTAGO_PASSWORD="your_password"`)
		fmt.Println("  go run . -login -session session.json")
		return
	}

	me, err := c.AccountInfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("logged in as", me.Username)
}

