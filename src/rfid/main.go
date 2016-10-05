package main

import (
	"fmt"
	"flag"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	evdev "github.com/gvalkov/golang-evdev"
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

	device, error := evdev.Open("/dev/input/by-id/usb-13ba_Barcode_Reader-event-kbd")

	fmt.Println(device)
	fmt.Println(error)

	for true {
		var action = ""
		for {
		  //ie, error := device.ReadOne()
			ie,_ := device.Read()
			input := convertToCharacter(ie)

			if (input == "\n") {
				fmt.Println("Received RFID read:"+action)
				if !Clocal.IsConnected() {
					if token := Clocal.Connect(); token.Wait() && token.Error() != nil {
						panic(token.Error())
					}
				}
				if ptoken := Clocal.Publish(intopic, 0, false,	action); ptoken.Wait() && ptoken.Error() != nil {
					panic(ptoken.Error())
				}
				action = ""
		  } else {
				action = action + input
			}


		}
	}

}

func convertToCharacter(ies []evdev.InputEvent) string {
  var input = ""
	for _, v := range ies {
		// Only get key up events
		if (v.Type == 0x01 && v.Value == 0x01) {
			switch v.Code {
			// new line read
			case 28:
				input = input + "\n"
				break
			case 0x02:
        input = input + "1"
			case 0x03:
        input = input + "2"
			case 0x04:
        input = input + "3"
			case 0x05:
        input = input + "4"
			case 0x06:
        input = input + "5"
			case 0x07:
        input = input + "6"
			case 0x08:
        input = input + "7"
			case 0x09:
        input = input + "8"
			case 0x10:
        input = input + "9"
			case 11:
        input = input + "0"
			// ignore
			default:
			}
		}
	}
	return input
}
