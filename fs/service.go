package fs

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/stillwondering/minar"
	"github.com/stillwondering/minar/xml"
)

type MinutesRepository struct {
	BaseDir    string
	GenerateID minar.IDGeneratorFunc
}

func (repo *MinutesRepository) FindAll() ([]minar.Minutes, error) {
	var minutes []minar.Minutes

	files, err := filepath.Glob(filepath.Join(repo.BaseDir, "/*.xml"))
	if err != nil {
		return minutes, errors.Wrap(err, "glob files")
	}

	for _, file := range files {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			return minutes, errors.Wrap(err, "read file")
		}

		m, err := xml.Decode(b)
		if err != nil {
			return minutes, errors.Wrap(err, "decode file content")
		}

		minutes = append(minutes, m)
	}

	return minutes, nil
}

func (repo *MinutesRepository) Create(data minar.CreateMinutesData) (minar.MinutesID, error) {
	id := repo.GenerateID()

	minutes := minar.Minutes{
		ID:           id,
		Title:        data.Title,
		Participants: data.Participants,
		Topics:       data.Topics,
	}

	encoded, err := xml.Encode(minutes)
	if err != nil {
		return minar.MinutesID(""), errors.Wrap(err, "encode minutes")
	}

	path := repo.createFilename(id)

	err = ioutil.WriteFile(path, encoded, 0644)
	if err != nil {
		return minar.MinutesID(""), errors.Wrap(err, "write file")
	}

	return id, nil
}

func (repo *MinutesRepository) Find(id minar.MinutesID) (*minar.Minutes, error) {
	path := repo.createFilename(id)
	if !fileExists(path) {
		return nil, nil
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "read file")
	}

	m, err := xml.Decode(b)
	if err != nil {
		return nil, errors.Wrap(err, "decode file content")
	}

	return &m, nil
}

func (repo *MinutesRepository) createFilename(id minar.MinutesID) string {
	return filepath.Join(repo.BaseDir, string(id)+".xml")
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
