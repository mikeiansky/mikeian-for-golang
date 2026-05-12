package test

import "testing"

func Test_Hello2(t *testing.T) {
	//t.Errorf("print error start")
	if _, err := hello("good"); err != nil {
		t.Fatal(err.Error())
	}
	//t.Errorf("print error complete")
}
