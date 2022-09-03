#include "textflag.h"

TEXT ·get(SB), NOSPLIT, $0-8
    MOVQ main·a(SB), AX
    MOVQ AX, ret+0(FP)
    RET

TEXT ·outerfunc(SB), NOSPLIT, $0
    CALL ·gofunc(SB)
    RET
