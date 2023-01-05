package main

import (
	"fmt"

	"github.com/gorhill/cronexpr"
)

func main() {
	var (
		expr *cronexpr.Expression
		err  error
	)

	if expr, err = cronexpr.Parse("* * * * *123"); err != nil {
		fmt.Println(err)
		return
	}

	expr = expr
}
