package helpers

import (
	"math/rand"
	"sync"
	"time"
)

var once sync.Once

func RandomIntArray(n int) []int {
	once.Do(func() {
		rand.Seed(time.Now().Unix())
	})

	var a []int
	for i := 0; i < n; i++ {
		a = append(a, rand.Int())
	}
	return a
}
