package dao

import (
	"context"
	"garagesvc/model"
	"garagesvc/module/mongodb"

	"go.mongodb.org/mongo-driver/bson"
)

// ServiceCreate ...
func ServiceCreate(service model.ServiceCreateBSON) (err error) {
	var (
		serviceCol = mongodb.ServiceCol()
		ctx        = context.Background()
	)
	_, err = serviceCol.InsertOne(ctx, service)
	return
}

// ServiceFindOne ...
func ServiceFindOne(filter bson.M) (service model.Service, err error) {
	var (
		serviceCol = mongodb.ServiceCol()
		ctx        = context.Background()
	)
	err = serviceCol.FindOne(ctx, filter).Decode(&service)
	return
}

// ServiceFind ...
func ServiceFind(filter []bson.M) (serviceList []model.Service, err error) {
	var (
		serviceCol = mongodb.ServiceCol()
		ctx        = context.Background()
	)
	// Looking for services
	cur, err := serviceCol.Aggregate(ctx, filter)
	if err != nil {
		return
	}

	// Get data from cursor
	defer cur.Close(ctx)
	cur.All(ctx, &serviceList)
	return
}

// ServiceUpdateOne ...
func ServiceUpdateOne(filter bson.M, data bson.M) (err error) {
	var (
		serviceCol = mongodb.ServiceCol()
		ctx        = context.Background()
	)
	_, err = serviceCol.UpdateOne(ctx, filter, data)
	return
}

// ServiceDelete ...
func ServiceDelete(filter bson.M) (err error) {
	var (
		serviceCol = mongodb.ServiceCol()
		ctx        = context.Background()
	)
	_, err = serviceCol.DeleteOne(ctx, filter)
	return
}
