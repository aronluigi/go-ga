package ga

type Population struct {
	individuals []*Individual
	Fc          *FitnessCalc
}

func (pop *Population) Create(populationSize int, initialise bool) {
	pop.individuals = make([]*Individual, populationSize)

	if initialise {
		for i := 0; i < len(pop.individuals); i++ {
			ind := &Individual{}
			ind = ind.NewIndividual(pop.Fc)

			ind.GenerateIndividual()
			pop.SaveIndividual(i, ind)
		}
	}
}

func (pop *Population) GetFittest() *Individual {
	fittest := pop.individuals[0]

	for i := 0; i < pop.Size(); i++ {
		if fittest.GetFitness() <= pop.GetIndividual(i).GetFitness() {
			fittest = pop.GetIndividual(i)
		}
	}

	return fittest
}

func (pop *Population) GetIndividual(index int) *Individual {
	return pop.individuals[index]
}

func (pop *Population) SaveIndividual(index int, individual *Individual) {
	pop.individuals[index] = individual
}

func (pop *Population) Size() int {
	return len(pop.individuals)
}
