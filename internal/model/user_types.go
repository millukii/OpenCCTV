package model

/*
UserType

Administrator
The account of this user type is called "admin", which has the permissions to access all supported
resources. And it must be activated at any time.
Operator
The account of this user type has the permissions to access common resources and less advanced
resources, see details in the description of each resource.
Viewer
The account of this user type has the permissions to access common resources only, see details in
the description of each resource.

*/
type UserType struct {
	Name        string
	Description string
}
