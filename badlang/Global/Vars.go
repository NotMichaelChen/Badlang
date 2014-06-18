package Global

//This is the list of variables that is accessable to the whole program
//Any function is allowed to read and write to it
var Variables map[string]float64 = make(map[string]float64)

func VarType(contents string) byte {

	for _, char := range contents {
		if (65 <= char && char <= 90) || (97 <= char && char <= 122) {
			return 's'
		}
	}

	return 'n'

}
