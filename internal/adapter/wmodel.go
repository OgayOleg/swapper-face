package adapter

import (
	"bytes"
	"encoding/json"
	"faceSwapper/internal/dto"
	"fmt"
	"net/http"
	"time"
)

const VERSION = "537e83f7ed84751dc56aa80fb2391b07696c85a49967c72c64f002a0ca2bb224"

type Wmodel struct {
	apiKey string
}

func New(apiKey string) *Wmodel {
	return &Wmodel{
		apiKey: apiKey,
	}
}

// Post sends a POST request to the Wmodel API to create a new task.
func (W *Wmodel) GetID(face, target string) (string, error) {
	url := "https://api.vmodel.ai/api/tasks/v1/create"
	reqBodyStruct := dto.NewFaceSwapReq(face, target, VERSION)
	reqBodyToJSON, err := json.Marshal(reqBodyStruct)
	if err != nil {
		return "", fmt.Errorf("Get ID: failed marshal struct to JSON: %v", err)
	}
	reqBodytoReader := bytes.NewReader(reqBodyToJSON)

	req, err := http.NewRequest("POST", url, reqBodytoReader)
	if err != nil {
		return "", fmt.Errorf("Get ID: failed create request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprint("Bearer ", W.apiKey))
	req.Header.Set("Content-Type", "application/json")

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Get ID: failed send response. %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Get ID: bad status code %d.", resp.StatusCode)
	}

	var response dto.ResponseID
	err = json.NewDecoder(resp.Body).Decode(&response)
	fmt.Println(response)
	if err != nil {
		return "", fmt.Errorf("Get ID: failed parse response. %v", err)
	}

	return response.Result.TaskID, nil
}

func (W *Wmodel) FetchResourceWithRetry(id string) (string, error) {
	client := &http.Client{}
	endpoint := "https://api.vmodel.ai/api/tasks/v1/get/" + id
	for {
		response, err := W.retryGetRequest(client, endpoint)
		if err != nil {
			return "", err
		}
		switch response.Result.Status {
		case "succeeded":
			var urls []string
			err := json.Unmarshal(response.Result.Output, &urls)
			if err != nil {
				return "", fmt.Errorf("Fetch Resource With Retry: failed unmarshal urls. %v", err)
			}
			return urls[0], nil
		case "failed":
			var code string
			json.Unmarshal(response.Result.Error, &code)
			return "", fmt.Errorf("%w %s", dto.ErrStatusFailed, code)
		case "canceled":
			var code string
			json.Unmarshal(response.Result.Error, &code)
			return "", fmt.Errorf("%w %s", dto.ErrStatusCanceled, code)
		default:
			time.Sleep(time.Second)
		}
	}

}

func (W *Wmodel) retryGetRequest(client *http.Client, endpoint string) (dto.ResponseURL, error) {
	// create request
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return dto.ResponseURL{}, fmt.Errorf("Retry Get Request: failed create request. %v", err)
	}

	// add header
	req.Header.Set("Authorization", fmt.Sprint("Bearer ", W.apiKey))

	// send request
	resp, err := client.Do(req)
	if err != nil {
		return dto.ResponseURL{}, fmt.Errorf("Retry Get Request: failed send. %v", err)
	}
	defer resp.Body.Close()

	// unmarshal response
	response := dto.ResponseURL{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return dto.ResponseURL{}, fmt.Errorf("Retry Get Request: failed unmarshal. %v", err)
	}

	return response, nil
}
