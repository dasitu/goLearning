package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func getAnswer() chromosome {
	return "Nit7R6hviiyKHEuj"
}

type chromosome string

type individual struct {
	chrom chromosome
	score float64
}

type population []individual

func (a population) Len() int {
	return len(a)
}

func (a population) Less(i, j int) bool {
	return a[i].score < a[j].score
}

func (a population) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func createChromosome(n int) chromosome {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return chromosome(string(b))
}

func createPopulation(popSize, chromSize int) population {
	p := make([]individual, popSize)
	for i := range p {
		c := createChromosome(chromSize)
		p[i] = individual{c, getScore(c)}
	}
	return p
}

func getScore(c chromosome) float64 {
	answer := getAnswer()
	cCount := 0.0
	for i, v := range c {
		if byte(v) == answer[i] {
			cCount++
		}
	}
	return cCount / float64(len(answer))
}

func getMeanScore(p population) float64 {
	total := 0.0
	for _, i := range p {
		total += i.score
	}
	return total / float64(len(p))
}

func selection(p population) population {
	greadedRetainPercent := 0.3    // percentage of retained best fitting individuals
	NonGreadedRetainPercent := 0.2 // percentage of retained remaining individuals (randomly selected)
	totalCount := len(p)
	bestCount := int(float64(totalCount) * greadedRetainPercent)
	randCount := int(float64(totalCount) * NonGreadedRetainPercent)
	selected := make([]individual, 0)
	sort.Sort(sort.Reverse(p))
	selected = append(selected, p[:bestCount]...)
	for i := 0; i < randCount; i++ {
		min := bestCount
		max := totalCount - 1
		randIndex := rand.Intn(max-min+1) + min
		selected = append(selected, p[randIndex])
	}
	return population(selected)
}

func crossover(c1, c2 chromosome) chromosome {
	l := len(c1) / 2
	return c1[:l] + c2[l:]
}

func mutation(c chromosome) chromosome {
	newLetter := letters[rand.Intn(len(letters))]
	randIndex := rand.Intn(len(c))
	return c[:randIndex] + chromosome(string(newLetter)) + c[randIndex+1:]
}

func generation(p population) population {
	var newGeneration population
	selected := selection(p)
	newGeneration = append(newGeneration, selected...)
	for len(newGeneration) < len(p) {
		twoRandInt := rand.Perm(len(selected))
		p1 := selected[twoRandInt[0]]
		p2 := selected[twoRandInt[1]]
		c := mutation(crossover(p1.chrom, p2.chrom))
		child := individual{c, getScore(c)}
		newGeneration = append(newGeneration, child)
	}
	return newGeneration
}

func main() {
	start := time.Now()

	rand.Seed(time.Now().UnixNano())
	chromSize := len(getAnswer())
	popSize := 10
	p := createPopulation(popSize, chromSize)
	answer := chromosome("")
	c := 0
	for answer == "" {
		p = generation(p)
		fmt.Printf("[%v]mean score:%v\n", c, getMeanScore(p))
		for _, i := range p {
			if i.chrom == getAnswer() {
				answer = i.chrom
			}
		}
		c++
	}
	fmt.Println(answer)

	elapsed := time.Since(start)
	fmt.Printf("took %s", elapsed)
}
