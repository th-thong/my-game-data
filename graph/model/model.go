package model

type Character struct {
    ID      string `json:"id" bson:"_id,omitempty"`
    Name    string `json:"name" bson:"name"`
    RoundIcon string `json:"round_icon" bson:"round_icon"`
}

type NewCharacterData struct {
    Name    string `json:"name" bson:"name"`
    RoundIcon string `json:"round_icon" bson:"round_icon"`
}

type Weapon struct {
    ID      string `json:"id" bson:"_id,omitempty"`
    Name    string `json:"name" bson:"name"`
    RoundIcon string `json:"round_icon" bson:"round_icon"`
}

type NewWeaponData struct {
    Name    string `json:"name" bson:"name"`
    RoundIcon string `json:"round_icon" bson:"round_icon"`
}

