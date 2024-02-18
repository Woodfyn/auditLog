package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Woodfyn/auditLog/internal/config"
	"github.com/Woodfyn/auditLog/pkg/core"
	"github.com/streadway/amqp"
)

type Server struct {
	conn *amqp.Connection
	ch   *amqp.Channel

	service AuditService
	msgs    <-chan amqp.Delivery
}

func NewServer(service AuditService) *Server {
	return &Server{
		conn: new(amqp.Connection),
		ch:   new(amqp.Channel),

		service: service,
		msgs:    make(<-chan amqp.Delivery),
	}
}

func (s *Server) Close() error {
	if err := s.ch.Close(); err != nil {
		return err
	}

	if err := s.conn.Close(); err != nil {
		return err
	}

	return nil
}

func (s *Server) Listen(cfg *config.Config) error {
	addr := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		cfg.MQ.Username, cfg.MQ.Password, cfg.MQ.Host, cfg.MQ.Port)

	conn, err := amqp.Dial(addr)
	if err != nil {
		return err
	}
	s.conn = conn

	s.ch, err = s.conn.Channel()
	if err != nil {
		return err
	}

	q, err := s.ch.QueueDeclare(
		"log", // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	s.msgs, err = s.ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Serve(ctx context.Context) error {
	var msg core.LogItem

	for d := range s.msgs {
		json.Unmarshal(d.Body, &msg)

		log.Print(msg)

		if err := s.service.Insert(ctx, &msg); err != nil {
			return err
		}
	}

	return nil
}
