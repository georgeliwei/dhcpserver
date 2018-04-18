package dhcpdb

/*
* define the interface for different db
* objectName: this parameter describe the store object
* objectValue: this parameter describe the store value, it is an json structure
* for QueryRecord, the return value is an json structure according query condition
 */

type TobjectType struct {
	objectKey string
	objectVal string
}

type Dbdriver interface {
	ConnectDbDriver(connectUrl string) bool
	DisconnectDbDriver() bool
	AddRecord(objectName string, objectValue []TobjectType) error
	UpdateRecord(objectName string, objectValue []TobjectType, opCondition []TobjectType) error
	RemoveRecord(objectName string, opCondition []TobjectType) error
	QueryRecord(objectName string, opCondition []TobjectType) (string, error)
}
