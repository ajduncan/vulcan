package tracker

// WCSPayload provides the most flexibility to arbitrarily inject key/value pairs
// which ultimately go into a nosql store (a wide column store) for analytics.
type WCSPayload struct {
	wcs map[string]string
}

// Allocate and return a new WCSPayload, accepts a map[string][]string from
// net/url ParseQuery.
func NewPayload(urldict map[string][]string) *WCSPayload {
	p := &WCSPayload{
		wcs: make(map[string]string),
	}
	p.AddURLDict(urldict)
	return p
}

// Add accepts a key value pair and adds them to the store.
// If either key or value is empty, discard the pair.
func (p *WCSPayload) Add(key string, value string) {
	if key != "" && value != "" {
		// pre-validate certain keys?
		p.wcs[key] = value
	}
}

// AddURLDict takes a map[string][]string, (from net/url ParseQuery, for example)
// and adds the key value pair
func (p *WCSPayload) AddURLDict(m map[string][]string) {
	// convert m to map[string]string
	for key, _ := range m {
		if len(m[key]) > 0 {
			p.Add(key, m[key][0])
		}
	}
}

// AddDict accepts a map of key value pairs and adds them to the store
func (p *WCSPayload) AddDict(wcs map[string]string) {
	for key, value := range wcs {
		p.Add(key, value)
	}
}

// Get accepts a key and returns a value from the store.
func (p *WCSPayload) Get(key string) string {
	return p.wcs[key]
}
