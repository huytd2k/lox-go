package scanner

type Token struct {
}

type Scanner struct {
	source []byte
}

func (scanner *Scanner) readSource(source []byte) {
	scanner.source = source
}
