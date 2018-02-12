package singletons

import "testing"

func TestGetInstance(t *testing.T) {

	// Check GetInstance method

	counter1 := GetInstance()

	if counter1 == nil {
		t.Error("Singleton instance not returned")
	}

	// Check AddOne method

	currentCount := counter1.AddOne()

	if currentCount != 1 {
		t.Errorf("Count was %d when it should be 1", currentCount)
	}

	// Check that a single instance returned

	expectedCounter := counter1
	counter2 := GetInstance()

	if counter2 != expectedCounter {
		t.Error("Two different instances returned by GetInstance()")
	}

	// Double check AddOne again

	currentCount = counter2.AddOne()

	if currentCount != 2 {
		t.Errorf("Count was %d when it should be 2", currentCount)
	}
}
