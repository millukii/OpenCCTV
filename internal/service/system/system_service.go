package system

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"opencctv/internal/model"
	n "opencctv/internal/model/network"
	"strings"
)

/*
Get device capability or encryption capability to check if the device or channel supports stream
encryption by the request URL: GET /ISAPI/System/capabilities
• Get device capability
Request URL: GET /ISAPI/System/deviceInfo/capabilities
• Get device information
Request URL: GET /ISAPI/System/deviceInfo
*/
type SystemService interface {
	//	GetSystemCapabilities()
	//GetDeviceHardwareConfiguration()
	GetDeviceCapabilities()
	GetDeviceNetWorkInterfaces() (*n.NetworkInterfaceList, error)
	GetDeviceSystemStatus()
	//	GetDeviceInfo() (*model.DeviceInfo, error)
	//	ShutDownDevice()
	//	RebootDevice()
	//	RestoreDeviceDefaultSettings()
	//	GetConfigurationFile()
}

type systemService struct {
	Connection *model.Connection
}

func NewSystemService(connection *model.Connection) SystemService {
	return &systemService{
		Connection: &model.Connection{
			User:     connection.User,
			Password: connection.Password,
			Host:     connection.Host,
			Port:     connection.Port,
			Protocol: connection.Protocol,
			AbsPath:  "/ISAPI/System",
		},
	}
}

func (svc *systemService) GetDeviceSystemStatus()
func (svc *systemService) GetDeviceNetWorkInterfaces() (*n.NetworkInterfaceList, error) {
	var stringConnection strings.Builder
	stringConnection.WriteString(
		svc.Connection.Protocol +
			"://" +
			svc.Connection.User +
			":" +
			svc.Connection.Password +
			"@" +
			svc.Connection.Host +
			":" +
			svc.Connection.Port +
			svc.Connection.AbsPath +
			"/Network/interfaces",
	)
	fmt.Printf("String connection %+v\n", stringConnection.String())
	resp, err := http.Get(stringConnection.String())
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	var inter *n.NetworkInterfaceList
	err = xml.Unmarshal(body, &inter)

	if err != nil {
		fmt.Println(err.Error())
	}

	return inter, nil

}

func (svc *systemService) GetDeviceCapabilities() {
	var stringConnection strings.Builder
	stringConnection.WriteString(
		svc.Connection.Protocol +
			"://" +
			svc.Connection.User +
			":" +
			svc.Connection.Password +
			"@" +
			svc.Connection.Host +
			":" +
			svc.Connection.Port +
			svc.Connection.AbsPath +
			"/deviceInfo/capabilities",
	)
	fmt.Printf("String connection %+v\n", stringConnection.String())
	resp, err := http.Get(stringConnection.String())
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Printf("%+v\n", string(body))

}

func (svc *systemService) GetDeviceInfo() (*model.DeviceInfo, error) {

	var stringConnection strings.Builder
	stringConnection.WriteString(
		svc.Connection.Protocol +
			"://" +
			svc.Connection.User +
			":" +
			svc.Connection.Password +
			"@" +
			svc.Connection.Host +
			":" +
			svc.Connection.Port +
			svc.Connection.AbsPath +
			"/deviceInfo",
	)
	fmt.Printf("String connection %+v\n", stringConnection.String())
	resp, err := http.Get(stringConnection.String())
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	var dinfo *model.DeviceInfo
	err = xml.Unmarshal(body, &dinfo)

	if err != nil {
		fmt.Println(err.Error())
	}

	return dinfo, nil

}
