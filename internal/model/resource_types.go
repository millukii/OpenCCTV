package model

type ResourceType struct {
	Name        string
	Description string
}

var ResourceTypes = []ResourceType{
	{
		Name:        "System",
		Description: "System related resources",
	},
	{
		Name:        "Security",
		Description: "Security related resources",
	},
	{
		Name:        "Streaming",
		Description: "Video streaming and management related resources",
	},
	{
		Name:        "Event",
		Description: "Event/alarm related resources",
	},
	{
		Name:        "PTZCtrl",
		Description: "PTZ control related resources",
	},
	{
		Name:        "Image",
		Description: "Video encoding and image related resources",
	},
	{
		Name:        "ContentMgmt",
		Description: "Storage management related resources",
	},
}
