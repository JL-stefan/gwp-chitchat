package data

import (
	"testing"
	"time"
)

func TestThreads(t *testing.T) {
	res, _ := Threads()
	t.Log(res)
	t.Log(res[0].CreatedAt.Unix())
}

func TestTimestampType(t *testing.T) {
	timestamp := time.Now().Unix()
	t.Logf("timestamp type is %T", timestamp)
	t.Log(time.Now())
}
