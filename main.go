package main

import (
	"fmt"

	"os"
	"time"

	"github.com/flashmob/go-guerrilla"
	"github.com/rahulvramesh/mailbay-backend/drivers"
	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//forever loop
	forever := make(chan bool)

	// cfg := &guerrilla.AppConfig{LogFile: "mails.log"}
	// sc := guerrilla.ServerConfig{
	// 	ListenInterface: "127.0.0.1:2525",
	// 	IsEnabled:       true,
	// }

	// cfg.Servers = append(cfg.Servers, sc)

	d := guerrilla.Daemon{}
	_, err := d.LoadConfig("guerrillad.conf.json")

	if err != nil {
		fmt.Println("ReadConfig error", err)

	}

	//add
	d.AddProcessor("SqlLite", drivers.SqlLiteProcessor)

	err = d.Start()

	if err != nil {
		fmt.Println(err)
	}

	log.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	//**** ned to fix below area
	log.Debug("Trying to shutdown")

	go func() {
		select {
		// exit if graceful shutdown not finished in 60 sec.
		case <-time.After(time.Second * 60):
			log.Fatalf("graceful shutdown timed out")
			os.Exit(1)
		}
	}()
	d.Shutdown()
	log.Println("Shutdown completed, exiting.")

}
