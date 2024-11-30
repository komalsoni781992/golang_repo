package main

import "fmt"

/*q2. Create a program to store and print a person's and their project's details. Declare and initialize variables for the following details,
  Project name (string)
  Code lines written (uint8)
  Bugs found (int)
  Is the project complete? (bool)
  Average lines of code written per hour (float64)
  Team lead name (string)
  Project deadline in days (int)
  Additionally, demonstrate a uint overflow by initializing the largest possible value for uint and then adding 1 to it*/

func main() {
	var (
		projectName        string  = "ProjectName"
		codeLines          uint8   = 100
		bugsFound          int     = 20
		isComplete         bool    = false
		averageLinesOfCode float64 = 75.6
		teamLead           string  = "TeamLead"
		projectDeadline    int     = 25
	)
	fmt.Println(projectName, codeLines, bugsFound, isComplete, averageLinesOfCode, teamLead, projectDeadline)

	var b byte = 255
	fmt.Println(b)
	fmt.Println(b + 1)
}
