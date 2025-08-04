package main

import (
	"fmt"
	"library/services"
)

func main() {
	studentService := services.NewStudentService()
	issueService := services.NewIssueService()
	bookService := services.NewBookService()

	student1 := studentService.AddStudent("Mani", "CSE", 1001)
	book1 := bookService.AddBook("Concepts of Physics","HC Verma")

	issue1 := issueService.IssueBook(book1.ID, student1.ID, 14)
	fmt.Printf("Issued Book #%d to Student #%d\n", issue1.BookID, issue1.StudentID)

	success := issueService.ReturnBook(issue1.ID)
	if success {
		fmt.Println("Book returned successfully!")
	}

	student2 := studentService.AddStudent("Saurabh", "ECE", 1002)
	book2 := bookService.AddBook("Concepts of Maths","RD Pandey")

	issue2 := issueService.IssueBook(book2.ID, student2.ID, 28)
	fmt.Printf("Issued Book #%d to Student #%d\n", issue2.BookID, issue2.StudentID)

	success = issueService.ReturnBook(issue2.ID)
	if success {
		fmt.Println("Book returned successfully!")
	}
}
