package gma

import (
	"reflect"
	"testing"
)

func TestNewSimpleMovingAverageWithWindow(t *testing.T) {
	sma := NewSimpleMovingAverage(10)

	//Test to see if the window size matches what was passed as the parameter.
	if sma.windowSize != 10 {
		t.Errorf("Window Size expected to be 10. Actual: %d", sma.windowSize)
	}

	//Test to see if the window is defined as true, since we passed a positive integer.
	if sma.window != true {
		t.Errorf("Window expected to be true, Actual: %b", sma.window)
	}

	//Test to see if the simple moving average is full.
	if sma.full != false {
		t.Errorf("Full expected to be false, Actual: %b", sma.full)
	}
}

func TestNewSimpleMovingAverageWithoutWindow(t *testing.T) {
	sma := NewSimpleMovingAverage(0)
	smaN := NewSimpleMovingAverage(-10)

	//Window size should be 0 on a non-window average.
	if sma.windowSize != 0 {
		t.Errorf("Window Size expected to be 0. Actual: %d", sma.windowSize)
	}

	//Window size should be 0 even when parameter is negative (-10)
	if smaN.windowSize != 0 {
		t.Errorf("Window Size expected to be 0. Actual: %d", smaN.windowSize)
	}

	//Window should be false when passing 0 to windowSize.
	if sma.window != false {
		t.Errorf("Window expected to be false, Actual: %b", sma.window)
	}

	//Window should be false when passing negative integer to windowSize.
	if smaN.window != false {
		t.Errorf("Window expected to be false, Actual: %b", smaN.window)
	}

	//Average should not be full in windowless average.
	if sma.full != false {
		t.Errorf("Full expected to be false, Actual: %b", sma.full)
	}

	//Average should not be full in windowless average.
	if smaN.full != false {
		t.Errorf("Full expected to be false, Actual: %b", smaN.full)
	}
}

func TestSimpleMovingAverageAddFunction(t *testing.T) {

	smaWindow := NewSimpleMovingAverage(10)
	smaWindowless := NewSimpleMovingAverage(0)

	addValues := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	expectedSMAWindowValues := []float64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	expectedSMAWindowlessValues := addValues

	for _, value := range addValues {
		smaWindow.Add(value)
		smaWindowless.Add(value)
	}

	//smaWindow should be full.
	if smaWindow.full != true {
		t.Errorf("Window expected to be true, Actual: %b", smaWindow.full)
	}

	//smaWindow should NOT be full.
	if smaWindowless.full != false {
		t.Errorf("Window expected to be false, Actual: %b", smaWindowless.full)
	}

	//Compare smaWindow values to expected.
	if !reflect.DeepEqual(smaWindow.values, expectedSMAWindowValues) {
		t.Errorf("smaWindow values expected to match SMAWindowValues expected: %v", smaWindow.values)
	}

	//Compare smaWindowless values to expected.
	if !reflect.DeepEqual(smaWindowless.values, expectedSMAWindowlessValues) {
		t.Errorf("smaWindowless values expected to match SMAWindowlessValues")
	}
}

func TestSimpleMovingAverageValueFunction(t *testing.T) {

	windowSize := 5

	smaWindow := NewSimpleMovingAverage(windowSize)
	smaWindowless := NewSimpleMovingAverage(0)

	addValues := []float64{1, 2, 3, 4, 5, 6}

	total := float64(0)

	for _, value := range addValues {
		total += value
	}

	//Subtract 1 since the window only uses the last 5 values. (Not pretty, but does the job)
	expectedWindowValue := (total - 1) / float64(windowSize)

	expectedWindowlessValue := total / float64(len(addValues))

	//Add the values to the averages now.

	for _, value := range addValues {
		smaWindow.Add(value)
		smaWindowless.Add(value)
	}

	if smaWindow.Value() != expectedWindowValue {
		t.Errorf("Expected Window Value: %f, Actual: %f", expectedWindowValue, smaWindow.Value())
	}

	if smaWindowless.Value() != expectedWindowlessValue {
		t.Errorf("Expected Windowless Value: %f, Actual: %f", expectedWindowlessValue, smaWindowless.Value())
	}

}

func TestSimpleMovingAverageGetterFunctions(t *testing.T) {

	windowSize := 10

	smaWindow := NewSimpleMovingAverage(windowSize)
	smaWindowless := NewSimpleMovingAverage(0)

	addValues := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	for _, value := range addValues {
		smaWindow.Add(value)
		smaWindowless.Add(value)
	}

	//smaWindow should be full.
	if smaWindow.Full() != true {
		t.Errorf("Full expected to be true, Actual: %b", smaWindow.Full())
	}

	//Window should be true.
	if smaWindow.Window() != true {
		t.Errorf("Window expected to be true, Actual: %b", smaWindow.Window())
	}

	//Windowsize should match set window size.
	if smaWindow.WindowSize() != windowSize {
		t.Errorf("Window size expected to be %d, Actual: %d", windowSize, smaWindow.WindowSize())
	}

	//smaWindowless should be not full.
	if smaWindowless.Full() != false {
		t.Errorf("Full expected to be false, Actual: %b", smaWindowless.Full())
	}

	//Window should be false.
	if smaWindowless.Window() != false {
		t.Errorf("Window expected to be false, Actual: %b", smaWindowless.Window())
	}

	//Window size should be 0.
	if smaWindowless.WindowSize() != 0 {
		t.Errorf("Window size expected to be 0, Actual: %d", smaWindowless.WindowSize())
	}

}
