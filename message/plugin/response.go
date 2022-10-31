package plugin

import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"
)

func NewMessageResponseWriter(w http.ResponseWriter) MessageResponseWriter {
	return &msgResponseWriter{
		rw: w,
		MessageResponse: &MessageResponse{
			Header: http.Header{},
		},
	}
}

type MessageResponse struct {
	PushMsgs   []*PushMessage
	MailMsgs   []*MailMessage
	Response   string
	Header     http.Header
	StatusCode int
}

func (mr *MessageResponse) DecoreReponse() ([]byte, error) {
	return hex.DecodeString(mr.Response)
}

type MessageResponseWriter interface {
	http.ResponseWriter
	Encode(diKey string) error
	AddPushMsg(m *Message, title, body string, data map[string]any)
	AddPushTplMsg(m *Message, title, bodyTplKey string, data map[string]any, variables map[string]string)
	AddMailMsg(m *Message, title, plaint, html string)
	AddMailTplMsg(m *Message, title, plaintTplKey, htmlTplKey string, variables map[string]string)
}

type msgResponseWriter struct {
	rw      http.ResponseWriter
	msgSize int
	*MessageResponse
}

func (mr *msgResponseWriter) Header() http.Header {
	return mr.MessageResponse.Header
}

func (mr *msgResponseWriter) WriteHeader(statusCode int) {
	mr.StatusCode = statusCode
}

func (mr *msgResponseWriter) Encode(diKey string) error {
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
	mr.rw.Header().Add("X-DiKey", diKey)
	err := json.NewEncoder(mr.rw).Encode(mr.MessageResponse)
	return err
}

func (w *msgResponseWriter) Write(b []byte) (int, error) {
	w.Response = hex.EncodeToString(b)
	return 0, nil
}

func (w *msgResponseWriter) AddPushMsg(m *Message, title, body string, data map[string]any) {
	w.msgSize++
	w.PushMsgs = append(w.PushMsgs, &PushMessage{
		Message: m,
		Title:   title,
		Content: struct {
			Body string
			Data map[string]any
		}{
			Body: body,
			Data: data,
		},
	})
}

func (w *msgResponseWriter) AddPushTplMsg(m *Message, title, bodyTplKey string, data map[string]any, variables map[string]string) {
	w.msgSize++
	w.PushMsgs = append(w.PushMsgs, &PushMessage{
		Message: m,
		Title:   title,
		Content: struct {
			Body string
			Data map[string]any
		}{
			Body: GetTemplateTag(bodyTplKey),
			Data: data,
		},
		Variables: variables,
	})
}

func (w *msgResponseWriter) AddMailMsg(m *Message, title, plaint, html string) {
	w.msgSize++
	w.MailMsgs = append(w.MailMsgs, &MailMessage{
		Message: m,
		Title:   title,
		Content: struct {
			Plaint string
			Html   string
		}{
			Plaint: plaint,
			Html:   html,
		},
	})
}

func (w *msgResponseWriter) AddMailTplMsg(m *Message, title, plaintTplKey, htmlTplKey string, variables map[string]string) {
	w.msgSize++
	w.MailMsgs = append(w.MailMsgs, &MailMessage{
		Message: m,
		Title:   title,
		Content: struct {
			Plaint string
			Html   string
		}{
			Plaint: GetTemplateTag(plaintTplKey),
			Html:   GetTemplateTag(htmlTplKey),
		},
		Variables: variables,
	})
}
