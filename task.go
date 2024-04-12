package main

func task(c chan int, N int) {
	c <- N + 1
}

func task2(c chan string, s string) {
	for i := 0; i < 5; i++ {
		c <- s + " "
	}
}

func removeDuplicates(inputStream chan string, outputStream chan string) {
	prev := ""
	for v := range inputStream {
		if v != prev {
			outputStream <- v
		}
		prev = v
	}
	close(outputStream)
}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	output := make(chan int)
	go func(output chan int) {
		defer close(output)
		select {
		case x := <-firstChan:
			output <- x * x
		case x := <-secondChan:
			output <- x * 3
		case <-stopChan:
			return
		}
	}(output)

	return output
}

func calculator2(arguments <-chan int, done <-chan struct{}) <-chan int {
	output := make(chan int)
	go func(output chan int) {
		defer close(output)
		sum := 0
		for {
			select {
			case x := <-arguments:
				sum += x
			case <-done:
				output <- sum
				return
			}
		}

	}(output)
	return output
}

func merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	idx := make(chan int) 
	val := make(chan int) 

	go func() {
		for i := 0; i < n; i++ {
			x1 := <-in1
			x2 := <-in2

			done := make(chan int)
			go func() {
				done <- f(x1)
			}()
			go func() {
				done <- f(x2)
			}()

			go func(i int) {
				idx <- i
				val <- (<-done) + (<-done)
			}(i)
		}
	}()

	go func() {
		m := make(map[int]int)
		for i := 0; i < n; i++ {
			m[<-idx] = <-val
		}
		for i := 0; i < n; i++ {
			out <- m[i]
		}
	}()
}
