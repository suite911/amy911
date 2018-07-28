package maths

func ClampFloat32(n, min, max float32) float32 {
	switch {
	case n < min:
		return min
	case n > max:
		return max
	default:
		return n
	}
}

func ClampFloat64(n, min, max float64) float64 {
	switch {
	case n < min:
		return min
	case n > max:
		return max
	default:
		return n
	}
}

func ClampInt(n, min, max int) int {
	switch {
	case n < min:
		return min
	case n > max:
		return max
	default:
		return n
	}
}

func ClampInt8(n, min, max int8) int8 {
	switch {
	case n < min:
		return min
	case n > max:
		return max
	default:
		return n
	}
}

func ClampInt16(n, min, max int16) int16 {
	switch {
	case n < min:
		return min
	case n > max:
		return max
	default:
		return n
	}
}

func ClampInt32(n, min, max int32) int32 {
	switch {
	case n < min:
		return min
	case n > max:
		return max
	default:
		return n
	}
}

func ClampInt64(n, min, max int64) int64 {
	switch {
	case n < min:
		return min
	case n > max:
		return max
	default:
		return n
	}
}

func ClampUint(n, min, max uint) uint {
	switch {
	case n < min:
		return min
	case n > max:
		return max
	default:
		return n
	}
}

func ClampUint8(n, min, max uint8) uint8 {
	switch {
	case n < min:
		return min
	case n > max:
		return max
	default:
		return n
	}
}

func ClampUint16(n, min, max uint16) uint16 {
	switch {
	case n < min:
		return min
	case n > max:
		return max
	default:
		return n
	}
}

func ClampUint32(n, min, max uint32) uint32 {
	switch {
	case n < min:
		return min
	case n > max:
		return max
	default:
		return n
	}
}

func ClampUint64(n, min, max uint64) uint64 {
	switch {
	case n < min:
		return min
	case n > max:
		return max
	default:
		return n
	}
}
