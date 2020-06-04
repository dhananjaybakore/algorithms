package core

import "fmt"

/*
	Includes most common variable declaration styles in GO
*/
func BasicVariableDeclaration() {
	var i int
	i = 42
	println(i)

	var f float32 = 3.14
	println(f)

	firstName := "John"
	println(firstName)

	c := complex(3, 4)
	println(c)

	a, b := real(c), imag(c)
	println(a, b)
}

func PointersDeclaration() {
	firstName := new(string)
	*firstName = "John"
	//Print value at address
	fmt.Printf("Var Address: %v Var Value: %v", &firstName, *firstName)
}

func PointersDeclaration2() {
	var firstName string
	firstName = "Arthur"
	fmt.Println(firstName)
	ptr := &firstName
	fmt.Println(ptr, *ptr)
}

//const (
//	first = iota
//	second
//)

//const (
//	third = iota
//	fourth
//)

func UsingConstants() {
	const c = 3
	println(c + 1.2)
	// a bunch of code
	println(c + 5)
}

func Example() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[1])

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
	fmt.Println(arr2[1])
}

func SlicesExample1() {
	arr := [3]int{1, 2, 3}

	slice := arr[:]
	fmt.Println(slice)

	arr[0] = 42
	slice[2] = 27

	fmt.Println(arr, slice)
}

func SlicesExample2() {
	slice := []int{1, 2, 3}
	fmt.Println(slice)
	slice = append(slice, 4, 42, 27)
	fmt.Println(slice)
	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]
	fmt.Println(slice, s2, s3, s4)
}

func MapsExample1() {
	m := map[string]int{"foo": 42}

	fmt.Println(m)
	fmt.Println(m["foo"])

	m["bar"] = 27
	fmt.Println(m)
	delete(m, "bar")
	fmt.Println(m)
}

func StructsExample1() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	u := user{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}
	fmt.Println(u)
}

func FunctionWithParams(arg1, arg2 int) {
	fmt.Printf("Arguments received %v and %v", arg1, arg2)
}

func FunctionWithReturnParams(port int) (int, error) {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
	return port, nil
}

func BasicLoopTillCondition() {
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			// break
			continue
		}
		println("after continue")
	}
}

func LoopWithPostClause() {
	for i := 1; i < 5; i++ {
		println(i)
	}
}

func LoopOverCollection() {
	slice := []int{1, 2, 3}
	for i, v := range slice {
		println(i, v)
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for k, v := range wellKnownPorts {
		println(k, v)
	}
}

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func IfThenElseExample() {
	u1 := User{
		ID:        1,
		FirstName: "Arthur",
		LastName:  "Dent",
	}
	u2 := User{
		ID:        2,
		FirstName: "Ford",
		LastName:  "Prefect",
	}
	if u1 == u2 {
		println("Same user!")
	} else if u1.FirstName == u2.FirstName {
		println("Similar user.")
	} else {
		println("Different user!")
	}
}

type HTTPRequest struct {
	Method string
}

func SwitchStatementExample() {
	r := HTTPRequest{Method: "HEAD"}

	switch r.Method {
	case "GET":
		println("Get request")
	case "DELETE":
		println("Delete request")
	case "POST":
		println("Post request")
	case "PUT":
		println("Put request")
	default:
		println("Unhandled method")
	}
}
