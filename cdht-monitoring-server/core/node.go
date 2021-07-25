package core

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Node struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IP_address   string             `json:"ip_address,omitempty" bson:"ip_address,omitempty"`
	Node_id  string             `json:"node_id" bson:"node_id,omitempty"`
    Created_date time.Time   `json:"created_date" bson:"created_date,omitempty"`
}