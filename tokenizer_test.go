package main

import (
	"fmt"
	"testing"
)

func TestGetInteger(t *testing.T) {
	var tt = []struct {
		Input         string
		ExpectedToken TokenType
	}{
		{"1234;", Constant},
		{"0 ", Constant},
		{"0xEF ", Constant},
		{".1234 ", Constant},
		{"54u ", Constant},
		{"", Constant},
	}

	for _, tt := range tt {
		t.Run(tt.Input, func(t *testing.T) {
			t.Parallel()
			actualToken, i := getInteger("1234", 0)
			fmt.Println(i)
			//assert.Equal(actualToken, tt.ExpectedToken)
			if actualToken != tt.ExpectedToken {
				t.Errorf("Expected: %v, got: %v", tt.ExpectedToken, actualToken)
			}
		})
	}
}
