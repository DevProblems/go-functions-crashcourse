package anonymousfunction

import "math"

/**
 * @author Dev Problems(A Sarang Kumar Tak)
 * @YoutubeChannel <a href="https://www.youtube.com/channel/UCVno4tMHEXietE3aUTodaZQ">...</a>
 */

//Calculate function as parameter
func Calculate(a, b int, c func(int, int) int) {
	r := c(a, b)
	println(r)
}

//GetFunction return function from function
func GetFunction() func(string) string {
	return func(s string) string {
		return s
	}
}

//StateFul function
func StateFul() func() int64 {
	a := 0.0
	return func() int64 {
		a += 1
		return int64(math.Pow(a, 2))
	}
}
