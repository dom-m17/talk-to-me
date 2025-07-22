package cmd

type Trivia struct {
	ResponseCode int `json:"response_code"`
	Results      []struct {
		Question string `json:"question"`
	} `json:"results"`
}

type DadJoke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

type Fact struct {
	Fact string `json:"fact"`
}

type Quote struct {
	Quote    string `json:"quote"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

type MeowFact struct {
	Data []string `json:"data"`
}
