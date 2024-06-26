package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type ChatPrompt struct {
	Contents  []Content `json:"contents"`
	ModelType string    `json:"modelType"`
}

type Content struct {
	Role   string `json:"role"`
	Status string `json:"status"`
	Parts  []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type ChatResponse struct {
	Output string `json:"output"`
}

func main() {
	url := "https://ai.nanosync.dev/"
	SearchMainPage(url)
	SearchRobotsLeaked(url)
}

func SearchRobotsLeaked(url string) {
	html := GetHTML(url + "robots.txt")
	println("[+] Realizando verificação da Robots.txt")
	response := Prompt("Verifique esse(a) robots.txte veja se há vulnerabilidades ou riscos de segurança" + html)
	println(response)
}

func SearchMainPage(url string) {
	html := GetHTML(url)
	println("[+] Realizando verificação de página principal")
	response := Prompt("Verifique esse HTML principal e veja se há vulnerabilidades ou riscos de segurança" + html)
	println(response)
}

func GetHTML(url string) (outputReader string) {
	response, err := http.Get(url)
	if err != nil {
		println("Erro ao fazer a solicitação GET:", err.Error())
		return
	}
	defer response.Body.Close()

	output, err := io.ReadAll(response.Body)
	if err != nil {
		println("Erro ao ler a resposta:", err.Error())
		return
	}

	outputReader = string(output)
	return
}

func Prompt(ask string) (promptChatResponse string) {
	chatPrompt := ChatPrompt{
		Contents: []Content{
			{
				Role:   "user",
				Status: "success",
				Parts: []Part{
					{
						Text: ask + ": Diga FOUND[Caso encontrar diga na mesma linha qual vunerabilidade em resumo] para encontrado e NOT FOUND para se não encontrar:",
					},
				},
			},
		},
		ModelType: "NORMAL",
	}

	jsonData, _ := json.Marshal(chatPrompt)
	response, _ := http.Post("https://demo-ia.nanosync.dev/api/v1/chat", "application/json", bytes.NewBuffer(jsonData))
	defer response.Body.Close()
	readAll, _ := io.ReadAll(response.Body)
	var chatResponse ChatResponse
	json.Unmarshal(readAll, &chatResponse)
	return chatResponse.Output
}
