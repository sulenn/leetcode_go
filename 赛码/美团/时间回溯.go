package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func main() {
	var week int
	var time string
	var hour int
	var mi int
	var minutes int
	for {
		_, err := fmt.Scan(&week, &time, &minutes)
		if err == io.EOF {
			break
		}
		timeArr := strings.Split(time, ":")
		hour, err = strconv.Atoi(timeArr[0])
		mi, err = strconv.Atoi(timeArr[1])
		if mi == 60 {
			mi = 0
			if hour+1 > 23 {
				hour = (hour+1)%24
				if week+1 > 7 {
					week = (week+1)%7
				}
			}
		}
		if hour > 23 {
			hour = (hour)%24
			if week+1 > 7 {
				week = (week+1)%7
			}
		}
		week, time = solution(week, hour, mi, minutes)
		fmt.Println(week)
		fmt.Println(time)
	}
}

func solution(week, hour, mi, minutes int) (int, string) {
	//return 0, "0"
	newHour := minutes / 60
	newMi := minutes % 60
	newWeek := newHour / 24
	newHour = newHour % 24
	if mi >= newMi {  // 分钟
		mi -= newMi
	} else {
		mi += 60 - newMi
		hour--
		//if hour > 0 {
		//	hour--
		//} else {
		//	hour = 23
		//	if week == 1 {
		//		week = 7
		//	} else {
		//		week--
		//	}
		//}
	}
	if hour >= newHour {  // 小时
		hour -= newHour
	} else {
		hour += 24 - newHour
		week--
		//if week == 1 {
		//	week = 7
		//} else {
		//	week--
		//}
	}
	if week > newWeek {  // 星期几
		week -= newWeek
	} else {
		week += 7-newWeek%7
	}
	var backMin string
	var backHour string
	if hour < 10 {
		backHour = "0"+strconv.Itoa(hour)
	} else {
		backHour = strconv.Itoa(hour)
	}

	if mi < 10 {
		backMin = "0"+strconv.Itoa(mi)
	} else {
		backMin = strconv.Itoa(mi)
	}
	return week, backHour+":"+backMin
}
