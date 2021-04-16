package main

import (
	"fmt"
	"github.com/sujit-baniya/ts"
	"time"
)

func main() {
	u, _ := ts.Parse("first friday in 2021")
	t := time.Unix(u, 0)
	fmt.Println(t)
	fmt.Println(ts.Date("2021-01-12"))
}
