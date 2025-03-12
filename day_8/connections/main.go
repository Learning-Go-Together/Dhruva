package connections

func countConnections(groupSize int) int {
	connection := 0
	for i := 1; i < groupSize; i++ {
		connection = connection + i
	}
	return connection
}
