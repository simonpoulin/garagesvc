package dao

import (
	"context"
	"garagesvc/model"
	"garagesvc/module/mongodb"

	"go.mongodb.org/mongo-driver/bson"
)

// BookingCreate ...
func BookingCreate(booking model.BookingCreateBSON) (err error) {
	var (
		bookingCol = mongodb.BookingCol()
		ctx        = context.Background()
	)
	_, err = bookingCol.InsertOne(ctx, booking)
	return
}

// BookingFindOne ...
func BookingFindOne(filter bson.M) (booking model.Booking, err error) {
	var (
		bookingCol = mongodb.BookingCol()
		ctx        = context.Background()
	)
	err = bookingCol.FindOne(ctx, filter).Decode(&booking)
	return
}

// BookingFind ...
func BookingFind(filter []bson.M) (bookingList []model.Booking, err error) {
	var (
		bookingCol = mongodb.BookingCol()
		ctx        = context.Background()
	)
	// Looking for bookings
	cur, err := bookingCol.Aggregate(ctx, filter)
	if err != nil {
		return
	}

	// Get data from cursor
	defer cur.Close(ctx)
	cur.All(ctx, &bookingList)
	return
}

// BookingUpdateOne ...
func BookingUpdateOne(filter bson.M, data bson.M) (err error) {
	var (
		bookingCol = mongodb.BookingCol()
		ctx        = context.Background()
	)
	_, err = bookingCol.UpdateOne(ctx, filter, data)
	return
}

// BookingDelete ...
func BookingDelete(filter bson.M) (err error) {
	var (
		bookingCol = mongodb.BookingCol()
		ctx        = context.Background()
	)
	_, err = bookingCol.DeleteOne(ctx, filter)
	return
}
