// This package gets rid of the annoying pattern:
// if e != nil {
//
// }
//
// and replaces it with:
// import e "errortest/errorfatal"
// ....
// e.F(os.Chdir("123"))
// ....
// e.F(os.WriteFile("123.txt", e.F2(os.ReadFile("456.txt")), 0644))
package errorfatal

import (
	"log"
)

func F(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func F2[K any](r K, e error) K {
	F(e)
	return r
}
