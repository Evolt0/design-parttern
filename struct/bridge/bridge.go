package bridge

type ISendMessage interface {
	Send(msg string) error
}

type EmailMessageSender struct {
	emails []string
}

func NewEmailMessageSender(emails []string) *EmailMessageSender {
	return &EmailMessageSender{emails: emails}
}

func (e *EmailMessageSender) Send(msg string) error {
	// fmt.Printf("send %s to %v", msg, e.emails[0])
	return nil
}

// --------------------------

type INotification interface {
	Notify(msg string) error
}

type ErrorNotification struct {
	sender ISendMessage
}

func (e *ErrorNotification) Notify(msg string) error {
	return e.sender.Send(msg)
}

func NewErrorNotification(sender ISendMessage) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}