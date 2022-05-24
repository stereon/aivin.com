package util

import (
	"log"
	"math/rand"
	"time"
)

func RandomString(n int) string  {
	var letstrs = []byte("dsasdwrwefwerwrf3rfwefsfs")
	result := make([]byte,n)
	rand.Seed(time.Now().Unix())
	for i := range result{
		log.Println(len(letstrs))
		result[i] = letstrs[rand.Intn(len(letstrs))]
	}
	return string(result)
}
