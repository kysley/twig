package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var HeadphoneID = "{0.0.0.00000000}.{72abcb0f-8df8-49db-a7e7-557ca8b81fff}"
var SpeakersID = "{0.0.0.00000000}.{9791eeb2-85e5-48c4-a591-1b8342b0a042}"

func main() {
	argsWithoutProg := os.Args[1:]
	profile := argsWithoutProg[0]

	flag.Parse()
	if profile == "speakers" {
		out, err := exec.Command("powershell", "Set-AudioDevice", "-ID", fmt.Sprintf("'%s'", SpeakersID)).CombinedOutput()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + string(out))
		} else {
			fmt.Printf("%s", out)
		}
	} else if profile == "headphones" {
		out, err := exec.Command("powershell", "Set-AudioDevice", "-ID", fmt.Sprintf("'%s'", HeadphoneID)).CombinedOutput()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + string(out))
		} else {
			fmt.Printf("%s", out)
		}
	}
}
