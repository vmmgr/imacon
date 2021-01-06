package v0

import (
	"os"
)

func fileExistsCheck(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func AddCheck() {

}
