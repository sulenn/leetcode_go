package main

import (
	"fmt"
	"sort"
)

var count int64 = 1 // 发帖子时累加，表示帖子的先后关系

type Twitter struct {
	tweets          map[int][]Tweet // UserId 发表的帖子
	followRalations map[int][]int   // userId 关注的人
}

type Tweet struct {
	tweetId int
	time    int64
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	count = 1
	return Twitter{make(map[int][]Tweet), make(map[int][]int)}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.tweets[userId]; ok {
		tweets := this.tweets[userId]
		for i := 0; i < len(tweets); i++ {
			if tweets[i].tweetId == tweetId { // 该帖子已发表过，更新发表时间
				tweets[i].time = count
				count++
				return
			}
		}
		this.tweets[userId] = append(this.tweets[userId], Tweet{tweetId: tweetId, time: count}) // 发表新帖子
	} else {
		this.tweets[userId] = []Tweet{{tweetId: tweetId, time: count}} // 发表第一个帖子
	}
	count++
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	tweets := this.tweets[userId]                        // 该用户发表的帖子
	if userIds, ok := this.followRalations[userId]; ok { // 该用户关注的人发表的帖子
		for i := 0; i < len(userIds); i++ {
			tweets = append(tweets, this.tweets[userIds[i]]...)
		}
	}
	sort.Slice(tweets, func(i, j int) bool { // 按时间从大到小排序
		return tweets[i].time > tweets[j].time
	})
	if len(tweets) > 10 {
		result := make([]int, 10)
		for i := 0; i < 10; i++ {
			result[i] = tweets[i].tweetId
		}
		return result
	} else {
		result := make([]int, len(tweets))
		for i := 0; i < len(tweets); i++ {
			result[i] = tweets[i].tweetId
		}
		return result
	}
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	} // 自己关注自己
	if followeeIds, ok := this.followRalations[followerId]; ok { // 该用户已关注了其它人
		for i := 0; i < len(followeeIds); i++ {
			if followeeIds[i] == followeeId { // 该用户已关注
				return
			}
		}
		this.followRalations[followerId] = append(this.followRalations[followerId], followeeId) // 添加新关注用户
	} else { // 该用户第一次关注他人
		this.followRalations[followerId] = []int{followeeId}
	}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	} // 自己取消关注自己
	if followeeIds, ok := this.followRalations[followerId]; ok {
		for i := 0; i < len(followeeIds); i++ {
			if followeeIds[i] == followeeId {
				this.followRalations[followerId] = append(followeeIds[:i], followeeIds[i+1:]...)
			}
		}
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

func main() {
	twitter := Constructor()
	twitter.PostTweet(1, 5)

	// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
	fmt.Println(twitter.GetNewsFeed(1))

	// 用户1关注了用户2.
	twitter.Follow(1, 2)

	// 用户2发送了一个新推文 (推文id = 6).
	twitter.PostTweet(2, 6)

	// 用户1的获取推文应当返回一个列表，其中包含两个推文，id分别为 -> [6, 5].
	// 推文id6应当在推文id5之前，因为它是在5之后发送的.
	fmt.Println(twitter.GetNewsFeed(1))

	// 用户1取消关注了用户2.
	twitter.Unfollow(1, 2)

	// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
	// 因为用户1已经不再关注用户2.
	fmt.Println(twitter.GetNewsFeed(1))
	fmt.Println(twitter.GetNewsFeed(6))
}
