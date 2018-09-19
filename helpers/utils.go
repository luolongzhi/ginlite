package helpers

import (
    "math/rand"
    "time"
)

func InitRand() {
    rand.Seed(time.Now().Unix())
}

func RandIntRange(min, max int) int {
    if min >= max || min == 0 || max == 0 {
        return max
    }

    return rand.Intn(max - min) + min
}
