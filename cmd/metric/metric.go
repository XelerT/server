package metric

import "runtime"

type counter int64
type gauge float64

type Metrics struct {
	std         runtime.MemStats
	pollCount   counter
	randomValue gauge
}

func NewMetrics() *Metrics {
	var metrics Metrics
	runtime.ReadMemStats(&metrics.std)

	return &metrics
}

func (m *Metrics) UpdateAll(randomValue ...gauge) error {
	runtime.ReadMemStats(&m.std)
	m.pollCount += 1
	if len(randomValue) == 1 {
		m.randomValue = randomValue[0]
	} else if len(randomValue) != 0 {
		return nil
	}

	return nil
}

func (m *Metrics) GetStd() runtime.MemStats {
	return m.std
}

func (m *Metrics) GetPollCount() counter {
	return m.pollCount
}

func (m *Metrics) GetRandomValue() gauge {
	return m.randomValue
}
