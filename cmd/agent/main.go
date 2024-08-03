package agent

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/XelerT/server.git/cmd/metric"
)

const (
	pollInterval   = time.Second * 2
	reportInterval = time.Second * 10
)

func doRequest(client *http.Client, request *http.Request) {
	response, err := client.Do(request)
	if err != nil {
		return
	}
	fmt.Println("Response status", response.Status)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	response.Body.Close()
}

func createDoRequestMetric(client *http.Client, metrType string, metrName string, metrValue string) {
	requestURL := fmt.Sprintf("http://localhost:8080/%s/%s/%s", metrType, metrName, metrValue)

	request, err := http.NewRequest(http.MethodPost, requestURL, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "text/plain")

	doRequest(client, request)
}

func main() {
	polls2report := int(reportInterval / pollInterval)

	metrs := metric.NewMetrics()
	for {
		i := 0
		client := &http.Client{}
		stdMetrVal := reflect.ValueOf(metrs.GetStd())
		if i == polls2report {
			for j := 0; j < stdMetrVal.NumField(); j++ {
				valStr := fmt.Sprint(stdMetrVal.Field(j))
				typeStr := fmt.Sprint(stdMetrVal.Field(j).Type())
				createDoRequestMetric(client, typeStr, stdMetrVal.Type().Field(j).Name, valStr)
			}
			pollCountVal := strconv.FormatInt(int64(metrs.GetPollCount()), 10)
			randomValueVal := strconv.FormatFloat(float64(metrs.GetRandomValue()), 'g', -1, 64)
			createDoRequestMetric(client, "counter", "PollCount", pollCountVal)
			createDoRequestMetric(client, "gauge", "RandomValue", randomValueVal)
			i = 0
		}
		metrs.UpdateAll()
		time.Sleep(pollInterval)
		i++
	}
}
