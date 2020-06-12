package visitor

import (
	"errors"
	"fmt"
)

// Entryに統合でOK
//type Element interface {
//}

type Entry interface {
	Accept(v Visitor)
	Name() string
	Size() int
	Entries() []Entry
}

type File struct {
	name string
	size int
}

func (f *File) String() string {
	return fmt.Sprintf("%s (%d)", f.Name(), f.Size())
}

func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Size() int {
	return f.size
}

func (f *File) Entries() []Entry {
	err := errors.New("FileTreatmentError.s")
	fmt.Println(err)
	return nil
}

func (f *File) Accept(v Visitor) {
	v.Visit(f)
}

type Directory struct {
	name string
	dir  []Entry
}

func (d *Directory) String() string {
	return fmt.Sprintf("%s (%d)", d.Name(), d.Size())
}

func NewDirectory(name string) *Directory {
	return &Directory{name: name}
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Size() int {
	sizeVisitor := NewSizeVisitor()
	d.Accept(sizeVisitor)
	return sizeVisitor.Size()
}

func (d *Directory) Entries() []Entry {
	return d.dir
}

func (d *Directory) Accept(v Visitor) {
	v.Visit(d)
}

func (d *Directory) Add(entry Entry) Entry {
	d.dir = append(d.dir, entry)
	return d
}
