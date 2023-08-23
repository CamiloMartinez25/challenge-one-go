package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

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

var wg *sync.WaitGroup

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	wg = new(sync.WaitGroup)
	fileInfo := getFileInfo(filename)

	storage := tickets.Storage{
		Tickets: tickets.GetTickets(fileInfo),
	}

	//fmt.Println(storage)

	chTickets := make(chan string)
	chError := make(chan error)

	var destination string = "Colombia"

	wg.Add(3)

	go getTicketsInfo(chTickets, chError, byDestination, destination, morning, storage)
	go getTicketsInfo(chTickets, chError, byTime, destination, earlyMorning, storage)
	go getTicketsInfo(chTickets, chError, averageDestination, destination, afternoon, storage)

	go func() {
		wg.Wait()
		close(chTickets)
		close(chError)
	}()

	printChannelSInfo(chTickets, chError)
}

// getFileInfo is a function that reads a file and returns its information in a slice.
func getFileInfo(filename string) []string {

	file, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(file), "\n")
}

// getTicketsInfo is a function that loads the information consulted in a channel
func getTicketsInfo(chTicket chan string, chErr chan error, infoType string, destination string, time string, storage tickets.Storage) {
	defer wg.Done()
	switch infoType {
	case "ticketsByDestination":
		tickets, err := storage.GetTotalTicketsByDestination(destination)
		if err != nil {
			chErr <- err
		} else {
			chTicket <- fmt.Sprintf("The amount of tickets for the destination of %s is: %v", destination, strconv.Itoa(tickets))
		}

	case "ticketsByTime":
		tickets, err := storage.GetTotalTicketsByTime(time)
		if err != nil {
			chErr <- err
		} else {
			chTicket <- fmt.Sprintf("The amount of tickets for the %s time is: %v", time, strconv.Itoa(tickets))
		}

	case "ticketsAverageByDestination":
		tickets, err := storage.AveragePassengersByDestination(destination)
		if err != nil {
			chErr <- err
		} else {
			chTicket <- fmt.Sprintf("The average of tickets for the destination of %s is: %.2f%%", destination, tickets)
		}

	default:
		chErr <- errors.New("The info required is not available for the tickets")
	}

}

// printChannelSInfo is a function that I saw printing the information of a channel
func printChannelSInfo(chTickets chan string, chError chan error) {
	for {
		select {
		case x, ok := <-chTickets:
			if ok {
				fmt.Println("Ticket:", x)
			} else {
				fmt.Println("All tickets request are closed")
				chTickets = nil
			}
		case x, ok := <-chError:
			if ok {
				fmt.Println("Ticket Error:", x)
			} else {
				fmt.Println("All tickets error are closed")
				chError = nil
			}
		default:
			if chError == nil && chTickets == nil {
				os.Exit(1)
			}
		}
	}
}
