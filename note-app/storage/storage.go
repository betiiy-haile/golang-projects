package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"example.com/cli-notes/models"
)

var filePath = "notes.json"

func LoadNotes() ([]models.Note, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := ioutil.WriteFile(filePath, []byte("[]"), 0644)
		if err != nil {
			return nil, err
		}
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var notes []models.Note
	err = json.Unmarshal(data, &notes)
    if err != nil {
        return nil, err
    }

    return notes, nil

}

func SaveNotes(note []models.Note) error {
	data, err := json.MarshalIndent(note, "", " ")
	if err != nil {
		return err
	}


	return ioutil.WriteFile(filePath, data, 0644)
} 