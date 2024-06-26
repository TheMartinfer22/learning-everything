package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	apiKey := ""
	cx := ""
	numResults := 30
	for page := 1; page <= 10; page++ {
		start := (page - 1) * numResults
		url := fmt.Sprintf("https://customsearch.googleapis.com/customsearch/v1?q=@gmail.com+filetype:txt&cx=%s&key=%s&start=%d", cx, apiKey, start)
		response, _ := http.Get(url)
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		var data map[string]interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			fmt.Println("Erro ao decodificar JSON:", err)
			return
		}

		if items, ok := data["items"].([]interface{}); ok {
			for _, item := range items {
				if itemMap, ok := item.(map[string]interface{}); ok {
					if link, ok := itemMap["link"].(string); ok {
						fmt.Println(link)
					}
				}
			}
		} else {
			fmt.Println("Chave 'items' não encontrada no JSON")
		}
	}
}
