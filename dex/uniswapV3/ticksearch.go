package uniswapV3

import "math"

//Maybe not needed anymore?
func position(tick int) (int, uint8) {
	return int((tick >> 8) & 32767), uint8(tick % 256)
}

func (p *Pair) nextInitializedTickWithinOneWord(tick int, zeroForOne bool) (int, bool) {

	compressed := math.Floor(float64(tick) / float64(p.TickSpacing)) // matches rounding in the code

	if zeroForOne {
		wordPos := int(compressed) >> 8
		minimum := (wordPos << 8) * p.TickSpacing

		if tick < p.TickList[0].Index { // isBelowSmallest
			return minimum, false
		}

		initializedTickIndex := p.getNextInitializedTickIndex(tick, zeroForOne)
		nextInitializedTick := math.Max(float64(minimum), float64(initializedTickIndex))
		return int(nextInitializedTick), int(nextInitializedTick) == initializedTickIndex
	}

	wordPos := int(compressed+1) >> 8
	maximum := ((wordPos+1)<<8)*p.TickSpacing - 1

	if tick >= p.TickList[len(p.TickList)-1].Index { // isAtOrAboveLargest
		return maximum, false
	}

	initializedTickIndex := p.getNextInitializedTickIndex(tick, zeroForOne)
	nextInitializedTick := math.Min(float64(maximum), float64(initializedTickIndex))
	return int(nextInitializedTick), int(nextInitializedTick) == initializedTickIndex
}

//func (p *Pair) nextInitializedTickWithinOneWordOld(tick int, zeroForOne bool) (int, bool) {
//	//WARNING: THIS MAYBE STUPID AND LOSE PRECISION
//	//compressed := int64(math.Floor(float64(tick) / float64(p.TickSpacing))) // matches rounding in the code
//	compressed := tick / p.TickSpacing                                      //int(math.Floor(float64(tick) / float64(p.tickSpacing))) // matches rounding in the code
//
//	if tick < 0 && tick%p.TickSpacing != 0 {
//		compressed--
//	}
//
//	if zeroForOne {
//		wordPos, bitPos := position(compressed)
//		mask := new(uint256.Int).Lsh(OneUint256, uint(bitPos))
//		mask.Sub(mask, OneUint256)
//		mask.Add(mask, new(uint256.Int).Lsh(OneUint256, uint(bitPos)))
//		v := p.Tickbitmap[int16(wordPos)]
//		masked := new(uint256.Int).And(&v, mask)
//		initialized := !masked.IsZero()
//		var next int
//		if initialized {
//			next = (compressed - 16777215&int(uint64(bitPos)-mostSignificantBit(masked))) * p.TickSpacing
//		} else {
//			next = (compressed - 16777215&int(uint64(bitPos))) * p.TickSpacing
//		}
//		return next, initialized
//	} else {
//		wordPos, bitPos := position(compressed + 1)
//		mask := new(uint256.Int).Lsh(OneUint256, uint(bitPos))
//		mask.Sub(mask, OneUint256)
//		mask.Not(mask)
//		v := p.Tickbitmap[int16(wordPos)]
//		masked := new(uint256.Int).And(&v, mask)
//		initialized := !masked.IsZero()
//		var next int
//		if initialized {
//			next = (compressed + 1 + 16777215&int(leastSignificantBit(masked)-uint64(bitPos))) * p.TickSpacing
//		} else {
//			next = (compressed + 16777215&int(255-bitPos)) * p.TickSpacing
//		}
//		return next, initialized
//	}
//
//}
//

func (p *Pair) getNextInitializedTickIndex(tick int, zeroForOne bool) int {

	if zeroForOne {
		if tick >= p.TickList[len(p.TickList)-1].Index { // isAtOrAboveLargest
			return p.TickList[len(p.TickList)-1].Index
		}
		index := p.binarySearch(tick)
		return p.TickList[index].Index
	}

	if tick < p.TickList[0].Index { // isBelowSmallest
		return p.TickList[0].Index
	}

	index := p.binarySearch(tick)
	return p.TickList[index+1].Index
}

func (p *Pair) binarySearch(tick int) int {

	start := 0
	end := len(p.TickList) - 1
	for start <= end {
		mid := (start + end) / 2
		if p.TickList[mid].Index == tick {
			return mid
		} else if p.TickList[mid].Index < tick {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	// tick not found - we return closest tick
	if p.TickList[start].Index < tick {
		return start
	}
	return start - 1
}
