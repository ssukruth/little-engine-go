package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f1() {
	fmt.Println("f1 starts")
	fmt.Println("f1 exits")
}

func f2() {
	fmt.Println("f2 starts")
	fmt.Println("f2 exits")
}

func f1wg(wg *sync.WaitGroup) {
	fmt.Println("f1 starts")
	fmt.Println("f1 exits")
	wg.Done()
}

func runTask(t string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("starting task", t)
	s := rand.Intn(3)
	time.Sleep(time.Duration(s) * time.Second)
	fmt.Println("done with task", t)
}

func runTaskWithChan(t string, c chan string) {
	s := rand.Intn(3)
	time.Sleep(time.Duration(s) * time.Second)
	str := fmt.Sprintf("done with task %s", t)
	c <- str
}

func main() {
	// Concurrency is the first class citizen in go.
	// Go is the first major language released after multicore cpu was released.
	// Concurrency => loading more go routines at a time.
	// Go routines are multiple threads of execution. If one go routine blocks
	// and another is picked up. On a single core cpu, the go routines run
	// sequentially.
	// parallelism => multiple go routines execute at the same time.
	// parallelism requires multi-core cpu

	// Go routine is a lightweight thread of execution and is a key ingredient
	// to achieving concurrency in go.
	// Each goroutine is a function that is capable of running concurrently with
	// other functions. They are small threads which take upto 2kb of stack space
	// compared to threads which take around ~1-2Mb stack space.
	// Goroutines stacks can grow or shrink as needed.
	// Goroutines are cheaper to schedule compared to threads and the go scheduler
	// multiplexes m goroutines on n OS threads, m:n scheduling.
	// Goroutines have no notion of identity.

	// Goroutines can be spawned by using the "go" keyword
	fmt.Println("Starting main program execution")
	go f1()
	f2()
	fmt.Println("Finished calling f1 & f2")
	// f1 executes in parallel with main. Wait for f1 to complete
	// for now use sleep
	time.Sleep(2 * time.Second)
	fmt.Println()

	// The above synchronization problem can be solved using waitgroups
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go f1wg(&wg)
	}
	wg.Wait()
	fmt.Println("Finished executing all go routines")
	fmt.Println()

	// Let's say you need to perform n taks. Each task is represented by
	// the dummy function runTask. Assuming each task takes x seconds to complete,
	// executing them sequentially taks n*x seconds. On the other hand executing
	// them in parallel using goroutines would be faster
	var exWg sync.WaitGroup
	tasks := []string{"task1", "task2", "task3"}
	exWg.Add(len(tasks))
	for _, task := range tasks {
		go runTask(task, &exWg)
	}
	exWg.Wait()
	fmt.Println()

	// Unsynchronized access to memory leads to data races
	// For ex: Let's say we launch 100 goroutines to increment
	// a variable n and 100 more to decrement the same.
	// Ideally we'd expect n to return to it's original value
	// after all the goroutines are done but that's not the
	// case due to data race. All 200 goroutines work on the
	// same variable 'n', leading to goroutines reading value of
	// 'n' while it's being changed by other goroutines.

	orig, n := 42, 42
	var drWg sync.WaitGroup
	drWg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			defer drWg.Done()
			n++
		}()
		go func() {
			defer drWg.Done()
			n--
		}()
	}
	drWg.Wait()
	fmt.Println("n value is: ", n)
	fmt.Println("Is n value what we expected?", n == orig)
	fmt.Println()

	// go provides a tool to detect races
	// Just add "-race" flat during go run/build to enable race detection during runtime
	// go run -race concurrency.go
	// indicating there's a race where a value read at line 104 is being overwritten
	// at line 100
	/*

		WARNING: DATA RACE
		Read at 0x00c000192010 by goroutine 17:
		  main.main.func2()
		      /home/sukruth.sridharan/workspace/learn_go/0017_concurrency/concurrency.go:104 +0x6c

		Previous write at 0x00c000192010 by goroutine 16:
		  main.main.func1()
		      /home/sukruth.sridharan/workspace/learn_go/0017_concurrency/concurrency.go:100 +0x84

		Goroutine 17 (running) created at:
		  main.main()
		      /home/sukruth.sridharan/workspace/learn_go/0017_concurrency/concurrency.go:102 +0x490

		Goroutine 16 (finished) created at:
		  main.main()
		      /home/sukruth.sridharan/workspace/learn_go/0017_concurrency/concurrency.go:98 +0x464
		==================

	*/

	// Solving data races by Mutexes.
	// code blocks enclosed within lock and unlock are called critical sections
	// critical sections can be accessed by only one goroutine at any given time

	orig, n = 0, 0
	var mutexLock sync.Mutex
	var mdrWg sync.WaitGroup
	mdrWg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			defer mdrWg.Done()
			mutexLock.Lock()
			defer mutexLock.Unlock()
			n++
		}()
		go func() {
			defer mdrWg.Done()
			mutexLock.Lock()
			defer mutexLock.Unlock()
			n--
		}()
	}
	mdrWg.Wait()
	fmt.Println("n value is: ", n)
	fmt.Println("Is n value what we expected?", n == orig)
	fmt.Println()

	// Data races can also be solved using channels.
	// A channel in go provides a connection between two goroutines to communicate.
	// Channels in go can communicate data only of the type defined during declaration.

	// Channels are used in conjuntion with goroutines. Therefore the following code,
	// even though valid, leads to deadlock.
	/*
		// Declaring a channel
		var ch chan int // nil channel
		fmt.Println(ch)

		// Initialzing a channel
		ch = make(chan int)
		fmt.Println(ch) // stores an address, therefore passing channels has the same effect as passing pointers

		// '<-' is the channel operator

		// send operation
		ch <- 10

		// receive operation. This is a blocking operation
		num := <-ch

		// closing a channel
		close(ch)

		// ch is a bidirectional channel

		// channel only for sending
		chS := make(chan <- int)

		// channel only for receiving
		chR := make(<- chan int)

	*/

	// Calculating factorial of n using channels

	ch := make(chan int)

	for i := 1; i < 10; i++ {
		go func(n int, c chan int) {
			f := 1
			for j := 2; j <= n; j++ {
				f *= j
			}
			c <- f
		}(i, ch)
		fmt.Println("Factorial of", i, "is", <-ch)
	}
	fmt.Println()

	// Refactoring the task code uwing channels

	strChan := make(chan string)
	for _, task := range tasks {
		go runTaskWithChan(task, strChan)
	}

	for i := 0; i < len(tasks); i++ {
		fmt.Println(<-strChan)
	}
	fmt.Println()

	// Channels initialized without capacity are called unbuffered channels
	// Channels created with capacity are called buffered channels

	// c1 := make(chan int) // unbuffered
	// c2 := make(chan int) // buffered

	// For unbuffered channels, sender blocks on the channel until receiver
	// receives from the channel
	// Unbuffered channel give stronger synchronization because every send
	// operation is synchronized with receive operation.
	c1 := make(chan int) //unbuffered channel

	// Launching a goroutine
	go func(c chan int) {
		fmt.Println("anon func: before sending data")
		c <- 10
		fmt.Println("anon func: after sending data")
	}(c1)

	fmt.Println("main goroutine, sleeping for 2s")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine receive data")
	d := <-c1
	fmt.Println("main goroutine received data:", d)

	// we sleep for a second to give time to the goroutine to finish
	time.Sleep(2 * time.Second)
	fmt.Println()

	// For buffered channels, the sender blocks only when there's no
	// slot available on the channel
	c2 := make(chan int, 3) // buffered channel

	// Launching a goroutine
	go func(c chan int) {
		for i := 0; i < 5; i++ {
			fmt.Println("anon func: before sending data", i)
			c <- i
			fmt.Println("anon func: after sending data", i)
		}
		close(c)
	}(c2)

	fmt.Println("main goroutine, sleeping for 2s")
	time.Sleep(time.Second * 2)

	for v := range c2 {
		fmt.Println("main goroutine received data:", v)
	}
	fmt.Println()

	// Receiving value from closed channel leads to zero value of channel type
	// Sending value to closed channel causes panic

	// Using select statements
	// Select statement lets a goroutine wait on multiple communication operations

	chan1, chan2 := make(chan string), make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		chan1 <- "Hello!"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "Hey!"
	}()

	for i := 0; i < 3; i++ {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Done with 2 seconds")
		case msg := <-chan1:
			fmt.Println("chan1:", msg)
		case msg2 := <-chan2:
			fmt.Println("chan2:", msg2)
		}
	}
	fmt.Println()

	// Refactoring the data race example above using channels

	orig, n = 0, 0
	done := make(chan int, 1)
	var chWg sync.WaitGroup
	chWg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			defer chWg.Done()
			<-done
			n++
			done <- 1
		}()
		go func() {
			defer chWg.Done()
			<-done
			n--
			done <- 1
		}()
	}
	done <- 1
	chWg.Wait()
	fmt.Println("n value is: ", n)
	fmt.Println("Is n value what we expected?", n == orig)
	fmt.Println()

}
