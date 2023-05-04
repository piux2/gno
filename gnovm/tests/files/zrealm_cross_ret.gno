// PKGPATH: gno.land/r/returntype_test
package returntype_test

import (
	"gno.land/r/demo/tests"
)

func main() {

	a := tests.AbsAdd()
	b := tests.GetAbs()
	println(tests.AbsAdd())
	println(tests.GetAbs())
	println(a.Add())
	println(b.Add())

}

// Output:
// slice[(0 gno.land/r/demo/tests.Word)]
// slice[(0 gno.land/r/demo/tests.Word)]
// slice[(0 gno.land/r/demo/tests.Word)]
// slice[(0 gno.land/r/demo/tests.Word)]
