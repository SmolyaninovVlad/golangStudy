package cases

import (
	"fmt"
)

//InitCase3 ...
func InitCase3() {

	a := 42
	b := 153
	c := a

	a = b
	b = c

	fmt.Println("a:", a)

	fmt.Println("b:", b)

}
