package developers

import (
	"fmt"
	"initialcapacity.io/backlog-simulator/pkg/backlog"
	"initialcapacity.io/backlog-simulator/pkg/story"
	"time"
)

type Pair struct {
	Name string
}

func NewPair(name string) Pair {
	return Pair{Name: name}
}

func (d *Pair) WorkFrom(b *backlog.Backlog) {
	go func() {
		for s := range b.Firehose {
			if s.Status == story.NotStarted {
				d.complete(s)
			} else if s.Status == story.ReadyForReview {
				d.review(s)
			}
		}
	}()
}

func (d *Pair) complete(s *story.Story) {
	fmt.Printf("Pair '%s' starting story '%s'\n", d.Name, s.Name)
	s.Status = story.Started

	time.Sleep(time.Duration(s.Time) * time.Second)

	fmt.Printf("Pair '%s' finished story '%s'\n", d.Name, s.Name)
	s.CompletedBy = d.Name
	s.Status = story.Done
}

func (d *Pair) review(s *story.Story) {
	s.Status = story.InReview

	fmt.Printf("Pair '%s' reviewing story '%s'\n", d.Name, s.Name)
	time.Sleep(1 * time.Second)
	fmt.Printf("Pair '%s' reviewed story '%s'\n", d.Name, s.Name)

	s.Status = story.ReadyToResolve
}
