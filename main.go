package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ResponseData struct {
	SlackName string `json:"slack_name"`
	CurrentDay string `json:"current_day"`
	UTCTime string `json:"utc_time"`
	Track string `json:"track"`
	GithubFileURL string `json:"github_file_url"`
	GithubRepoURL string `json:"github_repo_url"`
	StatusCode int `json:"status_code"`
}

func main() {
	fmt.Println("Server running on port 8080")

    http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()

		dayName := currentTime.Weekday().String()
		const layout = "2006-01-02T15:04:05Z"
		formattedTime := currentTime.Format(layout)

		queryParams := r.URL.Query()
		slackName := queryParams.Get("slack_name")
		track := queryParams.Get("track")
		
		githubFileURL := "https://github.com/vicradon/hngx-task1-go/blob/main/main.py"
        githubRepoURL := "https://github.com/vicradon/hngx-task1-go"

		statusCode := 200
		
		response := ResponseData{
			CurrentDay: dayName, 
			UTCTime: formattedTime, 
			SlackName: slackName, 
			Track: track, 
			GithubFileURL: githubFileURL, 
			GithubRepoURL: githubRepoURL,
			StatusCode: statusCode,
		}

		jsonData, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
    })

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "It works!")
	})

    http.ListenAndServe(":8080", nil)
}