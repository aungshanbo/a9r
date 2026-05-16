package models

type Ec2instance struct {
	ID        string
	State     string
	Type      string
	Name      string
	AZ        string
	PrivateIP string
	PublicIP  string
}

