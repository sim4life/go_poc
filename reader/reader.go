package main

//import "code.google.com/p/go-tour/reader"
import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(b []byte) (i int, e error) {
	for x := range b {
		b[x] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
