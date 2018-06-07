package csv

import (
	"log"
	"strings"

	"github.com/shanesiebken/csv-util/util"
	"github.com/spf13/viper"
)

// Concatenation represents a mapping of field "From" to
// field "To"
type Concatenation struct {
	Infields  []string `mapstructure:"infields"`
	Separator string   `mapstructure:"separator"`
	Outfield  string   `mapstructure:"outfield"`
}

// DoConcatenation performs concatenations specified in configuration
func DoConcatenation(r [][]string) ([][]string, error) {
	concats := []Concatenation{}
	if err := viper.UnmarshalKey("concatenations", &concats); err != nil {
		log.Panicf("Error marshaling concatenations from configuration: %v", err)
	}

	log.Printf("Concatenations: %v+", concats)
	for _, concat := range concats {
		log.Printf("Concatenation: %v+", concat)
		concatInd := map[string][]int{}

		for _, f := range concat.Infields {
			i := util.Index(r[0], f)
			if i >= 0 {
				concatInd[concat.Outfield] = append(concatInd[concat.Outfield], i)
				continue
			}
			log.Printf("Field %s not found in source csv.", f)
		}
		r[0] = append(r[0], concat.Outfield)
		log.Printf("r[1:]: %s", r[1:])
		for i, record := range r[1:] {
			i = i + 1
			log.Printf("Index: %d", i)
			val := []string{}
			for _, ind := range concatInd[concat.Outfield] {
				val = append(val, record[ind])
			}
			log.Printf("val: %s", val)
			r[i] = append(r[i], strings.Join(val, concat.Separator))
		}
		log.Printf("r: %s", r)
	}

	if viper.GetBool("dropconcated") {
		for _, c := range concats {
			// Add fields to drop to viper global configuration
			dropfieldsConf := []string{}
			if err := viper.UnmarshalKey("dropfields", &dropfieldsConf); err != nil {
				log.Panicf("Error marshaling fields to drop from configuration: %v", err)
			}
			viper.Set("dropfields", append(c.Infields, dropfieldsConf...))
		}
	}
	return r, nil
}
