package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/beardedandnotmuch/google-sheets-observer/internal/app/cache"
	"github.com/beardedandnotmuch/google-sheets-observer/internal/app/models/sheets"
)

type Service struct {
	googleSheets sheets.SheetStorage
	cache        cache.GoogleSheetsCache
}

func New() *Service {
	return &Service{
		&sheets.Storage{},
		&cache.RedisCache{},
	}
}

func (s *Service) GetSheetsData(sId string, rng string) []string {
	if len(sId) == 0 {
		sId = os.Getenv("GOOGLE_SHEET_ID")
	}

	//TODO: checking data in cache getFromCache(sId)

	response, err := sendGoogleSheetsRequest(sId, rng)

	if err != nil {
		log.Fatal(err)
	}

	return parseResponse(response)
}

func (s *Service) InitCache() {

}

func sendGoogleSheetsRequest(sId string, rng string) ([]sheets.RowData, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", os.Getenv("GOOGLE_SHEETS_API_URL"), sId), nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()

	if len(rng) > 0 {
		q.Add("ranges", rng)
	}

	q.Add("includeGridData", "true")
	q.Add("key", os.Getenv("GOOGLE_API_KEY"))

	req.URL.RawQuery = q.Encode()

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		defer resp.Body.Close()

		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	googleSheetsResponse := &sheets.Response{}
	resbody, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(resbody, &googleSheetsResponse)

	defer resp.Body.Close()

	return googleSheetsResponse.Sheets[0].Data[0].RowData, nil
}

func parseResponse(r []sheets.RowData) []string {
	result := make([]string, len(r))

	if len(r) == 0 {
		return result
	}

	for i, d := range r {
		if d.Values[0].Value != "" {
			result[i] = d.Values[0].Value
		}
	}

	return result
}

func getFromCache(sId string) (string, error) {
	var data string

	return data, nil
}

func startPolling(intervalInSec int) {

}
