package flavor

import "api-gym/util"

func DataType(flavor, extra string) string {
	val := ""
	if flavor == "address" {
		val = "string"
	} else if flavor == "latitude" {
		val = "string"
	} else if flavor == "longitude" {
		val = "string"
	} else if flavor == "timestamp" {
		val = "string"
	} else if flavor == "few_words" {
		val = "string"
	} else if flavor == "float" {
		val = "float64"
	} else if flavor == "int" {
		val = "int"
	} else if flavor == "paragraph" {
		val = "string"
	} else if flavor == "first_name" {
		val = "string"
	} else if flavor == "last_name" {
		val = "string"
	} else if flavor == "pronouns" {
		val = "string"
	} else if flavor == "email" {
		val = "string"
	} else if flavor == "phone" {
		val = "string"
	} else if flavor == "uuid" {
		val = "string"
	} else if flavor == "bool" {
		val = "bool"
	} else if flavor == "model" {
		val = util.ToCamelCase(extra)
	} else if flavor == "[]model" {
		val = "[]" + util.ToCamelCase(extra)
	}
	return val
}
