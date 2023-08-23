package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/CamiloMartinez25/challenge-one-go/internal/tickets"
)

const (
	filename           = "./tickets.csv"
	earlyMorning       = "early morning"
	morning            = "morning"
	afternoon          = "afternoon"
	evening            = "evening"
	byDestination      = "ticketsByDestination"
	byTime             = "ticketsByTime"
	averageDestination = "ticketsAverageByDestination"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	//total, err := tickets.GetTotalTickets("Brazil")

	fileInfo := getFileInfo(filename)

	storage := tickets.Storage{
		Tickets: tickets.GetTickets(fileInfo),
	}

	//fmt.Println(storage)

	canalTickets := make(chan string)
	defer close(canalTickets)
	canalErr := make(chan error)
	defer close(canalErr)

	var destination string = "Brazil"

	go getTicketsInfo(canalTickets, canalErr, byDestination, destination, morning, storage)

	go getTicketsInfo(canalTickets, canalErr, byTime, destination, morning, storage)

	go getTicketsInfo(canalTickets, canalErr, averageDestination, destination, morning, storage)

	os.Exit(1)

	// for ticketMsj := range canalTickets {

	// 	fmt.Println(ticketMsj)
	// }

}

func getFileInfo(filename string) []string {

	file, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(file), "\n")
}

func getTicketsInfo(chTicket chan string, chErr chan error, infoType string, destination string, time string, storage tickets.Storage) {

	switch infoType {
	case "ticketsByDestination":
		tickets, err := storage.GetTotalTicketsByDestination(destination)
		if err != nil {
			chErr <- err
			return
		}

		chTicket <- "The quantity of tickets for " + destination + " is: " + strconv.Itoa(tickets)

	case "ticketsByTime":
		tickets, err := storage.GetTotalTicketsByTime(time)
		if err != nil {
			chErr <- err
			return
		}

		chTicket <- "The quantity of tickets for the " + time + " is: " + strconv.Itoa(tickets)

	case "ticketsAverageByDestination":
		tickets, err := storage.AveragePassengersByDestination(destination)
		if err != nil {
			chErr <- err
			return
		}

		chTicket <- "The average of tickets for " + destination + " is: " + fmt.Sprintf("%v", tickets) + "%"

	default:
		chErr <- errors.New("The info required is not available for the tickets")
	}

	printChanelsInfo(chTicket, chErr)

}

func printChanelsInfo(chTicket chan string, chErr chan error) {
	select {
	case pr := <-chTicket:
		fmt.Println(pr)
	case err := <-chErr:
		log.Fatal(err)
	}
}
