package main

import (
	"context"
	"fmt"
	"github.com/hikit-io/hikit-go/database/hkmongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
	Age  uint64             `bson:"age"`
	Addr string             `bson:"addr"`
}

var col *mongo.Collection
var ctx = context.Background()

var db *hkmongo.Database

func init() {
	ctx := context.Background()
	o := options.Client()
	o.ApplyURI(url)
	cli, _ := mongo.NewClient(o)
	cli.Connect(ctx)
	db = hkmongo.NewDB(cli, "test")
	//col = db.Collection("test")

}

type Test struct {
	Name string `json:"name" bson:"name"`
	Age  int64  `json:"age" bson:"age"`
	Addr string `json:"addr" bson:"addr"`
}

func (t Test) DocName() string {
	return "test"
}

func main() {

	builder := hkmongo.NewBuilder()
	//builder.Field("name").Regex("nieaowei")
	//builder.Field("age").LessThan(30)
	builder.OrFc(func(br *hkmongo.Builder) {
		br.Field("name").Equal("nieaowei")
	}).OrFc(func(br *hkmongo.Builder) {
		br.Field("age").Equal(11)
	})
	ts := []Test{}

	err := db.Col("test").Limit(1).HFind(ctx, builder, &ts)
	if err.Err() != nil {
		panic(err)
	}
	fmt.Println(ts)

	querySturct := &Test{
		Age: 11,
	}

	findRes := db.Col("test").HFind(ctx, querySturct, &ts)
	if findRes.Err() != nil {
		panic(err)
	}
	fmt.Println(ts)

	findOneRes := db.Col("test").HFindOne(ctx, builder, querySturct)
	if findOneRes.Err() != nil {
		panic(err)
	}
	fmt.Println(querySturct)

	r := db.Col("test").HUpdateOne(ctx, builder, querySturct)
	if r.Err() != nil {
		panic(r.Error())
	}

	if r.ExceptNoDocuments() != nil {

	}
	fmt.Println(querySturct)
}