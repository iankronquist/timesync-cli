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
	config := getConfig()
	if verb == "checkin" {
		parseCheckIn(config, arguments[1:])
	} else if verb == "somethingelse" {

	} else {
		fmt.Println("help text")
		os.Exit(EX_USAGE);
	}

}

func getConfig() map[string]string {
	user := ""
	if user == "" {
		user = os.Getenv("USER")
	}

	editor := os.Getenv("EDITOR");
	if editor == "" {
		editor = "vi"
	}

	config := make(map[string]string)
	// TODO: parse these fields from config file
	config["domain"] = "localhost"
	config["port"] = "3000"
	config["user"] = user
	config["editor"] = editor
	return config
}

func parseCheckIn(config map[string]string, arguments []string) {
	flagSet :=flag.FlagSet{}

	checkInData := timesync.CheckIn {
		Duration: flagSet.Int("duration", 0,
			"The time spent on the project in minutes, required"),
		User: flagSet.String("user", "", "The username, required"),
		Project: flagSet.String("project", "", "The project, required"),
		Activity: flagSet.String("activity", "", "The activity, required"),
		Notes: flagSet.String("notes", "", "The notes"),
		Issue: flagSet.String("issue", "", "A URI for the issue"),
		Date: flagSet.String("date", "", "The date work was done"),
	}

	flagSet.Parse(arguments)

	// TODO: validate date fields

	if *checkInData.Duration == 0 || *checkInData.Project == "" ||
			*checkInData.Activity == "" {
		fmt.Println("Duration, project, and activity are all required parameters");
		os.Exit(EX_USAGE);
	}

	err := timesync.PostCheckIn(config, checkInData)
	if err != nil {
		fmt.Println(err)
	}
}
