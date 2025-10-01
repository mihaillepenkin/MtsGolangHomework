package main

import (
	"fmt"
	"homework1/library"
	"math"
)


func main() {
	lib := library.NewLibrary(true)

	book1 := library.NewBook("War", "Long years ago....")
	book2 := library.NewBook("World", "Long month ago....")
	book3 := library.NewBook("Misha", "Long hours ago....")
	book4 := library.NewBook("Mipt&&Mts", "Long seconds ago....")
	books := []*library.Book{book2, book3, book4}

	err := lib.AddBooks(book1)
	if (err != nil) {
		fmt.Println(err)
	}
	err = lib.AddBooks(books...)
	if (err != nil) {
		fmt.Println(err)
	}
	

	if book, err := lib.GetBook("Misha"); err != nil {
		fmt.Println(err)
	} else {
		book.PrintText()
	}

	if book, err := lib.GetBook("War"); err != nil {
		fmt.Println(err)
	} else {
		book.PrintText()
	}

	if book, err := lib.GetBook("Anatoliy"); err != nil {
		fmt.Println(err)
	} else {
		book.PrintText()
	}

	lib.Clear()

	if book, err := lib.GetBook("Misha"); err != nil {
		fmt.Println(err)
	} else {
		book.PrintText()
	}

	err = lib.AddBooks(books...)
	if (err != nil) {
		fmt.Println(err)
	}

	if book, err := lib.GetBook("Misha"); err != nil {
		fmt.Println(err)
	} else {
		book.PrintText()
	}

	err = lib.RemoveBooks(books...)
	if (err != nil) {
		fmt.Println(err)
	}

	if book, err := lib.GetBook("Misha"); err != nil {
		fmt.Println(err)
	} else {
		book.PrintText()
	}

	err = lib.RemoveBooks(book3)
	if (err != nil) {
		fmt.Println(err)
	}

	lib.SetFunc(func (b *library.Book) int {
			name := []rune(b.GetName())
			id := 0
			for ind, elem := range name {
				id += int(elem) * (int(math.Pow(float64(17), float64(ind))))
			}
			id = id % 100003
			return id
		})

	err = lib.AddBooks(books...)
	if (err != nil) {
		fmt.Println(err)
	}

	if book, err := lib.GetBook("Misha"); err != nil {
		fmt.Println(err)
	} else {
		book.PrintText()
	}

	fmt.Println("Теперь проверим либу на мапе")

	lib = library.NewLibrary(false)

	err = lib.AddBooks(book1)
	if (err != nil) {
		fmt.Println(err)
	}
	err = lib.AddBooks(books...)
	if (err != nil) {
		fmt.Println(err)
	}
	

	if book, err := lib.GetBook("Misha"); err != nil {
		fmt.Println(err)
	} else {
		book.PrintText()
	}

	if book, err := lib.GetBook("War"); err != nil {
		fmt.Println(err)
	} else {
		book.PrintText()
	}

	if book, err := lib.GetBook("Anatoliy"); err != nil {
		fmt.Println(err)
	} else {
		book.PrintText()
	}

	lib.Clear()

	if book, err := lib.GetBook("Misha"); err != nil {
		fmt.Println(err)
	} else {
		book.PrintText()
	}

	err = lib.AddBooks(books...)
	if (err != nil) {
		fmt.Println(err)
	}

	if book, err := lib.GetBook("Misha"); err != nil {
		fmt.Println(err)
	} else {
		book.PrintText()
	}

	err = lib.RemoveBooks(books...)
	if (err != nil) {
		fmt.Println(err)
	}

	if book, err := lib.GetBook("Misha"); err != nil {
		fmt.Println(err)
	} else {
		book.PrintText()
	}

	err = lib.RemoveBooks(book3)
	if (err != nil) {
		fmt.Println(err)
	}

	lib.SetFunc(func (b *library.Book) int {
			name := []rune(b.GetName())
			id := 0
			for ind, elem := range name {
				id += int(elem) * (int(math.Pow(float64(17), float64(ind))))
			}
			id = id % 100003
			return id
		})

	err = lib.AddBooks(books...)
	if (err != nil) {
		fmt.Println(err)
	}

	if book, err := lib.GetBook("Misha"); err != nil {
		fmt.Println(err)
	} else {
		book.PrintText()
	}


}
