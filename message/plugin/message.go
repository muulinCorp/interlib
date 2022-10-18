package plugin

import (
	"fmt"
	"regexp"
)

type Channel string

const (
	Push = Channel("push")
	Mail = Channel("mail")
)

var (
	tplReqexp, _ = regexp.Compile("{{@([a-z_A-Z]+)}}")
)

func GetTemplateTag(tplName string) string {
	return fmt.Sprintf("{{@%s}}", tplName)
}

func ParserTemplateName(c string) string {
	s := tplReqexp.FindString(c)
	return s[3 : len(s)-2]
}

func ReplaceTemplateTag(c, replace string) string {
	return tplReqexp.ReplaceAllString(c, replace)
}

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
	Title    string
	Content  struct {
		Body string
		Data map[string]any
	}
	Variables map[string]string
}

type MailMessage struct {
	*Message `json:",inline"`
	Title    string
	Content  struct {
		Plaint string
		Html   string
	}
	Variables map[string]string
}

type MqttMessage struct {
	*Message  `json:",inline"`
	Content   map[string]any
	IsUseTpl  bool
	Variables map[string]string
}
