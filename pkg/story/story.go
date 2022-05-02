package story

type Status int

const (
	NotStarted Status = iota
	Started
	ReadyForReview
	InReview
	ReadyToResolve
	Resolving
	Done
)

type Story struct {
	Name        string
	Time        int
	DependsOn   []*Story
	Status      Status
	CompletedBy string
}

func New(name string, time int, dependsOn []*Story) Story {
	return Story{
		Name:      name,
		Time:      time,
		DependsOn: dependsOn,
		Status:    NotStarted,
	}
}

func (s Story) IsReadyForAction() bool {
	return s.Status == NotStarted || s.Status == ReadyForReview || s.Status == ReadyToResolve
}

func (s Story) IsBlocked() bool {
	for _, story := range s.DependsOn {
		if story.Status != Done {
			return true
		}
	}

	return false
}
