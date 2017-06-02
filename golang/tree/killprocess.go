/* Given n processes, each process has a unique PID (process id) and its PPID (parent process id).
 *
 * Each process only has one parent process, but may have one or more children processes. This is just like a tree structure. Only one process has PPID that is 0, which means this process has no parent process. All the PIDs will be distinct positive integers.
 *
 * We use two list of integers to represent a list of processes, where the first list contains PID for each process and the second list contains the corresponding PPID.
 *
 * Now given the two lists, and a PID representing a process you want to kill, return a list of PIDs of processes that will be killed in the end. You should assume that when a process is killed, all its children processes will be killed. No order is required for the final answer.
 *
 * Example 1:
 * Input:
 * pid =  [1, 3, 10, 5]
 * ppid = [3, 0, 5, 3]
 * kill = 5
 * Output: [5,10]
 * Explanation:
 *            3
 *          /   \
 *         1     5
 *              /
 *             10
 * Kill 5 will also kill 10.
 * Note:
 * The given kill id is guaranteed to be one of the given PIDs.
 * n >= 1.
 * Subscribe to see which companies asked this question. */

package main

import "fmt"

func killProcess(pid []int, ppid []int, kill int) []int {
	parent_tree := make(map[int][]int)
	visited := make(map[int]int)

	for i, p := range ppid {
		parent_tree[p] = append(parent_tree[p], pid[i])
	}
	var result []int
	killAll(visited, parent_tree, kill, &result)
	return result
}

func killAll(visited map[int]int, parent_tree map[int][]int, kill int, result *[]int) {
	*result = append(*result, kill)
	visited[kill] = 1
	for _, p := range parent_tree[kill] {
		if visited[p] == 0 {
			visited[p] = 1
			killAll(visited, parent_tree, p, result)
		}
	}
}

func main() {
	pid := make([]int, 50000)
	ppid := make([]int, 50000)
	for i := 0; i < 50000; i++ {
		pid[i] = i + 1
		ppid[i] = 1
	}
	ppid[0] = 0
	kill := 1
	pid = []int{1, 2, 3, 4}
	ppid = []int{0, 1, 2, 3}
	kill = 1
	fmt.Println(killProcess(pid, ppid, kill))
}
