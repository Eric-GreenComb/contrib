package uuid

import (
	"fmt"
	flake "github.com/zheng-ji/goSnowFlake"
)

// GetFlakeID GetFlakeID
func GetFlakeID() int64 {
	// Params: Given the workerId, 0 < workerId < 1024
	var id int64
	iw, err := flake.NewIdWorker(1)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	id, err = iw.NextId()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return id
}
