package tickets

import (
	"errors"
	"strconv"
	"strings"
)

const (
	increment = 1
)

var (
	countError   = "Ticket count is 0 for "
	timeNotFound = errors.New("The time provided doesn't match with any time existing")
	earlyMorning = []int{0, 6}
	morning      = []int{6, 12}
	afternoon    = []int{12, 19}
	evening      = []int{19, 0}
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
func (s *Storage) GetTotalTicketsByDestination(destination string) (int, error) {
	ticketsQ := 0

	for _, ticket := range s.Tickets {
		if ticket.Destination == destination {
			ticketsQ += increment
		}
	}

	if ticketsQ == 0 {
		err := errors.New(countError + "the destination of " + destination)
		return ticketsQ, err
	}

	return ticketsQ, nil
}

// Get total tickets by specific time
func (s *Storage) GetTotalTicketsByTime(time string) (int, error) {
	switch time {
	case "early morning":
		return getTicketsByTimeRange(earlyMorning, s.Tickets)
	case "morning":
		return getTicketsByTimeRange(morning, s.Tickets)
	case "afternoon":
		return getTicketsByTimeRange(afternoon, s.Tickets)
	case "evening":
		return getTicketsByTimeRange(evening, s.Tickets)
	default:
		return 0, timeNotFound
	}
}

// Get the average of passengers (%) which travel to specific destination
func (s *Storage) AveragePassengersByDestination(destination string) (float64, error) {
	totalPassengersByDestination, err := s.GetTotalTicketsByDestination(destination)
	if err != nil {
		return 0, err
	}

	var average = (float64(totalPassengersByDestination) / float64(len(s.Tickets))) * 100.0
	return average, nil

}

// getTickets return a Slice of Ticket
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

func getTicketsByTimeRange(timeRange []int, tickets []Ticket) (int, error) {
	ticketsQ := 0

	for _, ticket := range tickets {
		time := strings.Split(string(ticket.Time), ":")[0]
		timeInt, err := strconv.Atoi(time)

		if err != nil {
			return ticketsQ, err
		}

		if timeInt > timeRange[0] && timeInt <= timeRange[1] {
			ticketsQ += increment
		}
	}

	if ticketsQ == 0 {
		err := errors.New(countError + "the time range of " + strconv.Itoa(timeRange[0]) + " - " + strconv.Itoa(timeRange[1]) + " hours")
		return ticketsQ, err
	}

	return ticketsQ, nil
}
