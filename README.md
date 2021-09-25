# go-pool

A Go library for goroutine pools.

## Installation

```bash
go get github.com/spiretechnology/go-pool
```

## Problem

A common pattern in Go is to split up work into `N` parallel goroutines:

```go
// Create wait group so we can wait for all work to finish
var wg sync.WaitGroup
wg.Add(len(jobs))

// Loop through all the jobs that need to be done
for _, job := range jobs {
    job := job
    go func() {
        defer wg.Done()
        job()
    }()
}

// Wait for all jobs to finish
wg.Wait()
```

This is a great and extremely useful pattern. However, there is a downside to this pattern. As the number of jobs `N` increases, the number of active goroutines increases, reducing the resources available to other concurrent users of your application.

In systems handling many concurrent requests, long-running work spawning thousands of goroutines can cause slow response times for other clients who have simpler requests.

## Solution

This library provides a simple solution: **goroutine pools**. Instead of launching `N` goroutines for an input of that size, goroutine pools allow you to launch `min(N, limit)` where the `limit` value caps resource utilization.

```go
// Create a new worker pool
p := pool.New(5)

// Loop through the jobs that need to be done
for _, job := range jobs {
    job := job
    p.Go(func() {
        // Run the job
        job()
    })
}

// Wait for all jobs to finish
p.Wait()
```

This works *nearly* identical to the prior example, but the number of simultaneous goroutines spawned is limited to `5` instead of being infinite.

## Priority vs magic numbers

The pattern of creating worker pools with a hard-coded maximum (such as `pool.New(5)`) is not ideal, as it doesn't scale with the amount of available resources.

Because of this, there are a few built-in functions that are preferred:

```go
// Creates a pool with a maximum number of goroutines equal to the number
// of CPUs on the machine
pool.Default()
pool.WithPriority(pool.NORMAL)

// Creates a pool with a HIGH priority, which allows many times more goroutines to
// be scheduled than the number of CPUs
pool.WithPriority(pool.HIGH)

// Creates a pool with a LOW priority, which allows only a fraction of the number of 
// goroutines to be scheduled simultaneously.
pool.WithPriority(pool.LOW)
```

## Concluding notes

The Go runtime does a great job scheduling goroutines, but it has no concept of how we, as the systems engineers, want to prioritize concurrent tasks. Here's an example:

On a highly concurrent API server, some requests may spawn thousands of goroutines, and some requests may spawn only a handful or less. If these requests are received simultaneously, the larger one may dramatically slow down the response time for the smaller request by choking it of resources.

Using goroutine pools (this library) we're able to throttle down the larger tasks so all clients can have their requests fulfilled in a reasonable time.
