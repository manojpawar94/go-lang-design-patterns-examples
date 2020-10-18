package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter := GetInstance()

	if counter == nil {
		//Test of acceptance criterai 1 is failed
		t.Error("expected pointer to Counter after calling GetInstance(), not nil.")
	} else {
		t.Log("GetInstance() has return pointer to Counter as expected.")
	}

	//expectedCounter := counter1

	currentCount := counter.Increment()

	if currentCount != 1 {
		t.Errorf("after calling first time increment method current count must be 1, but it is %d", currentCount)
	}
}
