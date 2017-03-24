package configprovider

import "testing"

// @TODO add more complex JSON data to test

// Message JSON bytes
var MessageBytes = []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
var MessageStruct = Message{
	Name: "Alice",
	Body: "Hello",
	Time: 1294706395881547000,
}

// Message struct
type Message struct {
	Name string
	Body string
	Time int64
}

func TestJsonConfig_Get(t *testing.T) {
	jb := &JsonBackend{
		connector: NewBufferedConnector("key", "scope", MessageBytes),
		usage:     &AllBackendUsage{},
	}

	if !jb.Handles("key", "scope") {
		t.Error("JsonBackend didn't Handle() properly")
	} else if c, err := jb.Get("key", "scope"); err != nil {
		t.Error("JsonBackend gave an error when retreiving valid key-scope Config")
	} else {
		var m Message
		res := c.Get(&m)
		<-res.Finished()

		if !res.Success() {
			t.Error("JsonBackend Config reported failure in Get() : ", res.Errors())
		} else {

			if m.Name != MessageStruct.Name {
				t.Error("JsonBackend provided incorrect data ==> Name : ", m.Name)
			}
			if m.Body != MessageStruct.Body {
				t.Error("JsonBackend provided incorrect data ==> Body : ", m.Body)
			}
			if m.Time != MessageStruct.Time {
				t.Error("JsonBackend provided incorrect data ==> Time : ", m.Time)
			}

		}
	}
}

func TestJsonConfig_Set(t *testing.T) {
	con := NewBufferedConnector("key", "scope", []byte{})
	jb := &JsonBackend{
		connector: con,
		usage:     &AllBackendUsage{},
	}

	if !jb.Handles("key", "scope") {
		t.Error("JsonBackend didn't Handle() properly")
	} else if c, err := jb.Get("key", "scope"); err != nil {
		t.Error("JsonBackend gave an error when retreiving valid key-scope Config")
	} else {
		res := c.Set(MessageStruct)
		//<-res.Finished()

		if !res.Success() {
			t.Error("JsonBackend Config reported a failure when Set(): ", res.Errors())
		} else {

		}
	}
}
