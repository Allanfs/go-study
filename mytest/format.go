package mytest

import "fmt"

func formatName(st, nd string) string {
	return fmt.Sprintf("Dear %s %s", st, nd)
}

func g() int {
	return 4
}
