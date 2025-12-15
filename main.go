package main

import (
	"fmt"
	"math"
	"time"
)

const s string = "constant"

func HelloWorld() {
	fmt.Println("Print greeting:")
	var h = "Hello"
	fmt.Println(h + " World!")
}

// VALUE TYPES (string, int, float, bool)
func ValueTypes() {
	fmt.Println("Value Types:")
	// can be added together with +
	fmt.Println("go" + "lang")
	// for integers and floats
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println("7/3 =", 7/3)
	// boolean, with boolean operators
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

// VARIABLES
func Variables() {
	fmt.Println("Variables:")
	var a = "begin"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "awesome"
	fmt.Println(f)

	fmt.Println("combined a and f: ", a, f)
}

// CONSTANTS
func Constants() {
	fmt.Println("Constants:")
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}

// FOR LOOPS
func ForLoops() {
	fmt.Println("For Loops")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("-----")
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	fmt.Println("-----")
	for i := range 3 {
		fmt.Println("range", i)
	}

	fmt.Println("-----")
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("-----")
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

// IF/ELSE
func IfElse() {
	fmt.Println("If/Else Statements")
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// SWITCH STATEMENTS
func SwitchStatements() {
	fmt.Println("Switch Statements")
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("Don't know type", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

// ARRAYS
func Arrays() {
	fmt.Println("Arrays:")
	var a [5]int // holds exactly 5 integers
	fmt.Println("empty array:", a)

	a[4] = 100 // set index 4 to 100
	fmt.Println("set index 4:", a)
	fmt.Println("get index 4:", a[4])

	fmt.Println("array length:", len(a))

	b := [5]int{1, 2, 3, 4, 5} // initialize array
	fmt.Println("initialized array:", b)

	b = [...]int{1, 2, 3, 4, 5} // ellipsis infers length
	fmt.Println("inferred length array:", b)

	var twoD [2][3]int // 2D array
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D array:", twoD)

	twoD = [2][3]int{ // initialized 2D array
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("initialized 2D array:", twoD)
}

// SLICES
func Slices() {
	fmt.Println("Slices:")
	s := make([]string, 3) // create slice of strings with length 3
	fmt.Println("empty slice:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set slice:", s)
	fmt.Println("get slice index 2:", s[2])
	fmt.Println("slice length:", len(s))
}

// MAIN FUNCTION
func main() {
	HelloWorld()
	// ValueTypes()
	// Variables()
	// Constants()
	// ForLoops()
	// IfElse()
	// SwitchStatements()
	// Arrays()
}
