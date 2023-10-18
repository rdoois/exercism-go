package chessboard

type File [8]bool

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    files, exists := cb[file]
    if !exists {
        return 0
    }
    var total int
    for _, f := range files {
        if f {
            total++
        }
    }

    return total
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    if rank > 8 || rank < 1 {
        return 0
    }
    var total int
    for file := range cb {
        if cb[file][rank - 1] {
            total++ 
        }
    }
    return total
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    var total int
    for _, file := range cb {
        for range file {
            total++
        }
    }

    return total
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    var total int
    for _, file := range cb {
        for _, f := range file {
            if f {

                total++
            } 
        }
    }

    return total
}
