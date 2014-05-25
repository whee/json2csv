// Copyright (c) 2014 Brian Hetro <whee@smaertness.net>
// Use of this source code is governed by the ISC
// license which can be found in the LICENSE file.

package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
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
		switch f := d[k].(type) {
		default:
			log.Fatalf("Unsupported type %T. Aborting.\n", f)
		case string:
			record = append(record, f)
		case float64:
			record = append(record, strconv.FormatFloat(f, 'f', -1, 64))
		}
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
