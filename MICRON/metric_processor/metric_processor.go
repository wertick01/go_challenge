package processors

import (
	"context"
	"errors"
	"fmt"
	"./checker"
	"./db"
	model "./metric_model"
	"log"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
)

type MetricProcessor struct {
	storage *db.Storage
	checker checker.Checkable
}

func CreateMetricProcessor(storage *db.Storage) *MetricProcessor {
	processor := new(MetricProcessor)
	processor.storage = storage
	return processor
}

func (processor *MetricProcessor) Start(cont context.Context) {
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Seconds().Do(func(){
		go processor.GetSmbMetrics(cont)
	})

	for {
		select {
		case <- cont.Done():
			fmt.Println("--> Context is Done")
			return
		}
	}
}

func (processor *MetricProcessor) Stop(cancel context.CancelFunc) {
	fmt.Println("--> Processor has been stopped")
	cancel()
}

func (processor *MetricProcessor) GetSmbMetrics(cont context.Context) {
	str := processor.checker.GetMetrics()
	for _, value := range strings.Split(str, "\n") {
		value = value[1:len(value)-1]
		value = strings.Split(value, ">")[1]
		value = strings.Split(value, "<")[0]
	}
	// прошу прощения, я ещё не умею получать метрики с веб страниц без помощи каких-нибудь прометеев
	// помогите написать эту функцию, если не сложно
}

func (processor *MetricProcessor) ListMetrics(ctx context.Context, filter models.Filters) ([]models.Filters, error) {
	return processor.storage.List(ctx, filter), nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}