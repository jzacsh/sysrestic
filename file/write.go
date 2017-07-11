package file

import "io/ioutil"

// WriteASCIILines writes each element of 'strs' as ASCII bytes, each separated
// by a newline, into a new file at path 'file'
func WriteASCIILines(strs []string, file string) error {
	var lines []byte
	for _, ln := range strs {
		lines = append(lines, []byte(ln)...)
		lines = append(lines, '\n')
	}

	return ioutil.WriteFile(file, lines, 0644)
}
