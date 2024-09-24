package event

type Event struct {
	Time        int64       `json:"time"`
	SelfID      int64       `json:"self_id"`
	PostType    string      `json:"post_type"`
	MessageType string      `json:"message_type"`
	SubType     string      `json:"sub_type"`
	TempSource  int         `json:"temp_source"`
	MessageID   int32       `json:"message_id"`
	UserID      int64       `json:"user_id"`
	GroupID     int64       `json:"group_id"`
	NoticeType  string      `json:"notice_type"`
	Message     interface{} `json:"message"`
	RawMessage  string      `json:"raw_message"`
	Font        int32       `json:"font"`
	Sender      Sender      `json:"sender"`
	Anonymous   Anonymous   `json:"anonymous"`
	OperatorID  int64       `json:"operator_id"`
	File        File        `json:"file"`
	RequestType string      `json:"request_type"`
}

type PrivateEvent struct {
	Event
}

type GroupEvent struct {
	Event
}

type PerMessageData struct {
	Text string `json:"text"`
}

type PerMessage struct {
	Type string         `json:"type"`
	Data PerMessageData `json:"data"`
}

type Message struct {
}

type Sender struct {
	Nickname string `json:"nickname"`
	UserID   int64  `json:"user_id"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
	Card     string `json:"card"`
	Area     string `json:"area"`
	Level    string `json:"level"`
	Role     string `json:"role"`
	Title    string `json:"title"`
}

type Anonymous struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Flag int64  `json:"flag"`
}

type File struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	Busid int64  `json:"busid"`
	Url   string `json:"url"`
}
