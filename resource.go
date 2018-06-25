package stingray

// Resourcer is the interface implemented by objects that represent
// Stingray configuration resources.
type Resourcer interface {
	Name() string
	setName(string)
	endpoint() string
	pathType() pathType
	contentType() string
	decode([]byte) error
	Bytes() []byte
}

type baseResource struct {
	name        string
	contentType string
}


func (r *baseResource) Name() string {
	return r.name
}

func (r *baseResource) setName(name string) {
	r.name = name
}


type configResource struct {
	baseResource
}

func (r *configResource) pathType() pathType {
	return configPath
}

type statsResource struct {
	baseResource
}

func (r *statsResource) pathType() pathType {
	return statsPath
}

