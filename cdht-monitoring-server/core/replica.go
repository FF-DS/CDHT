package core

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type ReplicaInformation struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedDate time.Time `json:"created_date" bson:"created_date,omitempty"`
	IpAddress string `json:"ip_address" bson:"ip_address,omitempty"`
	ParentNodeIpAddress string `json:"parent_node_ip_address" bson:"parent_node_ip_address,omitempty"`
	ParentNodeId string `json:"parent_node_id" bson:"parent_node_id,omitempty"`
	Status string `json:"status" bson:"status,omitempty"`
}