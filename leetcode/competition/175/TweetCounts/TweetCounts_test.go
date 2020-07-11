package TweetCounts

import (
	"testing"
)

//func Test_TweetCounts_1(t *testing.T)  {
//	obj := Constructor()
//	obj.RecordTweet("tweet3",0)
//	obj.RecordTweet("tweet3",60)
//	obj.RecordTweet("tweet3",10)
//	temp := []int {}
//
//	temp = obj.GetTweetCountsPerFrequency("minute", "tweet3", 0, 59)
//	if temp[0] != 2 {
//		t.Error("fail")
//	}
//
//	temp = obj.GetTweetCountsPerFrequency("minute", "tweet3", 0, 60)
//	if temp[0] != 2 && temp[1] != 1 {
//		t.Error("fail")
//	}
//
//	obj.RecordTweet("tweet3",120)
//	temp = obj.GetTweetCountsPerFrequency("hour", "tweet3", 0,210)
//	if temp[0] != 4 {
//		t.Error("fail")
//	}
//}

func Test_TweetCounts_2(t *testing.T) {
	obj := Constructor()
	obj.RecordTweet("tweet0", 43)
	obj.RecordTweet("tweet1", 65)
	obj.RecordTweet("tweet2", 39)
	obj.RecordTweet("tweet3", 16)
	obj.RecordTweet("tweet4", 82)
	obj.RecordTweet("tweet1", 25)
	obj.RecordTweet("tweet3", 66)
	temp := []int{}

	temp = obj.GetTweetCountsPerFrequency("hour", "tweet3", 43, 1838)
	if temp[0] != 1 {
		t.Error("fail")
	}
}
