package CodeParser

import "io/ioutil"

func CodeParse(filename string) []string {

	code, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Something is wrong with the file being opened (CodeParse)")
	}

	var codeslist []string
	var tempcode string

	for _, letter := range code {

		if letter == ':' {
			codeslist = append(codeslist, tempcode)
			tempcode = ""

		} else if letter == '\n' {
			continue

		} else {
			tempcode += string(letter)

		}
	}
	codeslist = append(codeslist, tempcode)

	return codeslist

}
