package avg_test

import (
	"avg"
	"testing"
)

func TestAvg(t *testing.T) {
	avg := avg.Average([]int{4320, 28, 438, 498})
	if avg != (4320+28+438+498)/4 {
		t.Error("incorrect average function")
	}
}

func TestIncorrectAvg(t *testing.T) {
	avg := avg.Average([]int{1, 2, 3, 3, 1})
	if avg == 2 {
		t.Log("correct average function")
	}
}
