package stringutil

func Reverse(s string) string {

	return reverseTwo(s)
}

/*
go build
	go build reverse.go reverseTwo.go
	won't produce an output file.

go install
	will place the package
 */