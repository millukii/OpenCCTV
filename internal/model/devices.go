package model

type Device struct {
	DeviceInfo   *DeviceInfo `xml:"DeviceInfo"`
	Capabilities []Capability
}

type DeviceInfo struct {
	DeviceName          string `xml:"deviceName"`
	DeviceID            string `xml:"deviceID"`
	Model               string `xml:"model"`
	SerialNumber        string `xml:"serialNumber"`
	MacAddress          string `xml:"madAddress"`
	firmwareVersion     string `xml:"firmwareVersion"`
	firmwareReleaseDate string `xml:"firmwareReleaseDate"`
	EncoderVersion      string `xml:"encoderVersion"`
	DeviceType          string `xml:"deviceType"`
	TelecontrolID       string `xml:"telecontrolID"`
	HardwareVersion     string `xml:"hardwareVersion"`
}
