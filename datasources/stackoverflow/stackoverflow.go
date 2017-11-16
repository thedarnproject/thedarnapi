package stackoverflow

import (
	"io/ioutil"
	"net/http"

	"encoding/json"

	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
)

func getResults(query string) (*responseStruct, error) {

	httpClient := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		"https://api.stackexchange.com/2.2/search/advanced",
		nil)
	if err != nil {
		return nil, errors.Wrap(err, "error building request for StackExchange API")
	}

	parameters := req.URL.Query()
	parameters.Set("order", "desc")
	parameters.Set("sort", "relevance")
	parameters.Set("q", query)
	parameters.Set("site", "stackoverflow")
	parameters.Set("answers", "1")
	parameters.Set("accepted", "True")
	req.URL.RawQuery = parameters.Encode()

	response, err := httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error querying StackExchange API")
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading response body")
	}

	unmarshalled := responseStruct{}
	json.Unmarshal(contents, &unmarshalled)

	return &unmarshalled, nil
}

func getAnswerFromId(answerId int) (*answerStruct, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://api.stackexchange.com/2.2/answers/%v", answerId),
		nil)
	if err != nil {
		return nil, errors.Wrap(err, "error building request for StackExchange API")
	}

	parameters := req.URL.Query()
	parameters.Set("order", "desc")
	parameters.Set("sort", "activity")
	parameters.Set("site", "stackoverflow")
	parameters.Set("filter", "!9YdnSLiq6")
	req.URL.RawQuery = parameters.Encode()

	response, err := httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error querying StackExchange API")
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading response body")
	}

	unmarshalled := answerStruct{}
	json.Unmarshal(contents, &unmarshalled)

	return &unmarshalled, nil
}

func DarnIt(query string) (string, error) {

	var fix string

	// Get results
	results, err := getResults(query)
	if err != nil {
		return "", errors.Wrap(err, "error getting results")
	}

	// Return error if no results
	if len(results.Items) == 0 {
		return "", errors.New("the query did not return any results")
	}

	for priority, result := range results.Items[0:1] {
		if result.AcceptedAnswerID != 0 {
			answer, err := getAnswerFromId(result.AcceptedAnswerID)
			if err != nil {
				return "", errors.Wrap(err, "unable to get answer from ID")
			}
			logrus.Infof("suggesting answer:", answer.Items[0].BodyMarkdown)
			fix = fix +
				fmt.Sprintf("Priority %v fix -", priority) +
				"\n" +
				answer.Items[0].BodyMarkdown +
				"\n"
		}
	}
	return fix, nil
}
