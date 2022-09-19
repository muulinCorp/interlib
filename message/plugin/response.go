package plugin

import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"
)

func NewMessageResponseWriter(w http.ResponseWriter) MessageResponseWriter {
	return &msgResponseWriter{
		rw:              w,
		header:          w.Header(),
		MessageResponse: &MessageResponse{},
	}
}

type MessageResponse struct {
	PushMsgs []*PushMessage
	MailMsgs []*MailMessage
	Response string
}

func (mr *MessageResponse) DecoreReponse() ([]byte, error) {
	return hex.DecodeString(mr.Response)
}

type MessageResponseWriter interface {
	Encode() error
	Write(b []byte) (int, error)
	AddPushMsg(m *Message, title, body string)
	AddMailMsg(m *Message, title, plaint, html string)
}

type msgResponseWriter struct {
	rw      http.ResponseWriter
	header  http.Header
	msgSize int
	*MessageResponse
}

func (mr *msgResponseWriter) Encode() error {
	if mr.msgSize == 0 {
		out, err := mr.DecoreReponse()
		if err != nil {
			return err
		}
		_, err = mr.rw.Write(out)
		return err
	}
	mr.rw.Header().Add("X-Message", strconv.Itoa(mr.msgSize))
	mr.rw.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(mr.rw).Encode(mr.MessageResponse)
	return err
}

func (w *msgResponseWriter) Write(b []byte) (int, error) {
	w.Response = hex.EncodeToString(b)
	return 0, nil
}

func (w *msgResponseWriter) AddPushMsg(m *Message, title, body string) {
	w.msgSize++
	w.PushMsgs = append(w.PushMsgs, &PushMessage{
		Message: m,
		Content: struct {
			Title string
			Body  string
		}{
			Title: title,
			Body:  body,
		},
	})
}

func (w *msgResponseWriter) AddMailMsg(m *Message, title, plaint, html string) {
	w.msgSize++
	w.MailMsgs = append(w.MailMsgs, &MailMessage{
		Message: m,
		Content: struct {
			Title  string
			Plaint string
			Html   string
		}{
			Title:  title,
			Plaint: plaint,
			Html:   html,
		},
	})
}
