package model



type BrokerConfig struct {
	Id       string   `json:"id,omitempty" bson:"_id,omitempty"`
	ClientId string   `json:"clientId" bson:"client_id"`
	Topics   []string `json:"topics,omitempty" bson:"topics,omitempty"`
	Broker   string   `json:"broker" bson:"broker"`
	User     string   `json:"user" bson:"user"`
	Passwd   string   `json:"passwd" bson:"passwd"`
	Qos      byte     `json:"qos" bson:"qos`
	Store    string   `json:"store" bson:"cleanses" bson:"cleanses"`
	State    int      `json:"state" bson:"state"`
}


