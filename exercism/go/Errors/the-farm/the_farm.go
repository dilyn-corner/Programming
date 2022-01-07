package thefarm

import (
    "errors"
    "fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
    cowCount int
}

func (shenanigans *SillyNephewError) Error() string {
    return fmt.Sprintf("silly nephew, there cannot be %d cows", shenanigans.cowCount)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
    foodWeight, err := weightFodder.FodderAmount()

    if err != nil && err != ErrScaleMalfunction {
        return 0.0, err
    }

    if foodWeight < 0 {
        return 0.0, errors.New("negative fodder")
    }
    if cows == 0 {
        return 0.0, errors.New("division by zero")
    }
    if cows < 0 {
        return 0.0, &SillyNephewError{cowCount: cows}
    }

    if err == ErrScaleMalfunction && foodWeight > 0 {
        return foodWeight / float64(cows) * 2.0, nil
    }

    return foodWeight / float64(cows), err
}
