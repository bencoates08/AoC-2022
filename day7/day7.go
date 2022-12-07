package day7

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_SPACE      = 70000000
	REQUIRED_SPACE = 30000000
)

type Directory struct {
	Name     string
	Parent   *Directory
	Children []*Directory
	Files    []File
}

type File struct {
	Name string
	Size int
}

func NewDirectory(name string, parent *Directory, children []*Directory, files []File) *Directory {
	return &Directory{name, parent, children, files}
}

func NewFile(name string, size int) File {
	return File{name, size}
}

func (p *Directory) parseOutput(command string, output []string) *Directory {
	cmd := strings.Split(command, " ")
	switch cmd[1] {
	case "cd":
		return p.cd(cmd[2])
	case "ls":
		p.ls(output)
	}
	return p
}

func (p *Directory) cd(cd string) *Directory {
	if cd == "/" {
		root := p
		for root.Parent != nil {
			root = root.Parent
		}
		return root
	}
	if cd == ".." {
		return p.Parent
	} else {
		for _, v := range p.Children {
			if v.Name == cd {
				return v
			}
		}
	}
	return p
}

func (d *Directory) ls(output []string) {
	if len(d.Children) > 0 {
		return
	}
	for _, v := range output {
		w := strings.Split(v, " ")
		if w[0] == "dir" {
			d.Children = append(d.Children, NewDirectory(w[1], d, nil, nil))
		} else {
			size, _ := strconv.Atoi(w[0])
			d.Files = append(d.Files, NewFile(w[1], size))
		}
	}
}

func (d *Directory) totalFilesOverSize(max int, total int) (int, int) {
	size := 0

	for _, v := range d.Files {
		size += v.Size
	}
	for _, v := range d.Children {
		childrenSize := 0
		childrenSize, total = v.totalFilesOverSize(max, total)
		size += childrenSize
	}

	if size <= max {
		return size, total + size
	}
	return size, total
}

func (d *Directory) findDirsOverSize(min int, dirSizes []int) (int, []int) {
	size := 0

	for _, v := range d.Files {
		size += v.Size
	}
	for _, v := range d.Children {
		childrenSize := 0
		childrenSize, dirSizes = v.findDirsOverSize(min, dirSizes)
		size += childrenSize
	}

	if size >= min {
		return size, append(dirSizes, size)
	}
	return size, dirSizes
}

func Part1and2() (int, int) {
	readFile, _ := os.Open("day7/input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()

	directory := NewDirectory("/", nil, nil, nil)

	command := ""
	output := []string{}

	eof := true
	for fileScanner.Err() == nil {
		command = fileScanner.Text()

		for eof {
			eof = fileScanner.Scan()
			line := fileScanner.Text()

			if len(line) > 0 && line[0] != '$' {
				output = append(output, line)
				continue
			}

			break
		}

		directory = directory.parseOutput(command, output)

		output = []string{}

		if !eof {
			break
		}
	}

	directory = directory.cd("/")
	size, total := directory.totalFilesOverSize(100000, 0)

	// Part 2
	availableSpace := MAX_SPACE - size
	spaceNeeded := REQUIRED_SPACE - availableSpace

	_, dirsSizes := directory.findDirsOverSize(spaceNeeded, []int{})

	min := 0
	for i, e := range dirsSizes {
		if i == 0 || e < min {
			min = e
		}
	}

	return total, min
}
