package flavor

import (
	"api-gym/util"
	"fmt"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
)

func Generate(flavor, extra string) string {
	val := ""
	if flavor == "address" {
		a := gofakeit.Address()
		val = fmt.Sprintf("%s, %s, %s %s %s", a.Street, a.City, a.State, a.Zip, a.Country)
	} else if flavor == "latitude" {
		val = fmt.Sprintf("%f", gofakeit.Latitude())
	} else if flavor == "longitude" {
		val = fmt.Sprintf("%f", gofakeit.Longitude())
	} else if flavor == "timestamp" {
		val = "2022-04-18T06:52:29.940Z"
	} else if flavor == "few_words" {
		val = gofakeit.Word() + " " + gofakeit.Word()
	} else if flavor == "float" {
		val = fmt.Sprintf("%d.%d", rand.Intn(30), rand.Intn(10))
	} else if flavor == "int" {
		val = fmt.Sprintf("%d", rand.Intn(65000))
	} else if flavor == "paragraph" {
		val = gofakeit.LoremIpsumParagraph(1, 3, 33, ".")
	} else if flavor == "first_name" {
		val = gofakeit.FirstName()
	} else if flavor == "last_name" {
		val = gofakeit.LastName()
	} else if flavor == "pronouns" {
		val = "she/her"
		if rand.Intn(2) == 0 {
			val = "he/him"
		}
		if rand.Intn(2) == 0 {
			val = "they/them"
		}
	} else if flavor == "email" {
		val = gofakeit.Email()
	} else if flavor == "phone" {
		val = gofakeit.PhoneFormatted()
	} else if flavor == "uuid" {
		val = util.PseudoUuid()
	}
	return val
}
