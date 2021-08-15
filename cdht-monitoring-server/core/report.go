package core

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

const (
	TYPE_CORRUPTED string = "CORRUPTED"
	TYPE_LEAVING  = "LEAVING"
	TYPE_INCOMING  = "INCOMING"
)


type NormalPacket struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedDate time.Time `json:"created_date" bson:"created_date,omitempty"`
	PacketType string `json:"packet_type,omitempty" bson:"packet_type,omitempty"`
	SourceIP string `json:"source_ip,omitempty" bson:"source_ip,omitempty"`
	DestinationIP string `json:"destination_ip,omitempty" bson:"destination_ip,omitempty"`
	SourceNodeID string `json:"source_node_id,omitempty" bson:"source_node_id,omitempty"`
	DestinationNodeID string `json:"destination_node_id,omitempty" bson:"destination_node_id,omitempty"`
}

type TestPacket struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedDate time.Time `json:"created_date" bson:"created_date,omitempty"`
	PacketType string `json:"packet_type,omitempty" bson:"packet_type,omitempty"`
	SourceIP string `json:"source_ip,omitempty" bson:"source_ip,omitempty"`
	DestinationIP string `json:"destination_ip,omitempty" bson:"destination_ip,omitempty"`
	SourceNodeID string `json:"source_node_id,omitempty" bson:"source_node_id,omitempty"`
	DestinationNodeID string `json:"destination_node_id,omitempty" bson:"destination_node_id,omitempty"`
	RTT int `json:"rtt,omitempty" bson:"rtt,omitempty"`
	Latency int `json:"latency,omitempty" bson:"latency,omitempty"`
	SRTT int `json:"srtt,omitempty" bson:"srtt,omitempty"`
	TTL int `json:"ttl,omitempty" bson:"ttl,omitempty"`
}