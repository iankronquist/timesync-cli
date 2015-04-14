package timesync

import (
	"fmt"
	"net/http"
	"net/url"
)

func CheckIn(config map[string]string, issue_uri string, activity string,
			project string, notes string, duration int) error {
	domain := config["domain"]
	port := config["port"]
	user := config["user"]
	issue_uri = url.QueryEscape(issue_uri)
	activity = url.QueryEscape(activity)
	project = url.QueryEscape(project)
	notes = url.QueryEscape(notes)
	query := fmt.Sprintf("activity=%s&issue_uri=%s&project=%s&notes=%s&user=%s&duration=%d", activity, issue_uri, project, notes, user, duration)
	resp, err := http.Get(fmt.Sprintf("http://%s:%s/time/add?%s", domain, port, query))
	if err != nil {
		panic(err)
		//return err
	} else {
		fmt.Print(resp.Body);
		fmt.Print("success!");
		return nil
	}
}
