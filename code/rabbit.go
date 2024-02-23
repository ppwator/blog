package main

import "fmt"

// 兔子在出生两个月后，就有繁殖能力，一对兔子每个月能生出一对小兔子来。如果所有兔子都不死，那么一年以后可以繁殖多少对兔子？
// 初始为小兔子。1个月后长成大兔子
func main() {
	months := 12
	rabbitPairs := calculateRabbitPairs(months)
	fmt.Printf("%d月后共有 %d 对兔子\n", months, rabbitPairs)
}

func calculateRabbitPairs(months int) int {
	if months <= 1 {
		return 1
	}
	return calculateRabbitPairs(months-1) + calculateRabbitPairs(months-2)
}
