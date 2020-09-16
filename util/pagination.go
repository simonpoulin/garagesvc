package util

import (
	"errors"
	"math"
	"reflect"
)

// PagedList ...
type PagedList struct {
	Data        []interface{} `json:"data"`
	CurrentPage int           `json:"currentpage"`
	TotalPage   int           `json:"totalpage"`
	Limit       int           `json:"limit"`
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

	// Calculate total pages
	totalPage := int(math.Ceil(float64(len(unpagedList)) / float64(limit)))
	if pagedList.TotalPage > page {
		err = errors.New("out of pages")
		return
	}

	// Calculate page objects
	if page == totalPage {
		pageObjects = len(unpagedList) % limit
		if pageObjects == 0 {
			pageObjects = limit
		}
	} else {
		pageObjects = limit
	}

	// Set list data
	firstObject := (page - 1) * limit
	lastObject := firstObject + pageObjects
	for i := firstObject; i < lastObject; i++ {
		data = append(data, unpagedList[i])
	}

	// Set remaining fields
	pagedList.Data = data
	pagedList.TotalPage = totalPage
	pagedList.CurrentPage = page
	pagedList.Limit = limit
	return
}

func interfaceSlice(slice interface{}) (itf []interface{}, err error) {

	//Check valid slice
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		err = errors.New("not a slice")
	}

	//Convert slice into interface slice
	itf = make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		itf[i] = s.Index(i).Interface()
	}

	return
}
