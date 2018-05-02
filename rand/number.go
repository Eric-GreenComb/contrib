package rand

import (
	"fmt"
	"math/rand"
	"time"
)

// Gen6Number gen 6 rand number
func Gen6Number() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}
