package leetcode

import "container/list"

/*
	649. Dota2 Senate
	- thanks to neetcode for explaination, couldn't get this one
	- https://www.youtube.com/watch?v=zZA5KskfMuQ
*/

func predictPartyVictory(senate string) string {
	rQueue := list.New()
	dQueue := list.New()
	for i, c := range senate {
		if c == 'R' {
			rQueue.PushBack(i)
		} else {
			dQueue.PushBack(i)
		}
	}
	for rQueue.Len() > 0 && dQueue.Len() > 0 {
		dTurn, _ := dQueue.Remove(dQueue.Front()).(int)
		rTurn, _ := rQueue.Remove(rQueue.Front()).(int)
		if rTurn < dTurn {
			rQueue.PushBack(dTurn + len(senate))
		} else {
			dQueue.PushBack(rTurn + len(senate))
		}
	}
	if rQueue.Len() > 0 {
		return "Radiant"
	}
	return "Dire"
}
