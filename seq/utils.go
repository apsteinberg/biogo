package seq

func Reverse(s []byte) []byte {
	bs := []byte{}
	for i := 0; i < len(s); i++ {
		bs = append(bs, s[i])
	}

	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}

	return bs
}

func Complement(s []byte) []byte {
	m := make(map[byte]byte)
	m['A'] = 'T'
	m['T'] = 'A'
	m['G'] = 'C'
	m['C'] = 'G'
	bs := []byte{}
	for i := 0; i < len(s); i++ {
		bs = append(bs, m[s[i]])
	}

	return bs
}
