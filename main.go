package main

import (
	"fmt"
	"strconv"
)

func greetings() {
	const BOT_NAME = "SeregaBot"
	const BIRTH_YEAR = 2022
	fmt.Println("Hello! My name is ", BOT_NAME)
	fmt.Println("I was created in ", BIRTH_YEAR)
}
func showName() {
	var name string
	fmt.Println("Please, remind me your name.")
	fmt.Scan(&name)
	fmt.Println("Hello " + name + "!")
}
func guessAge() {
	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")

	var remainder3, remainder5, remainder7, age int

	fmt.Scan(&remainder3, &remainder5, &remainder7)
	age = (remainder3*70 + remainder5*21 + remainder7*15) % 105

	fmt.Println("Your age is " + strconv.Itoa(age) + "; that's a good time to start programming!")
}
func count() {
	fmt.Println("Now I will prove to you that I can count to any number you want.")

	var number int
	fmt.Scan(&number)
	for i := 0; i < number+1; i++ {
		fmt.Println(i, " !")
	}
}
func startQuiz() {
	fmt.Println("Let's test your programming knowledge.")
	fmt.Println("Why do we use methods?")
	fmt.Println("1. To repeat a statement multiple times.")
	fmt.Println("2. To decompose a program into several small subroutines.")
	fmt.Println("3. To determine the execution time of a program.")
	fmt.Println("4. To interrupt the execution of a program.")
	var ans int
Check:
	fmt.Scan(&ans)
	if ans != 2 {
		fmt.Println("Please, Try again.")
		goto Check
	}
}

func main() {
	greetings()
	showName()
	guessAge()
	count()
	startQuiz()

	fmt.Println("Completed, have a nice day!")

}
