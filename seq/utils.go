package seq

func Reverse(s string) string {
	bs := []byte{}
	for i := 0; i < len(s); i++ {
		bs = append(bs, s[i])
	}

	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}

	return string(bs)
}

func Complement(s string) string {
	m := make(map[byte]byte)
	m['A'] = 'T'
	m['T'] = 'A'
	m['G'] = 'C'
	m['C'] = 'G'
	bs := []byte{}
	for i := 0; i < len(s); i++ {
		bs = append(bs, m[s[i]])
	}

	return string(bs)
}
