package kafka

import (
	"context"
	"encoding/json"

	"github.com/developeerz/restorio-reserving/reserving-service/pkg/models"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	reader   *kafka.Reader
	listener Listener
}

func New(listener Listener, brokers []string, topic string) *Kafka {
	return &Kafka{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:   brokers,
			Topic:     topic,
			Partition: 0,
			MaxBytes:  10e6,
		}),
		listener: listener,
	}
}

func (k *Kafka) read(ctx context.Context) ([]byte, error) {
	msg, err := k.reader.ReadMessage(ctx)
	if err != nil {
		return nil, err
	}

	return msg.Value, nil
}

func (k *Kafka) ReadLoop(ctx context.Context) {
	for {
		data, err := k.read(ctx)
		if err != nil {
			log.Error().AnErr("error", err).Send()
			continue
		}

		var payload models.PayloadTelegram

		err = json.Unmarshal(data, &payload)
		if err != nil {
			log.Error().AnErr("error", err).Send()
			continue
		}

		if k.listener != nil {
			k.listener.Notify(&payload)
		}
	}
}
