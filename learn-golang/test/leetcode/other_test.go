package leetcode

/**
 * 快速排序
 */
// func Test_Quick_Sort(t *testing.T) {
// 	defer timeCost()()
// 	arr := []int{5, 2, 3, 1}
// 	left, right := 0, len(arr)-1
// 	quickSort(left, right, &arr)
// 	fmt.Println(arr)
// }

// func quickSort(left, right int, arr *[]int) {
// 	if left >= right {
// 		return
// 	}
// 	mid := left + (left+right+1)/2

// 	quickSort(left, mid-1, arr)
// 	quickSort(mid+1, right, arr)
// }

// // 第一种写法
// func quickSort(values []int, left, right int) {
// 	temp := values[left]
// 	p := left
// 	i, j := left, right

// 	for i <= j {
// 		for j >= p && values[j] >= temp {
// 			j--
// 		}
// 		if j >= p {
// 			values[p] = values[j]
// 			p = j
// 		}

// 		for i <= p && values[i] <= temp {
// 			i++
// 		}
// 		if i <= p {
// 			values[p] = values[i]
// 			p = i
// 		}
// 	}
// 	values[p] = temp
// 	if p-left > 1 {
// 		quickSort(values, left, p-1)
// 	}
// 	if right-p > 1 {
// 		quickSort(values, p+1, right)
// 	}
// }

// func QuickSort(values []int) {
// 	if len(values) <= 1 {
// 		return
// 	}
// 	quickSort(values, 0, len(values)-1)
// }

// // 第二种写法
// func Quick2Sort(values []int) {
// 	if len(values) <= 1 {
// 		return
// 	}
// 	mid, i := values[0], 1
// 	head, tail := 0, len(values)-1
// 	for head < tail {
// 		fmt.Println(values)
// 		if values[i] > mid {
// 			values[i], values[tail] = values[tail], values[i]
// 			tail--
// 		} else {
// 			values[i], values[head] = values[head], values[i]
// 			head++
// 			i++
// 		}
// 	}
// 	values[head] = mid
// 	Quick2Sort(values[:head])
// 	Quick2Sort(values[head+1:])
// }

// // 第三种写法
// func Quick3Sort(a []int, left int, right int) {

// 	if left >= right {
// 		return
// 	}

// 	explodeIndex := left

// 	for i := left + 1; i <= right; i++ {

// 		if a[left] >= a[i] {

// 			//分割位定位++
// 			explodeIndex++
// 			a[i], a[explodeIndex] = a[explodeIndex], a[i]

// 		}

// 	}

// 	//起始位和分割位
// 	a[left], a[explodeIndex] = a[explodeIndex], a[left]

// 	Quick3Sort(a, left, explodeIndex-1)
// 	Quick3Sort(a, explodeIndex+1, right)

// }

func edit_distance(s string, t string) int {
	lenS, lenT := len(s), len(t)
	var dp = make([][]int, lenT+1)
	for i := 0; i <= lenS; i++ {
		dp[i] = make([]int, lenS+1)
	}
	// 空字符串编辑为空字符串的编辑代价为0
	dp[0][0] = 0
	// dp[i][0]表示编辑为空串的代价，即为将所有字符删除的代价
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	// dp[0][i]表示将空串编辑为str2[:i]的代价，即插入字符的代价
	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}
	// 下面是动态规划的主方法
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] // 如果 word1[i-1] 与 word2[i-1]相等
			} else {
				dp[i][j] = dp[i-1][j-1] + 1 // 替换代价
			}
			dp[i][j] = min(dp[i][j], dp[i-1][j]+1) // 删除代价
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
