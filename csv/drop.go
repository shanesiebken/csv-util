package csv

import (
	"log"

	"github.com/shanesiebken/csv-util/util"
	"github.com/spf13/viper"
)

// DoDrop drops fields as specified in configuration
func DoDrop(r [][]string) [][]string {
	dropfields := []string{}
	if err := viper.UnmarshalKey("dropfields", &dropfields); err != nil {
		log.Panicf("Error marshaling fields to drop from configuration: %v", err)
	}

	if len(dropfields) <= 0 {
		return r
	}

	log.Printf("Dropfields: %s", dropfields)
	records := [][]string{}
	records = append(records, []string{})
	// Indices to drop
	ind := []int{}
	for i, val := range r[0] {
		if util.ContainsString(dropfields, val) {
			ind = append(ind, i)
			continue
		}
		records[0] = append(records[0], val)
	}

	for i, record := range r[1:] {
		i = i + 1
		records = append(records, []string{})
		for j, val := range record {
			if util.ContainsInt(ind, j) {
				continue
			}
			records[i] = append(records[i], val)
		}
	}
	return records
}
