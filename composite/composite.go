package composite

import "fmt"

type Entry interface {
	Name() string
	Size() int
	Add(entry Entry) Entry
	PrintList(prefix string)
}

type FileEntry struct {
	name string
	size int
}

func (e *FileEntry) Add(entry Entry) Entry {
	err := fmt.Errorf("%s", "FileTreatmentException")
	fmt.Println(err, e, entry)
	return nil
}

func NewFileEntry(name string, size int) *FileEntry {
	return &FileEntry{
		name: name,
		size: size,
	}
}

func (e *FileEntry) Name() string {
	return e.name
}

func (e *FileEntry) Size() int {
	return e.size
}

func (e *FileEntry) PrintList(prefix string) {
	fmt.Printf("%s/%s\n", prefix, e)
}

func (e *FileEntry) String() string {
	return fmt.Sprintf("%s (%d)", e.Name(), e.Size())
}

type DirectoryEntry struct {
	name      string
	directory []Entry
}

func (e *DirectoryEntry) Name() string {
	return e.name
}

func (e *DirectoryEntry) Size() int {
	size := 0
	for _, entry := range e.directory {
		size += entry.Size()
	}
	return size
}

func (e *DirectoryEntry) PrintList(prefix string) {
	fmt.Printf("%s/%s\n", prefix, e)
	for _, entry := range e.directory {
		entry.PrintList(fmt.Sprintf("%s/%s", prefix, e.name))
	}
}

func (e *DirectoryEntry) Add(entry Entry) Entry {
	e.directory = append(e.directory, entry)
	return e
}

func NewDirectoryEntry(name string) *DirectoryEntry {
	return &DirectoryEntry{
		name: name,
	}
}
func (e *DirectoryEntry) String() string {
	return fmt.Sprintf("%s", e.Name())
}
