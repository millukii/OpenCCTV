package content

type ContentMgmtService interface {
	ActivateNetWorkDeviceByNVR()
}

type contentMgmtService struct {
}

func NewContentMgmtService() ContentMgmtService {

	return &contentMgmtService{}

}

func (svc *contentMgmtService) ActivateNetWorkDeviceByNVR() {

	/*
	   5.1.2 Activate Network Camera via NVR
	   For network devices, except directly activating them, you can also activate them via NVR if they can
	   be searched on the same network domain of NVR.
	   Steps
	   1. Call /ISAPI/ContentMgmt/InputProxy/search by GET method to search for the online network
	   devices in the same network domain with the NVR.
	   2. Optional: Call /ISAPI/ContentMgmt/InputProxy/channels/activate/capabilities by GET
	   method to get the activation capability of network devices for reference.
	   3. Call /ISAPI/ContentMgmt/InputProxy/channels/activate by PUT method to activate the
	   searched online devices via NVR.
	*/
}
