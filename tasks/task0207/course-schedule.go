package task0207

// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
// You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must
// take course bi first if you want to take course ai.
// For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
// Return true if you can finish all courses. Otherwise, return false.
//
// Constraints:
//   1 <= numCourses <= 2000
//   0 <= prerequisites.length <= 5000
//   prerequisites[i].length == 2
//   0 <= ai, bi < numCourses
//   All the pairs prerequisites[i] are unique.
//
// Results:
//   Runtime: 6 ms, faster than 97.12% of Go online submissions for Course Schedule.
//   Memory Usage: 6.7 MB, less than 35.07% of Go online submissions for Course Schedule.
func canFinish(numCourses int, prerequisites [][]int) bool {
	childCoursesCount := make([]int, numCourses)
	parentCourses := make(map[int][]int, 0)

	var p0 int
	var p1 int
	for _, p := range prerequisites {
		p0, p1 = p[0], p[1]
		if p0 == p1 {
			return false
		}
		if _, ok := parentCourses[p1]; !ok {
			parentCourses[p1] = []int{}
		}
		parentCourses[p1] = append(parentCourses[p1], p0)
		childCoursesCount[p0]++
	}

	// Add to queue all end courses
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if childCoursesCount[i] == 0 {
			queue = append(queue, i)
		}
	}

	visitedCount := 0
	for len(queue) > 0 {
		p0 = queue[0]
		queue = queue[1:]
		visitedCount++

		parents := parentCourses[p0]
		if len(parents) > 0 {
			for _, parent := range parents {
				childCoursesCount[parent]--
				if childCoursesCount[parent] == 0 {
					queue = append(queue, parent)
				}
			}
		}
	}
	return visitedCount == numCourses
}
