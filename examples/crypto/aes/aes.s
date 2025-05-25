# AES-128 usando AES-NI em Assembly (GAS syntax)
# Este é um exemplo didático: criptografa 1 bloco de 16 bytes com 128 bits de chave fixa
# Exibe entrada e saída com syscall write

.intel_syntax noprefix
.section .data
    plaintext: .byte 0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d
               .byte 0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x07, 0x34
    key:       .byte 0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6
               .byte 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c
    msg_in:    .asciz "Entrada:  "
    msg_out:   .asciz "\nSaida:    "

.section .bss
    output: .space 16

.section .text
.globl _start
_start:
    # Print "Entrada:"
    mov rax, 1          # syscall write
    mov rdi, 1          # stdout
    lea rsi, msg_in[rip]
    mov rdx, 9
    syscall

    # Print dados da entrada
    lea rsi, plaintext[rip]
    mov rdx, 16
    call print_hex

    # Carrega bloco e chave
    movdqu xmm0, xmmword ptr [plaintext]
    movdqu xmm1, xmmword ptr [key]

    # AddRoundKey inicial
    pxor xmm0, xmm1

    # Rodadas do AES (sem gerar subchaves reais)
    mov ecx, 9
.round_loop:
    aesenc xmm0, xmm1
    loop .round_loop

    aesenclast xmm0, xmm1

    # Salva resultado em 'output'
    movdqu xmmword ptr [output], xmm0

    # Print "\nSaida:"
    mov rax, 1
    mov rdi, 1
    lea rsi, msg_out[rip]
    mov rdx, 10
    syscall

    # Print dados da saida
    lea rsi, output[rip]
    mov rdx, 16
    call print_hex

    # Exit
    mov rax, 60
    xor rdi, rdi
    syscall

# --------------------------------------------
# Função print_hex
# Entrada: RSI = ponteiro para dados, RDX = tamanho
# Imprime os bytes em hex no stdout
.type print_hex, @function
print_hex:
    push rbx
.next:
    cmp rdx, 0
    je .done
    movzx rax, byte ptr [rsi]
    mov rbx, offset hex_chars
    mov cl, al
    shr al, 4
    mov al, byte ptr [rbx + rax]
    call write_char
    mov al, cl
    and al, 0x0F
    mov al, byte ptr [rbx + rax]
    call write_char
    inc rsi
    dec rdx
    jmp .next
.done:
    mov al, 0x0A   # newline
    call write_char

    pop rbx
    ret

write_char:
    push rax
    mov rsi, rsp
    mov rdi, 1
    mov rdx, 1
    mov rax, 1
    syscall
    pop rax
    ret

.section .rodata
hex_chars: .ascii "0123456789ABCDEF"
