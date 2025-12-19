package broker

import (
	"fmt"
	"os"
	"router-gostashlg-template/entities/app"
	"router-gostashlg-template/entities/common/logger"
	"time"

	"github.com/randyardiansyah25/libpkg/security/aes"
	"github.com/randyardiansyah25/libpkg/util/env"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	IsUse          = false
	BrokerChannel  *amqp.Channel
	BrokerPrepared = false
	CloseChan      chan *amqp.Error
)

func ConnectToRabbit() {
	var er error
	//amqp://<user>:<password>@localhost:5672/
	rabbitUser := env.GetString("rabbit.user")
	rabbitHassPwd := env.GetString("rabbit.password")
	var rabbitPwd string
	if rabbitHassPwd != "" {
		pvKey := []byte(app.PrivateKey)
		rabbitPwd, er = aes.Decrypt(pvKey, pvKey, rabbitHassPwd)
		if er != nil {
			logger.PrintError("Failed to decode rabbit password : ", er)
			os.Exit(91)
			return
		}
	}
	rabbitHost := env.GetString("rabbit.host")
	url := fmt.Sprintf("amqp://%s:%s@%s/", rabbitUser, rabbitPwd, rabbitHost)

	heartbeatPeriod := env.GetInt("rabbit.heartbeat_period", 60)
	reconnectPeriod := env.GetInt("rabbit.reconnect_period", 6)

	for {
		logger.PrintLog("Connecting to rabbit..")
		conn, er := amqp.DialConfig(url, amqp.Config{
			Heartbeat: time.Duration(heartbeatPeriod) * time.Second,
		})
		if er == nil {
			CloseChan = conn.NotifyClose(make(chan *amqp.Error))
			BrokerChannel, er = conn.Channel()
			if er == nil {
				if !BrokerPrepared {
					prepareBroker()
					BrokerPrepared = true
				}
				break
			}
			logger.PrintError("Closing rabbit connection..")
			_ = conn.Close()
		}
		logger.PrintWarnf("Failed to connect to Broker: %s. Retrying in %d second...", er, reconnectPeriod)
		time.Sleep(time.Duration(reconnectPeriod) * time.Second)
	}
	logger.PrintLog("Connected to rabbit..")
}

func prepareBroker() {

	xchange := env.GetString("rabbit.exchange")
	xchangeType := env.GetString("rabbit.exchange_type")

	er := BrokerChannel.ExchangeDeclare(
		xchange,     //name
		xchangeType, //type
		true,        //durable,
		false,       //auto-delete
		false,       //internal
		false,       //no-wait,
		nil,         //arguments
	)

	if er != nil {
		logger.PrintError("error prepare exchange : ", er)
		os.Exit(99)
	}
}

func BrokerClosedChannelObserver() {
	for {
		<-CloseChan
		reconnectPeriod := env.GetInt("rabbit.reconnect_period", 10)
		logger.PrintError(fmt.Sprintf("disconnected from rabbit : re-connection in %d seconds..", reconnectPeriod))
		ConnectToRabbit()
	}
}
