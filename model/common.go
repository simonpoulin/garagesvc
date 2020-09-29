package model

import (
	"garagesvc/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AppQuery ...
type AppQuery struct {
	Status     string
	Name       string
	Active     string
	Page       int
	Phone      string
	ServiceID  primitive.ObjectID
	CompanyID  primitive.ObjectID
	CustomerID primitive.ObjectID
}

// GenerateFindQuery ...
func (query AppQuery) GenerateFindQuery() (findQuery []bson.M) {
	var (
		blankObjectID primitive.ObjectID
		filters       []bson.M
	)

	//Set filter parts
	if query.Status != "" {
		filters = append(filters, bson.M{"status": query.Status})
	}

	if query.Phone != "" {
		filters = append(filters, bson.M{"phone": bson.M{"$regex": query.Phone}})
	}

	if query.Name != "" {
		filters = append(filters, bson.M{"searchstring": bson.M{"$regex": util.ConvertToHex(query.Name)}})
	}

	if query.Active != "" && query.Active != "all" {
		filters = append(filters, bson.M{"active": query.Active})
	}

	if query.ServiceID != blankObjectID {
		filters = append(filters, bson.M{"serviceid": query.ServiceID})
	}

	if query.CompanyID != blankObjectID {
		filters = append(filters, bson.M{"companyid": query.CompanyID})
	}

	if query.CustomerID != blankObjectID {
		filters = append(filters, bson.M{"customerid": query.CustomerID})
	}

	//Set filter query from parts
	findQuery = append(findQuery, bson.M{"$match": func() bson.M {
		if filters != nil {
			if len(filters) > 0 {
				return bson.M{"$and": filters}
			}
		}
		return bson.M{}
	}()})

	return
}
