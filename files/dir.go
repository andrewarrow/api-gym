package files

import "os"

func MkdirJson() {
	os.Mkdir("json", 0755)
}
