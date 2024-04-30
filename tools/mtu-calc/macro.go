package main

func macroSKBWithOverhead(x int, v vars) int {
	return x - macroSKBDataAlign(v.SizeOfSKBSharedInfo(), v)
}

func macroSKBDataAlign(x int, v vars) int {
	return macroAlign(x, v.SMPCacheBytes())
}

func macroAlign(x int, align int) int {
	return (x + align - 1) & ^(align - 1)
}
