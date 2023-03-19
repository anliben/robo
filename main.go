package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)

	app.Post("/enviar", func(ctx *fiber.Ctx) error {
		// Recupera o valor do campo de texto
		texto := ctx.FormValue("contacts")

		fmt.Println(texto)

		err := read_file(texto)
		
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Arquivo gerado com sucesso!")
			return ctx.Render("download", fiber.Map{
				"Title": "Hello, World!",
			})
		}

	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Static("/", "./public")

	log.Fatal(app.Listen(":3000"))
}

func read_file(texto string) error {
	// Abrir o arquivo .txt

	// Ler o conteúdo do arquivo e dividir os números em uma slice
	var numeros []string
	// scanner := bufio.NewScanner(file)
	scanner := bufio.NewScanner(strings.NewReader(texto))
	for scanner.Scan() {
		numeros = append(numeros, strings.Split(scanner.Text(), ",")...)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Abrir o arquivo .csv para escrita
	csvFile, err := os.Create("public/arquivo.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	// Criar o escritor CSV
	writer := csv.NewWriter(csvFile)

	// Escrever o cabeçalho das colunas
	header := []string{"Nome", "Celular"}
	writer.Write(header)

	// Escrever cada número com um nome aleatório incrementado a cada contato
	count := 1
	for _, numero := range numeros {
		nome := fmt.Sprintf("Contato %d", count)
		count++
		row := []string{nome, numero}
		writer.Write(row)
	}

	// Finalizar a escrita e verificar erros
	writer.Flush()
	if err := writer.Error(); err != nil {
		panic(err)
	}

	fmt.Println("Arquivo gerado com sucesso! no read file")

	return nil

}
