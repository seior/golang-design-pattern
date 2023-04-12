package builder

import "errors"

type CarType string

// CarType is a type of car
const (
	Normal  CarType = "normal"
	Speed   CarType = "speed"
	Offroad CarType = "offroad"
)

// CarDirector is a director for car builder
type IBuilder interface {
	SetEngine()
	SetWheel()
	SetColor()
	GetCar() Car
}

// GetCarBuilder create car builder based on car type
func GetCarBuilder(carType CarType) (IBuilder, error) {
	if carType == Normal {
		return NewNormalCarBuilder(), nil
	}

	if carType == Speed {
		return NewSpeedCarBuilder(), nil
	}

	if carType == Offroad {
		return NewOffroadCarBuilder(), nil
	}

	return nil, errors.New("invalid car type")
}

type NormalBuilder struct {
	Engine string
	Wheel  string
	Color  string
}

func NewNormalCarBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (builder *NormalBuilder) SetEngine() {
	builder.Engine = "normal engine"
}

func (builder *NormalBuilder) SetWheel() {
	builder.Wheel = "normal wheel"
}

func (builder *NormalBuilder) SetColor() {
	builder.Color = "normal color"
}

func (builder *NormalBuilder) GetCar() Car {
	return Car{
		Engine: builder.Engine,
		Wheel:  builder.Wheel,
		Color:  builder.Color,
	}
}

type SpeedBuilder struct {
	Engine string
	Wheel  string
	Color  string
}

func NewSpeedCarBuilder() *SpeedBuilder {
	return &SpeedBuilder{}
}

func (builder *SpeedBuilder) SetEngine() {
	builder.Engine = "speed engine"
}

func (builder *SpeedBuilder) SetWheel() {
	builder.Wheel = "speed wheel"
}

func (builder *SpeedBuilder) SetColor() {
	builder.Color = "speed color"
}

func (builder *SpeedBuilder) GetCar() Car {
	return Car{
		Engine: builder.Engine,
		Wheel:  builder.Wheel,
		Color:  builder.Color,
	}
}

type OffroadBuilder struct {
	Engine string
	Wheel  string
	Color  string
}

func NewOffroadCarBuilder() *OffroadBuilder {
	return &OffroadBuilder{}
}

func (builder *OffroadBuilder) SetEngine() {
	builder.Engine = "offroad engine"
}

func (builder *OffroadBuilder) SetWheel() {
	builder.Wheel = "offroad wheel"
}

func (builder *OffroadBuilder) SetColor() {
	builder.Color = "offroad color"
}

func (builder *OffroadBuilder) GetCar() Car {
	return Car{
		Engine: builder.Engine,
		Wheel:  builder.Wheel,
		Color:  builder.Color,
	}
}

type CarDirector struct {
	builder IBuilder
}

func NewCarDirector(builder IBuilder) *CarDirector {
	return &CarDirector{
		builder,
	}
}

func (director *CarDirector) SetBuilder(b IBuilder) {
	director.builder = b
}

func (director *CarDirector) BuildCar() Car {
	director.builder.SetEngine()
	director.builder.SetColor()
	director.builder.SetWheel()

	return director.builder.GetCar()
}
