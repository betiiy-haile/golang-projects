package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"library-management/models"
	"library-management/services"
)

func StartConsole(library services.LibraryManger) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Library Management System ---")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter Book ID: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Enter Title: ")
			scanner.Scan()
			title := scanner.Text()

			fmt.Print("Enter Author: ")
			scanner.Scan()
			author := scanner.Text()

			book := models.Book{Id: id, Title: title, Author: author, Status: "Available"}
			library.AddBook(book)
			fmt.Println("Book added!")

		case "2":
			fmt.Print("Enter Book ID to remove: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			library.RemoveBook(id)
			fmt.Println("Book removed!")

		case "3":
			fmt.Print("Enter Book ID: ")
			scanner.Scan()
			bookID, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Enter Member ID: ")
			scanner.Scan()
			memberID, _ := strconv.Atoi(scanner.Text())

			err := library.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book borrowed!")
			}

		case "4":
			fmt.Print("Enter Book ID: ")
			scanner.Scan()
			bookID, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Enter Member ID: ")
			scanner.Scan()
			memberID, _ := strconv.Atoi(scanner.Text())

			err := library.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book returned!")
			}

		case "5":
			books := library.ListAvailableBooks()
			fmt.Println("Available Books:")
			for _, b := range books {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n", b.Id, b.Title, b.Author)
			}

		case "6":
			fmt.Print("Enter Member ID: ")
			scanner.Scan()
			memberID, _ := strconv.Atoi(scanner.Text())

			books := library.ListBorrowedBooks(memberID)
			fmt.Printf("Borrowed Books for Member %d:\n", memberID)
			for _, b := range books {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n", b.Id, b.Title, b.Author)
			}

		case "7":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
