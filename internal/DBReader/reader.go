package dbreader

type Resipe struct {
	Name        string   `json:"name" xml:"name"`
	Ingredients []string `json:"ingredients" xml:"ingredients>ingredient"`
	Time        int      `json:"stovetime" xml:"time"`
}

type DBReader interface {
	Read(filename string) ([]Resipe, error)
	Print(resipes []Resipe) error
}

type JSONReader struct{}

type XMLReader struct{}

func (r JSONReader) Read(filename string) ([]Resipe, error) {
	return nil, nil
}

func (r XMLReader) Read(filename string) ([]Resipe, error) {
	return nil, nil
}

func (r JSONReader) Print(resipes []Resipe) {

}

func (r XMLReader) Print(resipes []Resipe) {

}
