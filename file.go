package main

type file struct {
	filetName    string
	fileType     string
	size         int
	lastUpdate   string
	creationDate string
	data         []byte
}

/* even though I pass a non pointer as a receiver, go infere the pointer for us
Canonical way of doing this call:
var f := file {...}
var p_f = &f
p_f.updateName("whatever")

Go infered way of doing it:
var f := file {...}
f.updateName("whatever")
*/
func (p_file *file) updateName(newName string) {
	/* same happens here, although the Canonical way is doing this:
	(*p_file).fileName = ...
	we can do it like this:
	p_file.fileName = ...
	Go inferes the pointer for us */
	p_file.filetName = newName
}
