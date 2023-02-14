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
}

type Option struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
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
					{Name: "Зевса"},
					{Name: "Рысь"},
				},
			},
		},
	}
}

func (survey *Survey) Increment(i int) {
	(*survey).Questions[0].Options[i].Count++
}

func (survey *Survey) LatestQuestion() *Question {
	return &(*survey).Questions[len((*survey).Questions)-1]
}
