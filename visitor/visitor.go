package visitor

import (
	"fmt"
	"strings"
)

type Visitor interface {
	Visit(entry Entry)
}

type ListVisitor struct {
	currentDir string
}

func (v *ListVisitor) Visit(entry Entry) {
	fmt.Printf("%s/%s\n", v.currentDir, entry)
	if _, ok := entry.(*Directory); ok {
		saveDir := v.currentDir
		v.currentDir = fmt.Sprintf("%s/%s", v.currentDir, entry.Name())
		for _, e := range entry.Entries() {
			e.Accept(v)
		}
		v.currentDir = saveDir
	}
}

func NewListVisitor() *ListVisitor {
	return &ListVisitor{currentDir: ""}
}

type FileFindVisitor struct {
	filetype string
	found    []Entry
}

func NewFileFindVisitor(filetype string) *FileFindVisitor {
	return &FileFindVisitor{filetype: filetype}
}

func (v *FileFindVisitor) Visit(entry Entry) {
	if _, ok := entry.(*Directory); ok {
		for _, e := range entry.Entries() {
			e.Accept(v)
		}
	}
	if strings.HasSuffix(entry.Name(), v.filetype) {
		v.found = append(v.found, entry)
	}
}

func (v *FileFindVisitor) FoundFiles() []Entry {
	return v.found
}

type SizeVisitor struct {
	size int
}

func NewSizeVisitor() *SizeVisitor {
	return &SizeVisitor{}
}

func (v *SizeVisitor) Size() int {
	return v.size
}

func (v *SizeVisitor) Visit(entry Entry) {
	if _, ok := entry.(*Directory); ok {
		for _, entry := range entry.Entries() {
			entry.Accept(v)
		}
	} else {
		v.size += entry.Size()
	}
}
