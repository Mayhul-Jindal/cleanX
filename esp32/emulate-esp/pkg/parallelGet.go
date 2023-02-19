package pkg

import (
	"log"
	"net/http"
	"time"
)

func ParallelGet(url string, ch chan <- time.Duration){
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- 0
		log.Fatal(err)
	}else{
		timeTaken := time.Since(start).Round(time.Millisecond)
		ch <- timeTaken
		resp.Body.Close()
	}
}