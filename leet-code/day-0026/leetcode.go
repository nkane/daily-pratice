package leetcode

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
