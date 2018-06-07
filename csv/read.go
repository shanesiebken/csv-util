package csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/spf13/viper"
)

// Read reads in a csv into a 2d string array for further processing
func Read(filename string) ([][]string, error) {
	// Read in source file as specified in source file in configuration
	in, err := os.Open(viper.GetString("source"))
	if err == io.EOF {
		log.Print("Source file is empty.")
		os.Exit(1)
	}
	if err != nil {
		log.Printf("Error opening source file: %s", err)
	}

	r := csv.NewReader(in)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}
