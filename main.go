package main

import "fmt"

type Stringer interface {
	String() string
}

type Book struct {
	Title  string
	Author string
}

func (b *Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

type Library struct {
	Name  string
	Books []Book
}

func (l *Library) String() string {
	result := fmt.Sprintf("Library: %s\n", l.Name)
	for _, book := range l.Books {
		result += book.String() + "\n"
	}
	return result
}

func main() {
	library := &Library{
		Name: "City Library",
		Books: []Book{
			{Title: "كتاب الحجة على أهل المدينة", Author: "محمد بن حسن الشيباني"},
			{Title: "المبسوط", Author: "محمد بن حسن الشيباني"},
			{Title: "كتاب السير الكبير", Author: "محمد بن حسن الشيباني"},
		},
	}

	fmt.Println(library.Books)
}
