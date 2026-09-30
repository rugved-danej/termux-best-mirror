package tester

import (
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/pterm/pterm"
	"github.com/termux/termux-best-mirror/internal/mirror"
)

func TestMirrors(mirrors []string) []mirror.Result {
	var results []mirror.Result
	var mu sync.Mutex
	var wg sync.WaitGroup

	semaphore := make(chan struct{}, 20)

	pb, _ := pterm.DefaultProgressbar.
		WithTotal(len(mirrors)).
		WithTitle("Benchmarking mirrors").
		WithShowElapsedTime(true).
		WithShowCount(true).
		Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		pb.Stop()
		pterm.Error.Println("Process interrupted by user")
		os.Exit(1)
	}()

	client := &http.Client{
		Timeout: 8 * time.Second,
	}

	for _, m := range mirrors {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			url := mirror.GetMirrorURL(path)
			if url == "" {
				mu.Lock()
				pb.Increment()
				mu.Unlock()
				return
			}

			testUrl := url + "dists/stable/Release"

			start := time.Now()
			req, err := http.NewRequest("HEAD", testUrl, nil)
			var duration time.Duration
			success := false

			if err == nil {
				resp, err := client.Do(req)
				if err == nil {
					resp.Body.Close()
					if resp.StatusCode == 200 {
						duration = time.Since(start)
						success = true
					}
				}
			}

			if !success {
				start = time.Now()
				resp, err := client.Get(testUrl)
				if err == nil {
					resp.Body.Close()
					if resp.StatusCode == 200 {
						duration = time.Since(start)
						success = true
					}
				}
			}

			if !success {
				duration = time.Since(start)
			}

			mu.Lock()
			results = append(results, mirror.Result{
				Path:     path,
				Name:     filepath.Base(path),
				Duration: duration,
				Success:  success,
			})
			pb.Increment()
			mu.Unlock()
		}(m)
	}

	wg.Wait()
	pb.Stop()

	return results
}
