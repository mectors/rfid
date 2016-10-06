package main

import (
	"fmt"
	"flag"
	"os"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var conns = flag.Int("conns", 10, "how many conns (0 means infinite)")
var host = flag.String("host", "localhost:1883", "hostname of broker")
var clientID = flag.String("clientid", "rfid", "the mqtt clientid")
var user = flag.String("user", "", "username")
var pass = flag.String("pass", "", "password")

var topic = "discovery"
var intopic = "sensor/rfid/in"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	flag.Parse()

  // Prepare the local MQTT connection
	opts2 := MQTT.NewClientOptions().AddBroker("tcp://"+*host)
  opts2.SetClientID(*clientID)
  opts2.SetDefaultPublishHandler(local)


  //create and start a client using the above ClientOptions
  Clocal := MQTT.NewClient(opts2)
  if token := Clocal.Connect(); token.Wait() && token.Error() != nil {
    panic(token.Error())
  }
	fmt.Println("Connected locally")
  defer Clocal.Disconnect(250)

	// Say we are ready for action
	Clocal.Publish(topic, 0, false,	intopic)
  fmt.Println("Published to:"+topic+" that we are listening on:"+intopic)
	startup = false

	// Publish to the discovery service
	Clocal.Publish(topic, 0, false,	intopic)

	for true {
		var action string
		fmt.Scanln(&action)
		result, mid = Clocal.Publish(intopic, 0, false,	action)
		if result == MQTT.MQTT_ERR_SUCCESS
		{
			if token := Clocal.Reconnect(); token.Wait() && token.Error() != nil {
				panic(token.Error())
			}
			Clocal.Publish(intopic, 0, false,	action)
		}
	}

}
