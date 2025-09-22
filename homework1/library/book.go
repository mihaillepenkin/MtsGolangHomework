package library

import (
	"fmt"
)

const simpleNumber = 13

type Book struct {
	name string
	text string
	exist bool
}

func (b *Book) GetText() string {
	return b.text
}

func (b *Book) GetExist() bool {
	return b.exist
}

func (b *Book) GetName() string {
	return b.name
}

func (b *Book) SetText(text string) {
	b.text = text
}

func (b *Book) SetName(name string) {
	b.name = name
}

func (b *Book) PrintText() {
	fmt.Println(b.text)
}


func NewBook(name, text string) *Book{
	return &Book{name: name, text: text, exist: true}
}