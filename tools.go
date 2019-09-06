package pcircle

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
