package stingray

// fileResource represents a file resource.
type fileResource struct {
	configResource
	Content []byte
	Note    string
}

//Bytes will return back the content as a array of bytes
func (f *fileResource) Bytes() []byte {
	return f.Content
}

func (f *fileResource) String() string {
	return string(f.Content)
}

func (f *fileResource) decode(data []byte) error {
	f.Content = data
	return nil
}

func (f *fileResource) contentType() string {
	return "application/octet-stream"
}
