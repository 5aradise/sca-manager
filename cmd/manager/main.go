package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/5aradise/sca-manager/config"
	"github.com/5aradise/sca-manager/pkg/db/postgresql"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

var envPath = flag.String("env", "", "Path to env file")

func main() {
	flag.Parse()

	if *envPath != "" {
		err := config.Load(*envPath)
		if err != nil {
			log.Fatal("can't load env vars: ", err)
		}
	}

	cfg, err := config.New()
	if err != nil {
		log.Fatal("can't load config: ", err)
	}

	conn, err := postgresql.New(cfg.DB.Address, cfg.DB.User, cfg.DB.Password, cfg.DB.Port, cfg.DB.Name)
	if err != nil {
		log.Fatal("can't open sql: ", err)
	}
	defer conn.Close()

	// create service

	// create router

	// init router

	app := fiber.New(fiber.Config{
		ReadTimeout: cfg.Server.ReadTimeout,
		IdleTimeout: cfg.Server.IdleTimeout,

		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowMethods:  []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		ExposeHeaders: []string{"Link"},
	}))
	app.Use(logger.New())

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	serverErr := make(chan error)
	go func() {
		serverErr <- app.Listen(net.JoinHostPort("", cfg.Server.Port))
	}()

	select {
	case s := <-interrupt:
		log.Println("signal interrupt: ", s.String())
	case err := <-serverErr:
		log.Println("server error: ", err)
	}

	err = app.Shutdown()
	if err != nil {
		log.Fatal("can't shutdown server: ", err)
	}
}
