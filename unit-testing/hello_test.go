package main

import (
	"testing"
)

func TestFirst(t *testing.T) {

	//verify the Greeting Func
	actualResult := Greetings("")
	expectedResult := "Hello Guys Good morning1111"

	if actualResult != expectedResult {
		t.Errorf("Error in greetings  %v expected-actucal is %v", expectedResult, actualResult)
	}
}

// func TestSecond(t *testing.T) {

// 	//verify the Greeting Func
// 	actualResult := Greetings("Kiran")
// 	expectedResult := "Helloo Kiran"

// 	if actualResult != expectedResult {
// 		t.Errorf("Error in greetings  %v expected-actucal is %v", expectedResult, actualResult)
// 	}
// }
