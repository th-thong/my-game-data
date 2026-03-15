package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Element struct {
	Id   int    `json:"Id" bson:"Id"`
	Icon string `json:"Icon" bson:"Icon"`
	Name string `json:"Name" bson:"Name"`
}

type WeaponType struct {
	DBId bson.ObjectID `json:"DBId" bson:"_id,omitempty"`
	Id   int           `json:"Id" bson:"Id"`
	Name string        `json:"Name" bson:"Name"`
	Icon string        `json:"Icon" bson:"Icon"`
}

type Character struct {
	DBId         bson.ObjectID `json:"DBId" bson:"_id,omitempty"`
	Id           int           `json:"Id" bson:"Id"`
	Name         string        `json:"Name" bson:"Name"`
	QualityID    int           `json:"QualityId" bson:"QualityId"`
	Element      Element       `json:"Element" bson:"Element"`
	RoleHeadIcon string        `json:"RoleHeadIcon" bson:"RoleHeadIcon"`
	WeaponType   WeaponType    `json:"WeaponType" bson:"WeaponType"`
	BannerImg    string        `json:"BannerImg" bson:"BannerImg"`
}

type Weapon struct {
	DBId      bson.ObjectID `json:"DBId" bson:"_id,omitempty"`
	Id        int           `json:"Id" bson:"Id"`
	Name      string        `json:"Name" bson:"Name"`
	Icon      string        `json:"Icon" bson:"Icon"`
	Type      int           `json:"Type" bson:"Type"`
	QualityID int           `json:"QualityId" bson:"QualityId"`
	TypeName  string        `json:"TypeName" bson:"TypeName"`
	TypeIcon  string        `json:"TypeIcon" bson:"TypeIcon"`
}

type EchoSet struct {
	Id   int    `json:"Id" bson:"Id"`
	Icon string `json:"Icon" bson:"Icon"`
	Name string `json:"Name" bson:"Name"`
}

type Echo struct {
	DBId        bson.ObjectID `json:"DBId" bson:"_id,omitempty"`
	Id          int           `json:"Id" bson:"Id"`
	Name        string        `json:"Name" bson:"Name"`
	Rarity      int           `json:"Rarity" bson:"Rarity"`
	Icon        string        `json:"Icon" bson:"Icon"`
	IconMiddle  string        `json:"IconMiddle" bson:"IconMiddle"`
	IconSmall   string        `json:"IconSmall" bson:"IconSmall"`
	PhantomType int           `json:"PhantomType" bson:"PhantomType"`
	Element     Element       `json:"Element" bson:"Element"`
	Type        string        `json:"Type" bson:"Type"`
	Attributes  string        `json:"Attributes" bson:"Attributes"`
	EchoGroups  []EchoSet     `json:"FetterGroups" bson:"EchoGroups"`
}
