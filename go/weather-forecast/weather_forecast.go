// Package weather provides functionality to return the current weather conditions in current city in string form.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current location of the user.
var CurrentLocation string

// Forecast returns a string value containing the weather conditions in the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
