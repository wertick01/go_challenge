package main


import (
	"fmt"
	"time"
	"context"
)

type GoMetrClient struct {
	ServiceID string
	Timeout time.Duration
	Problem string
}

type HealthCheck struct {
	functionality bool
}

func (gmc *GoMetrClient) getHealth() *HealthCheck {

	if gmc.ServiceID != "" {
		return &HealthCheck{functionality: true}
	} else {
		return &HealthCheck{functionality: false}
	}
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

func Create_Client(s string, tm time.Duration) *GoMetrClient {
	return &GoMetrClient{ServiceID: s, Timeout: tm}
}

func (gmc *GoMetrClient) GetMetrics() string {
	return "Sorry, or site was hacked by Ukranian hackers and because of this we have lost your metrics :((("
}

func (gmc *GoMetrClient) Ping() string {
	return "Ping is so high, that we can't fit it in a string (stupid Ukranian hackers) :((("
}

func (gmc *GoMetrClient) GetID() string {
	return gmc.ServiceID
}
