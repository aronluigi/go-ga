package main

import (
	"fmt"
	GA "go-ga/ga"
)

func main() {
	fitnessCalc := &GA.FitnessCalc{}
	population := &GA.Population{}
	algorith := &GA.Algorithm{}
	algorith = algorith.Init()

	generationCount := 0
	fitnessSolution := []byte{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}

	fitnessCalc.SetSolution(fitnessSolution)

	algorith.Fc = fitnessCalc
	population.Fc = fitnessCalc

	population.Create(50, true)

	for population.GetFittest().GetFitness() < fitnessCalc.GetMaxFitness() {
		generationCount++

		s := fmt.Sprintf("Generation: %d Fittest: %d", generationCount, population.GetFittest().GetFitness())
		fmt.Println(s)

		population = algorith.EvolvePopulation(population)
	}

	fmt.Println("\nSolution found!")
	fmt.Printf("Generation: %d\n", generationCount)
	fmt.Printf("Genes: %v", population.GetFittest())
}
