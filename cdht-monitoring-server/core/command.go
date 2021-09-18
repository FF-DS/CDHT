package core

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	COMMAND_OPERATION_STATUS_SUCCESS 	string = "SUCCESS"
	COMMAND_OPERATION_STATUS_FAILED     string = "FIALED"

	COMMAND_TYPE_HOP_COUNT		      	string = "COMMAND_TYPE_HOP_COUNT"
	COMMAND_TYPE_LOOK_UP            	string = "COMMAND_TYPE_LOOK_UP"
	COMMAND_TYPE_PING               	string = "COMMAND_TYPE_PING"
	COMMAND_TYPE_CLOSE_TOOL         	string = "COMMAND_TYPE_CLOSE_TOOL"

)
type ToolCommand struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	CreatedDate time.Time `json:"created_date" bson:"created_date,omitempty"`
	Type  string `bson:"type" json:"type"`
	OperationID primitive.ObjectID `bson:"operation_id" json:"operation_id"`
	OperationStatus  string `bson:"operation_status" json:"operation_status"`
	NodeId string `bson:"node_id" json:"node_id"`
	NodeAddress string `bson:"node_address" json:"node_address"`
	Body interface{} `bson:"body" json:"body"`
}
