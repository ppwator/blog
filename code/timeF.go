package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func ToInt32(v string) int32 {
	n, _ := strconv.Atoi(v)
	return int32(n)
}
func main() {
	hhmi := int16(ToInt32(time.Now().Format(`1504`)))
	fmt.Println(hhmi)
	start := time.Now()
	reset := make(chan bool)
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		reset <- true
	})
	for time.Since(start) < 5*time.Second {
		<-reset
		t.Reset(randomDuration())
	}

}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
