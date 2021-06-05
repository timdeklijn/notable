package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"
)

// Config holds the configuration for notable
type Config struct {
	folder  string // folder to notes
	vimPath string // path to vim on local machine
}

// NewConfig creates a Config
func NewConfig(folder string) Config {
	vimPath := "/usr/local/bin/nvim"
	return Config{folder: folder, vimPath: vimPath}
}

// File holds the parameters to create a new file
type File struct {
	path   string
	date   string
	config Config
}

// NewFile creates a new file from the arguments
func NewFile(config Config) {
	date := getDate()
	p := path.Join(config.folder, date+".md")
	f := File{path: p, date: date, config: config}
	if !f.exists() {
		f.createNewNote()
		f.openFile()
	} else {
		f.openFile()
	}
}

// exists checks if the path in File exists and returns
// true or false.
func (f *File) exists() bool {
	if _, err := os.Stat(f.path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
	}
}

// createNewNote creates a file and prints a writes template
// to the file.
func (f *File) createNewNote() {
	ff, err := os.Create(f.path)
	if err != nil {
		panic(err)
	}
	s := fmt.Sprintf("# %s\n\n", f.date)
	ff.WriteString(s)
}

// opeFile simply opens a file in neovim when not in use
func (f *File) openFile() {
	// prepare the command
	cmd := exec.Command(f.config.vimPath, f.path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	// Run the command
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

// getDate returns the current date formatted like YYYY-MM-DD
func getDate() string {
	return time.Now().Format("2006-01-02")
}

func main() {
	config := NewConfig("/Users/timdeklijn/notes/logs")
	NewFile(config)
}
