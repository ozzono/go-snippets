package main

import "fmt"

type test struct {
	a string
	b int
}

func main() {
	type tt struct {
		m map[string]test
	}
	t1 := map[string]test{}
	t1["a"] = test{a: "a"}
	t2 := map[string]test{}
	t2["b"] = test{a: "b"}
	t := tt{m: t1}
	fmt.Printf("t: %+v\n", t)
	t = tt{m: t2}
	fmt.Printf("t: %+v\n", t)
}
