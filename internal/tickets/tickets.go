package tickets

import (
	"strings"
	"errors"
)

const (
	increment: 1
)

var (
	countError: errors.New("The count of tickets for the search is 0")
)

type Ticket struct {
	Id          string
	Name        string
	Email       string
	Destination string
	Time        string
	Price       string
}

type Storage struct {
	Tickets []Ticket
}

// Get total tickets booked by specific destination
func (s *Storage) GetTotalTickets(destination string) (int, error) {
	var ticketsQ := 0

	for _, ticket := range s.Tickets {
		if ticket.Destination == destination {
			ticketsQ += increment
		}
	}

	if ticketsQ == 0 {
		err := countError
	}

	return ticketsQ, err
}

// ejemplo 2
func GetMornings(time string) (int, error) {}

// ejemplo 3
func AverageDestination(destination string, total int) (int, error) {}

// getTickets es una funcion que retorna un slice de Ticket
func GetTickets(info []string) []Ticket {

	var tickets []Ticket
	for i := 0; i < len(info); i++ {

		if len(info[i]) > 0 {
			file := strings.Split(string(info[i]), ",")
			ticket := Ticket{
				Id:          file[0],
				Name:        file[1],
				Email:       file[2],
				Destination: file[3],
				Time:        file[4],
				Price:       file[5],
			}
			tickets = append(tickets, ticket)
		}
	}
	return tickets
}
