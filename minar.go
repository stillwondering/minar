package minar

// MinutesID is used to identify a single minutes document.
type MinutesID string

// The Minutes object contains all the meta information and notes taken during a meeting.
type Minutes struct {
	ID           MinutesID
	Title        string
	Participants []string
	Topics       []Topic
}

// Topic is something that was talked about in a meeting.
type Topic struct {
	Title   string
	Content string
}

// CreateMinutesData contains data that is necessary to create a new minutes record.
type CreateMinutesData struct {
	Title        string
	Participants []string
	Topics       []Topic
}

// IDGeneratorFunc is used to generate IDs for our meeting records.
type IDGeneratorFunc func() MinutesID

func (m Minutes) Equals(other Minutes) bool {
	idEquality := m.ID == other.ID
	titleEquality := m.Title == other.Title
	participantsCntEquality := len(m.Participants) == len(other.Participants)
	participantsEquality := true
	for _, p := range m.Participants {
		exists := false
		for _, op := range other.Participants {
			if p == op {
				exists = true
				break
			}
		}
		if !exists {
			participantsEquality = false
			break
		}
	}

	topicsCntEquality := len(m.Topics) == len(other.Topics)
	topicsEquality := true
	for _, t := range m.Topics {
		exists := false
		for _, ot := range other.Topics {
			if t == ot {
				exists = true
				break
			}
		}
		if !exists {
			topicsEquality = false
			break
		}
	}

	return idEquality && titleEquality && participantsCntEquality && participantsEquality && topicsCntEquality && topicsEquality
}
