package basicdatatypes

import "fmt"

// TestTypeConversion - type conversion is explicit
func TestTypeConversion() {
	// restrict bool values to be of type bool only
	var x int = 10
	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = x + int(y)
	fmt.Println(x, y, z, d)
}

// ConvertFeetIntoMeters - convert feet to meters
func ConvertFeetIntoMeters(feet float64) float64 {
	const oneFeet float64 = 0.3048
	return oneFeet * feet
}

// CalculateTimeToReachMars - calculate time
func CalculateTimeToReachMars() {
	const lightSpeed = 299792 // km/s
	var distance = 56000000   // km
	fmt.Println(distance, " km will take ", distance/lightSpeed, " seconds")

	distance = 401000000
	fmt.Println(distance, " km will take ", distance/lightSpeed, " seconds")

	const hoursPerDay, minutesPerHour = 24, 60
	var distanceMars, speed = 96300000, 100800
	fmt.Println(distance, " km will take ", distanceMars/speed, " hours and ", distanceMars/speed/hoursPerDay, " days with speed ", speed)

	//Write a program to determine how fast a ship would need to travel (in km/h) in order to reach Malacandra in 28 days. Assume a distance of 56,000,000 km.
	s := float64(56_000_000)
	t := float64(28 * hoursPerDay)
	v := s / t
	fmt.Println(v)
	fmt.Println(s, " km will need spped = ", v, " km/hr to reach in 28 days")
}
