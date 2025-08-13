package main

import (
	"flag"
	"fmt"
	"net/http"
    "strings"
    
	"github.com/fatih/color"
)

var version = "dev"

// colors
var (
	red   = color.New(color.FgRed).SprintFunc()
	green = color.New(color.FgGreen).SprintFunc()
)

func showBanner() {
	fmt.Printf(`
                 _                       
                | |                      
 __      ___   _| |_ ___  ___ _ ____   __
 \ \ /\ / / | | | __/ __|/ _ \ '__\ \ / /
  \ V  V /| |_| | |_\__ \  __/ |   \ V / 
   \_/\_/  \__,_|\__|___/\___|_|    \_/ v%s

   https://github.com/erhaem/wutserv

`, version)
}

func printErr(msg string) {
	fmt.Println(red("[err] " + msg))
}

func printOk(msg string) {
	fmt.Println(green("[ok] " + msg))
}

func main() {
	showBanner()

	url := flag.String("url", "", "url target")
	flag.Parse()

	if *url == "" {
		printErr("url is not defined -___-")
		printErr("usage: wutserv --url https://the-url-here.com")
		return
	}

	fmt.Println(fmt.Sprintf("[inf] url: %s, checking webserver..", *url))
	resp, err := http.Head(*url)
	if err != nil {
		printErr(err.Error())
		return
	}

	serv := resp.Header.Get("Server")
	if serv == "" {
		printErr(fmt.Sprintf("webserver not detected on %s", *url))
		return
	}
	if strings.Contains(strings.ToLower(serv), "cloudflare") {
	    printErr(fmt.Sprintf("no actual webserver detected on %s, the website is behind cloudflare", *url))
	    return
	}

	printOk(fmt.Sprintf("%s is using %s", *url, serv))
}
