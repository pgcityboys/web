package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Meeting struct {
	ID        int    `json:"id"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Category  string `json:"category"`
	Link      string `json:"link"`
	Day       int    `json:"day"`
	Month     int    `json:"month"`
	Year      int    `json:"year"`
}

func getMeetings(category string) ([]Meeting, error) {
	apiUrl := fmt.Sprintf("http://userio:8080/api/v1/meetings/%s", category)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var meetings []Meeting
	err = json.Unmarshal(body, &meetings)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return meetings, nil
}

func HandleGetMeeting(w http.ResponseWriter, r *http.Request) {
	category := r.PathValue("category")
	meetings, err := getMeetings(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	meetingsJSON, err := json.Marshal(meetings)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal JSON: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(meetingsJSON)
}

func HandleCreateMeeting(w http.ResponseWriter, r *http.Request) {
	// Odczytanie ciała żądania
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("błąd odczytu ciała żądania: %v", err), http.StatusInternalServerError)
		return
	}
	targetURL := "http://userio:8080/api/v1/meetings"

	req, err := http.NewRequest("POST", targetURL, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, fmt.Sprintf("błąd tworzenia nowego żądania: %v", err), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("błąd wykonania żądania: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("błąd odczytu odpowiedzi: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}
