package timesync

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
)

type CheckIn struct {
	Duration *int `json:"duration"`
	User     *string `json:"user"`
	Project  *string `json:"project"`
	Activity *string `json:"activity"`
	Notes    *string `json:"notes"`
	Issue    *string `json:"issue_uri"`
	Date     *string `json:"date"`
}

func PostCheckIn(config Config, checkInData CheckIn) error {
	url := config.URL

	checkinAsJson, err := json.Marshal(checkInData)
	if err != nil {
		panic(err)
	}
	body := bytes.NewReader(checkinAsJson)

	resp, err := http.Post(url, "text/json", body)

	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		fmt.Print("Uh oh!! The request wasn't OK.")
		fmt.Print(resp.Body)
		return errors.New("Request failed");
	} else {
		return nil
	}
}

func ParseCheckIn(config Config, arguments []string) error {
	flagSet := flag.FlagSet{}

	checkInData := CheckIn{
		Duration: flagSet.Int("duration", 0,
			"The time spent on the project in minutes, required"),
		User:     flagSet.String("user", "", "The username, required"),
		Project:  flagSet.String("project", "", "The project, required"),
		Activity: flagSet.String("activity", "", "The activity, required"),
		Notes:    flagSet.String("notes", "", "The notes"),
		Issue:    flagSet.String("issue", "", "A URI for the issue"),
		Date:     flagSet.String("date", "", "The date work was done"),
	}

	flagSet.Parse(arguments)

	// TODO: validate date fields

	if *checkInData.Duration == 0 || *checkInData.Project == "" ||
		*checkInData.Activity == "" {
		fmt.Println("Duration, project, and activity are all required parameters")
		return errors.New("Required parameter missing")
	}

	err := PostCheckIn(config, checkInData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
