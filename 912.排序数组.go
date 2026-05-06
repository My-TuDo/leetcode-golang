/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
	// 种子初始化，确保随机性
	rand.Seed(time.Now().UnixNano())
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	// 1. 随机选择基准点并换到最左边
	randomIndex := rand.Intn(right-left+1) + left
	nums[left], nums[randomIndex] = nums[randomIndex], nums[left]

	pivot := nums[left]
	// 2. 三路划分指针
	// lt 指向小于 pivot 区间的最后一个
	// gt 指向大于 pivot 区间的第一个
	// i 是当前扫描指针
	lt := left
	i := left + 1
	gt := right

	for i <= gt {
		if nums[i] < pivot {
			nums[lt], nums[i] = nums[i], nums[lt]
			lt++
			i++
		} else if nums[i] > pivot {
			nums[gt], nums[i] = nums[i], nums[gt]
			gt--
			// 注意：换回来的数还没看，i 不动
		} else {
			i++
		}
	}

	// 3. 递归解决左边和右边（等于 pivot 的中间部分已经排好了，不用管）
	quickSort(nums, left, lt-1)
	quickSort(nums, gt+1, right)
}

// @lc code=end

