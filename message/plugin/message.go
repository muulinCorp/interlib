package plugin

type Channel string

const (
	Push = Channel("push")
	Mail = Channel("mail")
)

type Message struct {
	Channel   Channel
	Receivers []*Receiver
}

type Receiver struct {
	Name    string
	Address string
}

type PushMessage struct {
	*Message `json:",inline"`
	Content  struct {
		Title string
		Body  string
	}
}

type MailMessage struct {
	*Message `json:",inline"`
	Content  struct {
		Title  string
		Plaint string
		Html   string
	}
}

type MqttMessage struct {
	*Message `json:",inline"`
	Content  map[string]any
}
