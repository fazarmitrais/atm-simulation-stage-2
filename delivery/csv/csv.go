package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"
)

type CSV struct {
	Path string `json:"path"`
}

func NewCSV(path string) *CSV {
	return &CSV{path}
}

func (c *CSV) Import() error {
	if strings.TrimSpace(c.Path) == "" {
		return errors.New("path must not be empty")
	}
	file, err := os.Open(c.Path)
	if err != nil {
		return errors.New("file not found: " + c.Path)
	}
	defer func() {
		file.Close()
	}()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return err
	}

	for _, r := range records {
		fmt.Println(r)
	}

	return nil
}
