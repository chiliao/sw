#include "apollo.h"

%%

nop:
    nop.e
    nop

.align
erspan:
    nop.e
    nop

/*****************************************************************************/
/* error function                                                            */
/*****************************************************************************/
.align
.assert $ < ASM_INSTRUCTION_OFFSET_MAX
mirror_error:
    nop.e
    nop
