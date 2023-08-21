package main

import (
	"os"
	"strings"
	"github.com/CamiloMartinez25/challenge-one-go/internal/tickets"
)

const (
	filename = "./tickets.csv"
)

func main() {
	total, err := tickets.GetTotalTicketsByDestination("Brazil")

	fileInfo := getFileInfo(filename)

	storage := tickets.Storage{
		Tickets: tickets.GetTickets(fileInfo),
	}

}

func getFileInfo(filename string) []string {

	file, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(file), "\n")
}
