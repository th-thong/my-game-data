package graph
import "go.mongodb.org/mongo-driver/v2/mongo"

type Resolver struct{
	DB *mongo.Database
}
