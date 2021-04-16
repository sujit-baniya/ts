package main

import (
	"fmt"
	"github.com/sujit-baniya/ts"
	"time"
)

func main() {
	u, _ := ts.Parse("tomorrow 4pm")
	t := time.Unix(u,0)
	fmt.Println(t)
}
