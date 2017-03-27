package ga

type FitnessCalc struct {
	solution []byte
}

// SetSolution - Set a candidate solution as a byte array
func (fc *FitnessCalc) SetSolution(newSolution []byte) {
	fc.solution = newSolution
}

// GetFitness - Calculate individuals fitness by comparing it to our candidate solution
func (fc *FitnessCalc) GetFitness(individual *Individual) int {
	fitness := 0

	for i := 0; i < individual.Size() && i < len(fc.solution); i++ {
		indGene := individual.GetGene(i)

		if indGene == fc.solution[i] {
			fitness++
		}
	}

	return fitness
}

// GetMaxFitness - Get optimum fitness
func (fc *FitnessCalc) GetMaxFitness() int {
	return len(fc.solution)
}
