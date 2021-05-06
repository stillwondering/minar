package xml

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stillwondering/minar"
)

func TestEncode(t *testing.T) {
	tt := []struct {
		Name        string
		Minutes     minar.Minutes
		ExpectedXML func() ([]byte, error)
	}{
		{
			Name: "Test1",
			Minutes: minar.Minutes{
				ID:           minar.MinutesID("awesomeID"),
				Title:        "First meeting minutes",
				Participants: []string{"Me", "You"},
				Topics: []minar.Topic{
					{
						Title:   "First topic",
						Content: "That's what we discussed in here",
					},
				},
			},
			ExpectedXML: func() ([]byte, error) {
				return ioutil.ReadFile("testdata/Test1.xml")
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			expected, err := tc.ExpectedXML()
			if err != nil {
				t.Fatalf("expected to load expected XML, got: %v", err)
			}

			enc, err := Encode(tc.Minutes)
			if err != nil {
				t.Fatalf("expected no error, vot %v", err)
			}

			if !bytes.Equal(enc, expected) {
				t.Errorf("expected: %s\ngot: %s", expected, enc)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tt := []struct {
		Name            string
		Input           func() ([]byte, error)
		ExpectedMinutes minar.Minutes
	}{
		{
			Name: "Test1",
			Input: func() ([]byte, error) {
				return ioutil.ReadFile("testdata/Test1.xml")
			},
			ExpectedMinutes: minar.Minutes{
				ID:           minar.MinutesID("awesomeID"),
				Title:        "First meeting minutes",
				Participants: []string{"Me", "You"},
				Topics: []minar.Topic{
					{
						Title:   "First topic",
						Content: "That's what we discussed in here",
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			input, err := tc.Input()
			if err != nil {
				t.Fatalf("expected to load input XML, got: %v", err)
			}

			dec, err := Decode(input)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if !dec.Equals(tc.ExpectedMinutes) {
				t.Errorf("expected: %v, got: %v", tc.ExpectedMinutes, dec)
			}
		})
	}
}
