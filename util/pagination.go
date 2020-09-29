package util

import (
	"errors"
	"math"
	"reflect"
)

// PagedList ...
type PagedList struct {
	List  []interface{} `json:"list"`
	Page  int           `json:"page"`
	Total int           `json:"total"`
	Limit int           `json:"limit"`
}

// Paging ...
func Paging(list interface{}, page int, limit int) (pagedList PagedList, err error) {
	var (
		data        []interface{}
		pageObjects int
	)

	// Convert interface to interface slice
	unpagedList, err := interfaceSlice(list)
	if err != nil {
		return
	}

	// Check empty list
	listSize := len(unpagedList)
	if listSize == 0 {
		err = errors.New("empty list")
		return
	}

	// Calculate total pages
	totalPage := int(math.Ceil(float64(listSize) / float64(limit)))
	if totalPage < page+1 {
		err = errors.New("out of pages")
		return
	}

	// Calculate page objects
	if totalPage == page+1 {
		pageObjects = listSize % limit
		if pageObjects == 0 {
			pageObjects = limit
		}
	} else {
		pageObjects = limit
	}

	// Set list data
	firstObject := page * limit
	lastObject := firstObject + pageObjects
	for i := firstObject; i < lastObject; i++ {
		data = append(data, unpagedList[i])
	}

	// Set remaining fields
	pagedList.List = data
	pagedList.Total = listSize
	pagedList.Page = page
	pagedList.Limit = limit
	return
}

func interfaceSlice(slice interface{}) (itf []interface{}, err error) {

	//Check valid slice
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		err = errors.New("not a slice")
		return
	}

	//Convert slice into interface slice
	itf = make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		itf[i] = s.Index(i).Interface()
	}

	return
}
