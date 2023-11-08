package main

import (
	"fmt"
)

type Rate struct {
	Rate    float32
	Subject string
}

type Student struct {
	Name  string
	Rates []Rate
}

func initStudents() []Student {
	return []Student{
		{
			Name: "Jhon",
			Rates: []Rate{
				{Rate: 10, Subject: "English"},
				{Rate: 11, Subject: "Math"},
				{Rate: 11.5, Subject: "Music"},
			},
		},
		{
			Name: "Mark",
			Rates: []Rate{
				{Rate: 8, Subject: "English"},
				{Rate: 7.5, Subject: "Math"},
				{Rate: 9.5, Subject: "Music"},
			},
		},
		{
			Name: "Richard",
			Rates: []Rate{
				{Rate: 3.5, Subject: "English"},
				{Rate: 4, Subject: "Math"},
				{Rate: 4, Subject: "Music"},
			},
		},
	}
}

func getSubjectRates(students []Student, subject string) []float32 {
	rates := make([]float32, 0, len(students))
	for _, student := range students {
		for _, rate := range student.Rates {
			if rate.Subject == subject {
				rates = append(rates, rate.Rate)
			}
		}
	}
	return rates
}

func getAverageRate(rates []float32) float32 {
	var average float32
	for _, rate := range rates {
		average = average + rate/float32(len(rates))
	}
	return average
}

func main() {
	students := initStudents()

	subjects := []string{"English", "Math", "Music"}
	for _, subject := range subjects {
		rates := getSubjectRates(students, subject)
		avgRate := getAverageRate(rates)
		fmt.Printf("Subject: %s, Average rate: %2.1f\n", subject, avgRate)
	}
}
