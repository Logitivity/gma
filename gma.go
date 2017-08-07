package gma

//MovingAverage is the interface that all other moving average types will conform to.
type MovingAverage interface {
	Add(float64)
	Value() float64
	Full() boolean
	Window()
	WindowSize()
}
