package publisher

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Publisher struct {
	client   mqtt.Client
	nodeName string
}

func NewPublisher(brokerURL, nodeName, caPath, username, password string) (*Publisher, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURL)
	opts.SetClientID(nodeName)
	opts.SetUsername(username)
	opts.SetPassword(password)
	caCert, err := os.ReadFile(caPath)
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	tlsConfig := &tls.Config{RootCAs: caCertPool}
	opts.SetTLSConfig(tlsConfig)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &Publisher{client, nodeName}, nil
}

func (p *Publisher) Publish(metric string, value float64) error {
	payload := fmt.Sprintf("%.1f", value)
	token := p.client.Publish(fmt.Sprintf("homelab/%s/%s", p.nodeName, metric), 0, false, payload)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
