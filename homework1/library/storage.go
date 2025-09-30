package library

import "fmt"

type storage struct {
	useSlice bool
	slice []Book
	index map[int]*Book
	getID func(b *Book) int
}

func (s *storage) setFunc(f func(b *Book) int) {
	s.getID = f
}

func (s *storage) makeSlice(cap int) {
	s.slice = append(s.slice, make([]Book, cap + 1)...)
}

func (s *storage) addBook(b *Book) int {
	id := s.getID(b)
	if (s.useSlice) {
		if (cap(s.slice) <= id) {
			s.makeSlice(id)
		}
		s.slice[id] = *b
	} else {
		if (s.index == nil ) {
			s.index = make(map[int]*Book)
		}
		s.index[id] = b
	}
	return id
}

func (s *storage) removeBook(id int) error {
	if (s.useSlice) {
		if (cap(s.slice) <= id) {
			return fmt.Errorf("Нет книги с таким id в хранилище")
		}
		s.slice[id] = *new(Book)
	} else {
		if _, ok := s.index[id]; !ok {
			return fmt.Errorf("Нет книги с таким id в хранилище")
		} else {
			delete(s.index, id)
		}

	}
	return nil
}

func (s *storage) getBook(id int) (*Book, error) {
	if (s.useSlice) {
		if (cap(s.slice) <= id) {
			return nil, fmt.Errorf("Нет книги с таким id в хранилище")
		}
		if (!s.slice[id].GetExist()) {
			return nil, fmt.Errorf("Нет книги с таким id в хранилище")
		}
		return &s.slice[id], nil
	} else {
		if _, ok := s.index[id]; !ok {
			return nil, fmt.Errorf("Нет книги с таким id в хранилище")
		} else {
			return s.index[id], nil
		}
	}
}

func newStorage(useSlice bool) *storage {
	a := new(storage)
	a.useSlice = useSlice
	return a
}