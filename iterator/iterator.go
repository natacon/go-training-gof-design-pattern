package iterator

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func NewBookShelfIterator(bookShelf *BookShelf) Iterator {
	return &BookShelfIterator{bookShelf: bookShelf}
}

func (b *BookShelfIterator) HasNext() bool {
	if b.index < b.bookShelf.Length() {
		return true
	} else {
		return false
	}
}

func (b *BookShelfIterator) Next() interface{} {
	book := b.bookShelf.GetBookAt(b.index)
	b.index++
	return *book
}
