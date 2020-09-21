package main

import (
	"fmt"
	"testing"
)

/*func TestMain(t *testing.T) {
	//t.Errorf("hi")
	//t.Logf("2")
	dat, err := ioutil.ReadFile("files/foo.hpp")
	check(err)
	datStr := string(dat)

	tokenSlice := GetTokens(datStr)
	for _, token := range tokenSlice {
		fmt.Println(token)
		//t.Error(token)
		//t.Log("1")
	}
}*/

/*func TestIntTokens(t *testing.T) {
	t.Error("hmm")
	var tt = []struct {
		Input         string
		ExpectedToken TokenType
	}{
		{"1234;", Constant},
		/*{"0 ", Constant},
		{"0xEF ", Constant},
		{".1234 ", Constant},
		{"54u ", Constant},
	}

	for _, tt := range tt {
		t.Run(tt.Input, func(t *testing.T) {
			//t.Parallel()
			got := GetTokens(tt.Input)
			_, i := getInteger("1234", 0)
			fmt.Println(i)
			if got[0].tokenType != tt.ExpectedToken {
				t.Errorf("Expected: %v, got: %v", tt.ExpectedToken, got)
			}
		})
	}

}*/

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
			if actualToken != tt.ExpectedToken {
				t.Errorf("Expected: %v, got: %v", tt.ExpectedToken, actualToken)
			}
		})
	}
}
