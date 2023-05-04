To use this application you need to have Go installed on your machine.

To play a game: 
1. copy all the files into a new directory and `cd` into it.
2. run `go mod init sudoku`
3. run `go mod tidy`
4. run `go run .`

To run the benchmark tests:
1. follow steps 1 to 3 above
2. run `go test -bench=. -count 5`