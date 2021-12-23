package task0210

// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
// You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you
// must take course bi first if you want to take course ai.
// For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
// Return the ordering of courses you should take to finish all courses. If there are many
// valid answers, return any of them. If it is impossible to finish all courses, return an empty array.
//
// Constraints:
//   1 <= numCourses <= 2000
//   0 <= prerequisites.length <= numCourses * (numCourses - 1)
//   prerequisites[i].length == 2
//   0 <= ai, bi < numCourses
//   ai != bi
//   All the pairs [ai, bi] are distinct.
//
// Results:
//   Runtime: 8 ms, faster than 96.37% of Go online submissions for Course Schedule II.
//   Memory Usage: 6.4 MB, less than 68.32% of Go online submissions for Course Schedule II.
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make(map[int][]int)
	incomings := make([]int, numCourses)
	for _, p := range prerequisites {
		in, out := p[0], p[1]
		graph[out] = append(graph[out], in)
		incomings[in] += 1
	}

	// Find all starting points: the points that doesn't have incoming nodes
	var starts []int
	for i, incoming := range incomings {
		if incoming == 0 {
			starts = append(starts, i)
		}
	}

	var path []int
	var visits int
	for _, start := range starts {
		// Iterate through all starting points and start new queue for each of them
		queue := []int{start}
		path = append(path, start)
		visits++
		for len(queue) > 0 {
			currVal := queue[0]
			queue = queue[1:]
			for _, neighbor := range graph[currVal] {
				incomings[neighbor]--
				if incomings[neighbor] == 0 {
					// Put into the queue on the last incoming
					queue = append(queue, neighbor)
					path = append(path, neighbor)
					visits++
				}
			}
		}
	}
	if visits == numCourses {
		return path
	}
	return []int{}
}
