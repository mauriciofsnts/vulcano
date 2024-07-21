package utils

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"os"
	"time"
)

type Holiday struct {
	Name string `json:"name"`
	Date string `json:"date"`
	Type string `json:"type"`
}

func GetNextHoliday() (string, error) {
	holidayList, err := os.Open("./internal/providers/utils/holidays.json")

	if err != nil {
		slog.Error("Error opening holidays file", err)
		return "", err
	}

	defer holidayList.Close()

	bytes, err := io.ReadAll(holidayList)

	if err != nil {
		slog.Error("Error reading holidays file", err)
		return "", err
	}

	var holidays []Holiday

	if err := json.Unmarshal(bytes, &holidays); err != nil {
		slog.Error("Error parsing holidays file", err)
		return "", err
	}

	today := time.Now()

	for _, holiday := range holidays {
		date, err := time.Parse("02/01", holiday.Date)

		if err != nil {
			slog.Error("Error parsing date", err)
			return "", err
		}

		date = date.AddDate(time.Now().Year(), 0, 0)

		if date.After(today) || date.Equal(today) {
			if date.Equal(today) {
				return "Today is " + holiday.Name + "." + "\n" + "Date: " + holiday.Date + "." + "\n" + "Type: " + holiday.Type + ".", nil
			} else {
				return "The next holiday is " + holiday.Name + "." + "\n" + "Date: " + holiday.Date + "." + "\n" + "Type: " + holiday.Type + ".", nil
			}
		}
	}

	return "", errors.New("no holidays found")
}
