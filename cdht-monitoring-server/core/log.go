package core

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
    "time"
	"math/big"
)

type LogEntry struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedDate time.Time   `json:"created_date" bson:"created_date,omitempty"`
	Type string `json:"type" bson:"type,omitempty"`
	OperationStatus string `json:"operation_status" bson:"operation_status,omitempty"`
	LogLocation string  `json:"log_location" bson:"log_location,omitempty"`
	NodoeId *big.Int `json:"node_id" bson:"node_id,omitempty"`
	NodeAddress string `json:"node_address" bson:"node_address,omitempty"`
	LogBody interface{} `json:"log_body" bson:"log_body,omitempty"`
}