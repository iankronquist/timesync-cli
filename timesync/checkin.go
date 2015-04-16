package timesync

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CheckIn struct {
	Duration *int
	User *string
	Project *string
	Activity *string
	Notes *string
	Issue *string
	Date *string
}

func PostCheckIn(config map[string]string, checkInDate CheckIn) error {
	domain := config["domain"]
	port := config["port"]

	url := fmt.Sprintf("http://%s:%s/time/add", domain, port)
	checkinAsJson, err := json.Marshal(checkInDate)
	if err != nil {
		panic(err)
	}
	body := bytes.NewReader(checkinAsJson)

	resp, err := http.Post(url, "text/json", body)

	if err != nil {
		panic(err)
		return err
	} else if resp.StatusCode != 200 {
		fmt.Print("Uh oh!! The request wasn't OK.");
		fmt.Print(resp.Body);
		return nil
	} else {
		return nil
	}
}
