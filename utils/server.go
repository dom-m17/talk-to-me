package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetDataFromAPI[T any](
	baseURL, apiKey string,
	res T,
) (T, error) {
	client := &http.Client{}
	request, err := http.NewRequest(
		http.MethodGet,
		baseURL,
		nil,
	)
	if err != nil {
		return res, fmt.Errorf("creating request: %w", err)
	}

	resp, err := client.Do(request)
	if err != nil {
		return res, fmt.Errorf("getting data: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error:", err)
		return res, fmt.Errorf("reading body: %w", err)
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("unmarshal error:", err)
		return res, fmt.Errorf("unmarshaling: %w", err)
	}

	return res, nil
}
