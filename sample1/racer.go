package sample1

import (
	"net/http"
	"time"
)

func Racer(slowUrl, fastUrl string) string {
	slowDuration := measureResponseTime(slowUrl)
	fastDuration := measureResponseTime(fastUrl)

	if slowDuration < fastDuration {
		return slowUrl
	}

	return fastUrl
}


func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}