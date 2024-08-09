package leetcode

/*
	841: Keys and Rooms
	There are n rooms labeled from 0 to n-1 and all the rooms are locked expect for the room 0. Your goal is to visit
	all the rooms. However, you cannot enter a locked room without having its key.

	When you visit a room, you may find a set of distinct keys in it. Each Key has a number on it, denoting which
	room it unlocks, and you can take all of them with you to unlock the other rooms.

	Given an array rooms where room[i] is the set of keys that you can obtain if you visited room i, return true if
	you can visit all the rooms, or false otherwise.
*/

func canVisitAllRooms(rooms [][]int) bool {
	visited := map[int]struct{}{}
	visited[0] = struct{}{}
	q := []int{0}

	for len(q) > 0 {
		// check all adj list
		idx := q[0]
		q = q[1:]
		keys := rooms[idx]
		for _, room := range keys {
			if _, v := visited[room]; !v {
				visited[room] = struct{}{}
				q = append(q, room)
			}
		}
	}

	return len(visited) == len(rooms)
}
