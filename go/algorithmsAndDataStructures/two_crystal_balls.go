package main

import "math"

func two_crystal_balls(breaks []bool) int32 {
	jmpAmnt :=int(math.Sqrt(float64(len(breaks))));
	i := jmpAmnt;
	for ; i < len(breaks); i += jmpAmnt {
		if(breaks[i]) {
			break;
		}
	}

	i -= jmpAmnt;
	for j:=0; j < jmpAmnt && i < len(breaks); {
		if(breaks[i]) {
			return int32(i);
		}
		j++;
		i++;
	}

	return -1;
} 	