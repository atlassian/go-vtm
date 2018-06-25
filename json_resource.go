package stingray

// jsonResource represents a JSON resource.
type jsonConfigResource struct {
	configResource
}

func (f *jsonConfigResource) contentType() string {
	return "application/json"
}

type jsonStatsResource struct {
	statsResource
}

func (f *jsonStatsResource) contentType() string {
	return "application/json"
}
