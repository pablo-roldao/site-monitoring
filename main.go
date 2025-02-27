package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoringQuantity = 5
const monitoringDelay = 4
const sitesPath = "sites.txt"
const logsPath = "log.txt"

func main() {
	for {
		printMenu()
		option := readOption()

		switch option {
		case 1:
			startMonitoring(monitoringQuantity, monitoringDelay, sitesPath)
		case 2:
			printLogs(logsPath)
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option!")
			os.Exit(-1)
		}
	}
}

func printMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Quit program")
}

func readOption() int {
	var option int
	fmt.Print("Insert an option: ")
	fmt.Scan(&option)

	return option
}

func startMonitoring(quantity int, delay int, sitesPath string) {
	fmt.Println("\nMonitoring...")
	sites := readSitesFromFile(sitesPath)

	for range quantity {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(time.Duration(delay) * time.Second)
	}

	fmt.Print("\n")
}

func readSitesFromFile(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := bufio.NewReader(file)

	sites := []string{}
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if err == io.EOF {
			break
		}

		sites = append(sites, line)
	}

	file.Close()

	return sites
}

func testSite(site string) {
	response, err := http.Get(site)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(site, response.Status, response.StatusCode)
		registerResponse(*response)
	}
}

func registerResponse(response http.Response) {
	file, err := os.OpenFile(logsPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("Errror:", err)
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	file.WriteString(now + "\t" + response.Request.URL.String() + "\t" + response.Status + "\n")

	file.Close()
}

func printLogs(path string) {
	file, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(file))
}
