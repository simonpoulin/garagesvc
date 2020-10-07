package dao

import (
	"context"
	"garagesvc/model"
	"garagesvc/module/mongodb"

	"go.mongodb.org/mongo-driver/bson"
)

// ResourceCreate ...
func ResourceCreate(resource model.Resource) (err error) {
	var (
		resourceCol = mongodb.ResourceCol()
		ctx         = context.Background()
	)
	_, err = resourceCol.InsertOne(ctx, resource)
	return
}

// ResourceFindOne ...
func ResourceFindOne(filter bson.M) (resource model.Resource, err error) {
	var (
		resourceCol = mongodb.ResourceCol()
		ctx         = context.Background()
	)
	err = resourceCol.FindOne(ctx, filter).Decode(&resource)
	return
}

// ResourceDelete ...
func ResourceDelete(filter bson.M) (err error) {
	var (
		resourceCol = mongodb.ResourceCol()
		ctx         = context.Background()
	)
	_, err = resourceCol.DeleteOne(ctx, filter)
	return
}
