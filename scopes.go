package netatgo

type (
	Scope = string
)

const (
	ReadStation Scope = "read_station" //: to retrieve weather station data (Getstationsdata, Getmeasure)

	ReadThermostats Scope = "read_thermostat" //: to retrieve thermostat data ( Homestatus, Getroommeasure...)

	WriteSThermostat Scope = "write_thermostat" //: to set up the thermostat (Synchomeschedule, Setroomthermpoint...)

	ReadCamera Scope = "read_camera" //: to retrieve Smart Indoor Cameradata (Gethomedata, Getcamerapicture...)

	WriteCamera Scope = "write_camera" //: to inform the Smart Indoor Camera that a specific person or everybody has left the Home (Setpersonsaway, Setpersonshome)

	AccessCamera Scope = "access_camera" //: to access the camera, the videos and the live stream *

	ReadPrescene Scope = "read_presence" //: to retrieve Smart Outdoor Camera data (Gethomedata, Getcamerapicture...)

	AccessPresence Scope = "access_presence" //: to access the camera, the videos and the live stream *

	ReadSmokedetector Scope = "read_smokedetector " //: to retrieve the Smart Smoke Alarm informations and events (Gethomedata, Geteventsuntil...)

	ReadHomecoach Scope = "read_homecoach" //: to read data coming from Smart Indoor Air Quality Monitor (gethomecoachsdata)

//If no scope is provided during the token request, the default is "read_station"
)
