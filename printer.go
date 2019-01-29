package tree

import "fmt"

// Node is an abstraction of data string and tree topology.
type (
	Node interface {
		GetString() string
		GetChildren() []Node
	}
	Color int
)

const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
)

// Traverse prints out the whole tree structure.
func Traverse(root Node, indent string, depth int) {
	fmt.Println(root.GetString())
	if depth == 0 {
		return
	}
	children := root.GetChildren()
	for i, child := range children {
		add := "│   "
		if i == len(children)-1 {
			fmt.Print(indent + "└── ")
			add = "    "
		} else {
			fmt.Print(indent + "├── ")
		}
		Traverse(child, indent+add, depth-1)
	}
}

// Colorize add terminal color to string s.
func Colorize(s string, color Color) string {
	return fmt.Sprintf("\x1b[1;%vm%s\x1b[0m", color, s)
}
