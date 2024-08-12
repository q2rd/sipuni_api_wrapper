package sipuni_api_wrapper

import (
	"strings"
)

const (
	dateIndex      = 2
	nameIndex      = 4
	fromPhoneIndex = 6
	toPhoneIndex   = 7
	minFieldCount  = 8
	status         = 20
	callStatus     = "Не отвечен"
	recallStatus   = "Оператор не перезвонил"
)

type Record struct {
	Name      string `json:"operator_name"`
	FromPhone string `json:"from_phone"`
	ToPhone   string `json:"to_phone"`
	Date      string `json:"date"`
	Status    string `json:"status"`
}

func parseCSVResponse(byteArray []byte) ([]Record, error) {
	dataStr := string(byteArray)
	var sortedList []string

	lines := strings.Split(dataStr, "\n")
	for _, line := range lines[1:] {
		if strings.Contains(line, recallStatus) && strings.Contains(line, callStatus) {
			sortedList = append(sortedList, line)
		}
	}
	var result []Record
	for _, line := range sortedList {
		fields := strings.Split(line, ";")
		if len(fields) < minFieldCount {
			continue
		}
		if fields[14] != "" {
			continue
		}
		record := Record{
			Date:      fields[dateIndex],
			Name:      fields[nameIndex],
			FromPhone: fields[fromPhoneIndex],
			ToPhone:   fields[toPhoneIndex],
			Status:    fields[status],
		}
		result = append(result, record)
	}

	return result, nil
}
