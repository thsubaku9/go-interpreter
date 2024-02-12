package main

import (
	"fmt"
	"monkey-i/repl"
	"os"
	"os/user"
)

func main() {
	userRef, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", userRef.Username)

	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}

// type Reference struct {
// 	current string
// 	next    string
// 	s       *StringStream
// }

// type StringStream struct {
// 	stream  []string
// 	pointer int
// 	limit   int
// }

// func (s *StringStream) Next() string {
// 	if s.pointer < s.limit {
// 		s.pointer += 1
// 		return s.stream[s.pointer-1]
// 	}
// 	panic("out of scope")
// }

// func (r *Reference) RecurseDive(level int) {
// 	if level == 0 {
// 		return
// 	}
// 	r.RecurseDive(level - 1)

// }

// func (r *Reference) NextString() {
// 	r.current, r.next = r.next, r.s.Next()
// }

// var stringStream []string = []string{"one", "two", "three", "four"}

// func main() {

// }
