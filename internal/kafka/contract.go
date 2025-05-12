package kafka

import "github.com/developeerz/restorio-reserving/reserving-service/pkg/models"

type Listener interface {
	Notify(payload *models.PayloadTelegram)
}
