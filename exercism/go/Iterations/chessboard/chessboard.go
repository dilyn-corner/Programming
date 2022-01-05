package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
    var i int
    for _, value := range cb[rank] {
        if value {
            i++
        }
    }
    return i
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
    var i int

    if file < 1 || file > 8 {
        return 0
    }

    for _, value := range cb {
        if value[file-1] {
            i++
        }
    }
    return i
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
    var i int
    for _, rank := range cb {
        for range rank {
            i++
        }
    }
    return i
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
    var i int
    for rank := range cb {
        i += CountInRank(cb, rank)
    }
    return i
}
