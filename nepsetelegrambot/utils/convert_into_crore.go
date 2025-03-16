package utils

import (
	"fmt"
	"math"
)

func NumberToCroreArab(num float64) string {
	// Define constants for crore and arab
	const (
		lakh  = 1e5 // 100 thousand
		crore = 1e7 // 10 million
		arab  = 1e9 // 1 billion
	)

	// Handle negative numbers
	if num < 0 {
		return "-" + NumberToCroreArab(-num)
	}

	// Cap the number at 1000 arab (10^12)
	if num >= 1000*arab {
		return "1000 arab"
	}

	// Calculate arabs and crores
	arabCount := math.Floor(num / arab)
	remainingAfterArab := num - (arabCount * arab)
	croreCount := math.Floor(remainingAfterArab / crore)

	// Build the result string
	var result string
	if arabCount > 0 {
		// For numbers in arab range, show "X arab, Y crore+"
		result = fmt.Sprintf("%.0f arab", arabCount)
		if croreCount > 0 {
			result += fmt.Sprintf(", %.0f crore+", croreCount)
		}
	} else if croreCount > 0 {
		// For numbers in crore range, show "X crore+"
		result = fmt.Sprintf("%.0f crore+", croreCount)
	} else if num >= lakh {
		// For numbers in lakh range, show "X lakh+"
		lakhCount := math.Floor(num / lakh)
		result = fmt.Sprintf("%.0f lakh+", lakhCount)
	} else {
		// For smaller numbers, show as is
		result = fmt.Sprintf("%.0f", num)
	}

	if result == "" {
		return "0"
	}
	return result
}

func NumberToCroreArabFull(num float64) string {
	// Define constants for crore and arab
	const (
		crore = 1e7 // 10 million
		arab  = 1e9 // 1 billion
	)

	// Handle negative numbers
	if num < 0 {
		return "-" + NumberToCroreArabFull(-num)
	}

	// Cap the number at 1000 arab (10^12)
	if num >= 1000*arab {
		return "1000 arab"
	}

	// Cap at 100 arab if between 100 arab and 1000 arab
	if num >= 100*arab {
		return "100 arab"
	}

	// Calculate arabs, crores, and remaining lakhs
	arabCount := math.Floor(num / arab)
	remainingAfterArab := num - (arabCount * arab)
	croreCount := math.Floor(remainingAfterArab / crore)
	remaining := remainingAfterArab - (croreCount * crore)

	// Build the result string
	var result string
	if arabCount > 0 {
		result += fmt.Sprintf("%.0f arab", arabCount)
	}
	if croreCount > 0 {
		if result != "" {
			result += " "
		}
		result += fmt.Sprintf("%.0f crore", croreCount)
	}
	if remaining > 0 && num < crore {
		// For numbers less than a crore, show as is
		result = fmt.Sprintf("%.2f", num)
	} else if remaining > 0 {
		// For small remainders after crore, show in lakhs (1 lakh = 10^5)
		lakhCount := remaining / 1e5
		if result != "" {
			result += " "
		}
		result += fmt.Sprintf("%.2f lakh", lakhCount)
	}

	if result == "" {
		return "0"
	}
	return result
}
