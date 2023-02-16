package anonimousboard

type AnonBoard struct {
	Messages []string
}

func (board *AnonBoard) AddQuestion(question string) {
	(*board).Messages = append((*board).Messages, question)
}

func (board *AnonBoard) GetAllQuestions() []string {
	return (*board).Messages
}
