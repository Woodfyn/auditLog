package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Woodfyn/auditLog/internal/config"
	"github.com/Woodfyn/auditLog/internal/repository"
	"github.com/Woodfyn/auditLog/internal/server"
	"github.com/Woodfyn/auditLog/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CONF_FOLDER   = "configs"
	CONF_FILENAME = "main"
	CONF_ENVNAME  = ".main"
)

func main() {
	cfg, err := config.NewConfig(CONF_FOLDER, CONF_FILENAME, CONF_ENVNAME)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: cfg.Database.Username,
		Password: cfg.Database.Password,
	})

	opts.ApplyURI(cfg.Database.URI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	db := client.Database(cfg.Database.Database)

	auditRepo := repository.NewAuditMongo(db)
	auditService := service.NewAuditRepo(auditRepo)
	auditSrv := server.NewAuditServer(auditService)
	srv := server.New(auditSrv)

	log.Println("SERVER STARTED", time.Now())

	if err := srv.ListenAndServe(cfg.Server.Port); err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("SERVER STOPED", time.Now())

	client.Disconnect(ctx)
}
