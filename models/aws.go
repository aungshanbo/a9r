package models

import "time"

type Ec2instance struct {
	ID        string
	State     string
	Type      string
	Name      string
	AZ        string
	PrivateIP string
	PublicIP  string
}

type S3Bucket struct {
	Name         string
	CreationDate time.Time
}

type S3BucketDetail struct {
	Name            string
	Region          string
	CreationDate    time.Time
	ObjectCount     int64
	SizeBytes       int64
	Versioning      string
	Encryption      string
	ObjectLock      string
	PublicAccess    string
	ObjectOwnership string
	ACL             string
	Policy          string
	LifecycleRules  int
	Replication     string
	AccessLogging   string
	Tags            map[string]string
}
