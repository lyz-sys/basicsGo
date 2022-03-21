package main

import (
	"fmt"
	"reflect"
	"strings"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)

	//createQuery(90)
}

func createQuery(q interface{}) {
	v := reflect.ValueOf(q)
	if reflect.Struct != v.Kind() {
		fmt.Printf("%s is error type,not is %s", v.Kind(), reflect.Struct)
		return
	}

	var values []string
	table := v.Type().Name()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		field := fieldInfo.Name
		value := v.FieldByName(field)
		if value.Kind() == reflect.Int {
			values = append(values, fmt.Sprint(value))
		} else if value.Kind() == reflect.String {
			values = append(values, fmt.Sprintf("'%s'", fmt.Sprint(value)))
		}
	}
	sql := fmt.Sprintf("insert into %s values(%s)", table, strings.Join(values, string(',')))
	fmt.Println(sql)
}
