package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type aocFile struct {
	size int
	name string
}

type aocDirectory struct {
	parent *aocDirectory
	dirs []*aocDirectory
	files []*aocFile
	name string
}


func main() {

	var root aocDirectory
	var d, cwd *aocDirectory
	root.name = "/"
	root.parent = nil

	scanner := bufio.NewScanner(os.Stdin)

	for (scanner.Scan() ) {
		line := scanner.Text()
		//fmt.Println(line)
		word := strings.Split(line," ")
		if word[1] == "cd" {
			if word[2] == "/" {
				cwd = &root
				//fmt.Print(*cwd )
				//fmt.Printf("%p\n", cwd)
			} else if word[2] ==".." {
				if (cwd.parent == nil) {
					//fmt.Println("No parent!")
					os.Exit(1)
				}
				cwd = cwd.parent
				//fmt.Print(*cwd )
				//fmt.Printf("%p\n", cwd)
			} else {
				//fmt.Println("adding ",word[2], " in ",cwd.name)
				d = new(aocDirectory)
				d.parent = cwd
				d.name = word[2]
				cwd.dirs = append( cwd.dirs, d)
				//fmt.Print(*cwd )
				//fmt.Printf("%p\n", cwd)
				cwd = findDir(word[2],*cwd)
				//fmt.Print(*cwd )
				//fmt.Printf("%p\n", cwd)
			}
		}
		size, err := strconv.Atoi(word[0])
		if err == nil {
			AddFile(word[1], size, cwd)
			//fmt.Println("added ",word[1], " in ",cwd.name)
			//fmt.Print(*cwd )
			//fmt.Printf("%p\n", cwd)
		}
	}

	DirSize(&root)
}

func findDir(name string, dir aocDirectory) *aocDirectory {
	for _, d := range dir.dirs {
		if name == d.name {
			//fmt.Println("found ",name, " in ",dir.name)
			return d
		}
	}
	return &dir
}

func AddFile(name string, size int, dir *aocDirectory) *aocDirectory {
	file := new(aocFile)
	file.name = name
	file.size = size
	dir.files = append( dir.files, file)
	return dir
}

func DirSize(dir *aocDirectory) int {
	var size int
	for _, f := range dir.files {
		size += f.size
	}
	for _, d := range dir.dirs {
		size += DirSize(d)
	}
	fmt.Println("dir: ", dir.name, " size: ", size)
	return size
}