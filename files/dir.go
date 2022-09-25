package files

import "os"

func Mkdir(s string) {
	os.Mkdir(s, 0755)
}
