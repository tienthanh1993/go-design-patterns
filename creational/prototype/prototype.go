package main

import "fmt"

type Noder interface {
	Name() string
	Print(string)
	Clone() Noder
}

//Post a forum post
type Post struct {
	name string
}

func (f Post) Name() string {
	return f.name
}

func (f Post) Print(s string) {
	fmt.Println(s + f.name)
}

func (f Post) Clone() Noder {
	clone := &Post{name: f.name + "_clone"}
	return clone
}

//Forum contains all post in this subject
type Forum struct {
	name  string
	posts []Noder
}

func (f Forum) Name() string {
	return f.name
}

func (f Forum) Print(s string) {
	fmt.Println(s + f.name)
	for _, p := range f.posts {
		p.Print(s + s)
	}
}

func (f Forum) Clone() Noder {
	clone := &Forum{name: f.name + "_clone"}
	for _, p := range f.posts {
		clone.posts = append(clone.posts, p.Clone())
	}
	return clone
}

//Category only contains forum
type Category struct {
	name   string
	forums []Noder
}

func (c Category) Name() string {
	return c.name
}

func (c Category) Print(s string) {
	fmt.Println(s + c.name)
	for _, f := range c.forums {
		f.Print(s + s)
	}
}

func (c Category) Clone() Noder {
	clone := &Category{name: c.name + "_clone"}
	for _, f := range c.forums {
		clone.forums = append(clone.forums, f.Clone())
	}
	return clone
}
