package model

// Location ...
type Location struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates Coordinates `json:"coordinates" bson:"coordinates"`
}

//Coordinates ...
type Coordinates struct {
	Latitude  float32 `json:"lat" bson:"lat"`
	Longitude float32 `json:"lng" bson:"lng"`
	Accuracy  int     `json:"acc" bson:"acc"`
}
