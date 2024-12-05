package algo

var MF = 0

func MaxFlow() {
	totalFlow := 0

	for {
		// Find a path
		parent, found := BFS()
		if !found {
			// If no more augmenting path is found, break the loop
			break
		}

		// Update the flow along the path and add it to the total flow
		pathFlow := UpdateFlow(Farm, parent)
		totalFlow += pathFlow
	}

	MF = totalFlow
	// Output the total flow
}
