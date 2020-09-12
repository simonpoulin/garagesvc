package util

import (
	"errors"
	"math"
	"reflect"
)

// PagedList ...
type PagedList struct {
	Data        []interface{}
	CurrentPage int
	TotalPage   int
}

// Paging ...
func Paging(list interface{}, page int, limit int) (pagedList PagedList, err error) {

	//Convert interface to interface slice
	unpagedList, err := interfaceSlice(list)
	if err != nil {
		return
	}

	//Calculate total pages
	totalPage := int(math.Ceil(float64(len(unpagedList)) / float64(limit)))
	if pagedList.TotalPage > page {
		err = errors.New("out of pages")
		return
	}

	// Set list data
	for i := (page - 1) * limit; i < page*limit; i++ {
		pagedList.Data = append(pagedList.Data, unpagedList[i])
	}

	// Set remaining fields
	pagedList.TotalPage = totalPage
	pagedList.CurrentPage = page
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
