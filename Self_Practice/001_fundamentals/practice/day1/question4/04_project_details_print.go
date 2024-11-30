package main

import "fmt"

/*q4. Print default values and Type names of variables from question 2 using printf
// Quick Tip, Use %v if not sure about what verb should be used,
// but don't use it in this question :)
// but generally using %v should be fine*/

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
	fmt.Printf("%T, %T, %T, %T, %T, %T, %T", projectName, codeLines, bugsFound, isComplete, averageLinesOfCode, teamLead, projectDeadline)

}
