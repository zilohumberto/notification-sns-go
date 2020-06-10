package main

import (
	"flag"
	"fmt"
	"github.com/zilohumberto/notification-sns-go/pkg/http/rest"
	"github.com/zilohumberto/notification-sns-go/pkg/sending"
	"github.com/zilohumberto/notification-sns-go/version"
	"log"
	"net/http"
)

func main() {
	versionFlag := flag.Bool("version", false, "Version")
	flag.Parse()

	if *versionFlag {
		fmt.Println("Build Date:", version.BuildDate)
		fmt.Println("Git Commit:", version.GitCommit)
		fmt.Println("Version:", version.Version)
		fmt.Println("Go Version:", version.GoVersion)
		fmt.Println("OS / Arch:", version.OsArch)
		return
	}
	var sender sending.Service
	sender = sending.NewService()
	// set up the HTTP server
	router := rest.Handler(sender)
	fmt.Println("The Notification server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
