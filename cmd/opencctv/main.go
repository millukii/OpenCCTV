package main

import (
	"fmt"
	"opencctv/internal/model"
	"opencctv/internal/service/system"
)

func main() {

	var con = &model.Connection{
		User:     "admin",
		Host:     "",
		Password: "",
		Port:     "8086",
		Protocol: "http",
	}
	svc := system.NewSystemService(con)

	/* 	resp, err := svc.GetDeviceInfo()

	   	if err != nil {
	   		panic(err)
	   	}
			 fmt.Printf("Device Info %+v\n", resp) */
	//	svc.GetDeviceCapabilities()
	//	svc.GetDeviceHardwareConfiguration()

	resp, err := svc.GetDeviceNetWorkInterfaces()

	if err != nil {
		panic(err)
	}
	fmt.Printf("Device Network Interfaces %+v\n", resp.NetworkInterface[0].IPAddress)
}
