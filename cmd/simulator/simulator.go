package main

import (
	"initialcapacity.io/backlog-simulator/pkg/backlog"
	"initialcapacity.io/backlog-simulator/pkg/developers"
	"initialcapacity.io/backlog-simulator/pkg/story"
)

func main() {
	createPerson := story.New("Create person", 4, []*story.Story{})
	viewPerson := story.New("View person", 4, []*story.Story{&createPerson})
	createProject := story.New("Create project", 4, []*story.Story{})
	viewProject := story.New("View project", 4, []*story.Story{&createProject})

	b := backlog.New([]*story.Story{
		&createPerson,
		&viewPerson,
		&createProject,
		&viewProject,
	})

	pairA := developers.NewPair("A")
	pairB := developers.NewPair("B")

	pairA.WorkFrom(&b)
	pairB.WorkFrom(&b)

	//soloA := developers.NewSolo("A")
	//soloB := developers.NewSolo("B")
	//soloC := developers.NewSolo("C")
	//soloD := developers.NewSolo("D")

	//soloA.WorkFrom(&b)
	//soloB.WorkFrom(&b)
	//soloC.WorkFrom(&b)
	//soloD.WorkFrom(&b)

	b.StartWork()
}
