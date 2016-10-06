# RF ID Reader Snap

Snap for the usb-13ba_Barcode_Reader-event-kbd RFID reader.

Build:
git clone ...
snapcraft

Run:

sudo snap install rfid_0.1_... --force-dangerous --devmode

This will publish to localhost:1833 to the MQTT topic sensor/rfid/in the digital code of the security card that is read.
