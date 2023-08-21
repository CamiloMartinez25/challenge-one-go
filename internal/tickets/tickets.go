package tickets

import (
	"strings"
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

// ejemplo 1
func GetTotalTickets(destination string) (int, error) {}

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
