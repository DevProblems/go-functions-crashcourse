package main

import (
	"fmt"
	"go-function-crashcourse/anonymousfunction"
	"go-function-crashcourse/basicfunction"
	"go-function-crashcourse/genericfunction"
	"go-function-crashcourse/receiverfunction"
)

/**
 * @author Dev Problems(A Sarang Kumar Tak)
 * @YoutubeChannel <a href="https://www.youtube.com/channel/UCVno4tMHEXietE3aUTodaZQ">...</a>
 */

var Name string

//init function
func init() {
	Name = "Sarang"
}

func main() {
	println(Name)

	basicfunction.Add(1, 2)

	r := basicfunction.Subtract(4, 2)
	println(r)

	s, d, p := basicfunction.Subscribe()
	println(s, d, p)

	ans, err := basicfunction.Divide(4, 0)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Printf("division result : %f", ans)
	}

	sum := basicfunction.Sum(1, 2, 3, 4, 5)
	fmt.Printf("sum result : %f", sum)
	fmt.Println()

	var a = 5
	basicfunction.PassByValue(a)
	println(a)

	basicfunction.PassByReference(&a)
	println(a)

	basicfunction.Defer()
	fmt.Println()

	person := receiverfunction.Person{}
	println("address of person object", &person)
	person.SetAge(20)
	fmt.Println(person)

	fun := func(name string) {
		println(name)
	}

	fun("Sarang")

	c := func(a, b int) int {
		return a + b
	}

	anonymousfunction.Calculate(1, 2, c)

	function := anonymousfunction.GetFunction()
	str := function("DevProblems")
	println(str)

	var res int64
	sf := anonymousfunction.StateFul()
	res = sf()
	println(res)
	res = sf()
	println(res)
	res = sf()
	println(res)

	m := map[string]float64{
		"first":  12.00,
		"second": 14.00,
	}
	re := genericfunction.SumIntsOrFloats(m)
	println(re)
}
