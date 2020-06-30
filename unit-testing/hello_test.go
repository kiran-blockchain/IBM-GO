package main

import (
	"testing"
)

func TestXYZ(t *testing.T) {

	//verify the Greeting Func
	actualResult := Greetings("")
	expectedResult := "Hello Guys Good morning1111"

	if actualResult != expectedResult {
		t.Errorf("Error in greetings  %v expected-actucal is %v", expectedResult, actualResult)
	}
}
