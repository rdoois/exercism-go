package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var count int
    for _, birds := range birdsPerDay {
       count += birds 
    }
    return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    min := (week - 1) * 7
    max := week * 7
    return TotalBirdCount(birdsPerDay[min:max])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i, birds := range birdsPerDay {
        if i == 0 || i % 2 == 0 {
            birdsPerDay[i] = birds + 1
        }
    }
    return birdsPerDay
}
