package main

import "fmt"

/*
给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
示例:
输入: [2,3,1,1,4,1,4,2,3,1]
输出: 3
解释: 跳到最后一个位置的最小跳跃数是 2。
从下标为 0 跳到下标为 1 的位置，跳1步，然后跳3步到达数组的最后一个位置。
*/
/*
贪心还是动态规划呢？
其实就是把它分为很多步，从0的位置开始，开始的最大步数是2，所以就定义了一个区间[1,3),就是两种走法，
走一步到1的位置或者走两步到2的位置，我们要计算的数字就是这个数本身的位置加上这个数前进的值就是这个
数的值谁更大，因为更大代表走更远，就是步数更少，nums[1]+1=1+3>nums[2]+2=1+2,所以从1开始走三步
走到四的位置是最远的，我们第二轮从nums[3]开始，区间在[3,5),然后根据后面的值计算3+1<4+4,所以第三
轮就是到nums[5]开始区间在[5,9),这个区间最大的是8+3,所以下一轮就到[9,12),这个时候12已经大于nums
的长度10了，所以无需再往下算了，得出来了步数
O(N) O(1)
*/
func jump(nums []int) int {
	start, end, bj := 0, 1, 0
	for end < len(nums) {
		maxB := 0
		for i := start; i < end; i++ {
			maxB = max(maxB, nums[i]+i)
		}
		start = end
		end = maxB + 1
		bj++
	}
	return bj
}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	a := []int{2, 3, 1, 1, 4, 3, 3, 1, 1, 1}
	b := jump(a)
	fmt.Print(b)
}
