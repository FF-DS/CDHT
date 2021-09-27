package core

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	LOG_OPERATION_STATUS_SUCCESS 	 string = "SUCCESS"
	LOG_OPERATION_STATUS_FAILED      string = "FIALED"

	LOG_TYPE_ROUTING_TABLE           string = "TYPE_ROUTING_TABLE"
	LOG_TYPE_NODE_INFORMATION        string = "TYPE_NODE_INFORMATION"
	LOG_TYPE_NETWORK_TOOL            string = "TYPE_NETWORK_TOOL"
	LOG_TYPE_APP_SERVICE             string = "TYPE_APP_SERVICE"
	LOG_TYPE_CONFIGURATION_SERVICE   string = "TYPE_CONFIGURATION_SERVICE"


	LOG_LOCATION_TYPE_INCOMMING      string = "LOCATION_TYPE_INCOMMING"
	LOG_LOCATION_TYPE_LEAVING        string = "LOCATION_TYPE_LEAVING"
	LOG_LOCATION_TYPE_SELF           string = "LOCATION_TYPE_SELF"
)

type LogEntry struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedDate time.Time   `json:"created_date" bson:"created_date,omitempty"`
	Type string `json:"type" bson:"type,omitempty"`
	OperationStatus string `json:"operation_status" bson:"operation_status,omitempty"`
	LogLocation string  `json:"log_location" bson:"log_location,omitempty"`
	NodoeId string `json:"node_id" bson:"node_id,omitempty"`
	NodeAddress string `json:"node_address" bson:"node_address,omitempty"`
	LogBody interface{} `json:"log_body" bson:"log_body,omitempty"`
}