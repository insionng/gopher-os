package irq

import "github.com/achilleasa/gopher-os/kernel/kfmt/early"

// Regs contains a snapshot of the register values when an interrupt occurred.
type Regs struct {
	RAX uint64
	RBX uint64
	RCX uint64
	RDX uint64
	RSI uint64
	RDI uint64
	RBP uint64
	R8  uint64
	R9  uint64
	R10 uint64
	R11 uint64
	R12 uint64
	R13 uint64
	R14 uint64
	R15 uint64
}

// Print outputs a dump of the register values to the active console.
func (r *Regs) Print() {
	early.Printf("RAX = %16x RBX = %16x\n", r.RAX, r.RBX)
	early.Printf("RCX = %16x RDX = %16x\n", r.RCX, r.RDX)
	early.Printf("RSI = %16x RDI = %16x\n", r.RSI, r.RDI)
	early.Printf("RBP = %16x\n", r.RBP)
	early.Printf("R8  = %16x R9  = %16x\n", r.R8, r.R9)
	early.Printf("R10 = %16x R11 = %16x\n", r.R10, r.R11)
	early.Printf("R12 = %16x R13 = %16x\n", r.R12, r.R13)
	early.Printf("R14 = %16x R15 = %16x\n", r.R14, r.R15)
}

// Frame describes an exception frame that is automatically pushed by the CPU
// to the stack when an exception occurs.
type Frame struct {
	RIP    uint64
	CS     uint64
	RFlags uint64
	RSP    uint64
	SS     uint64
}

// Print outputs a dump of the exception frame to the active console.
func (f *Frame) Print() {
	early.Printf("RIP = %16x CS  = %16x\n", f.RIP, f.CS)
	early.Printf("RSP = %16x SS  = %16x\n", f.RSP, f.SS)
	early.Printf("RFL = %16x\n", f.RFlags)
}
