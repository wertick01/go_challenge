package main

type GoMetrClient struct {
	ServiceID string
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

func (gmc *GoMetrClient) Health() bool {
	return gmc.getHealth().functionality
}

func Create_Client(s string) *GoMetrClient {
	return &GoMetrClient{ServiceID: s}
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
