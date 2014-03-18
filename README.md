## Intro

Go package to compare the performance of different solutions for the same problem.

**Problem**: Given a non-empty string, return true if the string is a palindrome. Punctuation and spaces should be ignored.

### Install and Run

```
go get github.com/hgfischer/go-palindrome
cd $GOPATH/src/github.com/hgfischer/go-palindrome
go test -test.bench="."
```

### My Results

Results in a Intel Core-i7 laptop:

```
PASS
BenchmarkPalindromeA	   20000	     79121 ns/op
BenchmarkPalindromeB	 1000000	      1700 ns/op
ok  	github.com/hgfischer/gopalindrome	4.118s
```