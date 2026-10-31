package embedbinwrapper

/*
An embed executable source.
*/
type Src struct {
	// The raw executable in binary.
	bin []byte

	os   string
	arch string
}

// NewSrc creates new Src instance
func NewSrc() *Src {
	return &Src{}
}

// Os tie the source to a specific OS. Possible values are same as runtime.GOOS
func (s *Src) Os(value string) *Src {
	s.os = value
	return s
}

// Arch tie the source to a specific arch. Possible values are same as runtime.GOARCH
func (s *Src) Arch(value string) *Src {
	s.arch = value
	return s
}

// Sets the raw binary for this Src.
func (s *Src) Bin(value []byte) *Src {
	s.bin = value
	return s
}
