package day7

// comFile represents a single file stored on an elf's communication device
type comFile struct {
	name string
	size int
}

// comDirectory represents a single directory on an elf's communication device
type comDirectory struct {
	name        string
	parent      *comDirectory
	directories []*comDirectory
	files       []*comFile
}

// totalSize will return the size of all comFile objections in a comDirectory
func (dir *comDirectory) totalSize() (totalSize int) {
	for _, file := range dir.files {
		totalSize += file.size
	}
	for _, childDir := range dir.directories {
		totalSize += childDir.totalSize()
	}
	return totalSize
}

// path will return the absolute path of the current directory
func (dir *comDirectory) path() string {
	path := "/" + dir.name
	if dir.name == "/" {
		path = dir.name
	}

	if dir.parent != nil && dir.parent.name != "/" {
		path = dir.parent.path() + path
	}
	return path
}

// comTerminal represents a single terminal session on an elf's communication device
type comTerminal struct {
	curDir *comDirectory
}

func newTerminal() *comTerminal {
	term := &comTerminal{curDir: nil}
	return term
}

// ExeReceipt provides some helpful information regarding the result of a command execution
type ExeReceipt struct {
	nextCommand string // this only gets populated for ls commands
	ok          bool   // Will be false when there are no more commands to execute from the channel
}
