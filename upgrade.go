package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	v8 "rogchap.com/v8go"
)

func main() {
	type Value struct {
		String json.RawMessage `json:"string"`
	}

	type KeyValue map[string]Value

	js, err := os.ReadFile("figma_app.min.js")
	if err != nil {
		panic(err)
	}

	pn := len([]byte("JSON.parse(`"))
	allDictionary := KeyValue{}
	for {
		n := bytes.Index(js, []byte(`"string":`))
		if n == -1 {
			break
		}
		p := bytes.LastIndex(js[:n], []byte("JSON.parse("))
		if p == -1 {
			break
		}

		n2 := n
		for {
			p2 := js[p+pn-1]
			e1 := bytes.Index(js[n2:], append([]byte("}"), p2))
			bb := js[p : n2+e1+2]

			ctx := v8.NewContext()
			val, err := ctx.RunScript(fmt.Sprintf(`JSON.stringify(%s))`, string(bb)), "main.js")
			if err == nil {
				dictionary := KeyValue{}
				err = json.Unmarshal([]byte(val.String()), &dictionary)
				if err == nil {
					js = js[n2+e1+1:]
					for k, value := range dictionary {
						allDictionary[k] = value
					}
					break
				}
			}

			n2 += e1 + 1
		}
	}

	b, err := json.MarshalIndent(allDictionary, "", "    ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("dictionary.json", b, 0655)
	if err != nil {
		panic(err)
	}
}
