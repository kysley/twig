package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var HeadphoneID = "{0.0.0.00000000}.{c8cd3c1f-1362-4155-a94e-1dc6de192397}"
var SpeakersID = "{0.0.0.00000000}.{b1455689-75ac-44d8-bd9a-3a133140e8c3}"

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
