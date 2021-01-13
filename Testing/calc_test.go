package calc_test

import (
	"testing"

	calc "github.com/wonos/goCrawler/Testing"
)

func TestSum(t *testing.T) {
	s := calc.Sum(1, 2, 3)

	if s != 6 {
		t.Error("Wrong result")
	}
}
