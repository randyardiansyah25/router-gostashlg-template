package consumer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"router-gostashlg-template/entities/brokermessage"
	"router-gostashlg-template/entities/common/logger"
	"router-gostashlg-template/repository/built_in/broker"
	"time"

	"github.com/randyardiansyah25/libpkg/util/env"

	amqp "github.com/rabbitmq/amqp091-go"
)

// ** Modul untuk consume notifikasi yang dipublish dari service-service yang memproses transaksi
func Start() {
	reconnectPeriod := env.GetInt("rabbit.reconnect_period", 10)

	retryCfg := env.GetInt("rabbit.consumer_add_retry_periode", 5)
	if retryCfg == 0 {
		retryCfg = 5
	}

	retryPeriode := time.Duration(reconnectPeriod+retryCfg) * time.Second

	for {
		if er := startConsumer(); er != nil {
			logger.PrintLog(fmt.Sprintf("error consume : %s. \n\nretry in %d second...", er.Error(), reconnectPeriod+retryCfg))
			time.Sleep(retryPeriode)
		}
	}
}

func startConsumer() (er error) {
	logger.PrintLog("Start Consume()...")
	queueName := env.GetString("rabbit.queue_name")
	messages, er := broker.BrokerChannel.Consume(
		queueName, // queue name
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)

	if er != nil {
		return er
	}

	for message := range messages {
		go messageProcessor(message)
	}

	return errors.New("force closed consume channel")
}

func messageProcessor(message amqp.Delivery) {
	var prettyJson bytes.Buffer
	var metaData brokermessage.Metadata[any]
	if er := json.Unmarshal(message.Body, &metaData); er != nil {
		logger.PrintErrorf("We got message, it seems the message is not in json format : %s, error : %v\n", string(message.Body), er)
		message.Reject(false)
	} else {
		_ = json.Indent(&prettyJson, message.Body, "", "    ")
		logger.PrintLogf("A notification message was received :\n%s\n", prettyJson.String())

		//* {
		//*   your code to handle the message here
		//* }
		message.Ack(false)

	}
}
