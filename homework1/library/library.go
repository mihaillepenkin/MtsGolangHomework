package library

import (
	"fmt"
	"math"
)

type library struct {
	mapa map[string]int
	stor storage
}

func (l *library) makeMap() {
	l.mapa = make(map[string]int)
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


func (l *library) AddBooks(books ...*Book) {
	if l.mapa == nil {
		l.makeMap()
	}
	if l.stor.getId == nil {
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
	for _, b := range books {
		if _, ok := l.mapa[b.GetName()]; ok {
			fmt.Printf("Книга %s уже есть в библиотеке \n", b.GetName())
		} else {
			id := l.stor.addBook(b)
			l.mapa[b.GetName()] = id
			fmt.Printf("Книга %s добавлена в библиотеку \n", b.GetName())
		}
	}
}

func (l *library) RemoveBooks(books ...*Book) {
	for _, b := range books {
		if _, ok := l.mapa[b.GetName()]; ok {
			err := l.stor.removeBook(l.mapa[b.GetName()])
			if err != nil {
				fmt.Println(err)
			} else {
				delete(l.mapa, b.GetName())
				fmt.Printf("Книга %s удалена из библиотеки \n", b.GetName())
			}
		} else {
			fmt.Printf("Книги %s нет в библиотеке \n", b.GetName())
		}
	}
}

func (l *library) GetBook(name string) (*Book, error) {
	ans, err := l.stor.getBook(l.mapa[name])
	return ans, err
}

func NewLibrary() *library{
	return new(library)
}