package core

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConfigurationProfile struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedDate time.Time `json:"created_date" bson:"created_date,omitempty"`
	ConfigurationName string `json:"configuration_name" bson:"configuration_name,omitempty"`
	JumpSpaceBalancing int `json:"jump_space_balancing" bson:"jump_space_balancing,omitempty"`
	ConfigurationDescription string `json:"configuration_description" bson:"configuration_description,omitempty"`
}

