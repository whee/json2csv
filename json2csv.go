package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
	"sort"
)

func sortedKeys(m map[string]interface{}) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func printRow(w *csv.Writer, keys []string, d map[string]interface{}) error {
	var record []string
	for _, k := range keys {
		record = append(record, d[k].(string))
	}
	return w.Write(record)
}

func main() {
	dec := json.NewDecoder(os.Stdin)
	enc := csv.NewWriter(os.Stdout)
	var keys []string

	for {
		var jsd map[string]interface{}

		if err := dec.Decode(&jsd); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		if keys == nil {
			keys = sortedKeys(jsd)
			if err := enc.Write(keys); err != nil {
				log.Fatal(err)
			}
		}

		if err := printRow(enc, keys, jsd); err != nil {
			log.Fatal(err)
		}
	}

	enc.Flush()
}
