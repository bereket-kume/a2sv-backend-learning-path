package main

import (
	"fmt"
)

func calculateAverage(val map[string]int, n int) (float32, int) {
	var total int
	for _, grade := range val {
		total += grade
	}
	average := float32(total) / float32(n)
	return average, total
}

func main() {
	var studentName string
	var numberOfCourse int

	fmt.Print("Enter your name: ")
	fmt.Scan(&studentName)

	fmt.Print("Enter the number of courses you are taking: ")
	fmt.Scan(&numberOfCourse)

	studentInfo := make(map[string]int)

	for i := 0; i < numberOfCourse; i++ {
		var courseName string
		var grade int

		fmt.Print("Enter course name: ")
		fmt.Scan(&courseName)

		fmt.Print("Enter grade: ")
		fmt.Scan(&grade)

		studentInfo[courseName] = grade
	}

	average, total := calculateAverage(studentInfo, numberOfCourse)

	fmt.Println("\nCourse         Grade")
	for course, grade := range studentInfo {
		fmt.Printf("%-15s %d\n", course, grade)
	}

	fmt.Printf("\nTotal: %d\n", total)
	fmt.Printf("Average: %.2f\n", average)
}
