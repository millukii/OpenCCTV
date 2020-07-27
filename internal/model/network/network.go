package network

type Network struct {
	NetworkInterfaceList *NetworkInterfaceList `xml:"NetWorkInterfaceLists"`
}
type NetworkInterfaceList struct {
	NetworkInterface []*NetworkInterface `xml:"NetworkInterface"`
}

type NetworkInterface struct {
	ID        string     `xml:"id"`
	IPAddress *IPAddress `xml:"IPAddress"`
	Discovery *Discovery `xml:"Discovery"`
	Link      *Link      `xml:"Link"`
}

type Link struct {
	MacAddress      string `xml:"MacAddress"`
	AutoNegotiation bool   `xml:"autoNegotiation"`
	Speed           int    `xml:"speed"`
	Duplex          string `xml:"full"`
	MTU             int    `xml:"MTU"`
}
type Discovery struct {
	UPnP     *Parameter `xml:"UPnP"`
	Zeroconf *Parameter `xml:"Zeroconf"`
}
type Parameter struct {
	Enabled bool `xml:"enabled"`
}

type IPAddress struct {
	IPVersion      string     `xml:"ipVersion"`
	AddressingType string     `xml:"addressingType"`
	IPAddress      string     `xml:"ipAddress"`
	SubnetMask     string     `xml:"subnetMask"`
	IPV6Address    string     `xml:"ipv6Address"`
	BitMask        string     `xml:"bitMask"`
	DefaultGateway *IP        `xml:"DefaultGateway"`
	PrimaryDNS     *IP        `xml:"PrimaryDNS"`
	SecondaryDNS   *IP        `xml:"SecondaryDNS"`
	DNSEnable      *Parameter `xml:"DNSEnable"`
}
type IP struct {
	IPAddress string `xml:"ipAddress"`
}
