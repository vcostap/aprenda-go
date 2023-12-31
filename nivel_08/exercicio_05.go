package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type SortByAge []user

func (a SortByAge) Len() int           { return len(a) }
func (a SortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type SortByLast []user

func (a SortByLast) Len() int           { return len(a) }
func (a SortByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	for _, u := range users {
		sort.Strings(u.Sayings)
	}

	printUsers(users)
	fmt.Println("---------------------------------------------------------")

	sort.Sort(SortByAge(users))
	printUsers(users)
	fmt.Println("---------------------------------------------------------")

	sort.Sort(SortByLast(users))
	printUsers(users)
	fmt.Println("---------------------------------------------------------")

}

func printUsers(users []user) {
	for _, u := range users {
		fmt.Printf("Name: %v\tLastname: %v\tAge: %v\n", u.First, u.Last, u.Age)
		fmt.Println("Sayings:")
		for _, saying := range u.Sayings {
			fmt.Println("\t", saying)
		}
		fmt.Println("------------------")
	}
}