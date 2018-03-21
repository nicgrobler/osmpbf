package osmpbf

// Make tags map from stringtable and two parallel arrays of IDs.
func extractTags(stringTable []string, keyIDs, valueIDs []uint32, validKeys map[string]bool) map[string]string {
	tags := make(map[string]string, len(keyIDs))
	for index, keyID := range keyIDs {
		key := stringTable[keyID]
		if !validKeys[key] {
			//we don't care about this key, so move on
			continue
		}
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
func (tu *tagUnpacker) next(validKeys map[string]bool) map[string]string {
	tags := make(map[string]string)
	for tu.index < len(tu.keysVals) {
		keyID := tu.keysVals[tu.index]
		tu.index++
		if keyID == 0 {
			break
		}

		valID := tu.keysVals[tu.index]
		tu.index++

		key := tu.stringTable[keyID]
		if !validKeys[key] {
			//we don't care about this key, so move on
			continue
		}
		val := tu.stringTable[valID]
		tags[key] = val
	}
	return tags
}
