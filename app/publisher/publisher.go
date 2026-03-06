package publisher

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Publisher struct {
	client   mqtt.Client
	nodeName string
}

func NewPublisher(brokerURL, nodeName, username, password string) (*Publisher, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURL)
	opts.SetClientID(nodeName)
	opts.SetUsername(username)
	opts.SetPassword(password)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &Publisher{client, nodeName}, nil
}

func (p *Publisher) Publish(metric string, value float64) error {
	payload := fmt.Sprintf("%.2f", value)
	token := p.client.Publish(fmt.Sprintf("homelab/%s/%s", p.nodeName, metric), 0, false, payload)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
