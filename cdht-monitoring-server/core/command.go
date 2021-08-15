package core

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

const (
	PRIORITY_HIGH string = "HIGH"
	PRIORITY_MEDIUM = "MEDIUM"
	PRIORITY_LOW = "LOW"
)

type Command struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedDate time.Time `json:"created_date" bson:"created_date,omitempty"`
	CommandString string `json:"command_string" bson:"command_string,omitempty"`
	CommandDestinationNodeId string `json:"command_destination_node_id" bson:"command_destination_node_id,omitempty"`
	CommandSourceNodeId string `json:"command_source_node_id" bson:"command_source_node_id,omitempty"`
	CommandPriority string `json:"command_priority" bson:"command_priority,omitempty"`
}

type CommandResult struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedDate time.Time `json:"created_date" bson:"created_date,omitempty"`
	CommandId primitive.ObjectID `json:"command_id,omitempty" bson:"command_id,omitempty"`
	ResultMessage string `json:"result_message,omitempty" bson:"result_message,omitempty"`
}


