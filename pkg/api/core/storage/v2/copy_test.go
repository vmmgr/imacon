package v2

import (
	"strings"
	"testing"
)

func Test(t *testing.T) {
	dstPath := "/home/yonedayuto/vm-image/test0/1.img"
	dstPathArray := strings.Split(dstPath, "/")
	t.Log(dstPath[:len(dstPath)-(len(dstPathArray))])
}
