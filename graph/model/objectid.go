package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ObjectID = bson.ObjectID

func MarshalObjectID(id ObjectID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(strconv.Quote(id.Hex())))
	})
}

func UnmarshalObjectID(v interface{}) (ObjectID, error) {
	switch v := v.(type) {
	case string:
		return bson.ObjectIDFromHex(v)
	default:
		return bson.ObjectID{}, fmt.Errorf("%T is not a valid ObjectID", v)
	}
}
