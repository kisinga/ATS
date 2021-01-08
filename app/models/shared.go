package models

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MarshalObjectID(b primitive.ObjectID) graphql.Marshaler {
	// return graphql.WriterFunc(func(w io.Writer) {
	// 	w.Write([]byte(b.Hex()))
	// })
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(b.Hex()))
	})
}

func UnmarshalObjectID(v interface{}) (primitive.ObjectID, error) {
	switch v := v.(type) {
	case string:
		return primitive.ObjectIDFromHex(v)
	default:
		return primitive.NilObjectID, fmt.Errorf("%T is not a string", v)
	}
}
