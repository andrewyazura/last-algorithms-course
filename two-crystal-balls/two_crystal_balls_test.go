package two_crystal_balls

import "testing"

func TestTwoCrystalBalls(t *testing.T) {
	breaks := []bool{false, false, false, false, false, true, true, true, true, true}

	i, err := two_crystal_balls(breaks)

	if err != nil || i != 5 {
		t.Errorf("two_crystal_balls(%v) == (%d, %v), expected 5, nil", breaks, i, err)
	}
}
