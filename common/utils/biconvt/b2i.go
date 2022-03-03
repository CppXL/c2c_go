package biconv

// by https://gist.github.com/chiro-hiro/2674626cebbcb5a676355b7aaac4972d

func Btoi64(val []byte) uint64 {

	return btoi64(val)
}

func Btoi32(val []byte) uint32 {
	return btoi32(val)
}

func btoi64(val []byte) uint64 {
	r := uint64(0)
	for i := uint64(0); i < 8; i++ {
		r |= uint64(val[i]) << (8 * i)
	}
	return r
}

func btoi32(val []byte) uint32 {
	r := uint32(0)
	for i := uint32(0); i < 4; i++ {
		r |= uint32(val[i]) << (8 * i)
	}
	return r
}
