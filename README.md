Concat junit reports to csv
===========================

##Installation
```bash
go get github.com/sagix/concat-junit-to-csv
```
To install the executable in you *$GOBIN*
From your **src** directory execute:
```bash
go install github.com/sagix/concat-junit-to-csv
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
