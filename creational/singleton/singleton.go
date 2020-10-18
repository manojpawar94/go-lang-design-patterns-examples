package singleton

import "sync"

//Counter interface to count the steps
type Counter interface {
	Increment() int
}

//counter struct
type counter struct {
	Count int
}

var instance *counter

var once sync.Once

//GetInstance returns the Counter interface instance
func GetInstance() Counter {
	// thread safe and clean code
	once.Do(func() {
		//function guarantees the uniqueness of our instance, our code can now have 100 goroutines or more according to its need, puts them in competition not we will have Thread Safe problems or aggressive checking
		//instance = &counter{count: 0}
		//or
		instance = new(counter) //new keyword create a pointer to an instance of type passed to it
	})
	return instance
}

//GetInstance1 returns the Counter interface instance
func GetInstance1() Counter {
	//not thread safe
	if instance == nil {
		instance = new(counter) //new keyword create a pointer to an instance of type passed to it
	}
	return instance
}

var lock = &sync.Mutex{}

//GetInstance2 returns the Counter interface instance
//problem with this approach is excessive blocking even when it would not be necessary to do so in case the instance has already been created and should simply have returned the singleton instance. If our program is written to become highly concurrent, this can generate a bottleneck, since only one goroutine routine can get the singleton instance at a time, making it our slowest solution.
func GetInstance2() Counter {
	//thread safe
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = new(counter) //new keyword create a pointer to an instance of type passed to it
	}
	return instance
}

func (c *counter) Increment() int {
	println("Inside Increment()")
	c.Count++
	return c.Count
}
