package biconv

// by https://gist.github.com/chiro-hiro/2674626cebbcb5a676355b7aaac4972d

func I32tob(val uint32) []byte {
	return i32tob(val)
}

func I64tob(val uint64) []byte {

	return i64tob(val)
}

func i32tob(val uint32) []byte {
	r := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		r[i] = byte((val >> (8 * i)) & 0xff)
	}
	return r
}
func i64tob(val uint64) []byte {
	r := make([]byte, 8)
	for i := uint64(0); i < 8; i++ {
		r[i] = byte((val >> (i * 8)) & 0xff)
	}
	return r
}
