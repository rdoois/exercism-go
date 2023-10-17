package lasagna

func PreparationTime(layers []string, minutes int) (total int) {
    if minutes == 0 {
        minutes = 2
    }

    for range layers {
        total += minutes
    }

    return 
}

func Quantities(layers []string) (noodles int, sauce float64) {
    for _, layer := range layers {
        if layer == "noodles" {
            noodles += 50
        } else if layer == "sauce" {
            sauce += 0.2
        }
    }
    return
}

func AddSecretIngredient(friendsList, myList []string) {
    myIndex := len(myList) - 1
    friendIndex := len(friendsList) - 1
    myList[myIndex] = friendsList[friendIndex]
}

func ScaleRecipe(quantities []float64, portions int) (newQty []float64) {
    for _, qty := range quantities {
        newQty = append(newQty, qty * (float64(portions) / 2))
    }
    return

}

