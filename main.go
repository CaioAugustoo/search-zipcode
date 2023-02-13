package main

import (
	"io"
	"net/http"
	"os"
)

func getRequestURL(zipcode string) string {
	return "https://viacep.com.br/ws/" + zipcode + "/json/"
}

func main() {
	for _, v := range os.Args[1:] {
		url := getRequestURL(v)

		res, err := http.Get(url)

		if err != nil {
			panic(err)
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)

		if err != nil {
			panic(err)
		}

		file, err := os.Create("result.json")

		if err != nil {
			panic(err)
		}

		file.WriteString(string(data))

		defer file.Close()
	}
}
