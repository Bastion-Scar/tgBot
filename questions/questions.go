package questions

type Question struct {
	Text    string
	Options []string
	Answer  int
}

type Service struct {
	questions []Question
}

func NewService() *Service {
	return &Service{
		questions: []Question{
			{
				Text:    "В каком году началась Великая Отечественная война?",
				Options: []string{"1914", "1939", "1941", "1945"},
				Answer:  2,
			},
			// Добавь другие вопросы
		},
	}
}

func (s *Service) GetQuestion(index int) Question {
	return s.questions[index]
}
