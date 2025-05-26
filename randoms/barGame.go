package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Function to check if there are duplicate birthdays in a slice
func hasDuplicateBirthDays(birthdays []int) bool {
	seen := make(map[int]bool)
	for _, day := range birthdays {
		if seen[day] {
			return true
		}
		seen[day] = true
	}
	return false
}

// Run a single trial for a given number of people
func runTrial(numPeople int) bool {
	// Generate random birthdays for numPeople individuals
	birthdays := make([]int, numPeople)
	for i := range birthdays {
		// Generate a random birthday between 1 and 365
		birthdays[i] = rand.Intn(365) + 1 // 1-based indexing for dates
	}

	// Check for duplicate birthdays
	return hasDuplicateBirthDays(birthdays)
}

// Calculate probability for a specific group size
func calculateProbability(numPeople, numTrials int) float64 {
    successes := 0
    
    // Run specified number of trials
    for i := 0; i < numTrials; i++ {
        if runTrial(numPeople) {
            successes++
        }
    }
    
    // Return probability as decimal
    return float64(successes) / float64(numTrials)
}

func main() {
    const (
        minPeople = 10
        maxPeople = 100
        increment = 10
        numTrials = 10000
    )
    
    // Seed random number generator
    rand.Seed(time.Now().UnixNano())
    
    fmt.Printf("%-20s %s\n", "People in room", "Probability")
    fmt.Println("----------------------------------------")
    
    // Calculate and display probabilities for each group size
    for people := minPeople; people <= maxPeople; people += increment {
        prob := calculateProbability(people, numTrials)
        fmt.Printf("%-20d %.3f\n", people, prob)
        
        // Mark the point where probability exceeds 50%
        if people == 23 {
            fmt.Println("(*) Around 23 people, probability crosses 50%")
        }
    }
}
