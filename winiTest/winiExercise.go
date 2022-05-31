package main

type Book struct {
	pageMap map[int]Page
}

type Page struct {
	info string
	infoLenght int
}

func main() {
	testBook := Book{}
	testBook.pageMap[0] = Page{}
}