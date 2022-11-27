package main

import (
	"encoding/json"
	"os"
)

type Value struct {
	String string `json:"string"`
}

type KeyValue map[string]Value

const localFilename = "locales/zh-hans/locale.json"
const localDictionary = "locales/zh-hans/dictionary.json"

func main() {
	b, err := os.ReadFile("dictionary.json")
	if err != nil {
		panic(err)
	}

	dictionary := KeyValue{}
	err = json.Unmarshal(b, &dictionary)
	if err != nil {
		panic(err)
	}

	localFile, err := os.ReadFile(localFilename)
	if err != nil {
		panic(err)
	}
	kv := map[string]string{}
	err = json.Unmarshal(localFile, &kv)
	if err != nil {
		panic(err)
	}

	for _, value := range dictionary {
		_, ok := kv[value.String]
		if !ok {
			kv[value.String] = value.String
		}
	}

	b, err = json.MarshalIndent(kv, "", "    ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(localFilename, b, 0655)
	if err != nil {
		panic(err)
	}

	for k, value := range dictionary {
		val, ok := kv[value.String]
		if ok {
			dictionary[k] = Value{String: val}
		}
	}

	b, err = json.MarshalIndent(dictionary, "", "    ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(localDictionary, b, 0655)
	if err != nil {
		panic(err)
	}
}
