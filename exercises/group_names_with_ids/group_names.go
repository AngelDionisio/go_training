package main

import "fmt"

type user struct {
	name string
	id   int
}

type groupedUsers struct {
	name string
	ids  []int
}

var people = []user{
	{name: "John", id: 1},
	{name: "Jacob", id: 2},
	{name: "Josh", id: 3},
	{name: "Wes", id: 4},
	{name: "Eve", id: 5},
	{name: "Daniel", id: 6},
	{name: "John", id: 7},
	{name: "Wes", id: 8},
	{"John", 9}, // just to show you do not need to specify keys, but its good practice to do so
}

func main() {
	groupedUsers := groupIdsWithNames(people)

	fmt.Println(groupedUsers)
}

// groupIdsWithNames iterates over a list of users, creates a map which adds an entry per user
// and appends any 1 or more ids, and then calls the helper function usersWithMultipleIds
// to return users with more than one id
func groupIdsWithNames(listOfUsers []user) []groupedUsers {
	ls := make(map[string][]int) // holds map of users and their list of possible multipleIds

	for _, user := range listOfUsers {
		ls[user.name] = append(ls[user.name], user.id)
	}

	// filter only users with more than one Id
	result := usersWithMultipleIds(ls)

	return result
}

// returns users with more than one set of id from a map whose keys are slices of int
func usersWithMultipleIds(users map[string][]int) []groupedUsers {
	var result []groupedUsers

	for key, value := range users {
		if len(value) > 1 {
			result = append(result, groupedUsers{name: key, ids: value})
		}
	}

	return result

}

//You have a text file with name and id comma separated. Print out all the names which are duplicates and their corresponding ids.
// first step is to build the data object that will be passed to the function
//
//eg.
//Name, Id
//John, 1
//Jacob, 2
//Josh, 3
//Eve, 4
//Daniel, 5
//John, 6
//Wes, 7
//John, 10
//
//Output:
//John 1, 6, 10
