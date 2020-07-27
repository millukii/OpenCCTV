package model

/*
Connection

- Protocol string * HTTP, HTTPS, RTSP *
- Host string     * Hostname IP Address DNS *
- Port string     * 80 443 *
- AbsPath string  Resource URI: /ServiceName/ResourceType/resource. * /ISAPI/System/... *
*/
type Connection struct {
	User     string
	Password string
	Protocol string
	Host     string
	Port     string
	AbsPath  string
}
