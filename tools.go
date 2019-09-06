package pcircle

import (
	"errors"
	"strings"
)

func bool2int(b bool) (i int) {
	if b {
		i = 1
	}
	return i
}

func bool2int2string(b bool) (s string) {
	s = "0"
	if b {
		s = "1"
	}
	return s
}

func string2int2bool(s string) (b bool, err error) {
	switch s {
	case "0":
		return false, nil
	case "1":
		return true, nil
	}
	return false, errors.New("failed to convert string to int to bool: invalid string " + s)
}

func tokenize(str string) (head, data string) {
	sep := strings.IndexRune(str, ':')
	return strings.TrimSpace(str[:sep]), strings.TrimSpace(str[sep+1:])
}
