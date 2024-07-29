package leetcode

/*
933. Number of Recent Calls
You have a RecentCounter class which counts the number of recent requests within a certain time frame.

Implement the RecentCounter class:

RecentCounter() Initializes the counter with zero recent requests.
int ping(int t) Adds a new request at time t, where t represents some time in milliseconds, and returns the number of requests that has happened in the past 3000 milliseconds (including the new request). Specifically, return the number of requests that have happened in the inclusive range [t - 3000, t].
It is guaranteed that every call to ping uses a strictly larger value of t than the previous call.
*/

// NOTE(nick): not solved, needed help

type RecentCounter struct {
	Pings []int
}

func Constructor() RecentCounter {
	pings := make([]int, 0)
	return RecentCounter{pings}
}

func (this *RecentCounter) Ping(t int) int {
	for len(this.Pings) > 0 && (t-this.Pings[0]) > 3000 {
		this.Pings = this.Pings[1:]
	}
	this.Pings = append(this.Pings, t)
	return len(this.Pings)
}
