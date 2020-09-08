package websock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AutoGenerated struct {
	Heartbeat struct {
		Interval int `json:"interval"`
	} `json:"heartbeat"`
}

func getHeartBeat() AutoGenerated {

	req, err := http.NewRequest("GET", "http://api.sportradar.us/nfl/official/simulation/stream/en/clock/subscribe?api_key=3h4vt4muu2d4qqn6mek23f2j", nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var example AutoGenerated

	err = json.Unmarshal(body, &example)
	if err != nil {
		panic(err)
	}

	return example

}
