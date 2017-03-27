package ga

import (
	"math/rand"
	"time"
)

type Algorithm struct {
	uniformRate    float64
	mutationRate   float64
	tournamentSize int
	elitism        bool
	Fc             *FitnessCalc
}

func (al *Algorithm) EvolvePopulation(pop *Population) *Population {
	var elitismOffset int
	newPopulation := &Population{Fc: al.Fc}
	newPopulation.Create(pop.Size(), false)

	if al.elitism {
		newPopulation.SaveIndividual(0, pop.GetFittest())
	}

	if al.elitism {
		elitismOffset = 1
	} else {
		elitismOffset = 0
	}

	for i := elitismOffset; i < pop.Size(); i++ {
		indv1 := al.tournamentSelection(pop)
		indv2 := al.tournamentSelection(pop)
		newIndv := al.crossover(indv1, indv2)

		newPopulation.SaveIndividual(i, newIndv)
	}

	for i := elitismOffset; i < newPopulation.Size(); i++ {
		al.mutate(newPopulation.GetIndividual(i))
	}

	return newPopulation
}

func (al *Algorithm) mutate(indv *Individual) {
	for i := 0; i < indv.Size(); i++ {
		if RandomFloat() <= al.mutationRate {
			gene := byte(RoundFloat(RandomFloat()))
			indv.SetGene(i, gene)
		}
	}
}

func (al *Algorithm) crossover(indv1 *Individual, indv2 *Individual) *Individual {
	newSo1 := new(Individual).NewIndividual(al.Fc)
	newSo1.GenerateIndividual()

	for i := 0; i < indv1.Size(); i++ {
		if RandomFloat() <= al.uniformRate {
			newSo1.SetGene(i, indv1.GetGene(i))
		} else {
			newSo1.SetGene(i, indv2.GetGene(i))
		}
	}

	return newSo1
}

func (al *Algorithm) tournamentSelection(pop *Population) *Individual {
	tournament := &Population{Fc: al.Fc}
	tournament.Create(al.tournamentSize, false)

	for i := 0; i < al.tournamentSize; i++ {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)

		randomID := r.Intn(pop.Size())
		tournament.SaveIndividual(i, pop.GetIndividual(randomID))
	}

	return tournament.GetFittest()
}

func (al *Algorithm) Init() *Algorithm {
	al.uniformRate = 0.5
	al.mutationRate = 0.015
	al.tournamentSize = 5
	al.elitism = true

	return al
}
