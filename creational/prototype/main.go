package main

import "fmt"

func main() {
	//the Prototype pattern lets us copy data structure, without coupling the code to their classes (Post, Forum, Category).
	var (
		p1 = &Post{name: "Array"}
		p2 = &Post{name: "Linked List"}
		p3 = &Post{name: "Stack"}
		p4 = &Post{name: "Queue"}
		p5 = &Post{name: "Tree"}
		p6 = &Post{name: "Graph"}

		f1 = &Forum{name: "Static", posts: []Noder{p1}}
		f2 = &Forum{name: "Dynamic", posts: []Noder{p2, p3, p4}}
		f3 = &Forum{name: "None Linear", posts: []Noder{p5, p6}}
		f4 = &Forum{name: "Linear", posts: []Noder{f1, f2}}
		f5 = &Forum{name: "Primitive Data Structure", posts: []Noder{}}
		f6 = &Forum{name: "Non Primitive Data Structure", posts: []Noder{f3, f4}}
		f7 = &Category{name: "Data structure", forums: []Noder{f5, f6}}
	)

	fmt.Println(`Original tree:`)
	f7.Print("-")
	fmt.Println(`Cloned tree`)
	f8 := f7.Clone()
	f8.Print("-")
}
