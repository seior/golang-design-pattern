package builder_test

import (
	"sdp/builder"
	"testing"
)

func TestCarWithoutBuilder(t *testing.T) {
	normalCar := builder.Car{
		Engine: "normal engine",
		Wheel:  "normal wheel",
		Color:  "normal color",
	}

	speedCar := builder.Car{
		Engine: "speed engine",
		Wheel:  "speed wheel",
		Color:  "speed color",
	}

	offroadCar := builder.Car{
		Engine: "offroad engine",
		Wheel:  "offroad wheel",
		Color:  "offroad color",
	}

	t.Log("=== Normal car: ===")
	printCar(t, normalCar)

	t.Log("=== Speed car: ===")
	printCar(t, speedCar)

	t.Log("=== Offroad car: ===")
	printCar(t, offroadCar)
}

func TestCarWithBuilder(t *testing.T) {
	// create builder for each car
	normalCarBuilder, _ := builder.GetCarBuilder(builder.Normal)
	speedCarBuilder, _ := builder.GetCarBuilder(builder.Speed)
	offroadCarBuilder, _ := builder.GetCarBuilder(builder.Offroad)

	director := builder.NewCarDirector(normalCarBuilder)
	normalCar := director.BuildCar()

	director.SetBuilder(speedCarBuilder)
	speedCar := director.BuildCar()

	director.SetBuilder(offroadCarBuilder)
	offroadCar := director.BuildCar()

	t.Log("=== Normal car: ===")
	printCar(t, normalCar)

	t.Log("=== Speed car: ===")
	printCar(t, speedCar)

	t.Log("=== Offroad car: ===")
	printCar(t, offroadCar)
}

func printCar(t *testing.T, car builder.Car) {
	t.Log("Engine: ", car.Engine)
	t.Log("Wheel: ", car.Wheel)
	t.Log("Color: ", car.Color)
}
