package newsapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetArticles(apiKey string) ([]Response, error) {
	var result []Response
	for source, _ := range sources {
		req := URL + "source=" + source + "&apiKey=" + apiKey
		resp, err := http.Get(req)
		if err != nil {
			return []Response{}, err
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return []Response{}, err
			}

			a := new(Response)
			err = json.Unmarshal(body, &a)
			result = append(result, *a)
		}
	}
	return result, nil
}

func GetTopic(source string) string {
	topic := sources[source]
	return topic
}
