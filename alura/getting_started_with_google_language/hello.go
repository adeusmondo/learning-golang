package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring_times = 3
const delay = 5
const version = "1.1.1"

func main() {
	displaysIntroduction()

	for {
		displaysMenu()

		command := readInputCommand()
		switch command {
		case 1:
			startMonitoring()
		case 2:
			printLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não reconheço este comando...")
			fmt.Println("Saindo do programa...")
			os.Exit(-1)
		}
	}
}

func displaysIntroduction() {
	fmt.Println("Este programa está na versão:", version)
}

func displaysMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func readInputCommand() int {
	var input int
	fmt.Scan(&input)
	fmt.Println("\nO comando escolhido foi:", input)
	fmt.Println("")

	return input
}

func startMonitoring() {
	fmt.Println("Monitorando...")

	sites, err := readTextFileSites()
	if err != nil {
		os.Exit(-1)
	}

	for i := 0; i < monitoring_times; i++ {
		for i, site := range sites {
			fmt.Println("Site:", i+1, "-", site)
			testSiteConnection(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testSiteConnection(site string) error {
	resp, err := http.Get(site)
	time_now := time.Now()
	if err != nil {
		fmt.Println("Ocorreu um erro:", err.Error())
		return err
	}

	if resp != nil && resp.StatusCode == 200 {
		fmt.Println("Site:", site, "acessado com sucesso.")
		err = log(time_now, site, true)
	} else {
		fmt.Println("Site:", site, "apresentou um problema. Status:",
			resp.StatusCode)
		err = log(time_now, site, false)
	}

	return err
}

func readTextFileSites() ([]string, error) {
	var sites []string

	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Erro ao ler arquivo:", err)
		return nil, err
	}

	reader := bufio.NewReader(file)
	for {
		site, err := reader.ReadString('\n')
		site = strings.TrimSpace(site)
		sites = append(sites, site)

		if err == io.EOF {
			break
		}
	}
	file.Close()

	return sites, nil
}

func log(time_now time.Time, site string, status bool) error {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Erro ao escrever log arquivo:", err)
		return err
	}

	file.WriteString(time_now.Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()

	return nil
}

func printLogs() {
	fmt.Println("Exibindo Logs...")

	file, err := ioutil.ReadFile("logs.txt")
	if err != nil {
		fmt.Println("Um erro ocorreu:", err)
	}

	fmt.Println(string(file))
}
