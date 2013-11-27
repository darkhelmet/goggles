package plugins

/*
#include <stdlib.h>
*/
import "C"

func getloadavg() (float64, float64, float64) {
    avg := []C.double{0, 0, 0}

    C.getloadavg(&avg[0], C.int(len(avg)))

    one := float64(avg[0])
    five := float64(avg[1])
    fifteen := float64(avg[2])

    return one, five, fifteen
}
