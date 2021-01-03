package security

import (
	"log"

	"github.com/sethvargo/go-password/password"
)

// MinuteTTL is a minute in EPOCH seconds
const MinuteTTL uint64 = 60

// HourTTL is an hour in EPOCH seconds
const HourTTL uint64 = 3600

// DayTTL is a day in EPOCH seconds
const DayTTL uint64 = 86400

// WeekTTL is a day in EPOCH seconds
const WeekTTL uint64 = 604800

// MonthTTL is a month of 30.44 days in EPOCH seconds
const MonthTTL uint64 = 2629743

// YearTTL is a YEAR of 365.24 days in EPOCH seconds
const YearTTL uint64 = 31556926

//GenPassword128 creates a password that is 128 characters long with 32 digits,
// 32 symbols, allowing upper and lower case letters, disallowing repeat characters.
func GenPassword128() (string, error) {
	res1, err := password.Generate(64, 10, 22, false, false)
	if err != nil {
		log.Fatal(err)
	}
	res2, err := password.Generate(64, 10, 22, false, false)
	if err != nil {
		log.Fatal(err)
	}
	res := res1 + res2

	return res, err
}

//GenPassword256 creates a password that is 256 characters long with 64 digits,
// 64 symbols, allowing upper and lower case letters, disallowing repeat characters.
func GenPassword256() (string, error) {
	res1, err := GenPassword128()
	if err != nil {
		log.Fatal(err)
	}
	res2, err := GenPassword128()
	if err != nil {
		log.Fatal(err)
	}
	res := res1 + res2

	return res, err
}
