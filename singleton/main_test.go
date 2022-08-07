package main

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		t.Error("A new connection object must have been made")
	}
	expectedCounter := counter1

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		//Test 2 failed
		t.Error("Singleton instances must be different")
	}

	currentCount := counter2.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling 'AddOne' using the second counter the current count must be 1 but was %d\n", currentCount)
	}

}
