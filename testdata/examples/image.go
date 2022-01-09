package examples

//go:generate foggo foc --struct Image
type Image struct {
	Width  int
	Height int
	// don't want to create option, specify `foggo:"-"` as the structure tag
	Src string `foggo:"-"`
	Alt string
}
