// fizzybuzzy project main.go
package main

// function version

//func main() {

//	for i := 1; i <= 100; i++ {

//		x := fb(i)
//		if len(x) > 0 {
//			println(x)
//		} else {
//			println(i)
//		}
//	}
//}

//func fb(i int) string {
//	var returnMe string
//	if i%3 == 0 {
//		returnMe = "Fizz"
//	}
//	if i%5 == 0 {
//		returnMe += "Buzz"
//	}
//	return returnMe
//}

// straightforward version

//func main() {

//	for i := 1; i <= 100; i++ {

//		if i%3 == 0 && i%5 == 0 {
//			println("FizzBuzz")
//		} else if i%3 == 0 {
//			println("Fizz")
//		} else if i%5 == 0 {
//			println("Buzz")
//		} else {
//			println(i)
//		}
//	}
//}

// parameterized version with support for more than just fizz and buzz

func main() {

	MIN := 0
	MAX := 100
	STEP := 1

	fizzMap := make(map[int]string)
	fizzMap[3] = "Fizz"
	fizzMap[5] = "Buzz"
	fizzMap[7] = "Boom"

	for i := MIN; i <= MAX; i += STEP {
		//tried to have this in a sep func but had trouble passing in map
		var r string
		for k, v := range fizzMap {
			if i%k == 0 {
				r += v
			}
		}

		if len(r) > 0 {
			println(r)
		} else {
			println(i)
		}
	}
}
