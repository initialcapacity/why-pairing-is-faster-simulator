package main

import (
	"initialcapacity.io/backlog-simulator/pkg/backlog"
	"initialcapacity.io/backlog-simulator/pkg/developers"
	"initialcapacity.io/backlog-simulator/pkg/story"
)

func main() {

	story41 := story.New("Create person", 4, []*story.Story{})
	story42 := story.New("List people", 3, []*story.Story{&story41})
	story43 := story.New("Show person", 2, []*story.Story{&story42})
	story44 := story.New("Edit person", 4, []*story.Story{&story42})

	story51 := story.New("Create project", 4, []*story.Story{&story41})
	story52 := story.New("List projects", 3, []*story.Story{&story51})
	story53 := story.New("Show project", 2, []*story.Story{&story52})
	story54 := story.New("Edit project", 4, []*story.Story{&story52})

	story101 := story.New("Create timesheet", 4, []*story.Story{&story41})
	story102 := story.New("List timesheets", 3, []*story.Story{&story101})
	story103 := story.New("Show timesheet", 2, []*story.Story{&story102})
	story104 := story.New("Delete timesheet", 2, []*story.Story{&story102})
	story105 := story.New("Edit timesheet", 4, []*story.Story{&story102})
	story106 := story.New("Assign project to timesheet", 4, []*story.Story{&story42, &story102})
	story107 := story.New("Assign project to timesheet", 4, []*story.Story{&story52, &story102})

	story201 := story.New("Submit timesheet", 4, []*story.Story{&story106})
	story202 := story.New("Review timesheet", 2, []*story.Story{&story201})
	story203 := story.New("Accept timesheet", 3, []*story.Story{&story202})
	story204 := story.New("Reject timeheet", 3, []*story.Story{&story202})

	b := backlog.New([]*story.Story{
		&story41,
		&story42,
		&story43,
		&story44,
		&story51,
		&story52,
		&story53,
		&story54,
		&story101,
		&story102,
		&story103,
		&story104,
		&story105,
		&story106,
		&story107,
		&story201,
		&story202,
		&story203,
		&story204,
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
