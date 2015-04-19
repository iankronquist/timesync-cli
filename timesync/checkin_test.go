package timesync

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
)

func Test_PostCheckIn_200(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			expectedURL := "/checkin/add"
			if r.URL.String() != expectedURL {
				t.Error("Unexpected URL", r.URL.String(),
					"expected", expectedURL)
			}
			bodyData, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				t.Error("Error reading body")
			}
			expectedBodyData := []byte("{\"duration\":12,\"user\":\"iankronquist\",\"project\":\"ww\",\"activity\":\"review\",\"notes\":\"reviewed PR #14\",\"issue_uri\":\"github.com/osu-cass/whats_fresh_api/pulls/14\",\"date\":\"2015-04-18\"}");

			if bytes.Compare(bodyData, expectedBodyData) != 0 {
				t.Error("Request body did not match. Got: ", string(bodyData),
					"expected", string(expectedBodyData))
			}
		}))
	defer ts.Close()

	checkInData := CheckIn {
		Duration: 12,
		User:  "iankronquist",
		Activity: "review",
		Notes: "reviewed PR #14",
		Project: "ww",
		Issue: "github.com/osu-cass/whats_fresh_api/pulls/14",
		Date: "2015-04-18",
	}

	config := Config {
		URL: ts.URL + "/checkin/add",
	}

	err := PostCheckIn(config, checkInData)

	if err != nil {
		t.Error("Error in PostCheckIn", err)
	}
}
