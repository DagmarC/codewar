package arrays


func findPoisonedDuration(timeSeries []int, duration int) int {
    var total int
	
	if len(timeSeries) == 0 {
		return total
	}
	for i := 0; i < len(timeSeries)-1; i++ {
		if timeSeries[i+1] <= timeSeries[i]+duration {
			total += timeSeries[i+1] - timeSeries[i]
		} else {
			total += duration
		}
	}
	return total + duration
}