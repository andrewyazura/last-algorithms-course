package binary_search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = i
	}

	cases := []struct {
		in      int
		want    int
		wantErr bool
	}{
		{0, 0, false},
		{1, 1, false},
		{2, 2, false},
		{12, 0, true},
	}

	for _, c := range cases {
		value, err := BinarySearch(arr, c.in)

		if err != nil && !c.wantErr {
			t.Errorf("error occured: %v", err)
		}

		if c.wantErr && err == nil {
			t.Errorf("expected error, got nil")
		}

		if value != c.want {
			t.Errorf("BinarySearch(%d) == (%d, %v), expected (%d, %v)", c.in, value, err, c.want, c.wantErr)
		}
	}
}
