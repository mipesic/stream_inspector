package domain

type WatchPointsRepo interface {
	func WatchPointsForChannel(ch Channel) []WatchPoint, error
}