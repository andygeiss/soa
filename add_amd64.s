#include "textflag.h"

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
