package thefarm

import (
	"errors"
	"fmt"
)


func DivideFood(fc FodderCalculator, cows int) (float64, error) {
    amount, err := fc.FodderAmount(cows)
    if err != nil {
        return 0, err
    }

    fat, err := fc.FatteningFactor()
    if err != nil {
        return 0, err
    }

    return (amount * fat) / float64(cows), nil
} 

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
    if cows > 0 {
        return DivideFood(fc, cows)
    }

    return 0, errors.New("invalid number of cows")
} 

type InvalidCowsError struct {
    cows int
    message string
}

func (e *InvalidCowsError) Error() error {
    return fmt.Errorf("%d cows are invalid: %s", e.cows, e.message)
}

func ValidateNumberOfCows(cows int) error {
    if cows > 0 {
        return nil
    }

    invalidCowsError := &InvalidCowsError{cows: cows}
    if cows < 0 {
        invalidCowsError.message = "there are no negative cows"
    } else if cows == 0 {
        invalidCowsError.message = "no cows don't need food"
    } 

    return invalidCowsError.Error()
}

