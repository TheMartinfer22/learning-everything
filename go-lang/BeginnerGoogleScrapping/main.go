package main

import (
	"github.com/gocolly/colly/v2"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.ru/")
	if err != nil {
		println("Erro ao fazer a requisição HTTP:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		println("Erro ao ler o corpo da resposta:", err)
		return
	}

	if body != nil {
		println("[Scrapping Test] Funcionando corretamente!")
	} else {
		return
	}

	c := colly.NewCollector()

	c.OnHTML("div", func(e *colly.HTMLElement) {
		if e.Text != "" {
			println("[Scrapped Info] -> ", e.Text)
		}
	})

	err = c.Visit("https://google.ru/")
	if err != nil {
		return
	}
}
