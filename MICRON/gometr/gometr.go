package gometr


import (
	"fmt"
	"time"
	"context"
	"net/http"
	"io"
	"log"
	"github.com/go-ping/ping"
)

type GoMetrClient struct {
	ServiceID string
	Problem string
	Status int
	Timeout time.Duration
	cookies []*http.Cookie
}

type HealthCheck struct {
	functionality bool
}

func (gmc *GoMetrClient) getHealth() *HealthCheck {

	if gmc.ServiceID != "" {
		return &HealthCheck{functionality: true}
	} 
	return &HealthCheck{functionality: false}
}

func (gmc *GoMetrClient) Health(cont context.Context) bool {

	select {
	case <- cont.Done():
		gmc.getHealth().functionality = false
		gmc.Problem = "-->Context is DONE but smth is wrong:(\n"
	case <- time.After(gmc.Timeout):
		gmc.getHealth().functionality = false
		gmc.Problem = "-->Time is UP:(\n"
	default:
		gmc.getHealth().functionality = true
		gmc.Problem = "NO PROBLEMS in gmc.Problem"
	}

	fmt.Println(gmc.Problem)

	return gmc.getHealth().functionality
}

func CreateNewClient(s string, tm time.Duration) *GoMetrClient {
	return &GoMetrClient{
		ServiceID: s, 
		Problem: "",
		Timeout: tm,
	}
}

const st = "https://go.dev/"

func (gmc *GoMetrClient) GetMetrics() string {
	req, err := http.Get(st)
	gmc.Status = req.StatusCode
	gmc.cookies = req.Cookies()
	checker(err)
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	checker(err)

	return string(body)
}

func (gmc *GoMetrClient) Ping()  {
	pinger, err := ping.NewPinger(st)
	checker(err)
	pinger.Count = 3
	pinger.Run() // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
	fmt.Println(stats)
}

func (gmc *GoMetrClient) GetID() string {
	return gmc.ServiceID
}

func checker(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
