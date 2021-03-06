package cpu6502

// Operation models a CPU instruction
type Operation struct {
	Opcode      byte
	Length      uint8 // in bytes
	Cycles      int
	Documented  bool
	Addressing  func(*CPU) uint16
	Instruction func(*CPU, uint16)
}

// Only BRK (00) operation has no cycles here, because they're accounted for in the IRQ interrupt handler
var Opcodes = [0x100] Operation {
	{ Opcode: 0x00, Length: 1, Cycles: 0, Documented: true , Instruction: (*CPU).brk,   Addressing: (*CPU).implied},
	{ Opcode: 0x01, Length: 2, Cycles: 6, Documented: true , Instruction: (*CPU).ora,   Addressing: (*CPU).indexedIndirectX},
	{ Opcode: 0x02, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x03, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x04, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x05, Length: 2, Cycles: 3, Documented: true , Instruction: (*CPU).ora,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0x06, Length: 2, Cycles: 5, Documented: true , Instruction: (*CPU).asl,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0x07, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x08, Length: 1, Cycles: 3, Documented: true , Instruction: (*CPU).php,   Addressing: (*CPU).implied},
	{ Opcode: 0x09, Length: 2, Cycles: 2, Documented: true , Instruction: (*CPU).ora,   Addressing: (*CPU).immediate},
	{ Opcode: 0x0A, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).asla,  Addressing: (*CPU).implied},
	{ Opcode: 0x0b, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x0c, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x0D, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).ora,   Addressing: (*CPU).absolute},
	{ Opcode: 0x0E, Length: 3, Cycles: 6, Documented: true , Instruction: (*CPU).asl,   Addressing: (*CPU).absolute},
	{ Opcode: 0x0f, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x10, Length: 2, Cycles: 2, Documented: true , Instruction: (*CPU).bpl,   Addressing: (*CPU).immediate},
	{ Opcode: 0x11, Length: 2, Cycles: 5, Documented: true , Instruction: (*CPU).ora,   Addressing: (*CPU).indirectIndexedY},
	{ Opcode: 0x12, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x13, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x14, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x15, Length: 2, Cycles: 4, Documented: true , Instruction: (*CPU).ora,   Addressing: (*CPU).zeroPageX},
	{ Opcode: 0x16, Length: 2, Cycles: 6, Documented: true , Instruction: (*CPU).asl,   Addressing: (*CPU).zeroPageX},
	{ Opcode: 0x17, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x18, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).clc,   Addressing: (*CPU).implied},
	{ Opcode: 0x19, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).ora,   Addressing: (*CPU).absoluteY},
	{ Opcode: 0x1a, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x1b, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x1c, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x1D, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).ora,   Addressing: (*CPU).absoluteX},
	{ Opcode: 0x1E, Length: 3, Cycles: 7, Documented: true , Instruction: (*CPU).asl,   Addressing: (*CPU).absoluteX},
	{ Opcode: 0x1f, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x20, Length: 3, Cycles: 6, Documented: true , Instruction: (*CPU).jsr,   Addressing: (*CPU).absolute},
	{ Opcode: 0x21, Length: 2, Cycles: 6, Documented: true , Instruction: (*CPU).and,   Addressing: (*CPU).indexedIndirectX},
	{ Opcode: 0x22, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x23, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x24, Length: 2, Cycles: 3, Documented: true , Instruction: (*CPU).bit,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0x25, Length: 2, Cycles: 3, Documented: true , Instruction: (*CPU).and,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0x26, Length: 2, Cycles: 5, Documented: true , Instruction: (*CPU).rol,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0x27, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x28, Length: 1, Cycles: 4, Documented: true , Instruction: (*CPU).plp,   Addressing: (*CPU).implied},
	{ Opcode: 0x29, Length: 2, Cycles: 2, Documented: true , Instruction: (*CPU).and,   Addressing: (*CPU).immediate},
	{ Opcode: 0x2A, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).rola,  Addressing: (*CPU).implied},
	{ Opcode: 0x2b, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x2C, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).bit,   Addressing: (*CPU).absolute},
	{ Opcode: 0x2D, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).and,   Addressing: (*CPU).absolute},
	{ Opcode: 0x2E, Length: 3, Cycles: 6, Documented: true , Instruction: (*CPU).rol,   Addressing: (*CPU).absolute},
	{ Opcode: 0x2f, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x30, Length: 2, Cycles: 2, Documented: true , Instruction: (*CPU).bmi,   Addressing: (*CPU).immediate},
	{ Opcode: 0x31, Length: 2, Cycles: 5, Documented: true , Instruction: (*CPU).and,   Addressing: (*CPU).indirectIndexedY},
	{ Opcode: 0x32, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x33, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x34, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x35, Length: 2, Cycles: 4, Documented: true , Instruction: (*CPU).and,   Addressing: (*CPU).zeroPageX},
	{ Opcode: 0x36, Length: 2, Cycles: 6, Documented: true , Instruction: (*CPU).rol,   Addressing: (*CPU).zeroPageX},
	{ Opcode: 0x37, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x38, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).sec,   Addressing: (*CPU).implied},
	{ Opcode: 0x39, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).and,   Addressing: (*CPU).absoluteY},
	{ Opcode: 0x3a, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x3b, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x3c, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x3D, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).and,   Addressing: (*CPU).absoluteX},
	{ Opcode: 0x3E, Length: 3, Cycles: 7, Documented: true , Instruction: (*CPU).rol,   Addressing: (*CPU).absoluteX},
	{ Opcode: 0x3f, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x40, Length: 1, Cycles: 6, Documented: true , Instruction: (*CPU).rti,   Addressing: (*CPU).implied},
	{ Opcode: 0x41, Length: 2, Cycles: 6, Documented: true , Instruction: (*CPU).eor,   Addressing: (*CPU).indexedIndirectX},
	{ Opcode: 0x42, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x43, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x44, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x45, Length: 2, Cycles: 3, Documented: true , Instruction: (*CPU).eor,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0x46, Length: 2, Cycles: 5, Documented: true , Instruction: (*CPU).lsr,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0x47, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x48, Length: 1, Cycles: 3, Documented: true , Instruction: (*CPU).pha,   Addressing: (*CPU).implied},
	{ Opcode: 0x49, Length: 2, Cycles: 2, Documented: true , Instruction: (*CPU).eor,   Addressing: (*CPU).immediate},
	{ Opcode: 0x4A, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).lsra,  Addressing: (*CPU).implied},
	{ Opcode: 0x4b, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x4C, Length: 3, Cycles: 3, Documented: true , Instruction: (*CPU).jmp,   Addressing: (*CPU).absolute},
	{ Opcode: 0x4D, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).eor,   Addressing: (*CPU).absolute},
	{ Opcode: 0x4E, Length: 3, Cycles: 6, Documented: true , Instruction: (*CPU).lsr,   Addressing: (*CPU).absolute},
	{ Opcode: 0x4f, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x50, Length: 2, Cycles: 2, Documented: true , Instruction: (*CPU).bvc,   Addressing: (*CPU).immediate},
	{ Opcode: 0x51, Length: 2, Cycles: 5, Documented: true , Instruction: (*CPU).eor,   Addressing: (*CPU).indirectIndexedY},
	{ Opcode: 0x52, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x53, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x54, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x55, Length: 2, Cycles: 4, Documented: true , Instruction: (*CPU).eor,   Addressing: (*CPU).zeroPageX},
	{ Opcode: 0x56, Length: 2, Cycles: 6, Documented: true , Instruction: (*CPU).lsr,   Addressing: (*CPU).zeroPageX},
	{ Opcode: 0x57, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x58, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).cli,   Addressing: (*CPU).implied},
	{ Opcode: 0x59, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).eor,   Addressing: (*CPU).absoluteY},
	{ Opcode: 0x5a, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x5b, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x5c, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x5D, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).eor,   Addressing: (*CPU).absoluteX},
	{ Opcode: 0x5E, Length: 3, Cycles: 7, Documented: true , Instruction: (*CPU).lsr,   Addressing: (*CPU).absoluteX},
	{ Opcode: 0x5f, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x60, Length: 1, Cycles: 6, Documented: true , Instruction: (*CPU).rts,   Addressing: (*CPU).implied},
	{ Opcode: 0x61, Length: 2, Cycles: 6, Documented: true , Instruction: (*CPU).adc,   Addressing: (*CPU).indexedIndirectX},
	{ Opcode: 0x62, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x63, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x64, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x65, Length: 2, Cycles: 3, Documented: true , Instruction: (*CPU).adc,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0x66, Length: 2, Cycles: 5, Documented: true , Instruction: (*CPU).ror,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0x67, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x68, Length: 1, Cycles: 4, Documented: true , Instruction: (*CPU).pla,   Addressing: (*CPU).implied},
	{ Opcode: 0x69, Length: 2, Cycles: 2, Documented: true , Instruction: (*CPU).adc,   Addressing: (*CPU).immediate},
	{ Opcode: 0x6A, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).rora,  Addressing: (*CPU).implied},
	{ Opcode: 0x6b, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x6C, Length: 3, Cycles: 5, Documented: true , Instruction: (*CPU).jmp,   Addressing: (*CPU).indirect},
	{ Opcode: 0x6D, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).adc,   Addressing: (*CPU).absolute},
	{ Opcode: 0x6E, Length: 3, Cycles: 6, Documented: true , Instruction: (*CPU).ror,   Addressing: (*CPU).absolute},
	{ Opcode: 0x6f, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x70, Length: 2, Cycles: 2, Documented: true , Instruction: (*CPU).bvs,   Addressing: (*CPU).immediate},
	{ Opcode: 0x71, Length: 2, Cycles: 5, Documented: true , Instruction: (*CPU).adc,   Addressing: (*CPU).indirectIndexedY},
	{ Opcode: 0x72, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x73, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x74, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x75, Length: 2, Cycles: 4, Documented: true , Instruction: (*CPU).adc,   Addressing: (*CPU).zeroPageX},
	{ Opcode: 0x76, Length: 2, Cycles: 6, Documented: true , Instruction: (*CPU).ror,   Addressing: (*CPU).zeroPageX},
	{ Opcode: 0x77, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x78, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).sei,   Addressing: (*CPU).implied},
	{ Opcode: 0x79, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).adc,   Addressing: (*CPU).absoluteY},
	{ Opcode: 0x7a, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x7b, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x7c, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x7D, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).adc,   Addressing: (*CPU).absoluteX},
	{ Opcode: 0x7E, Length: 3, Cycles: 7, Documented: true , Instruction: (*CPU).ror,   Addressing: (*CPU).absoluteX},
	{ Opcode: 0x7f, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x80, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x81, Length: 2, Cycles: 6, Documented: true , Instruction: (*CPU).sta,   Addressing: (*CPU).indexedIndirectX},
	{ Opcode: 0x82, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x83, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x84, Length: 2, Cycles: 3, Documented: true , Instruction: (*CPU).sty,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0x85, Length: 2, Cycles: 3, Documented: true , Instruction: (*CPU).sta,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0x86, Length: 2, Cycles: 3, Documented: true , Instruction: (*CPU).stx,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0x87, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x88, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).dey,   Addressing: (*CPU).implied},
	{ Opcode: 0x89, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x8A, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).txa,   Addressing: (*CPU).implied},
	{ Opcode: 0x8b, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x8C, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).sty,   Addressing: (*CPU).absolute},
	{ Opcode: 0x8D, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).sta,   Addressing: (*CPU).absolute},
	{ Opcode: 0x8E, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).stx,   Addressing: (*CPU).absolute},
	{ Opcode: 0x8f, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x90, Length: 2, Cycles: 2, Documented: true , Instruction: (*CPU).bcc,   Addressing: (*CPU).immediate},
	{ Opcode: 0x91, Length: 2, Cycles: 6, Documented: true , Instruction: (*CPU).sta,   Addressing: (*CPU).indirectIndexedY},
	{ Opcode: 0x92, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x93, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x94, Length: 2, Cycles: 4, Documented: true , Instruction: (*CPU).sty,   Addressing: (*CPU).zeroPageX},
	{ Opcode: 0x95, Length: 2, Cycles: 4, Documented: true , Instruction: (*CPU).sta,   Addressing: (*CPU).zeroPageX},
	{ Opcode: 0x96, Length: 2, Cycles: 4, Documented: true , Instruction: (*CPU).stx,   Addressing: (*CPU).zeroPageY},
	{ Opcode: 0x97, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x98, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).tya,   Addressing: (*CPU).implied},
	{ Opcode: 0x99, Length: 3, Cycles: 5, Documented: true , Instruction: (*CPU).sta,   Addressing: (*CPU).absoluteY},
	{ Opcode: 0x9A, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).txs,   Addressing: (*CPU).implied},
	{ Opcode: 0x9b, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x9c, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x9D, Length: 3, Cycles: 5, Documented: true , Instruction: (*CPU).sta,   Addressing: (*CPU).absoluteX},
	{ Opcode: 0x9e, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0x9f, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xA0, Length: 2, Cycles: 2, Documented: true , Instruction: (*CPU).ldy,   Addressing: (*CPU).immediate},
	{ Opcode: 0xA1, Length: 2, Cycles: 6, Documented: true , Instruction: (*CPU).lda,   Addressing: (*CPU).indexedIndirectX},
	{ Opcode: 0xA2, Length: 2, Cycles: 2, Documented: true , Instruction: (*CPU).ldx,   Addressing: (*CPU).immediate},
	{ Opcode: 0xa3, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xA4, Length: 2, Cycles: 3, Documented: true , Instruction: (*CPU).ldy,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0xA5, Length: 2, Cycles: 3, Documented: true , Instruction: (*CPU).lda,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0xA6, Length: 2, Cycles: 3, Documented: true , Instruction: (*CPU).ldx,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0xa7, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xA8, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).tay,   Addressing: (*CPU).implied},
	{ Opcode: 0xA9, Length: 2, Cycles: 2, Documented: true , Instruction: (*CPU).lda,   Addressing: (*CPU).immediate},
	{ Opcode: 0xAA, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).tax,   Addressing: (*CPU).implied},
	{ Opcode: 0xab, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xAC, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).ldy,   Addressing: (*CPU).absolute},
	{ Opcode: 0xAD, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).lda,   Addressing: (*CPU).absolute},
	{ Opcode: 0xAE, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).ldx,   Addressing: (*CPU).absolute},
	{ Opcode: 0xaf, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xB0, Length: 2, Cycles: 2, Documented: true , Instruction: (*CPU).bcs,   Addressing: (*CPU).immediate},
	{ Opcode: 0xB1, Length: 2, Cycles: 5, Documented: true , Instruction: (*CPU).lda,   Addressing: (*CPU).indirectIndexedY},
	{ Opcode: 0xb2, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xb3, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xB4, Length: 2, Cycles: 4, Documented: true , Instruction: (*CPU).ldy,   Addressing: (*CPU).zeroPageX},
	{ Opcode: 0xB5, Length: 2, Cycles: 4, Documented: true , Instruction: (*CPU).lda,   Addressing: (*CPU).zeroPageX},
	{ Opcode: 0xB6, Length: 2, Cycles: 4, Documented: true , Instruction: (*CPU).ldx,   Addressing: (*CPU).zeroPageY},
	{ Opcode: 0xb7, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xB8, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).clv,   Addressing: (*CPU).implied},
	{ Opcode: 0xB9, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).lda,   Addressing: (*CPU).absoluteY},
	{ Opcode: 0xBA, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).tsx,   Addressing: (*CPU).implied},
	{ Opcode: 0xbb, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xBC, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).ldy,   Addressing: (*CPU).absoluteX},
	{ Opcode: 0xBD, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).lda,   Addressing: (*CPU).absoluteX},
	{ Opcode: 0xBE, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).ldx,   Addressing: (*CPU).absoluteY},
	{ Opcode: 0xbf, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xC0, Length: 2, Cycles: 2, Documented: true , Instruction: (*CPU).cpy,   Addressing: (*CPU).immediate},
	{ Opcode: 0xC1, Length: 2, Cycles: 6, Documented: true , Instruction: (*CPU).cmp,   Addressing: (*CPU).indexedIndirectX},
	{ Opcode: 0xc2, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xc3, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xC4, Length: 2, Cycles: 3, Documented: true , Instruction: (*CPU).cpy,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0xC5, Length: 2, Cycles: 3, Documented: true , Instruction: (*CPU).cmp,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0xC6, Length: 2, Cycles: 5, Documented: true , Instruction: (*CPU).dec,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0xc7, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xC8, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).iny,   Addressing: (*CPU).implied},
	{ Opcode: 0xC9, Length: 2, Cycles: 2, Documented: true , Instruction: (*CPU).cmp,   Addressing: (*CPU).immediate},
	{ Opcode: 0xCA, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).dex,   Addressing: (*CPU).implied},
	{ Opcode: 0xcb, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xCC, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).cpy,   Addressing: (*CPU).absolute},
	{ Opcode: 0xCD, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).cmp,   Addressing: (*CPU).absolute},
	{ Opcode: 0xCE, Length: 3, Cycles: 6, Documented: true , Instruction: (*CPU).dec,   Addressing: (*CPU).absolute},
	{ Opcode: 0xcf, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xD0, Length: 2, Cycles: 2, Documented: true , Instruction: (*CPU).bne,   Addressing: (*CPU).immediate},
	{ Opcode: 0xD1, Length: 2, Cycles: 5, Documented: true , Instruction: (*CPU).cmp,   Addressing: (*CPU).indirectIndexedY},
	{ Opcode: 0xd2, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xd3, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xd4, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xD5, Length: 2, Cycles: 4, Documented: true , Instruction: (*CPU).cmp,   Addressing: (*CPU).zeroPageX},
	{ Opcode: 0xD6, Length: 2, Cycles: 6, Documented: true , Instruction: (*CPU).dec,   Addressing: (*CPU).zeroPageX},
	{ Opcode: 0xd7, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xD8, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).cld,   Addressing: (*CPU).implied},
	{ Opcode: 0xD9, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).cmp,   Addressing: (*CPU).absoluteY},
	{ Opcode: 0xda, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xdb, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xdc, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xDD, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).cmp,   Addressing: (*CPU).absoluteX},
	{ Opcode: 0xDE, Length: 3, Cycles: 7, Documented: true , Instruction: (*CPU).dec,   Addressing: (*CPU).absoluteX},
	{ Opcode: 0xdf, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xE0, Length: 2, Cycles: 2, Documented: true , Instruction: (*CPU).cpx,   Addressing: (*CPU).immediate},
	{ Opcode: 0xE1, Length: 2, Cycles: 6, Documented: true , Instruction: (*CPU).sbc,   Addressing: (*CPU).indexedIndirectX},
	{ Opcode: 0xe2, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xe3, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xE4, Length: 2, Cycles: 3, Documented: true , Instruction: (*CPU).cpx,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0xE5, Length: 2, Cycles: 3, Documented: true , Instruction: (*CPU).sbc,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0xE6, Length: 2, Cycles: 5, Documented: true , Instruction: (*CPU).inc,   Addressing: (*CPU).zeroPage},
	{ Opcode: 0xe7, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xE8, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).inx,   Addressing: (*CPU).implied},
	{ Opcode: 0xE9, Length: 2, Cycles: 2, Documented: true , Instruction: (*CPU).sbc,   Addressing: (*CPU).immediate},
	{ Opcode: 0xEA, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).nop,   Addressing: (*CPU).implied},
	{ Opcode: 0xeb, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xEC, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).cpx,   Addressing: (*CPU).absolute},
	{ Opcode: 0xED, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).sbc,   Addressing: (*CPU).absolute},
	{ Opcode: 0xEE, Length: 3, Cycles: 6, Documented: true , Instruction: (*CPU).inc,   Addressing: (*CPU).absolute},
	{ Opcode: 0xef, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xF0, Length: 2, Cycles: 2, Documented: true , Instruction: (*CPU).beq,   Addressing: (*CPU).immediate},
	{ Opcode: 0xF1, Length: 2, Cycles: 5, Documented: true , Instruction: (*CPU).sbc,   Addressing: (*CPU).indirectIndexedY},
	{ Opcode: 0xf2, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xf3, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xf4, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xF5, Length: 2, Cycles: 4, Documented: true , Instruction: (*CPU).sbc,   Addressing: (*CPU).zeroPageX},
	{ Opcode: 0xF6, Length: 2, Cycles: 6, Documented: true , Instruction: (*CPU).inc,   Addressing: (*CPU).zeroPageX},
	{ Opcode: 0xf7, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xF8, Length: 1, Cycles: 2, Documented: true , Instruction: (*CPU).sed,   Addressing: (*CPU).implied},
	{ Opcode: 0xF9, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).sbc,   Addressing: (*CPU).absoluteY},
	{ Opcode: 0xfa, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xfb, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xfc, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
	{ Opcode: 0xFD, Length: 3, Cycles: 4, Documented: true , Instruction: (*CPU).sbc,   Addressing: (*CPU).absoluteX},
	{ Opcode: 0xFE, Length: 3, Cycles: 7, Documented: true , Instruction: (*CPU).inc,   Addressing: (*CPU).absoluteX},
	{ Opcode: 0xff, Length: 1, Cycles: 0, Documented: false, Instruction: (*CPU).undoc, Addressing: (*CPU).implied},
}

