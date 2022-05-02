package backlog

import (
	"fmt"
	"initialcapacity.io/backlog-simulator/pkg/story"
	"math"
	"time"
)

type Backlog struct {
	Stories  []*story.Story
	Firehose chan *story.Story
}

func New(stories []*story.Story) Backlog {
	return Backlog{
		Stories:  stories,
		Firehose: make(chan *story.Story, 0),
	}
}

func (b *Backlog) StartWork() {
	start := time.Now()
	fmt.Println("Starting work")
	for {
		for _, s := range b.Stories {
			time.Sleep(time.Duration(1) * time.Microsecond)
			if s.IsReadyForAction() && !s.IsBlocked() {
				b.Firehose <- s
			}
		}
		if b.isComplete() {
			close(b.Firehose)
			break
		}
	}
	elapsed := int(math.Round(time.Since(start).Seconds()))
	fmt.Printf("Backlog complete in %d hours\n", elapsed)
}

func (b *Backlog) isComplete() bool {
	for _, s := range b.Stories {
		if s.Status != story.Done {
			return false
		}
	}
	return true
}
