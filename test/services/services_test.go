package test

import (
	"testing"

	"github.com/gleison/peculi/internal/service"
	test "github.com/gleison/peculi/test/data"
)

func TestAveragePrice(t *testing.T) {
	var expected float32 = 57.33
	got := service.AveragePrice(test.Trades)
	if expected != got {
		t.Errorf("got %f, expected %f", got, expected)
	}
}
