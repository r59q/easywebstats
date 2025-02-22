package datastore

import (
	"r59q.com/easywebstats/internal/concurrent"
	"time"
)

var rateSmoothingFactor = 0.1

type NumberStore interface {
	Set(name string, label string, data float64) float64
	Get(name string, label string) float64
	GetRateEstimate(name string, label string) float64
	GetLabels(name string) map[string]float64
	GetRateEstimates(name string) map[string]float64
	GetNames() []string
	GetMean(name string) float64
}

type numberStore struct {
	numberMap StatMapper[float64]
	rateStore concurrent.Map[float64]
}

func (s *numberStore) Set(name string, label string, data float64) float64 {
	rateEstimate := s.estimateRate(name, label, data)
	s.rateStore.GetOrCreateInnerMap(name).Set(label, rateEstimate)
	newValue := s.numberMap.Set(name, label, data)
	return newValue
}

func (s *numberStore) Get(name string, label string) float64 {
	val, exists := s.numberMap.Get(name, label)
	if !exists {
		return 0
	}
	return val
}

func (s *numberStore) GetLabels(name string) map[string]float64 {
	return s.numberMap.GetLabels(name)
}

func (s *numberStore) GetRateEstimate(name string, label string) float64 {
	rate, exists := s.rateStore.GetOrCreateInnerMap(name).Get(label)
	if !exists {
		return 0
	}
	return rate
}

func (s *numberStore) estimateRate(name string, label string, newValue float64) float64 {
	updated, hasBeenUpdated := s.numberMap.GetLastUpdated(name, label)
	if !hasBeenUpdated {
		return 0
	}
	previousRate := s.GetRateEstimate(name, label)
	prevValue := s.Get(name, label)
	if prevValue == newValue {
		return previousRate
	}
	timeDiffSeconds := (float64(time.Now().Nanosecond()) - float64(updated.Nanosecond())) / 1000 / 1000
	if timeDiffSeconds == 0 {
		return prevValue
	}
	instRate := (newValue - prevValue) / timeDiffSeconds
	return rateSmoothingFactor*instRate + (1-rateSmoothingFactor)*previousRate
}

func (s *numberStore) GetRateEstimates(name string) map[string]float64 {
	return s.rateStore.GetOrCreateInnerMap(name).Values()
}

func (s *numberStore) GetNames() []string {
	return s.numberMap.GetNames()
}

func (s *numberStore) GetMean(name string) float64 {
	labels := s.numberMap.GetLabels(name)
	sum := 0.0
	for _, val := range labels {
		sum += val
	}
	numberOfLabels := float64(len(labels))
	if sum == 0.0 {
		return 0.0
	}
	return sum / numberOfLabels
}

var store = &numberStore{
	numberMap: CreateStatMap[float64](),
	rateStore: concurrent.Map[float64]{},
}

func GetNumberStore() NumberStore {
	return store
}
