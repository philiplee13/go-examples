## Unit tests

- simple example of how to write basic unit tests
- to run tests you can run `go test`
- to see all tests statuses use `go test -v`
- if you have multiple tests in sub directories - you can run all tests by using `go test ./...`
  - https://stackoverflow.com/questions/19200235/golang-tests-in-sub-directory
  - otherwise use `go test packagename`
    - so in our case - we would need to run `go test unittests/helpers`
