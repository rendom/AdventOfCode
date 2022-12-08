package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//go:embed data.txt
var input string

type FileI interface {
	GetSize() int
	GetParent() *Dir
}

type Dir struct {
	childs   []FileI
	parent   *Dir
	path     string
	filename string
}

func (d *Dir) GetParent() *Dir {
	return d.parent
}

func (d *Dir) GetSize() int {
	var s int
	for _, c := range d.GetChilds() {
		s += c.GetSize()
	}

	return s
}

func (d *Dir) GetChilds() []FileI {
	return d.childs
}

func (d *Dir) AddChild(f FileI) {
	d.childs = append(d.childs, f)
}

type File struct {
	size   int
	path   string
	parent *Dir
}

func (f *File) GetSize() int {
	return f.size
}

func (f *File) GetParent() *Dir {
	return f.parent
}

type State struct {
	filesystem  *Dir
	currentPath *Dir
}

func getSizeOfDirs(d *Dir) []int {
	sizes := []int{}
	sizes = append(sizes, d.GetSize())

	for _, v := range d.GetChilds() {
		if c, ok := v.(*Dir); ok {
			sizes = append(sizes, getSizeOfDirs(c)...)
		}
	}

	return sizes
}

func parseTermOutput() State {
	state := State{
		filesystem: &Dir{
			path:     "/",
			filename: "/",
		},
	}

	state.currentPath = state.filesystem

	// instruction starts with $
	output := []string{}
	cmd := ""
	args := ""
	for _, r := range strings.Split(input, "\n") {
		if len(r) == 0 {
			parseOutput(cmd, args, output, &state)
			continue
		}

		if r[0] == '$' {
			parseOutput(cmd, args, output, &state)
			cmd = r[2:4]
			args = r[4:]
			output = []string{}
			continue
		}

		output = append(output, r)
	}

	return state
}

func parseOutput(cmd string, args string, output []string, s *State) {
	args = strings.TrimSpace(args)

	switch cmd {
	case "cd":
		if args == ".." {
			if s.currentPath.parent != nil {
				s.currentPath = s.currentPath.parent
			}
			return
		}

		for _, c := range s.currentPath.GetChilds() {
			if d, ok := c.(*Dir); ok {
				if d.path == args {
					s.currentPath = d
					return
				}
			}
		}
	case "ls":
		re := regexp.MustCompile(`^(\w+) (.*)`)
		for _, r := range output {
			m := re.FindStringSubmatch(r)
			if m[1] == "dir" {
				f := Dir{}
				f.path = m[2]
				f.parent = s.currentPath
				s.currentPath.AddChild(&f)
			} else {
				f := File{}
				f.parent = s.currentPath
				f.size, _ = strconv.Atoi(m[1])
				f.path = m[2]
				s.currentPath.AddChild(&f)
			}
		}
	}
}

func answer1() {
	s := parseTermOutput()
	sum := 0
	for _, v := range getSizeOfDirs(s.filesystem) {
		if v <= 100000 {
			sum += v
		}
	}
	fmt.Printf("Answer 1: %d\n", sum)
}

func answer2() {
	diskSize := 70000000
	updateSize := 30000000

	s := parseTermOutput()
	freeSpace := diskSize - s.filesystem.GetSize()
	spaceNeeded := updateSize - freeSpace

	sizes := []int{}
	for _, v := range getSizeOfDirs(s.filesystem) {
		if v >= spaceNeeded {
			sizes = append(sizes, v)
		}
	}

	sort.SliceStable(sizes, func(i, j int) bool {
		return sizes[i] < sizes[j]
	})

	fmt.Printf("Answer 2: %d\n", sizes[0])
}

func main() {
	answer1()
	answer2()
}
