package integrations

import (
	"errors"
	"fmt"
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

func GetMqttClient() mqtt.Client {
	return client
}

func parsedCommand(command string) (string, string, error) {
	// Split topic and message
	// message should look like this: "topic/subtopic/Hello World!"
	lastSlashIndex := strings.LastIndex(command, "/")

	if lastSlashIndex == -1 {
		return "", "", errors.New("mqtt-parser> Topic missing!")
	}

	fmt.Printf("MQTTHANDLER: topic: [%s], message[%s]",
		command[:lastSlashIndex],
		command[lastSlashIndex:])

	return "", "", nil
}

func mqttIntegrationHanlder(command string) {
	fmt.Println("MQTT INTEGRATION HANDLER")
	// Send message
	// token := client.Publish(splitMessage[0], 0, false, splitMessage[1])
	// go func() {
	// 	_ = token.Wait()
	// 	if token.Error() != nil {
	// 		logs.Error.Println(token.Error())
	// 	}
	// }()
}

func mqttValidateCommand(command string) error {
	_, _, err := parsedCommand(command)
	return err
}
