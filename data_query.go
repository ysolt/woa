package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func getDatabaseEntries(queryresult *QueryResult) error {
	filename := "resources/cloudant_response_example.json"
	// Open our jsonFile
	jsonFile, err := os.Open(filename)

	byteValue, _ := ioutil.ReadAll(jsonFile)
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &queryresult)
	defer jsonFile.Close()
	return err
}
