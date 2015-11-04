package mutate

import (
	"fmt"
	"io/ioutil"
)

// Mask masks "program" at "nbyte" with "mask", and writes the result to "path"
func Mask(program []byte, nbyte int, mask byte, path string) (err error) {
	if nbyte > len(program) {
		err = fmt.Errorf("byte %d past end of file\n", nbyte)
		return
	}
	program[nbyte] ^= mask
	err = ioutil.WriteFile(path, program, 0777)
	return
}
