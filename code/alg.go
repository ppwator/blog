package main

import (
	"container/list"
	"fmt"
	"time"
)

/* 递归 */
func recur(n int) int {
	// 终止条件
	if n == 1 {
		return 1
	}
	// 递：递归调用
	res := recur(n - 1)
	// 归：返回结果
	return n + res
}

/* 尾递归 */
func tailRecur(n int, res int) int {
	// 终止条件
	if n == 0 {
		return res
	}
	// 尾递归调用
	return tailRecur(n-1, res+n)
}

/* 斐波那契数列：递归 */
func fib(n int) int {
	// 终止条件 f(1) = 0, f(2) = 1
	if n == 1 || n == 2 {
		return n - 1
	}
	// 递归调用 f(n) = f(n-1) + f(n-2)
	res := fib(n-1) + fib(n-2)
	// 返回结果 f(n)
	return res
}

func forLoopFibonacci(n int) int {
	if n <= 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	// 初始化前两个Fibonacci数
	prev, curr := 0, 1
	// 迭代n-1次计算Fibonacci数
	for i := 3; i <= n; i++ {
		// 计算下一个Fibonacci数
		next := prev + curr
		// 更新前两个数以进行下一次迭代
		prev, curr = curr, next
	}
	// 返回最终的Fibonacci数
	return curr
}

/* 使用迭代模拟递归 */
func forLoopRecur(n int) int {
	// 使用一个显式的栈来模拟系统调用栈
	stack := list.New()
	res := 0
	// 递：递归调用
	for i := n; i > 0; i-- {
		// 通过“入栈操作”模拟“递”
		stack.PushBack(i)
	}
	// 归：返回结果
	for stack.Len() != 0 {
		// 通过“出栈操作”模拟“归”
		res += stack.Back().Value.(int)
		stack.Remove(stack.Back())
	}
	// res = 1+2+3+...+n
	return res
}

// 动态规划
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func climbStairs(n int) int {
	// 创建一个长度为n+1的数组，用来存储每个台阶对应的爬法数
	dp := make([]int, n+1)

	// 初始情况：0阶楼梯有一种爬法，1阶楼梯有一种爬法
	dp[0], dp[1] = 1, 1

	// 使用动态规划递推每个台阶的爬法数
	for i := 2; i <= n; i++ {
		// 当前台阶的爬法数等于前一阶台阶的爬法数加上前两阶台阶的爬法数
		dp[i] = dp[i-1] + dp[i-2]
	}

	// 返回爬到楼顶的总爬法数
	return dp[n]
}

// 空间复杂度：O(1)
func climbStairs1(n int) int {
	res := 1
	n1 := 1
	n2 := 1

	// 使用动态规划递推每个台阶的爬法数
	for i := 2; i <= n; i++ {
		res = n1 + n2
		n1 = n2
		n2 = res
	}

	// 返回爬到楼顶的总爬法数
	return res
}

// 鸡翁一，值钱五；鸡母一，值钱三；鸡雏三，值钱一；百钱买百鸡，则翁、母、雏各几何？
func findChicken() {
	for x := 0; x <= 20; x++ {
		for y := 0; y <= 33; y++ {
			z := 100 - x - y
			if 5*x+3*y+z/3 == 100 && z%3 == 0 {
				fmt.Printf("公鸡：%d只，母鸡：%d只，小鸡：%d只\n", x, y, z)
			}
		}
	}
}

func main() {
	findChicken()

	//动态规划爬楼梯
	fmt.Println("-----", climbStairs(5))
	fmt.Println(climbStairs1(5))

	fmt.Println(fib(10))
	fmt.Println(forLoopFibonacci(2))

	// 递归求和
	start := time.Now().UnixNano()
	fmt.Println("rec", recur(100))
	t1 := time.Now().UnixNano()
	fmt.Println(t1 - start)
	fmt.Println("rec1", tailRecur(100, 0))
	fmt.Println(time.Now().UnixNano() - t1)
	//使用迭代模拟递归求和
	fmt.Println("===", forLoopRecur(3))
}
