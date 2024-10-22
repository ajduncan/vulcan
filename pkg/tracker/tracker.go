package tracker

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
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

// Serialize the tracker and return a string.
func (t *Tracker) Serialize() string {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	gob.Register(Tracker{})
	err := e.Encode(t)
	if err != nil {
		fmt.Printf("Encoding failed: %v\n", err)
	}
	return base64.StdEncoding.EncodeToString(b.Bytes())
}

// Deserialize takes an encoded string and returns a tracker.
func Deserialize(data string) *Tracker {
	t := &Tracker{}

	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		// sentry
		fmt.Printf("Decoding failed: %v\n", err)
	}

	b := bytes.Buffer{}
	b.Write(decoded)
	decoder := gob.NewDecoder(&b)

	gob.Register(Tracker{})
	err = decoder.Decode(&t)
	if err != nil {
		// sentry
		fmt.Printf("Decoding failed: %v\n", err)
	}

	return t
}

// Print prints diagnostic information about the tracker and payload.
func (t *Tracker) Print() {
	fmt.Printf("Tracker version: %v\n", t.Version)
	fmt.Printf("Tracker timestamp: %v\n", t.Timestamp)
	fmt.Printf("Tracker id: %v\n", t.Id)
	fmt.Printf("Tracker subject: %v\n", t.Subject)
	fmt.Printf("Payload: %v\n", t.Payload)
}
