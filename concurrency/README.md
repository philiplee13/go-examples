## Concurrency

- can specify goroutines either by `go f()` or `go func(...) {}` (anon-methods)
- channels are pipes that connect concurrent goroutines
  - you can send data in from one goroutine -> channel -> recieve it in another goroutine
- https://gobyexample.com/goroutines
- https://gobyexample.com/channels
- channel sync
- channel directions
- select
  - select allows you to wait on multiple channel operations
