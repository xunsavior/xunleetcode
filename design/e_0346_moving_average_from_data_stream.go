package design

/*
Company: Google(6), Microsoft(4), AppDynamics(3), Facebook(2); Amazon(5)

Given a stream of integers and a window size, calculate the moving average of all integers in the sliding window.

Example:
MovingAverage m = new MovingAverage(3);
m.next(1) = 1
m.next(10) = (1 + 10) / 2
m.next(3) = (1 + 10 + 3) / 3
m.next(5) = (10 + 3 + 5) / 3
*/

// MovingAverage ...
type MovingAverage struct {
	Size  int
	Sum   int
	Cache []int
}

// Constructor346 ...
func Constructor346(size int) MovingAverage {
	return MovingAverage{
		Cache: make([]int, 0, size),
		Size:  size,
	}
}

// Next346 ...
func (this *MovingAverage) Next346(val int) float64 {
	this.Sum += val
	this.Cache = append(this.Cache, val)
	if len(this.Cache) <= this.Size {
		return float64(this.Sum) / float64(len(this.Cache))
	}
	this.Sum -= this.Cache[0]
	this.Cache = this.Cache[1:]
	return float64(this.Sum) / float64(len(this.Cache))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
