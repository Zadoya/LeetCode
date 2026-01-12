package main

import (
	"fmt"
	"time"
	"strings"
)
const (
	dateLayout = "2006-01-02"
)

func loadAbFiles(objects []string) error {


	dates := make(map[string]struct{})

	for i := 0; i < len(objects); i++ {
		splitKey := strings.Split(objects[i], "/")
		if len(splitKey[0]) == 8 {
			name := splitKey[0]
			date := fmt.Sprintf("%s-%s-%s", name[:4], name[4:6], name[6:8])
			_, err := time.Parse(dateLayout, date)
			if err == nil {
				dates[date] = struct{}{}
			}
		}
	}

	var dir string

	if _, ok := dates[todayStr]; ok {
		dir = strings.Join(strings.Split(todayStr, "-"), "")
	} else {
		today, _ := time.Parse(dateLayout, todayStr)
		yesterdayTimeStr := today.Add(time.Hour - 24).Format(dateLayout)
		if _, ok := dates[yesterdayTimeStr]; ok {
			dir = strings.Join(strings.Split(yesterdayTimeStr, "-"), "")
		}
	}

	if dir == "" {
		return fmt.Errorf("directory empty", date)
	}

	if dir != "" {
		fmt.Println(dir)
	}
	return nil
}

func main() {

	for 
	loadAbFiles(objects)
}