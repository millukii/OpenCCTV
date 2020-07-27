package streaming

/*
StreamingService

/ISAPI/Streaming/
encryption/capabilities?format=json
*/
type StreamingService interface {
	GetStreamingCapabilities()
}
