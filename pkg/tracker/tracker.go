package tracker

import (
	"reflect"
	"time"
)

const version = 1 // internal tracker struct version.

// UAReferenceTracker is a reference struct that is not used.  This is not how
// we want to do things, and is only provided as a brief reference subset.
// For a similar set of values in katric, we may want to make use of gorilla
// scheme to save pairs to structs.
type UAReferenceTracker struct {
	// main
	cid string // client id
	de  string // document encoding type
	dl  string // document location
	dt  string // document title
	fl  string // flash version
	je  string // java enabled
	ni  string // Non-Interaction hit type
	_s  string // hit sequence - increments each time an event (inc pageview)
	sd  string // screen depth
	sr  string // screen resolution
	t   string // the Type of tracking call this is (e.g. pageview, event)
	tid string // tracking id / UA number
	_u  string // verification code
	ul  string // user language code
	_v  string // SDK version number
	v   string // protocol version
	vp  string // view port size
	z   string // cache buster

	// campaign
	cn string // campaign name
	cs string // campaign source
	cm string // campaign medium
	ck string // campaign keyword
	cc string // campaign content
	ci string // campaign id
}

type Tracker struct {
	Payload WCSPayload // payload structure kv pair

	Id        string    // id of what we're tracking.
	Subject   string    // ideally a way of associating page hits to a unique visit
	Timestamp time.Time // time stamp (of tracker creation)
	Version   int       // tracker version
}

// Allocate and return a new Tracker.  Requires a payload, id and subject.
func NewTracker(p *WCSPayload, id string, subject string) *Tracker {
	timestamp := time.Now()

	t := &Tracker{
		Payload:   *p,
		Id:        id,
		Subject:   subject,
		Timestamp: timestamp,
		Version:   version,
	}
	return t
}

// Add accepts a WCSPayload from a map[string]string kv pair and initializes
// the tracker's payload.
func (t *Tracker) AddPayload(p *WCSPayload) {
	t.Payload = *p
}

// AddProperty adds a property from the payload, if the property for the tracker
// exists.
// ex: t.AddProperty("Subject", "www.example.biz")
func (t *Tracker) AddProperty(key string, value string) {
	v := reflect.ValueOf(t).Elem().FieldByName(key)
	if v.IsValid() {
		v.SetString(value)
	}
}
