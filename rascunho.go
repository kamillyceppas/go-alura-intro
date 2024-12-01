// package main

// import "fmt"

// func main() {
// 	nome := "Kamilly"
// 	versao := 1.1
// 	fmt.Println("Olá, sra.", nome)
// 	fmt.Println("Este programa está na versão", versao)

// 	fmt.Println("1- Iniciar monitoramento")
// 	fmt.Println("2- Exibir Logs")
// 	fmt.Println("0- Sair do Programa")

// 	var comando int
// 	fmt.Scan(&comando)

// 	fmt.Println("O comando escolhido foi", comando)
// 	// if comando == 1 {
// 	// 	fmt.Println("Monitorando...")
// 	// } else if comando == 2 {
// 	// 	fmt.Println("Exibindo Logs...")
// 	// } else if comando == 0 {
// 	// 	fmt.Println("Saindo do programa...")
// 	// } else {
// 	// 	fmt.Println("Não conheço este comando")
// 	// }
// 	// if comando == 1 {
// 	// 	fmt.Println("Monitorando...")
// 	// } else if comando == 2 {
// 	// 	fmt.Println("Exibindo Logs...")
// 	// } else if comando == 0 {
// 	// 	fmt.Println("Saindo do programa...")
// 	// } else {
// 	// 	fmt.Println("Não conheço este comando")

// 	}

// }

package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

// func main() {

// 	exibeIntroducao()
// 	exibeMenu()
// 	comando := leComando()

//		switch comando {
//		case 1:
//			iniciarMonitoramento()
//		case 2:
//			fmt.Println("Exibindo Logs...")
//		case 0:
//			fmt.Println("Saindo do programa...")
//			os.Exit(0)
//		default:
//			fmt.Println("Não conheço este comando")
//			os.Exit(-1)
//		}
//	}
func exibeIntroducao() {
	nome := "Kamilly"
	versao := 1.1
	fmt.Println("Olá, sra.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	_, err := fmt.Scan(&comandoLido)
	if err != nil {
		fmt.Println("Entrada inválida. Por favor, insira um número.")
		return -1
	}
	return comandoLido
}

// CODIGO USANDO FOR

func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			return
		case -1:
			continue
		default:
			fmt.Println("Não conheço este comando")
		}
	}
}
func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	// sites := []string{"https://random-status-code.herokuapp.com/",
	//     "https://www.alura.com.br", "https://www.caelum.com.br"}

	sites := leSitesDoArquivo()
	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}

		time.Sleep(delay * time.Second)
	}

	fmt.Println("")
}

func testaSite(site string) {
	if site == "" {
		fmt.Println("URL vazia, pulando teste.")
		return
	}
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

// func exibeNomes() {
// 	nomes := []string{"Douglas", "Daniel", "Bernardo"}
// 	fmt.Println("O meu slice tem", len(nomes), "itens")
// 	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

//		nomes = append(nomes, "Aparecida")
//		fmt.Println("O meu slice tem", len(nomes), "itens")
//		fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
//	}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo:", err)
		os.Exit(1)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		// Verifica se a linha não está vazia e se começa com "http" para garantir uma URL válida.
		if linha != "" && strings.HasPrefix(linha, "http") {
			sites = append(sites, linha)
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Erro durante a leitura do arquivo:", err)
			break
		}
	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site +
		" - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

// restante do código omitido

func imprimeLogs() {
	arquivo, err := os.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de log:", err)
		return
	}
	if len(arquivo) == 0 {
		fmt.Println("Nenhum log registrado até o momento.")
	} else {
		fmt.Println("Logs encontrados:")
		fmt.Println(string(arquivo))
	}
}
