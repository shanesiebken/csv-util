package csv

import (
	"encoding/csv"
	"log"
	"os"
)

// Write writes out a csv to filename
func Write(filename string, records [][]string) error {
	out, err := os.Create(filename)
	if err != nil {
		return err
	}

	w := csv.NewWriter(out)

	if err := w.WriteAll(records); err != nil {
		log.Fatal(err)
	}
	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	return err
}
