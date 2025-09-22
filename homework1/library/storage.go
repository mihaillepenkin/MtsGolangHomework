package library

import "fmt"

type storage struct {
	slice []Book
	getId func(b *Book) int
}

func (s *storage) setFunc(f func(b *Book) int) {
	s.getId = f
}

func (s *storage) makeSlice(cap int) {
	s.slice = append(s.slice, make([]Book, cap + 1)...)
}

func (s *storage) addBook(b *Book) int {
	id := s.getId(b)
	if (cap(s.slice) <= id) {
		s.makeSlice(id)
	}
	s.slice[id] = *b
	return id
}

func (s *storage) removeBook(id int) error {
	if (cap(s.slice) <= id) {
		return fmt.Errorf("Нет книги с таким id в хранилище")
	}
	s.slice[id] = *new(Book)
	return nil
}

func (s *storage) getBook(id int) (*Book, error) {
	if (cap(s.slice) <= id) {
		return nil, fmt.Errorf("Нет книги с таким id в хранилище")
	}
	if (!s.slice[id].GetExist()) {
		return nil, fmt.Errorf("Нет книги с таким id в хранилище")
	}
	return &s.slice[id], nil
}