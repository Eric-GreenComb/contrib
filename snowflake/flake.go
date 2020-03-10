package snowflake

import (
	"fmt"

	flake "github.com/zheng-ji/goSnowFlake"
)

// GIwFlake GIwFlake
var GIwFlake Flake

// Flake Flake
type Flake struct {
	iw *flake.IdWorker
}

// Init Init
func (f *Flake) Init(workerid int64) error {
	_iw, err := flake.NewIdWorker(workerid)
	if err != nil {
		fmt.Println(err)
		return err
	}
	f.iw = _iw
	return nil
}

// NextID NextID
func (f *Flake) NextID() (int64, error) {
	_id, err := f.iw.NextId()
	if err != nil {
		fmt.Println("NextID error:", err)
		return -1, err
	}
	return _id, nil
}
