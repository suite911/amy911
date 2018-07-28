package bitfield

type Uint uint

const Zero Uint = 0
const Full Uint = ^Zero

func (bf Uint) All(mask uint) bool {
	return (uint(bf) & mask) == mask
}

func (bf Uint) Any(mask uint) bool {
	return (uint(bf) & mask) != 0
}

func (bf *Uint) Clr(mask uint) *Uint {
	*bf &= Uint(^mask)
	return bf
}

func (bf Uint) ClrCopy(mask uint) Uint {
	bf.Clr(mask)
	return bf
}

func (bf *Uint) Copy() Uint {
	return *bf
}

func (bf *Uint) Set(mask uint) *Uint {
	*bf |= Uint(mask)
	return bf
}

func (bf Uint) SetCopy(mask uint) Uint {
	bf.Set(mask)
	return bf
}

func (bf *Uint) Xor(mask uint) *Uint {
	*bf ^= Uint(mask)
	return bf
}

func (bf Uint) XorCopy(mask uint) Uint {
	bf.Xor(mask)
	return bf
}

type Uint8 uint8

const Zero8 Uint8 = 0
const Full8 Uint8 = ^Zero8

func (bf Uint8) All(mask uint8) bool {
	return (uint8(bf) & mask) == mask
}

func (bf Uint8) Any(mask uint8) bool {
	return (uint8(bf) & mask) != 0
}

func (bf *Uint8) Clr(mask uint8) *Uint8 {
	*bf &= Uint8(^mask)
	return bf
}

func (bf Uint8) ClrCopy(mask uint8) Uint8 {
	bf.Clr(mask)
	return bf
}

func (bf *Uint8) Copy() Uint8 {
	return *bf
}

func (bf *Uint8) Set(mask uint8) *Uint8 {
	*bf |= Uint8(mask)
	return bf
}

func (bf Uint8) SetCopy(mask uint8) Uint8 {
	bf.Set(mask)
	return bf
}

func (bf *Uint8) Xor(mask uint8) *Uint8 {
	*bf ^= Uint8(mask)
	return bf
}

func (bf Uint8) XorCopy(mask uint8) Uint8 {
	bf.Xor(mask)
	return bf
}

type Uint16 uint16

const Zero16 Uint16 = 0
const Full16 Uint16 = ^Zero16

func (bf Uint16) All(mask uint16) bool {
	return (uint16(bf) & mask) == mask
}

func (bf Uint16) Any(mask uint16) bool {
	return (uint16(bf) & mask) != 0
}

func (bf *Uint16) Clr(mask uint16) *Uint16 {
	*bf &= Uint16(^mask)
	return bf
}

func (bf Uint16) ClrCopy(mask uint16) Uint16 {
	bf.Clr(mask)
	return bf
}

func (bf *Uint16) Copy() Uint16 {
	return *bf
}

func (bf *Uint16) Set(mask uint16) *Uint16 {
	*bf |= Uint16(mask)
	return bf
}

func (bf Uint16) SetCopy(mask uint16) Uint16 {
	bf.Set(mask)
	return bf
}

func (bf *Uint16) Xor(mask uint16) *Uint16 {
	*bf ^= Uint16(mask)
	return bf
}

func (bf Uint16) XorCopy(mask uint16) Uint16 {
	bf.Xor(mask)
	return bf
}

type Uint32 uint32

const Zero32 Uint32 = 0
const Full32 Uint32 = ^Zero32

func (bf Uint32) All(mask uint32) bool {
	return (uint32(bf) & mask) == mask
}

func (bf Uint32) Any(mask uint32) bool {
	return (uint32(bf) & mask) != 0
}

func (bf *Uint32) Clr(mask uint32) *Uint32 {
	*bf &= Uint32(^mask)
	return bf
}

func (bf Uint32) ClrCopy(mask uint32) Uint32 {
	bf.Clr(mask)
	return bf
}

func (bf *Uint32) Copy() Uint32 {
	return *bf
}

func (bf *Uint32) Set(mask uint32) *Uint32 {
	*bf |= Uint32(mask)
	return bf
}

func (bf Uint32) SetCopy(mask uint32) Uint32 {
	bf.Set(mask)
	return bf
}

func (bf *Uint32) Xor(mask uint32) *Uint32 {
	*bf ^= Uint32(mask)
	return bf
}

func (bf Uint32) XorCopy(mask uint32) Uint32 {
	bf.Xor(mask)
	return bf
}

type Uint64 uint64

const Zero64 Uint64 = 0
const Full64 Uint64 = ^Zero64

func (bf Uint64) All(mask uint64) bool {
	return (uint64(bf) & mask) == mask
}

func (bf Uint64) Any(mask uint64) bool {
	return (uint64(bf) & mask) != 0
}

func (bf *Uint64) Clr(mask uint64) *Uint64 {
	*bf &= Uint64(^mask)
	return bf
}

func (bf Uint64) ClrCopy(mask uint64) Uint64 {
	bf.Clr(mask)
	return bf
}

func (bf *Uint64) Copy() Uint64 {
	return *bf
}

func (bf *Uint64) Set(mask uint64) *Uint64 {
	*bf |= Uint64(mask)
	return bf
}

func (bf Uint64) SetCopy(mask uint64) Uint64 {
	bf.Set(mask)
	return bf
}

func (bf *Uint64) Xor(mask uint64) *Uint64 {
	*bf ^= Uint64(mask)
	return bf
}

func (bf Uint64) XorCopy(mask uint64) Uint64 {
	bf.Xor(mask)
	return bf
}
