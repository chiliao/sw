//+build !noasm !appengine gc
// AUTO-GENERATED BY C2GOASM -- DO NOT EDIT

DATA LCDATA1<>+0x000(SB)/8, $0x5c5c5c5c5c5c5c5c
DATA LCDATA1<>+0x008(SB)/8, $0x5c5c5c5c5c5c5c5c
DATA LCDATA1<>+0x010(SB)/8, $0x5c5c5c5c5c5c5c5c
DATA LCDATA1<>+0x018(SB)/8, $0x5c5c5c5c5c5c5c5c
DATA LCDATA1<>+0x020(SB)/8, $0x2222222222222222
DATA LCDATA1<>+0x028(SB)/8, $0x2222222222222222
DATA LCDATA1<>+0x030(SB)/8, $0x2222222222222222
DATA LCDATA1<>+0x038(SB)/8, $0x2222222222222222
DATA LCDATA1<>+0x070(SB)/8, $0x0706050403020100
DATA LCDATA1<>+0x078(SB)/8, $0xffffffffffff0908
DATA LCDATA1<>+0x080(SB)/8, $0xff0f0e0d0c0b0aff
DATA LCDATA1<>+0x088(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x090(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x098(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x0a0(SB)/8, $0xff0f0e0d0c0b0aff
DATA LCDATA1<>+0x0a8(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x0b0(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x0b8(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x0c0(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x0c8(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x0d0(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x0d8(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x0e0(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x0e8(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x0f0(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x0f8(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x100(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x108(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x110(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x118(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x120(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x128(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x130(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x138(SB)/8, $0xffffffffffffffff
DATA LCDATA1<>+0x140(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x148(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x150(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x158(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x160(SB)/8, $0x0000000000220000
DATA LCDATA1<>+0x168(SB)/8, $0x2f00000000000000
DATA LCDATA1<>+0x170(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x178(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x180(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x188(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x190(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x198(SB)/8, $0x0000005c00000000
DATA LCDATA1<>+0x1a0(SB)/8, $0x000c000000080000
DATA LCDATA1<>+0x1a8(SB)/8, $0x000a000000000000
DATA LCDATA1<>+0x1b0(SB)/8, $0x00000009000d0000
DATA LCDATA1<>+0x1b8(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x1c0(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x1c8(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x1d0(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x1d8(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x1e0(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x1e8(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x1f0(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x1f8(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x200(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x208(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x210(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x218(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x220(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x228(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x230(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x238(SB)/8, $0x0000000000000000
GLOBL LCDATA1<>(SB), 8, $576

TEXT ·_parse_string_validate_only(SB), $0-40

    MOVQ src+0(FP), DI
    MOVQ maxStringSize+8(FP), SI
    MOVQ str_length+16(FP), DX
    MOVQ dst_length+24(FP), CX
    LEAQ LCDATA1<>(SB), BP

    WORD $0x8b4c; BYTE $0x1e     // mov    r11, qword [rsi]
    WORD $0x854d; BYTE $0xdb     // test    r11, r11
	JE LBB0_30
    WORD $0xf631                 // xor    esi, esi
    LONG $0x456ffdc5; BYTE $0x00 // vmovdqa    ymm0, yword 0[rbp] /* [rip + LCPI0_0] */
    LONG $0x4d6ffdc5; BYTE $0x20 // vmovdqa    ymm1, yword 32[rbp] /* [rip + LCPI0_1] */
    LONG $0x404d8d4c             // lea    r9, 64[rbp] /* [rip + __ZL10digittoval] */
    LONG $0x40958d4c; WORD $0x0001; BYTE $0x00 // lea    r10, 320[rbp] /* [rip + __ZL10escape_map] */
    WORD $0x8949; BYTE $0xfd     // mov    r13, rdi
    WORD $0x3145; BYTE $0xf6     // xor    r14d, r14d
    WORD $0x8948; BYTE $0xf8     // mov    rax, rdi
LBB0_2:
    LONG $0x106ffec5             // vmovdqu    ymm2, yword [rax]
    LONG $0xd874edc5             // vpcmpeqb    ymm3, ymm2, ymm0
    LONG $0xdbd7fdc5             // vpmovmskb    ebx, ymm3
    LONG $0xd174edc5             // vpcmpeqb    ymm2, ymm2, ymm1
    LONG $0xe2d77dc5             // vpmovmskb    r12d, ymm2
    WORD $0x438d; BYTE $0xff     // lea    eax, [rbx - 1]
    WORD $0x8544; BYTE $0xe0     // test    eax, r12d
	JNE LBB0_3
    LONG $0x24448d41; BYTE $0xff // lea    eax, [r12 - 1]
    WORD $0xd885                 // test    eax, ebx
	JE LBB0_28
    WORD $0xd889                 // mov    eax, ebx
    LONG $0xbc0f4cf3; BYTE $0xf8 // tzcnt    r15, rax
    LONG $0x74b60f43; WORD $0x013d // movzx    esi, byte [r13 + r15 + 1]
    LONG $0x75fe8348             // cmp    rsi, 117
	JNE LBB0_26
    WORD $0x8545; BYTE $0xe4     // test    r12d, r12d
	JE LBB0_8
    WORD $0x8944; BYTE $0xe0     // mov    eax, r12d
    LONG $0xbc0f48f3; BYTE $0xf0 // tzcnt    rsi, rax
    WORD $0x2944; BYTE $0xfe     // sub    esi, r15d
    WORD $0xfe83; BYTE $0x06     // cmp    esi, 6
	JAE LBB0_11
	JMP LBB0_30
LBB0_28:
    LONG $0x20c58349             // add    r13, 32
    LONG $0x20c68349             // add    r14, 32
    WORD $0x894c; BYTE $0xe8     // mov    rax, r13
	JMP LBB0_29
LBB0_26:
    LONG $0x163c8042; BYTE $0x00 // cmp    byte [rsi + r10], 0
	JE LBB0_30
    LONG $0x3d448d4b; BYTE $0x02 // lea    rax, [r13 + r15 + 2]
    WORD $0xff49; BYTE $0xc7     // inc    r15
    WORD $0x014d; BYTE $0xfe     // add    r14, r15
	JMP LBB0_29
LBB0_8:
    LONG $0x000020be; BYTE $0x00 // mov    esi, 32
    LONG $0x15ff8341             // cmp    r15d, 21
	JB LBB0_10
    LONG $0xec478d41             // lea    eax, [r15 - 20]
    LONG $0x7475c1c4; WORD $0x0554; BYTE $0x00 // vpcmpeqb    ymm2, ymm1, yword [r13 + rax]
    LONG $0xc2d7fdc5             // vpmovmskb    eax, ymm2
    LONG $0xbc0f48f3; BYTE $0xf0 // tzcnt    rsi, rax
    WORD $0xc085                 // test    eax, eax
    LONG $0x000020b8; BYTE $0x00 // mov    eax, 32
    WORD $0x440f; BYTE $0xf0     // cmove    esi, eax
    LONG $0x3e748d42; BYTE $0xec // lea    esi, [rsi + r15 - 20]
LBB0_10:
    WORD $0x2944; BYTE $0xfe     // sub    esi, r15d
    WORD $0xfe83; BYTE $0x06     // cmp    esi, 6
	JB LBB0_30
LBB0_11:
    WORD $0x014d; BYTE $0xfd     // add    r13, r15
    LONG $0x45b60f41; BYTE $0x02 // movzx    eax, byte [r13 + 2]
    LONG $0x04be0f46; BYTE $0x08 // movsx    r8d, byte [rax + r9]
    LONG $0x5db60f41; BYTE $0x03 // movzx    ebx, byte [r13 + 3]
    LONG $0x1cbe0f42; BYTE $0x0b // movsx    ebx, byte [rbx + r9]
    LONG $0x45b60f41; BYTE $0x04 // movzx    eax, byte [r13 + 4]
    LONG $0x24be0f46; BYTE $0x08 // movsx    r12d, byte [rax + r9]
    LONG $0x45b60f41; BYTE $0x05 // movzx    eax, byte [r13 + 5]
    LONG $0x04be0f42; BYTE $0x08 // movsx    eax, byte [rax + r9]
    LONG $0x0ce0c141             // shl    r8d, 12
    WORD $0xe3c1; BYTE $0x08     // shl    ebx, 8
    WORD $0x0944; BYTE $0xc3     // or    ebx, r8d
    LONG $0x04e4c141             // shl    r12d, 4
    WORD $0x0941; BYTE $0xc4     // or    r12d, eax
    WORD $0x0941; BYTE $0xdc     // or    r12d, ebx
    LONG $0x06458d49             // lea    rax, [r13 + 6]
    WORD $0x8944; BYTE $0xe3     // mov    ebx, r12d
    LONG $0xfc00e381; WORD $0xffff // and    ebx, -1024
    LONG $0xd800fb81; WORD $0x0000 // cmp    ebx, 55296
	JE LBB0_12
    LONG $0x80fc8141; WORD $0x0000; BYTE $0x00 // cmp    r12d, 128
	JAE LBB0_19
LBB0_18:
    LONG $0x000001be; BYTE $0x00 // mov    esi, 1
	JMP LBB0_25
LBB0_12:
    WORD $0xfe83; BYTE $0x0c     // cmp    esi, 12
	JB LBB0_30
    WORD $0x3880; BYTE $0x5c     // cmp    byte [rax], 92
	JNE LBB0_30
    LONG $0x077d8041; BYTE $0x75 // cmp    byte [r13 + 7], 117
	JNE LBB0_30
    LONG $0x45b60f41; BYTE $0x08 // movzx    eax, byte [r13 + 8]
    LONG $0x04be0f46; BYTE $0x08 // movsx    r8d, byte [rax + r9]
    LONG $0x75b60f41; BYTE $0x09 // movzx    esi, byte [r13 + 9]
    LONG $0x1cbe0f42; BYTE $0x0e // movsx    ebx, byte [rsi + r9]
    LONG $0x75b60f41; BYTE $0x0a // movzx    esi, byte [r13 + 10]
    LONG $0x34be0f42; BYTE $0x0e // movsx    esi, byte [rsi + r9]
    LONG $0x45b60f41; BYTE $0x0b // movzx    eax, byte [r13 + 11]
    LONG $0x04be0f42; BYTE $0x08 // movsx    eax, byte [rax + r9]
    LONG $0x0ce0c141             // shl    r8d, 12
    WORD $0xe3c1; BYTE $0x08     // shl    ebx, 8
    WORD $0x0944; BYTE $0xc3     // or    ebx, r8d
    WORD $0xe6c1; BYTE $0x04     // shl    esi, 4
    WORD $0xc609                 // or    esi, eax
    WORD $0xde09                 // or    esi, ebx
    WORD $0xf089                 // mov    eax, esi
    WORD $0x0944; BYTE $0xe0     // or    eax, r12d
    LONG $0x00ffff3d; BYTE $0x00 // cmp    eax, 65535
	JA LBB0_30
    LONG $0x0ae4c141             // shl    r12d, 10
    LONG $0x00c48141; WORD $0xa000; BYTE $0xfc // add    r12d, -56623104
    LONG $0x2400c681; WORD $0xffff // add    esi, -56320
    WORD $0x0944; BYTE $0xe6     // or    esi, r12d
    LONG $0x0000c681; WORD $0x0001 // add    esi, 65536
    LONG $0x0cc58349             // add    r13, 12
    WORD $0x894c; BYTE $0xe8     // mov    rax, r13
    WORD $0x8941; BYTE $0xf4     // mov    r12d, esi
    LONG $0x80fc8141; WORD $0x0000; BYTE $0x00 // cmp    r12d, 128
	JB LBB0_18
LBB0_19:
    LONG $0x00fc8141; WORD $0x0008; BYTE $0x00 // cmp    r12d, 2048
	JAE LBB0_21
    LONG $0x000002be; BYTE $0x00 // mov    esi, 2
	JMP LBB0_25
LBB0_21:
    LONG $0x00fc8141; WORD $0x0100; BYTE $0x00 // cmp    r12d, 65536
	JAE LBB0_23
    LONG $0x000003be; BYTE $0x00 // mov    esi, 3
	JMP LBB0_25
LBB0_23:
    LONG $0xfffc8141; WORD $0x10ff; BYTE $0x00 // cmp    r12d, 1114111
	JA LBB0_30
    LONG $0x000004be; BYTE $0x00 // mov    esi, 4
LBB0_25:
    WORD $0x014d; BYTE $0xfe     // add    r14, r15
    WORD $0x0149; BYTE $0xf6     // add    r14, rsi
LBB0_29:
    WORD $0x8948; BYTE $0xc6     // mov    rsi, rax
    WORD $0x2948; BYTE $0xfe     // sub    rsi, rdi
    WORD $0x8949; BYTE $0xc5     // mov    r13, rax
    WORD $0x394c; BYTE $0xde     // cmp    rsi, r11
	JB LBB0_2
LBB0_30:
    WORD $0xc031                 // xor    eax, eax
	JMP LBB0_31
LBB0_3:
    WORD $0x8944; BYTE $0xe0     // mov    eax, r12d
    LONG $0xbc0f48f3; BYTE $0xc0 // tzcnt    rax, rax
    WORD $0x0148; BYTE $0xc6     // add    rsi, rax
    WORD $0x8948; BYTE $0x32     // mov    qword [rdx], rsi
    WORD $0x0149; BYTE $0xc6     // add    r14, rax
    WORD $0x894c; BYTE $0x31     // mov    qword [rcx], r14
    WORD $0x01b0                 // mov    al, 1
LBB0_31:
    VZEROUPPER
    MOVQ AX, result+32(FP)
    RET

TEXT ·_parse_string(SB), $0-32

    MOVQ src+0(FP), DI
    MOVQ dst+8(FP), SI
    MOVQ pcurrent_string_buf_loc+16(FP), DX
    LEAQ LCDATA1<>(SB), BP

    LONG $0x076ffec5             // vmovdqu    ymm0, yword [rdi]
    LONG $0x067ffec5             // vmovdqu    yword [rsi], ymm0
    LONG $0x4d74fdc5; BYTE $0x00 // vpcmpeqb    ymm1, ymm0, yword 0[rbp] /* [rip + LCPI0_0] */
    LONG $0xc9d7fdc5             // vpmovmskb    ecx, ymm1
    LONG $0x4574fdc5; BYTE $0x20 // vpcmpeqb    ymm0, ymm0, yword 32[rbp] /* [rip + LCPI0_1] */
    LONG $0xf0d77dc5             // vpmovmskb    r14d, ymm0
    WORD $0x418d; BYTE $0xff     // lea    eax, [rcx - 1]
    WORD $0x8544; BYTE $0xf0     // test    eax, r14d
	JE LBB0_3
LBB0_1:
    WORD $0x8944; BYTE $0xf0     // mov    eax, r14d
    LONG $0xbc0f48f3; BYTE $0xc0 // tzcnt    rax, rax
    WORD $0x0148; BYTE $0xf0     // add    rax, rsi
    WORD $0x8948; BYTE $0x02     // mov    qword [rdx], rax
    LONG $0x000001b8; BYTE $0x00 // mov    eax, 1
LBB0_2:
    VZEROUPPER
    MOVQ AX, res+24(FP)
    RET
LBB0_3:
    LONG $0x456ffdc5; BYTE $0x00 // vmovdqa    ymm0, yword 0[rbp] /* [rip + LCPI0_0] */
    LONG $0x4d6ffdc5; BYTE $0x20 // vmovdqa    ymm1, yword 32[rbp] /* [rip + LCPI0_1] */
    WORD $0xc031                 // xor    eax, eax
    LONG $0x40658d4c             // lea    r12, 64[rbp] /* [rip + __ZL10digittoval] */
    LONG $0x0001b941; WORD $0x0000 // mov    r9d, 1
    LONG $0x0002ba41; WORD $0x0000 // mov    r10d, 2
    LONG $0x40bd8d4c; WORD $0x0001; BYTE $0x00 // lea    r15, 320[rbp] /* [rip + __ZL10escape_map] */
	JMP LBB0_4
LBB0_24:
    LONG $0xfff88141; WORD $0x00ff; BYTE $0x00 // cmp    r8d, 65535
	JA LBB0_26
    WORD $0x8944; BYTE $0xc1     // mov    ecx, r8d
    WORD $0xe9c1; BYTE $0x0c     // shr    ecx, 12
    LONG $0x00e0c181; WORD $0x0000 // add    ecx, 224
    WORD $0x0e88                 // mov    byte [rsi], cl
    WORD $0x8944; BYTE $0xc1     // mov    ecx, r8d
    WORD $0xe9c1; BYTE $0x06     // shr    ecx, 6
    WORD $0xe180; BYTE $0x3f     // and    cl, 63
    WORD $0xc980; BYTE $0x80     // or    cl, -128
    WORD $0x4e88; BYTE $0x01     // mov    byte [rsi + 1], cl
    LONG $0x3fe08041             // and    r8b, 63
    LONG $0x80c88041             // or    r8b, -128
    LONG $0x02468844             // mov    byte [rsi + 2], r8b
    LONG $0x000003b9; BYTE $0x00 // mov    ecx, 3
	JMP LBB0_28
LBB0_26:
    LONG $0xfff88141; WORD $0x10ff; BYTE $0x00 // cmp    r8d, 1114111
	JA LBB0_2
    WORD $0x8944; BYTE $0xc1     // mov    ecx, r8d
    WORD $0xe9c1; BYTE $0x12     // shr    ecx, 18
    LONG $0x00f0c181; WORD $0x0000 // add    ecx, 240
    WORD $0x0e88                 // mov    byte [rsi], cl
    WORD $0x8944; BYTE $0xc1     // mov    ecx, r8d
    WORD $0xe9c1; BYTE $0x0c     // shr    ecx, 12
    WORD $0xe180; BYTE $0x3f     // and    cl, 63
    WORD $0xc980; BYTE $0x80     // or    cl, -128
    WORD $0x4e88; BYTE $0x01     // mov    byte [rsi + 1], cl
    WORD $0x8944; BYTE $0xc1     // mov    ecx, r8d
    WORD $0xe9c1; BYTE $0x06     // shr    ecx, 6
    WORD $0xe180; BYTE $0x3f     // and    cl, 63
    WORD $0xc980; BYTE $0x80     // or    cl, -128
    WORD $0x4e88; BYTE $0x02     // mov    byte [rsi + 2], cl
    LONG $0x3fe08041             // and    r8b, 63
    LONG $0x80c88041             // or    r8b, -128
    LONG $0x03468844             // mov    byte [rsi + 3], r8b
    LONG $0x000004b9; BYTE $0x00 // mov    ecx, 4
	JMP LBB0_28
LBB0_4:
    LONG $0xff5e8d41             // lea    ebx, [r14 - 1]
    WORD $0xcb85                 // test    ebx, ecx
	JE LBB0_8
    WORD $0xc989                 // mov    ecx, ecx
    LONG $0xbc0f4cf3; BYTE $0xd9 // tzcnt    r11, rcx
    LONG $0x4cb60f42; WORD $0x011f // movzx    ecx, byte [rdi + r11 + 1]
    LONG $0x75f98348             // cmp    rcx, 117
	JNE LBB0_9
    WORD $0x8545; BYTE $0xf6     // test    r14d, r14d
	JE LBB0_11
    WORD $0x8944; BYTE $0xf1     // mov    ecx, r14d
    LONG $0xbc0f4cf3; BYTE $0xf1 // tzcnt    r14, rcx
    WORD $0x2945; BYTE $0xde     // sub    r14d, r11d
    LONG $0x06fe8341             // cmp    r14d, 6
	JAE LBB0_14
	JMP LBB0_2
LBB0_8:
    LONG $0x20c78348             // add    rdi, 32
    LONG $0x20c68348             // add    rsi, 32
    WORD $0x8949; BYTE $0xfd     // mov    r13, rdi
	JMP LBB0_29
LBB0_9:
    LONG $0x0cb60f42; BYTE $0x39 // movzx    ecx, byte [rcx + r15]
    WORD $0xc984                 // test    cl, cl
	JE LBB0_2
    LONG $0x1e0c8842             // mov    byte [rsi + r11], cl
    LONG $0x1f6c8d4e; BYTE $0x02 // lea    r13, [rdi + r11 + 2]
    LONG $0x014b8d49             // lea    rcx, [r11 + 1]
LBB0_28:
    WORD $0x0148; BYTE $0xce     // add    rsi, rcx
	JMP LBB0_29
LBB0_11:
    LONG $0x0020be41; WORD $0x0000 // mov    r14d, 32
    LONG $0x15fb8341             // cmp    r11d, 21
	JB LBB0_13
    LONG $0xec4b8d41             // lea    ecx, [r11 - 20]
    LONG $0x1474f5c5; BYTE $0x0f // vpcmpeqb    ymm2, ymm1, yword [rdi + rcx]
    LONG $0xcad7fdc5             // vpmovmskb    ecx, ymm2
    LONG $0xbc0f48f3; BYTE $0xd9 // tzcnt    rbx, rcx
    WORD $0xc985                 // test    ecx, ecx
    LONG $0x000020b9; BYTE $0x00 // mov    ecx, 32
    WORD $0x440f; BYTE $0xd9     // cmove    ebx, ecx
    LONG $0x1b748d46; BYTE $0xec // lea    r14d, [rbx + r11 - 20]
LBB0_13:
    WORD $0x2945; BYTE $0xde     // sub    r14d, r11d
    LONG $0x06fe8341             // cmp    r14d, 6
	JB LBB0_2
LBB0_14:
    WORD $0x014c; BYTE $0xdf     // add    rdi, r11
    LONG $0x024fb60f             // movzx    ecx, byte [rdi + 2]
    LONG $0x2cbe0f46; BYTE $0x21 // movsx    r13d, byte [rcx + r12]
    LONG $0x035fb60f             // movzx    ebx, byte [rdi + 3]
    LONG $0x1cbe0f42; BYTE $0x23 // movsx    ebx, byte [rbx + r12]
    LONG $0x044fb60f             // movzx    ecx, byte [rdi + 4]
    LONG $0x04be0f46; BYTE $0x21 // movsx    r8d, byte [rcx + r12]
    LONG $0x054fb60f             // movzx    ecx, byte [rdi + 5]
    LONG $0x0cbe0f42; BYTE $0x21 // movsx    ecx, byte [rcx + r12]
    LONG $0x0ce5c141             // shl    r13d, 12
    WORD $0xe3c1; BYTE $0x08     // shl    ebx, 8
    WORD $0x0944; BYTE $0xeb     // or    ebx, r13d
    LONG $0x04e0c141             // shl    r8d, 4
    WORD $0x0941; BYTE $0xc8     // or    r8d, ecx
    WORD $0x0941; BYTE $0xd8     // or    r8d, ebx
    LONG $0x066f8d4c             // lea    r13, [rdi + 6]
    WORD $0x8944; BYTE $0xc1     // mov    ecx, r8d
    LONG $0xfc00e181; WORD $0xffff // and    ecx, -1024
    LONG $0xd800f981; WORD $0x0000 // cmp    ecx, 55296
	JNE LBB0_20
    LONG $0x0cfe8341             // cmp    r14d, 12
	JB LBB0_2
    LONG $0x007d8041; BYTE $0x5c // cmp    byte [r13], 92
	JNE LBB0_2
    LONG $0x75077f80             // cmp    byte [rdi + 7], 117
	JNE LBB0_2
    LONG $0x084fb60f             // movzx    ecx, byte [rdi + 8]
    LONG $0x34be0f46; BYTE $0x21 // movsx    r14d, byte [rcx + r12]
    LONG $0x094fb60f             // movzx    ecx, byte [rdi + 9]
    LONG $0x2cbe0f46; BYTE $0x21 // movsx    r13d, byte [rcx + r12]
    LONG $0x0a4fb60f             // movzx    ecx, byte [rdi + 10]
    LONG $0x0cbe0f42; BYTE $0x21 // movsx    ecx, byte [rcx + r12]
    LONG $0x0b5fb60f             // movzx    ebx, byte [rdi + 11]
    LONG $0x1cbe0f42; BYTE $0x23 // movsx    ebx, byte [rbx + r12]
    LONG $0x0ce6c141             // shl    r14d, 12
    LONG $0x08e5c141             // shl    r13d, 8
    WORD $0x0945; BYTE $0xf5     // or    r13d, r14d
    WORD $0xe1c1; BYTE $0x04     // shl    ecx, 4
    WORD $0xd909                 // or    ecx, ebx
    WORD $0x0944; BYTE $0xe9     // or    ecx, r13d
    WORD $0xcb89                 // mov    ebx, ecx
    WORD $0x0944; BYTE $0xc3     // or    ebx, r8d
    LONG $0xfffffb81; WORD $0x0000 // cmp    ebx, 65535
	JA LBB0_2
    LONG $0x0ae0c141             // shl    r8d, 10
    LONG $0x00c08141; WORD $0xa000; BYTE $0xfc // add    r8d, -56623104
    LONG $0x2400c181; WORD $0xffff // add    ecx, -56320
    WORD $0x0944; BYTE $0xc1     // or    ecx, r8d
    LONG $0x0000c181; WORD $0x0001 // add    ecx, 65536
    LONG $0x0cc78348             // add    rdi, 12
    WORD $0x8949; BYTE $0xfd     // mov    r13, rdi
    WORD $0x8941; BYTE $0xc8     // mov    r8d, ecx
LBB0_20:
    WORD $0x014c; BYTE $0xde     // add    rsi, r11
    LONG $0x7ff88341             // cmp    r8d, 127
	JA LBB0_22
    WORD $0x8844; BYTE $0x06     // mov    byte [rsi], r8b
    WORD $0x014c; BYTE $0xce     // add    rsi, r9
	JMP LBB0_29
LBB0_22:
    LONG $0xfff88141; WORD $0x0007; BYTE $0x00 // cmp    r8d, 2047
	JA LBB0_24
    WORD $0x8944; BYTE $0xc1     // mov    ecx, r8d
    WORD $0xe9c1; BYTE $0x06     // shr    ecx, 6
    LONG $0x00c0c181; WORD $0x0000 // add    ecx, 192
    WORD $0x0e88                 // mov    byte [rsi], cl
    LONG $0x3fe08041             // and    r8b, 63
    LONG $0x80c88041             // or    r8b, -128
    LONG $0x01468844             // mov    byte [rsi + 1], r8b
    WORD $0x014c; BYTE $0xd6     // add    rsi, r10
LBB0_29:
    LONG $0x6f7ec1c4; WORD $0x0055 // vmovdqu    ymm2, yword [r13]
    LONG $0x167ffec5             // vmovdqu    yword [rsi], ymm2
    LONG $0xd874edc5             // vpcmpeqb    ymm3, ymm2, ymm0
    LONG $0xcbd7fdc5             // vpmovmskb    ecx, ymm3
    LONG $0xd174edc5             // vpcmpeqb    ymm2, ymm2, ymm1
    LONG $0xf2d77dc5             // vpmovmskb    r14d, ymm2
    WORD $0x598d; BYTE $0xff     // lea    ebx, [rcx - 1]
    WORD $0x894c; BYTE $0xef     // mov    rdi, r13
    WORD $0x8544; BYTE $0xf3     // test    ebx, r14d
	JE LBB0_4
	JMP LBB0_1
