Concat junit reports to csv
===========================

##Installation
```bash
go get github.com/sagix/concat-junit-to-csv
```
To install the executable in you *$GOBIN*
```bash
go install concat-junit-to-csv.go
```
##Usage

Prints the result in the standard output:
```bash
concat-junit-to-csv directory
```
So it is easy to print the result in a file:
```bash
concat-junit-to-csv directory > file.csv
```
Result

    test filename, "test name", duration
