.section .data
msg_success:
    .asciz "RDRAND succeeded: "
msg_fail:
    .asciz "RDRAND failed\n"

.section .bss
    .lcomm rand_val, 8

.section .text
.globl _start

_start:
    # Executa RDRAND em %rax
    rdrand %rax
    jc success          # carry set → sucesso
    jmp fail            # carry clear → falha

success:
    movq %rax, rand_val(%rip)

    # write(1, msg_success, 18)
    movq $1, %rax        # syscall: write
    movq $1, %rdi        # stdout
    movq $msg_success, %rsi
    movq $18, %rdx
    syscall

    # imprime o número
    movq rand_val(%rip), %rax
    call print_uint64
    jmp exit

fail:
    # write(1, msg_fail, 14)
    movq $1, %rax
    movq $1, %rdi
    movq $msg_fail, %rsi
    movq $14, %rdx
    syscall
    jmp exit

exit:
    movq $60, %rax       # syscall: exit
    xor %rdi, %rdi       # exit code 0
    syscall

# ----------------------------------------------------------
# Função: print_uint64
# Imprime %rax como número decimal, seguido de \n
# ----------------------------------------------------------
print_uint64:
    subq $32, %rsp
    movq $10, %rcx
    leaq 31(%rsp), %rsi
    movb $10, (%rsi)
    decq %rsi

print_loop:
    xorq %rdx, %rdx
    divq %rcx             # divide %rax por 10
    addb $'0', %dl
    movb %dl, (%rsi)
    decq %rsi
    testq %rax, %rax
    jnz print_loop

    incq %rsi
    movq $1, %rax         # syscall write
    movq $1, %rdi
    movq %rsi, %rsi
    movq $0, %rdx
    leaq 32(%rsp), %rcx
    subq %rsi, %rcx
    movq %rcx, %rdx
    syscall

    addq $32, %rsp
    ret
