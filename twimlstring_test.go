package twiml

import (
	"testing"
)

func TestTwimlStringParsing(t *testing.T) {

	s := "{say|Hello Test}"
	m := ParseString(s)

	if m == nil || len(m) == 0 {
		t.Errorf("Failed to parse string")
		return
	}

	if m[0].Type() != "Say" {
		t.Errorf("Failed to parse into correct parameters")
		return
	}
}

func TestTwimlStringAttributes(t *testing.T) {

	s := "{say|Hello Test,voice:man}"
	m := ParseString(s)

	if m == nil || len(m) == 0 {
		t.Errorf("Failed to parse string")
		return
	}

	if m[0].Type() != "Say" {
		t.Errorf("Failed to parse into correct type. Expected Say got %s", m[0].Type())
		return
	}

	d := m[0].(*Say)
	if d.Voice != "man" {
		t.Errorf("Failed to parse into correct attributes. Expected man got %s", d.Voice)
		return
	}
}


func TestTwimlStringOutput(t *testing.T) {

	s := "{p|This is a spoken paragraph}{strong|This is a strongly worded}"
	m := ParseString(s)

	if m == nil || len(m) == 0 {
		t.Errorf("Failed to parse string")
		return
	}

	if m[0].Type() != "Say" {
		t.Errorf("Failed to parse into correct type. Expected Say got %s", m[0].Type())
		return
	}

	d := m[0].(*Say)

	if len(d.Children) == 0 {
		t.Errorf("Failed to parse SSML")
		return
	}

	if d.Children[0].Type() != "SSMLText" {
		t.Errorf("Failed to parse into correct SSML type. Expected SSMLText got %s", d.Children[0].Type())
		return
	}

}


func TestTwimlQuotedString(t *testing.T) {

	s := "{say|`Quoted, String`,voice:female}"
	m := ParseString(s)

	if m == nil || len(m) == 0 {
		t.Errorf("Failed to parse string")
		return
	}

	if m[0].Type() != "Say" {
		t.Errorf("Failed to parse into correct type. Expected Say got %s", m[0].Type())
		return
	}

	d := m[0].(*Say)
	if d.Text != "Quoted, String" {
		t.Errorf("Failed to parse into correct quoted string. Got %s", d.Text)
		return
	}

	if d.Voice != "female" {
		t.Errorf("Failed to parse into correct attributes. Expected female got %s", d.Voice)
		return
	}
}

func TestTwimlMultiMixQuotedString(t *testing.T) {

	s := "{say|`Quoted, String`,voice:female}{say|`Another quote`}{say|Whats happening?}"
	m := ParseString(s)

	if m == nil || len(m) != 3 {
		t.Errorf("Failed to parse string")
		return
	}

	if m[0].Type() != "Say" {
		t.Errorf("Failed to parse into correct type. Expected Say got %s", m[0].Type())
		return
	}

	a := m[0].(*Say)
	if a.Text != "Quoted, String" {
		t.Errorf("Failed to parse into correct quoted string. Got %s", a.Text)
		return
	}

	b := m[1].(*Say)
	if b.Text != "Another quote" {
		t.Errorf("Failed to parse into correct quoted string. Got %s", b.Text)
		return
	}

	c := m[2].(*Say)
	if c.Text != "Whats happening?" {
		t.Errorf("Failed to parse into correct quoted string. Got %s", c.Text)
		return
	}
}