package main

type IMsgSender interface {
	Send(msg string) error
}

type EmailMsgSender struct {
	Emails []string
}

func NewEmailMsgSender(emails []string) (e *EmailMsgSender) {
	e.Emails = emails
	return
}

func (e *EmailMsgSender) Send(msg string) error {
	return nil
}

type INotification interface {
	Notify(msg string) error
}

type ErrorNotification struct {
	Sender IMsgSender
}

func NewErrorNotification(sender IMsgSender) (e *ErrorNotification) {
	e.Sender = sender
	return
}

func (e *ErrorNotification) Notify(msg string) error {
	return e.Sender.Send(msg)
}
