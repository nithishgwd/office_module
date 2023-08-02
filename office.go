package office

import (
	"fmt"

	team "github.com/nithishgwd/team_module"
)

// Office represents an office containing multiple teams.
type Office struct {
	teams map[string]*team.Team // Mapping of team names to team instances
}

// NewOffice creates a new instance of Office.
func NewOffice() *Office {
	return &Office{
		teams: make(map[string]*team.Team),
	}
}

// AddTeam adds a new team to the office.
func (o *Office) AddTeam(name string) {
	o.teams[name] = team.NewTeam(name)
}

// GetTeamNames returns a slice of all team names in the office.
func (o *Office) GetTeamNames() []string {
	var names []string
	for name := range o.teams {
		names = append(names, name)
	}
	return names
}

// UpdateEmployeesTodayForTeam updates the number of employees present today for a particular team.
func (o *Office) UpdateEmployeesTodayForTeam(teamName string, numEmployees int) {
	team, found := o.teams[teamName]
	if found {
		team.UpdateEmployeesToday(numEmployees)
	}
}

// GetTeamEmployeesToday retrieves the number of employees present today for a specific team.
func (o *Office) GetTeamEmployeesToday(teamName string) int {
	team, found := o.teams[teamName]
	if found {
		return team.GetEmployeesToday()
	}
	return 0
}

// DisplayOfficeInformation displays all office information.
func (o *Office) DisplayOfficeInformation() {
	fmt.Println("\nOffice Information:")
	for _, team := range o.teams {
		employeesToday := team.GetEmployeesToday()
		fmt.Printf("Team: %s, Employees Today: %d\n", team.GetName(), employeesToday)
	}
}
