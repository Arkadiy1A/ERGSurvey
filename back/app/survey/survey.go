package survey

type Survey struct {
	Id        int
	Name      string
	Questions []Question
}

type Question struct {
	Id          int      `json:"id"`
	Description string   `json:"description"`
	Options     []Option `json:"options"`
	Counter     int      `json:"counter"`
	Answers     Answers  `json:"answers"`
}

type Answers struct {
	IpList map[string]*Option
}

type Option struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	//IpList []string `json:"ip_list"`
}

type Result struct {
	Id       int
	question *Question
	result   int
}

type ResponsePayload struct {
	Id int
}

func CreateDummySurvey() Survey {
	return Survey{
		Id:   1,
		Name: "BTS Опрос!",
		Questions: []Question{
			{
				Id:          1,
				Description: "Кого вы видите перед собой?",
				Options: []Option{
					{Name: "Сына маминой подруги"},
					{Name: "Десантникова Турникмена Выходсиловича"},
					{Name: "Рысь"},
				},
				Answers: Answers{IpList: map[string]*Option{}},
			},
		},
	}
}

func (survey *Survey) Increment(i int, ip string) {
	lastQuestion := (*survey).LatestQuestion()
	if _, ok := lastQuestion.Answers.IpList[ip]; !ok {
		lastQuestion.Answers.IpList[ip] = &(*survey).Questions[0].Options[i]
		lastQuestion.Answers.IpList[ip].Count++
	} else {
		lastQuestion.Answers.IpList[ip].Count--
		lastQuestion.Answers.IpList[ip] = &(*survey).Questions[0].Options[i]
		lastQuestion.Answers.IpList[ip].Count++
	}
}

func (survey *Survey) LatestQuestion() *Question {
	return &(*survey).Questions[len((*survey).Questions)-1]
}
