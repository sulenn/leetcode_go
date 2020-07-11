package TweetCounts

import (
	"errors"
	"math"
)

type TweetCounts struct {
	nametoTime map[string]*NametoTime
}

type NametoTime struct { // 每个人对应一个 NameTime
	prev *Time // 头指针
	tail *Time // 尾指针
}

type Time struct { //链表结构
	time int
	next *Time
}

func Constructor() TweetCounts {
	newTweetCounts := *new(TweetCounts)
	newTweetCounts.nametoTime = make(map[string]*NametoTime)
	return newTweetCounts
}

//记录新发表的 Tweet
func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	if _, ok := this.nametoTime[tweetName]; ok { // name 已存在
		if time >= this.nametoTime[tweetName].tail.time { // 直接从尾部插入
			this.nametoTime[tweetName].tail.next = &Time{time: time}
			this.nametoTime[tweetName].tail = this.nametoTime[tweetName].tail.next
		} else { // 从头开始找点插入
			tempPrev := this.nametoTime[tweetName].prev
			for tempPrev.next != nil {
				if tempPrev.next.time >= time {
					newTime := &Time{time: time}
					newTime.next = tempPrev.next
					tempPrev.next = newTime
					break
				} else {
					tempPrev = tempPrev.next
				}
			}
		}
	} else { // 新 name
		nametoTime := new(NametoTime)
		nametoTime.prev = new(Time)
		nametoTime.tail = &Time{time: time}
		nametoTime.prev.next = nametoTime.tail
		this.nametoTime[tweetName] = nametoTime
	}
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	var result []int
	time := 0
	if freq == "minute" {
		time = 60
	} else if freq == "hour" {
		time = 3600
	} else if freq == "day" {
		time = 86400
	} else { // 异常，freq 不存在
		panic(errors.New("传入的时间频率错误！"))
	}
	times := int(math.Ceil((float64(endTime-startTime+1) / float64(time)))) // 某个频率的次数。如分钟，五次
	result = make([]int, times)
	if _, ok := this.nametoTime[tweetName]; !ok { // 异常，tweetName不存在
		panic(errors.New("传入的人暂未发表 tweet！"))
	}
	nametoTime := this.nametoTime[tweetName]
	pointer := nametoTime.prev.next
	for i := 0; i < times; i++ {
		result[i] = 0
		if endTime-startTime+1 >= time {
			for pointer != nil && pointer.time <= startTime+time-1 {
				// 注意:将 pointer.time >= startTime 判断放在 for 循环里面，而不和 for 判断条件同级是因为
				//pointer.time < startTime 不代表之后所有的 pointer.time 都小于 startTime
				if pointer.time >= startTime {
					result[i]++
				}
				pointer = pointer.next
			}
			startTime = startTime + time
		} else {
			for pointer != nil && pointer.time <= endTime {
				if pointer.time >= startTime { //同上
					result[i]++
				}
				pointer = pointer.next
			}
		}
	}
	return result
}

/**
 * Your TweetCounts object will be instantiated and called as such:
 * obj := Constructor();
 * obj.RecordTweet(tweetName,time);
 * param_2 := obj.GetTweetCountsPerFrequency(freq,tweetName,startTime,endTime);
 */

//func main() {
//	obj := Constructor();
//	obj.RecordTweet("tweet3",0)
//	obj.RecordTweet("tweet3",60)
//	obj.RecordTweet("tweet3",10)
//	fmt.Println(obj.GetTweetCountsPerFrequency("minute", "tweet3", 0, 59))
//	fmt.Println(obj.GetTweetCountsPerFrequency("minute", "tweet3", 0, 60))
//	obj.RecordTweet("tweet3",120)
//	fmt.Println(obj.GetTweetCountsPerFrequency("hour", "tweet3", 0,210))
//}
