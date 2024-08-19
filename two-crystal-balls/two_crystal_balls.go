package two_crystal_balls

import (
	"errors"
	"fmt"
	"math"
)

func two_crystal_balls(breaks []bool) (int, error) {
	n := len(breaks)
	step := int(math.Floor(math.Sqrt(float64(n))))

	j := step
	for ; j < n; j += step {
		if breaks[j] {
			break
		}
	}

	for i := j - step; i < j; i++ {
		if breaks[i] {
			return i, nil
		}
	}

	fmt.Printf("%d", j)
	return 0, errors.New("no breaking point in the list")
}
