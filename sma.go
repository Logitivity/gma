package gma

type SimpleMovingAverage struct {
	window     bool
	windowSize int
	values     []float64
	full       bool
}

func NewSimpleMovingAverage(windowSize int) SimpleMovingAverage {

	if windowSize > 0 {
		return &SimpleMovingAverage{
			window:     true,
			windowSize: windowSize,
			values:     make([]float64, windowSize),
			full:       false,
		}
	} else {
		return &SimpleMovingAverage{
			window:     false,
			windowSize: 0,
			values:     make([]float64, 0),
			full:       false,
		}
	}
}

func (sma *SimpleMovingAverage) Add(newValue float64) {

	if window {
		//If window is full, remove the first element, and add the new element to the end
		if sma.full {
			_, sma.values = sma.values[0], sma.values[1:]
			sma.values = append(sma.values, newValue)

			//If the window is not full, add the new value to the end of the array.
		} else {
			sma.values = append(sma.values, newValue)

			//If the value array equals the window size, mark this as full.
			if len(sma.values) == windowSize {
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

	for _, value = range values {
		sum = sum + value
	}

	value := sum / float64(len(values))

	return value
}

func (sma *SimpleMovingAverage) Full() bool {
	return full
}

func (sma *SimpleMovingAverage) Window() bool {
	return window
}

func (sma *SimpleMovingAverage) WindowSize() int {
	return windowSize
}
