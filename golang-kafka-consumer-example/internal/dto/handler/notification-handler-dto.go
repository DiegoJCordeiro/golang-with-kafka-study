package handler

type NotificationHandlerDTO struct {
	UUID    string   `json:"uuid"`
	Message string   `json:"message"`
	From    string   `json:"from"`
	To      []string `json:"to"`
	Opened  bool     `json:"opened"`
}

func NewNotificationHandlerDTO(uuid, message, from string, to []string, opened bool) *NotificationHandlerDTO {
	return &NotificationHandlerDTO{
		UUID:    uuid,
		Message: message,
		From:    from,
		To:      to,
		Opened:  opened,
	}
}
