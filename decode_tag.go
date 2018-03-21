package osmpbf

// Make tags map from stringtable and two parallel arrays of IDs.
func extractTags(stringTable []string, keyIDs, valueIDs []uint32) map[string]string {
	var tags map[string]string
	//only init a map if we have some data.
	if len(keyIDs) > 0 && len(stringTable) > 0 {
		//length, but is it empty - i.e. initialized with a size
		tags = make(map[string]string, len(keyIDs))
	}
	for index, keyID := range keyIDs {
		key := stringTable[keyID]
		val := stringTable[valueIDs[index]]
		tags[key] = val
	}
	return tags
}

type tagUnpacker struct {
	stringTable []string
	keysVals    []int32
	index       int
}

// Make tags map from stringtable and array of IDs (used in DenseNodes encoding).
func (tu *tagUnpacker) next() map[string]string {
	var tags map[string]string
	//only init a map if we have some data.
	if len(tu.keysVals) > 0 && len(tu.stringTable) > 0 {
		tags = make(map[string]string)
	}
	for tu.index < len(tu.keysVals) {
		keyID := tu.keysVals[tu.index]
		tu.index++
		if keyID == 0 {
			break
		}

		valID := tu.keysVals[tu.index]
		tu.index++

		key := tu.stringTable[keyID]
		val := tu.stringTable[valID]
		tags[key] = val
	}
	return tags
}
