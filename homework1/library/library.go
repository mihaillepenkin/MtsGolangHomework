package library

import (
	"fmt"
	"math"
)

type library struct {
	index map[string]int
	stor storage
}

func (l *library) makeMap() {
	l.index = make(map[string]int)
}

func (l *library) SetFunc(f func(b *Book) int) {
	l.stor.setFunc(f)
	l.Clear()
	fmt.Println("Функция для выдачи айди изменена")
}

func(l *library) Clear() {
	l.stor.slice = make([]Book, len(l.stor.slice))
	l.makeMap()
	fmt.Println("Библиотека очищена")
}


func (l *library) AddBooks(books ...*Book) error {
	if l.index == nil {
		l.makeMap()
	}
	if l.stor.getID == nil {
		l.SetFunc(func (b *Book) int {
			name := []rune(b.name)
			id := 0
			for ind, elem := range name {
				id += int(elem) * (int(math.Pow(float64(simpleNumber), float64(ind))))
			}
			id = id % 100003
			return id
		})
	}
	listOfErrors := make([]string, 0)
	for _, b := range books {
		if _, ok := l.index[b.GetName()]; ok {
			listOfErrors = append(listOfErrors, b.GetName())
		} else {
			id := l.stor.addBook(b)
			l.index[b.GetName()] = id
			fmt.Printf("Книга %s добавлена в библиотеку \n", b.GetName())
		}
	}
	if (len(listOfErrors) != 0) {
		nameForError := ""
		for _, name := range listOfErrors {
			nameForError = nameForError + name + ", "
		}
		if (len(listOfErrors) > 1) {
			return fmt.Errorf("Книги %s уже есть в библиотеке \n", nameForError)
		} else {
			return fmt.Errorf("Книга %s уже есть в библиотеке \n", listOfErrors[0])
		}
	}
	return nil	
}

func (l *library) RemoveBooks(books ...*Book) error {
	listOfErrors := make([]string, 0)
	for _, b := range books {
		if _, ok := l.index[b.GetName()]; ok {
			err := l.stor.removeBook(l.index[b.GetName()])
			if err != nil {
				fmt.Println(err)
			} else {
				delete(l.index, b.GetName())
				fmt.Printf("Книга %s удалена из библиотеки \n", b.GetName())
			}
		} else {
			listOfErrors = append(listOfErrors, b.GetName())
		}
	}
	if (len(listOfErrors) != 0) {
		nameForError := ""
		for _, name := range listOfErrors {
			nameForError = nameForError + name + ", "
		}
		if (len(listOfErrors) > 1) {
			return fmt.Errorf("Книг %s нет в библиотеке \n", nameForError)
		} else {
			return fmt.Errorf("Книги %s нет в библиотеке \n", listOfErrors[0])
		}
	}
	return nil
}

func (l *library) GetBook(name string) (*Book, error) {
	ans, err := l.stor.getBook(l.index[name])
	return ans, err
}

func NewLibrary(useSliceForStrorage bool) *library{
	a := new(library)
	a.stor = *newStorage(useSliceForStrorage)
	return a
}