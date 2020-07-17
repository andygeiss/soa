#include "textflag.h"

// func AddFloat32s(a, b, c []float32)
TEXT ·AddFloat32s(SB), NOSPLIT, $0
    MOVQ    a+0(FP),        AX
    MOVQ    b+24(FP),       BX
    MOVQ    c+48(FP),       DI
    MOVQ    len+8(FP),      CX
    XORQ    DX,             DX
    SHLQ    $2,             CX

LOOP:
    MOVAPS  0(AX),          X0
    MOVAPS  16(AX),         X1
    MOVAPS  32(AX),         X2
    MOVAPS  48(AX),         X3
    MOVAPS  64(AX),         X4
    MOVAPS  80(AX),         X5
    MOVAPS  96(AX),         X6
    MOVAPS  112(AX),        X7
    
    MOVAPS  0(BX),          X8
    MOVAPS  16(BX),         X9
    MOVAPS  32(BX),         X10
    MOVAPS  48(BX),         X11
    MOVAPS  64(BX),         X12
    MOVAPS  80(BX),         X13
    MOVAPS  96(BX),         X14
    MOVAPS  112(BX),        X15

    ADDPS   X0,             X8
    ADDPS   X1,             X9
    ADDPS   X2,             X10
    ADDPS   X3,             X11
    ADDPS   X4,             X12
    ADDPS   X5,             X13
    ADDPS   X6,             X14
    ADDPS   X7,             X15
    
    MOVAPS  X8,             0(DI)
    MOVAPS  X9,             16(DI)
    MOVAPS  X10,            32(DI)
    MOVAPS  X11,            48(DI)
    MOVAPS  X12,            64(DI)
    MOVAPS  X13,            80(DI)
    MOVAPS  X14,            96(DI)
    MOVAPS  X15,            112(DI)

    ADDQ    $128,           AX
    ADDQ    $128,           BX
    ADDQ    $128,           DI
    SUBQ    $128,           CX

    CMPQ    DX,             CX
    JLT     LOOP

    RET

// func AddFloat64s(a, b, c []int64)
TEXT ·AddFloat64s(SB), NOSPLIT, $0
    MOVQ    a+0(FP),        AX
    MOVQ    b+24(FP),       BX
    MOVQ    c+48(FP),       DI
    MOVQ    len+8(FP),      CX
    XORQ    DX,             DX
    SHLQ    $3,             CX

LOOP:
    MOVAPD  0(AX),          X0
    MOVAPD  16(AX),         X1
    MOVAPD  32(AX),         X2
    MOVAPD  48(AX),         X3
    MOVAPD  64(AX),         X4
    MOVAPD  80(AX),         X5
    MOVAPD  96(AX),         X6
    MOVAPD  112(AX),        X7
    
    MOVAPD  0(BX),          X8
    MOVAPD  16(BX),         X9
    MOVAPD  32(BX),         X10
    MOVAPD  48(BX),         X11
    MOVAPD  64(BX),         X12
    MOVAPD  80(BX),         X13
    MOVAPD  96(BX),         X14
    MOVAPD  112(BX),        X15

    ADDPD   X0,             X8
    ADDPD   X1,             X9
    ADDPD   X2,             X10
    ADDPD   X3,             X11
    ADDPD   X4,             X12
    ADDPD   X5,             X13
    ADDPD   X6,             X14
    ADDPD   X7,             X15
    
    MOVAPD  X8,             0(DI)
    MOVAPD  X9,             16(DI)
    MOVAPD  X10,            32(DI)
    MOVAPD  X11,            48(DI)
    MOVAPD  X12,            64(DI)
    MOVAPD  X13,            80(DI)
    MOVAPD  X14,            96(DI)
    MOVAPD  X15,            112(DI)

    ADDQ    $128,           AX
    ADDQ    $128,           BX
    ADDQ    $128,           DI
    SUBQ    $128,           CX

    CMPQ    DX,             CX
    JLT     LOOP

    RET

// func AddInt32s(a, b, c []int32)
TEXT ·AddInt32s(SB), NOSPLIT, $0
    MOVQ    a+0(FP),        AX
    MOVQ    b+24(FP),       BX
    MOVQ    c+48(FP),       DI
    MOVQ    len+8(FP),      CX
    XORQ    DX,             DX
    SHLQ    $3,             CX

LOOP:
    MOVAPD  0(AX),          X0
    MOVAPD  16(AX),         X1
    MOVAPD  32(AX),         X2
    MOVAPD  48(AX),         X3
    MOVAPD  64(AX),         X4
    MOVAPD  80(AX),         X5
    MOVAPD  96(AX),         X6
    MOVAPD  112(AX),        X7
    
    MOVAPD  0(BX),          X8
    MOVAPD  16(BX),         X9
    MOVAPD  32(BX),         X10
    MOVAPD  48(BX),         X11
    MOVAPD  64(BX),         X12
    MOVAPD  80(BX),         X13
    MOVAPD  96(BX),         X14
    MOVAPD  112(BX),        X15

    ADDPD   X0,             X8
    ADDPD   X1,             X9
    ADDPD   X2,             X10
    ADDPD   X3,             X11
    ADDPD   X4,             X12
    ADDPD   X5,             X13
    ADDPD   X6,             X14
    ADDPD   X7,             X15
    
    MOVAPD  X8,             0(DI)
    MOVAPD  X9,             16(DI)
    MOVAPD  X10,            32(DI)
    MOVAPD  X11,            48(DI)
    MOVAPD  X12,            64(DI)
    MOVAPD  X13,            80(DI)
    MOVAPD  X14,            96(DI)
    MOVAPD  X15,            112(DI)

    ADDQ    $128,           AX
    ADDQ    $128,           BX
    ADDQ    $128,           DI
    SUBQ    $128,           CX

    CMPQ    DX,             CX
    JLT     LOOP

    RET

// func AddInt64s(a, b, c []int64)
TEXT ·AddInt64s(SB), NOSPLIT, $0
    MOVQ    a+0(FP),        AX
    MOVQ    b+24(FP),       BX
    MOVQ    c+48(FP),       DI
    MOVQ    len+8(FP),      CX
    XORQ    DX,             DX
    SHLQ    $3,             CX

LOOP:
    MOVAPD  0(AX),          X0
    MOVAPD  16(AX),         X1
    MOVAPD  32(AX),         X2
    MOVAPD  48(AX),         X3
    MOVAPD  64(AX),         X4
    MOVAPD  80(AX),         X5
    MOVAPD  96(AX),         X6
    MOVAPD  112(AX),        X7
    
    MOVAPD  0(BX),          X8
    MOVAPD  16(BX),         X9
    MOVAPD  32(BX),         X10
    MOVAPD  48(BX),         X11
    MOVAPD  64(BX),         X12
    MOVAPD  80(BX),         X13
    MOVAPD  96(BX),         X14
    MOVAPD  112(BX),        X15

    ADDPD   X0,             X8
    ADDPD   X1,             X9
    ADDPD   X2,             X10
    ADDPD   X3,             X11
    ADDPD   X4,             X12
    ADDPD   X5,             X13
    ADDPD   X6,             X14
    ADDPD   X7,             X15
    
    MOVAPD  X8,             0(DI)
    MOVAPD  X9,             16(DI)
    MOVAPD  X10,            32(DI)
    MOVAPD  X11,            48(DI)
    MOVAPD  X12,            64(DI)
    MOVAPD  X13,            80(DI)
    MOVAPD  X14,            96(DI)
    MOVAPD  X15,            112(DI)

    ADDQ    $128,           AX
    ADDQ    $128,           BX
    ADDQ    $128,           DI
    SUBQ    $128,           CX

    CMPQ    DX,             CX
    JLT     LOOP

    RET
