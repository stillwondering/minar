package fs

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stillwondering/minar"
)

func TestFindAll(t *testing.T) {
	tt := []struct {
		name            string
		rootDir         string
		returnsError    bool
		expectedMinutes []minar.Minutes
	}{
		{
			name:            "EmptyRootDir",
			rootDir:         ".",
			returnsError:    false,
			expectedMinutes: []minar.Minutes{},
		},
		{
			name:            "InvalidXML",
			rootDir:         "testdata/InvalidXML",
			returnsError:    true,
			expectedMinutes: []minar.Minutes{},
		},
		{
			name:            "InvalidRecord",
			rootDir:         "testdata/InvalidRecord",
			returnsError:    true,
			expectedMinutes: []minar.Minutes{},
		},
		{
			name:         "OneRecord",
			rootDir:      "testdata/OneRecord",
			returnsError: false,
			expectedMinutes: []minar.Minutes{
				{
					ID:           minar.MinutesID("id"),
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
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			repo := MinutesRepository{
				BaseDir: tc.rootDir,
			}

			minutes, err := repo.FindAll()

			if tc.returnsError {
				returnedError := err != nil
				if returnedError != tc.returnsError {
					t.Fatalf("expected returnError = %v, got %v", tc.returnsError, returnedError)
				}

				return
			}

			if len(minutes) != len(tc.expectedMinutes) {
				t.Fatalf("expected lenght: %d, got: %d", len(tc.expectedMinutes), len(minutes))
			}

			for i := 0; i < len(tc.expectedMinutes); i++ {
				if !tc.expectedMinutes[i].Equals(minutes[i]) {
					t.Fatalf("expected: %v, got: %v", tc.expectedMinutes[i], minutes[i])
				}
			}
		})
	}
}

func TestCreate(t *testing.T) {
	tt := []struct {
		name        string
		data        minar.CreateMinutesData
		idGenerator minar.IDGeneratorFunc
		expected    func() ([]byte, error)
	}{
		{
			name: "",
			data: minar.CreateMinutesData{
				Title:        "Taking the hobbits to Isengard",
				Participants: []string{"Legolas", "Gimli"},
				Topics: []minar.Topic{
					{
						Title:   "First topic",
						Content: "Something we talked about.",
					},
					{
						Title:   "Second topic",
						Content: "Something we talked about.",
					},
				},
			},
			idGenerator: func() minar.MinutesID {
				return minar.MinutesID("1")
			},
			expected: func() ([]byte, error) {
				return ioutil.ReadFile("testdata/CreateMinutes/1.xml")
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			expected, err := tc.expected()
			if err != nil {
				t.Fatalf("expected to read golden file, got: %v", err)
			}

			root := t.TempDir()
			repo := MinutesRepository{
				BaseDir:    root,
				GenerateID: tc.idGenerator,
			}

			id, err := repo.Create(tc.data)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}

			resultingFilename := filepath.Join(root, string(id)+".xml")
			resultingContent, err := ioutil.ReadFile(resultingFilename)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}

			if !bytes.Equal(expected, resultingContent) {
				t.Errorf("expected: %s, got: %s", expected, resultingContent)
			}
		})
	}
}
