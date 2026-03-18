package calculations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	type args struct {
		a float64
		b float64
	}

	tests := []struct {
		name   string
		args   args
		expect float64
	}{
		{name: "should return sum", args: args{a: 5, b: 7}, expect: 12.0},
	}

	for _, testCase := range tests {
		t.Run("should return 6", func(t *testing.T) {
			res := Add(testCase.args.a, testCase.args.b)
			assert.Equal(t, testCase.expect, res)
		})
	}
}
