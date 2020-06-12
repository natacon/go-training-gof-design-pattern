package iterator

type Aggregate interface {
	Iterator() Iterator
}

type BookShelf struct {
	books []*Book
	last  int
}

func NewBookShelf() *BookShelf {
	return &BookShelf{}
}

func (bs *BookShelf) Iterator() Iterator {
	return NewBookShelfIterator(bs)
}

func (bs *BookShelf) GetBookAt(index int) *Book {
	return bs.books[index]
}

func (bs *BookShelf) Append(book *Book) {
	bs.books = append(bs.books, book)
	bs.books[bs.last] = book
	bs.last++
}

func (bs *BookShelf) Length() int {
	return bs.last
}
