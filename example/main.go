package main

import (
	"fmt"
	"time"

	"github.com/spiretechnology/go-pool"
)

func main() {

	// Create a new worker pool
	p := pool.Default()

	// Loop through a slice of values, which represent units of work
	for value := 1; value <= 10; value++ {

		// Make a copy of the value in this lexical scope
		value := value

		// Submit this unit of work to the worker pool
		p.Go(func() {

			// Sleep for a few seconds to simulate work, then log the value
			time.Sleep(time.Second * 2)
			fmt.Println("Done: ", value)

		})

	}

	// Wait for all the jobs in the pool to finish
	p.Wait()

	fmt.Println("All jobs are done!")

}
