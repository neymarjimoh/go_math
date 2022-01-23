package geometry

import "errors"

func CubeVolume(x float64) (float64, error) {
	if x != 0 {

		return x * x * x, nil
	}
	return 0, errors.New("invalid CubeVolume value: does'nt accept 0")
}
