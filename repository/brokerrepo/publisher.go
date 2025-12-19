package brokerrepo

import (
	"bytes"
	"encoding/json"

	"router-gostashlg-template/entities/brokermessage"
	"router-gostashlg-template/entities/common/logger"
	"router-gostashlg-template/repository/built_in/broker"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/randyardiansyah25/libpkg/util/env"
)

func PublishMessage(payload brokermessage.Metadata[any]) (er error) {
	xchange := env.GetString("rabbit.exchange")
	message, er := json.Marshal(payload)
	if er != nil {
		return
	}

	var bbPrettyJson bytes.Buffer
	prettyJson := string(message)
	if er = json.Indent(&bbPrettyJson, message, "", "    "); er == nil {
		prettyJson = bbPrettyJson.String()
	}

	logger.PrintLogf("Publishing message to [%s], message :\n%s", xchange, prettyJson)
	er = broker.BrokerChannel.Publish(
		xchange, // exchange
		"",      // routing key
		true,    // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})

	return
}
