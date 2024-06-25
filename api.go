package main

import (
	"encoding/json"
	"net/http"
)

var employees = []Employee{
	{ID: 1, Name: "Oleg Khristaforov", BirthDate: "1997-07-24"},
	{ID: 2, Name: "Jane Smith", BirthDate: "1985-05-23"},
}

var subscriptions = []Subscription{}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)
}

func Subscribe(w http.ResponseWriter, r *http.Request) {
	var subscription Subscription
	_ = json.NewDecoder(r.Body).Decode(&subscription)
	subscriptions = append(subscriptions, subscription)
	json.NewEncoder(w).Encode(subscription)
}

func Unsubscribe(w http.ResponseWriter, r *http.Request) {
	var subscription Subscription
	_ = json.NewDecoder(r.Body).Decode(&subscription)

	for i, sub := range subscriptions {
		if sub.UserID == subscription.UserID && sub.EmployeeID == subscription.EmployeeID {
			subscriptions = append(subscriptions[:i], subscriptions[i+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(subscription)
}
