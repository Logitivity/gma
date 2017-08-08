package gma

//SimpleMovingAverage is a struct that contains the moving average values and variables.
type SimpleMovingAverage struct {
	window     bool
	windowSize int
	values     []float64
	full       bool
}

//NewSimpleMovingAverage is the constructor for creating a new moving average. It takes
//a windowSize for it's only parameter. If windowSize is not above 0, then the SMA will be
//constructed with no window.
func NewSimpleMovingAverage(windowSize int) SimpleMovingAverage {

	//If window is a positive integer, create the struct with that size.
	if windowSize > 0 {
		return SimpleMovingAverage{
			window:     true,
			windowSize: windowSize,
			values:     make([]float64, 0),
			full:       false,
		}
		//If window size is 0 or undefined, create a windowless struct.
	} else {
		return SimpleMovingAverage{
			window:     false,
			windowSize: 0,
			values:     make([]float64, 0),
			full:       false,
		}
	}
}

//Add will add the defined value to the moving average.
func (sma *SimpleMovingAverage) Add(newValue float64) {

	if sma.window {
		//If window is full, remove the first element, and add the new element to the end
		if sma.full {
			_, sma.values = sma.values[0], sma.values[1:]
			sma.values = append(sma.values, newValue)

			//If the window is not full, add the new value to the end of the array.
		} else {
			sma.values = append(sma.values, newValue)

			//If the value array equals the window size, mark this as full.
			if len(sma.values) == sma.windowSize {
				sma.full = true
			}
		}
		//If we don't have a window, just append to the end of the array
	} else {
		sma.values = append(sma.values, newValue)
	}
}

//Value will return the average of the moving simple average - if the window is not full, it still returns a value.
func (sma *SimpleMovingAverage) Value() float64 {

	var sum = float64(0)

	for _, value := range sma.values {
		sum += value
	}

	value := sum / float64(len(sma.values))

	return value
}

//Full returns if the amount of values in the array matches the window.
//It is recommended to check full before using the value, as it may be an unexpected value
//If the array is not full yet.
func (sma *SimpleMovingAverage) Full() bool {
	return sma.full
}

//Window returns whether or not this moving average has a window.
func (sma *SimpleMovingAverage) Window() bool {
	return sma.window
}

//WindowSize returns the size of the window (or 0 if no window exists)
func (sma *SimpleMovingAverage) WindowSize() int {
	return sma.windowSize
}
