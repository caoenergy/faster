package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	input_type := "github"
	if len(os.Args) >= 2 {
		input_type = os.Args[1]
	}
	if input_type == "" {
		input_type = "github"
	}
	if input_type != "github" && input_type != "docker" {
		fmt.Println("Usage: faster [github|docker]")
		return
	}
	if r, e := Get(fmt.Sprintf("https://api.akams.cn/%s", input_type)); e == nil {
		fmt.Println(fmt.Sprintf("%s加速器: %s", input_type, r.URL))
	} else {
		fmt.Println("没有可用的节点")
	}
}

type Result struct {
	URL        string
	Latency    time.Duration
	StatusCode int
	Error      string
}

func Get(link string) (*Result, error) {
	resp, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := make(map[string]interface{})
	err = json.Unmarshal(bytes, &r)
	if err != nil {
		return nil, err
	}
	code := r["code"].(float64)
	msg := r["msg"].(string)
	if code == 200 && msg == "success" {
		data := r["data"].([]interface{})
		results := make(chan Result, len(data))
		var wg sync.WaitGroup
		for _, v := range data {
			item := v.(map[string]interface{})
			wg.Add(1)
			go testURL(item["url"].(string), &wg, results)
		}
		go func() {
			wg.Wait()
			close(results) // 关闭 channel
		}()
		var faster Result
		isFirst := true
		for result := range results {
			if result.Error == "" && result.StatusCode >= 200 && result.StatusCode < 300 {
				if isFirst || result.Latency < faster.Latency {
					faster = result
					isFirst = false
				}
			}
		}
		if faster.URL != "" {
			return &faster, nil
		}
	}
	return nil, fmt.Errorf("没有可用的节点")
}

func testURL(url string, wg *sync.WaitGroup, results chan<- Result) {
	defer wg.Done()
	start := time.Now()
	client := &http.Client{
		Timeout: 2 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		results <- Result{URL: url, Error: err.Error()}
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		results <- Result{URL: url, Error: err.Error()}
		return
	}
	defer resp.Body.Close()
	duration := time.Since(start)
	results <- Result{URL: url, Latency: duration, StatusCode: resp.StatusCode}
}
