package ga

type Individual struct {
	GeneLength int
	fitness    int
	genes      []byte
	fc         *FitnessCalc
}

func (ind *Individual) GenerateIndividual() {
	for i := 0; i < ind.GeneLength; i++ {
		rand := RandomFloat()
		gene := byte(RoundFloat(rand))
		ind.genes[i] = byte(gene)
	}

	// fmt.Println(ind.genes)
}

func (ind *Individual) GetGenes() []byte {
	return ind.genes
}

func (ind *Individual) GetGene(index int) byte {
	return ind.genes[index]
}

func (ind *Individual) SetGene(index int, val byte) {
	ind.genes[index] = val
	ind.fitness = 0
}

func (ind *Individual) Size() int {
	return len(ind.genes)
}

func (ind *Individual) GetFitness() int {
	if ind.fitness == 0 {
		ind.fitness = ind.fc.GetFitness(ind)
	}

	return ind.fitness
}

func (ind *Individual) NewIndividual(fc *FitnessCalc) *Individual {
	ind.GeneLength = 64
	ind.genes = make([]byte, ind.GeneLength)
	ind.fitness = 0
	ind.fc = fc

	return ind
}
