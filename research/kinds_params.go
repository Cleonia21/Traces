package research

/*
type param struct {
	identPairsAmount_M  int
	identPairsMinDist_M int
	identNumsAmount_C   int
	identNumsMinDist_C  int

	identTrios_M          bool
	identPairs_R          bool
	threeAscendingOrder_R bool
	identNums_P           bool
}
*/

func fingersParamGet(columnNum int) fingers {
	m := map[int]param{
		3:  param{},
		4:  param{},
		5:  param{identPairsAmount_M: 1, identPairsMinDist_M: 3, identNumsAmount_C: 1, identNumsMinDist_C: 3},
		6:  param{identPairsAmount_M: 1, identPairsMinDist_M: 3, identNumsAmount_C: 1, identNumsMinDist_C: 3},
		7:  param{identPairsAmount_M: 1, identPairsMinDist_M: 2, identNumsAmount_C: 1, identNumsMinDist_C: 3},
		8:  param{identPairsAmount_M: 1, identPairsMinDist_M: 2, identNumsAmount_C: 1, identNumsMinDist_C: 3},
		9:  param{identPairsAmount_M: 2, identPairsMinDist_M: 2, identNumsAmount_C: 1, identNumsMinDist_C: 2},
		10: param{identPairsAmount_M: 2, identPairsMinDist_M: 1, identNumsAmount_C: 1, identNumsMinDist_C: 2},
	}
	return fingers{m[columnNum]}
}

func wordsParamGet(columnNum int) words {
	m := map[int]param{
		5:  param{identPairsAmount_M: 1, identPairsMinDist_M: 2},
		6:  param{},
		7:  param{},
		8:  param{},
		9:  param{},
		10: param{},
	}
	return words{m[columnNum]}
}
