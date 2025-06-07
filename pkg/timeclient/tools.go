package timeclient

import (
	"math/rand"
	"time"
)


var src = rand.NewSource(time.Now().UnixNano())
func gernerateString(n int) string{
	b := make([]byte, n)
	for i := range b {
		b[i] = token[src.Int63() % int64(len(token))]
	}
	return string(b)
}