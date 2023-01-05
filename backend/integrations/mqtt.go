package integrations

import (
	"errors"
	"log"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gookit/config/v2"
	"rlesjak.com/ha-scheduler/logs"
)

var client mqtt.Client

func initialiseMqttIntegration(config config.Config) {

	// Initialise logging files
	mqttLogFile := logs.CreateCustomLoggerFile("mqtt")
	mqtt.DEBUG = log.New(mqttLogFile, "[DEBUG]", log.Ldate|log.Ltime)
	mqtt.ERROR = log.New(mqttLogFile, "[ERROR]", log.Ldate|log.Ltime|log.Lshortfile)

	opts := mqtt.NewClientOptions().AddBroker(config.String("env_mqtt_host"))
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Send message
	token := client.Publish("go-mqtt/debug", 0, false, "Hello from ha-scheduler!")
	go func() {
		_ = token.Wait()
		if token.Error() != nil {
			logs.Error.Println(token.Error())
		}
	}()

}

func getParsedCommand(command string) (string, string, error) {
	// Split topic and message
	// message should look like this: "topic/subtopic/Hello World!"
	lastSlashIndex := strings.LastIndex(command, "/")

	if lastSlashIndex == -1 {
		return "", "", errors.New("mqtt-parser> Topic missing!")
	}

	topic := command[:lastSlashIndex]
	message := command[lastSlashIndex+1:]

	return topic, message, nil
}

func mqttIntegrationHandler(command string) {
	// Should already be validated, so there is no need to check for errors
	// Lets se how this comment ages xD
	topic, message, _ := getParsedCommand(command)
	// Send message
	token := client.Publish(topic, 0, false, message)
	go func() {
		_ = token.Wait()
		if token.Error() != nil {
			// log error
			logs.Error.Println(token.Error())
		}
	}()
}

func mqttValidateCommand(command string) error {
	_, _, err := getParsedCommand(command)
	return err
}
