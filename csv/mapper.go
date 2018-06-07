package csv

import (
	"log"

	"github.com/spf13/viper"
)

// Mapping represents a mapping of field "From" to
// field "To"
type Mapping struct {
	From string `mapstructure:"from"`
	To   string `mapstructure:"to"`
}

// DoMapping performs a field mapping on the incoming records
func DoMapping(r [][]string) ([][]string, error) {
	mappings := []Mapping{}
	if err := viper.UnmarshalKey("mappings", &mappings); err != nil {
		log.Panicf("Error marshaling mappings from configuration: %v", err)
	}

	log.Printf("Mappings: %v+", mappings)

	dropfields := []string{}
	for i, val := range r[0] {
		mapVal := mappingsContains(mappings, val)
		if mapVal != "" {
			r[0][i] = mapVal
		} else {
			dropfields = append(dropfields, val)
		}
	}

	// Add fields to drop to viper global configuration
	if viper.GetBool("dropunmapped") {
		dropfieldsConf := []string{}
		if err := viper.UnmarshalKey("dropfields", &dropfieldsConf); err != nil {
			log.Panicf("Error marshaling fields to drop from configuration: %v", err)
		}
		viper.Set("dropfields", append(dropfields, dropfieldsConf...))
	}

	return r, nil
}

// utility function to see if value is contained in mappings
func mappingsContains(m []Mapping, e string) string {
	for _, mapping := range m {
		if mapping.From == e {
			return mapping.To
		}
	}
	return ""
}
