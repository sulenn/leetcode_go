package main

type Time struct {
	Timestamp int
	Value int
}

func longestSubArr(arr []*Time, threshhold int) int {
	longestTime := 0
	var item *Time = nil
	for _, v := range arr {
		if v.Value > threshhold {
			if item == nil {
				item = v
			} else {
				if longestTime < v.Timestamp - item.Timestamp {
					longestTime = v.Timestamp - item.Timestamp
				}
			}
		} else {
			item = nil
		}
	}
	return longestTime
}

