package steps

import "fmt"

// --- Directions
// Write a function that accepts a positive number N.
// The function should write out a step shape
// with N levels using the # character.  Make sure the
// step has spaces on the right hand side!
// --- Examples
//   steps(2)
//       '# '
//       '##'
//   steps(3)
//       '#  '
//       '## '
//       '###'
//   steps(4)
//       '#   '
//       '##  '
//       '### '
//       '####'

//recursion with optional variable
func stepsRecursion(n int, rowOptional ...int) (str string) {
	row := 0
	if len(rowOptional) > 0 {
		row = rowOptional[0]
	}

	if row >= n {
		return str
	}

	for i := 0; i < n; i++ {
		if i <= row {
			str += "#"
		} else {
			str += " "
		}
	}
	if row < (n - 1) {
		str += "\n"
	}
	str += stepsRecursion(n, row+1)
	return

}

func steps(n int) string {
	str := ""
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j <= i {
				str += "#"
			} else {
				str += " "
			}
		}
		if i < n-1 {
			str += "\n"
		}
	}
	fmt.Println(str)
	return str
}
