package models

import (
	"fmt"
	"strconv"
	"strings"
)

// Estructura que representa los datos del API
type Stock struct {
	Ticker        string `json:"ticker"`
	TargetFromRaw string `json:"target_from" gorm:"-"`
	TargetFrom    float64
	TargetToRaw   string `json:"target_to" gorm:"-"`
	TargetTo      float64
	Company       string `json:"company"`
	Action        string `json:"action"`
	Brokerage     string `json:"brokerage"`
	RatingFrom    string `json:"rating_from"`
	RatingTo      string `json:"rating_to"`
	Time          string `json:"time"`
}

func ProcesarTargets(i *Stock) error {
	convert := func(s string) (float64, error) {
		s = strings.TrimSpace(s)
		s = strings.ReplaceAll(s, "$", "")
		s = strings.ReplaceAll(s, ",", "")
		if s == "" {
			return 0, nil
		}
		f, err := strconv.ParseFloat(s, 64)
		return float64(f), err
	}

	var err error

	if i.TargetFrom, err = convert(i.TargetFromRaw); err != nil {
		return fmt.Errorf("error al convertir TargetFrom (%s): %v", i.TargetFromRaw, err)
	}

	if i.TargetTo, err = convert(i.TargetToRaw); err != nil {
		return fmt.Errorf("error al convertir TargetTo (%s): %v", i.TargetToRaw, err)
	}

	return nil
}
