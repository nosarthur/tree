package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/nosarthur/tree"
)

type path string

func (p *path) GetString() string {
	// FIXME: It's wasteful to do os.Stat twice.
	//	      But I don't really care about this path display tree.
	// 		  Maybe won't fix.
	fname := string(*p)
	fi, err := os.Stat(fname)
	if err != nil {
		log.Fatalf("cannot stat %s: %v", fname, err)
	}

	if fi.IsDir() {
		return tree.Colorize(filepath.Base(fname), tree.Blue)
	}
	return filepath.Base(fname)
}

func (p *path) GetChildren() []tree.Node {
	fname := string(*p)
	fi, err := os.Stat(fname)
	if err != nil {
		log.Fatalf("cannot stat %s: %v", fname, err)
	}

	if fi.IsDir() {
		fis, err := ioutil.ReadDir(fname)
		if err != nil {
			fmt.Printf("failed to read dir %s: v\n", fname, err)
			return nil
		}
		children := []tree.Node{}
		for _, fi := range fis {
			// skip hidden file/folder
			if fi.Name()[0] == '.' {
				continue
			}
			kid := path(filepath.Join(fname, fi.Name()))
			children = append(children, &kid)
		}
		return children
	}
	return nil
}

func main() {
	depthPtr := flag.Int("depth", -1, "tree depth")
	flag.Parse()

	args := flag.Args()
	names := []string{"./"}
	if len(args) > 0 {
		names = args
	}
	for _, name := range names {
		root := path(name)
		tree.Traverse(&root, "", *depthPtr)
	}
}
