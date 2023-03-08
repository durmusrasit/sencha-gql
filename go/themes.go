package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Themes struct {
	data []*Theme
}

type Theme struct {
	ID         int
	Name       string
	Foreground string
	Background string
	Enabled    bool
}

func (t *Themes) ListThemes() ([]*Theme, error) {
	return t.data, nil
}

func (t *Themes) GetTheme(id int) (*Theme, error) {
	for _, t := range t.data {
		if t.ID == id {
			return t, nil
		}
	}
	return nil, errors.New("Theme not found")
}

func (t *Themes) ReadThemes() {
	content, err := os.ReadFile("./themes.json")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = json.Unmarshal(content, &t.data)

	if err != nil {
		fmt.Println("An error occurred while reading themes. Error:", err)
	}
}
