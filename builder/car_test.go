package builder_test

import (
	"sdp/builder"
	"strconv"
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

// benchmark performance between normal and builder pattern
func BenchmarkCarWithoutBuilder(b *testing.B) {
	normalCar := builder.Car{
		Engine: "normal engine",
		Wheel:  "normal wheel",
		Color:  "normal color",
	}

	for i := 0; i < b.N; i++ {
		_ = normalCar
	}
}

func BenchmarkCarWithBuilder(b *testing.B) {
	normalCarBuilder, _ := builder.GetCarBuilder(builder.Normal)
	director := builder.NewCarDirector(normalCarBuilder)

	for i := 0; i < b.N; i++ {
		_ = director.BuildCar()
	}
}

// benchmark using many car type
func BenchmarkManyCarWithoutBuilder(b *testing.B) {
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

	for i := 0; i < b.N; i++ {
		_ = normalCar
		_ = speedCar
		_ = offroadCar
	}
}

func BenchmarkManyCarWithBuilder(b *testing.B) {
	// create builder for each car
	normalCarBuilder, _ := builder.GetCarBuilder(builder.Normal)
	speedCarBuilder, _ := builder.GetCarBuilder(builder.Speed)
	offroadCarBuilder, _ := builder.GetCarBuilder(builder.Offroad)

	director := builder.NewCarDirector(normalCarBuilder)

	for i := 0; i < b.N; i++ {
		director.SetBuilder(normalCarBuilder)
		_ = director.BuildCar()

		director.SetBuilder(speedCarBuilder)
		_ = director.BuildCar()

		director.SetBuilder(offroadCarBuilder)
		_ = director.BuildCar()
	}
}

// bencmark with memory alloc
func BenchmarkManyCarWithoutBuilderWithAlloc(b *testing.B) {
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

	alloc := make(map[string]builder.Car)

	for i := 0; i < b.N; i++ {
		alloc[strconv.Itoa(i)+"normal"] = normalCar
		alloc[strconv.Itoa(i)+"speed"] = speedCar
		alloc[strconv.Itoa(i)+"offroad"] = offroadCar
	}
}

func BenchmarkManyCarWithBuilderWithAlloc(b *testing.B) {
	// create builder for each car
	normalCarBuilder, _ := builder.GetCarBuilder(builder.Normal)
	speedCarBuilder, _ := builder.GetCarBuilder(builder.Speed)
	offroadCarBuilder, _ := builder.GetCarBuilder(builder.Offroad)

	director := builder.NewCarDirector(normalCarBuilder)

	alloc := make(map[string]builder.Car)

	for i := 0; i < b.N; i++ {
		director.SetBuilder(normalCarBuilder)
		alloc[strconv.Itoa(i)+"normal"] = director.BuildCar()

		director.SetBuilder(speedCarBuilder)
		alloc[strconv.Itoa(i)+"speed"] = director.BuildCar()

		director.SetBuilder(offroadCarBuilder)
		alloc[strconv.Itoa(i)+"offroad"] = director.BuildCar()
	}
}
