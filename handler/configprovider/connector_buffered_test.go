package configprovider

import (
	"bytes"
	"io"
	"testing"
)

func TestBufferedConnector_Keys(t *testing.T) {
	bc := NewBufferedConnector("key", "scope", []byte{})

	k := bc.Keys()
	if len(k) != 1 {
		t.Error("BufferedConnector return the wrong number of ket", k)
	} else if k[0] != "key" {
		t.Error("BufferedConnector return incorrect key values")
	}
}

func TestBufferedConnector_Scopes(t *testing.T) {
	bc := NewBufferedConnector("key", "scope", []byte{})

	s := bc.Scopes()
	if len(s) != 1 {
		t.Error("BufferedConnector return the wrong number of scopes", s)
	} else if s[0] != "scope" {
		t.Error("BufferedConnector return incorrect scope values")
	}
}

func TestBufferedConnector_Get(t *testing.T) {
	bc := NewBufferedConnector("key", "scope", []byte("test"))

	if _, err := bc.Get("no", "no"); err == nil {
		t.Error("BufferredConnector did not return an error when incorrect key-scope was requested")
	} else if rc, err := bc.Get("key", "scope"); err != nil {
		t.Error("BufferredConnector returned an error when asked for a valid key-scope pair")
	} else {
		b := bytes.NewBufferString("")
		b.ReadFrom(rc)
		rc.Close()
		val := b.String()
		if val != "test" {
			t.Error("BufferedConnector reader has the wrong string: ", val)
		}
	}
}

func TestBufferedConnector_Set(t *testing.T) {
	bc := NewBufferedConnector("key", "scope", []byte("one"))

	buf := BufferCloser{Buffer: *bytes.NewBufferString("two")}
	if err := bc.Set("key", "scope", io.ReadCloser(&buf)); err != nil {
		t.Error("BufferedConnector returned an error when setting a new value: ", err.Error())
	} else {
		if rc, err := bc.Get("key", "scope"); err != nil {
			t.Error("BufferredConnector returned an error when asked for a valid key-scope pair")
		} else {
			b := bytes.NewBufferString("")
			b.ReadFrom(rc)
			rc.Close()
			val := b.String()
			if val != "two" {
				t.Error("BufferedConnector reader has the wrong string: ", val)
			} else if val == "one" {
				t.Error("BufferedConnector reader has the wrong string, it still has it's original value: ", val)
			}
		}
	}
}
