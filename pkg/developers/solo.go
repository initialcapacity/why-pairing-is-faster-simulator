package developers

import (
	"fmt"
	"initialcapacity.io/backlog-simulator/pkg/backlog"
	"initialcapacity.io/backlog-simulator/pkg/story"
	"time"
)

type Solo struct {
	Name string
}

func NewSolo(name string) Solo {
	return Solo{Name: name}
}

func (d *Solo) WorkFrom(b *backlog.Backlog) {
	go func() {
		for s := range b.Firehose {
			if s.Status == story.NotStarted {
				d.createPR(s)
			} else if s.Status == story.ReadyForReview && s.CompletedBy != d.Name {
				d.review(s)
			} else if s.Status == story.ReadyToResolve && s.CompletedBy == d.Name {
				d.resolveReview(s)
			}
		}
	}()
}

func (d *Solo) createPR(s *story.Story) {
	s.Status = story.Started

	fmt.Printf("Solo '%s' starting story '%s'\n", d.Name, s.Name)
	time.Sleep(time.Duration(s.Time) * time.Second)
	fmt.Printf("Solo '%s' finished story '%s'\n", d.Name, s.Name)

	s.CompletedBy = d.Name
	s.Status = story.ReadyForReview
}

func (d *Solo) review(s *story.Story) {
	s.Status = story.InReview

	fmt.Printf("Solo '%s' reviewing story '%s'\n", d.Name, s.Name)
	time.Sleep(1 * time.Second)
	fmt.Printf("Solo '%s' reviewed story '%s'\n", d.Name, s.Name)

	s.Status = story.ReadyToResolve
}

func (d *Solo) resolveReview(s *story.Story) {
	s.Status = story.Resolving

	fmt.Printf("Solo '%s' resolving story '%s'\n", d.Name, s.Name)
	time.Sleep(1 * time.Second)
	fmt.Printf("Solo '%s' resolved story '%s'\n", d.Name, s.Name)

	s.Status = story.Done
}
