package main

import (
	"flag"
	"fmt"
	"os"
	"./timesync"
)

// See man 3 sysexits, BSD
const EX_USAGE = 64

func main() {

	flag.Parse()
	arguments := flag.Args()

	if len(arguments) == 0 {
		fmt.Println("helptext")
		os.Exit(EX_USAGE);
	}
	verb := arguments[0]
	config, err:= timesync.GetConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(EX_USAGE);
	}
	if verb == "checkin" {
		timesync.ParseCheckIn(config, arguments[1:])
	} else if verb == "delete" {
		fmt.Println("delete an object...")
	} else if verb == "update" {
		fmt.Println("update an object...")
	} else if verb == "report" {
		fmt.Println("Generate a report...")
	} else {
		fmt.Println("help text")
		os.Exit(EX_USAGE);
	}

}


