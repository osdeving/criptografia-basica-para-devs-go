.section .data
    num:        .quad 0x0102030405060708  # Número de 8 bytes (64 bits) para testar endianness
    msg_little: .asciz "Este sistema é Little-endian\n"
    msg_big:    .asciz "Este sistema é Big-endian\n"

.section .text
    .global _start

_start:
    movq num, %rax    # Carrega o número na CPU
    movb %al, %bl     # Copia o primeiro byte para %bl

    cmpb $0x08, %bl   # Compara se o primeiro byte é 0x08 (Little-endian)
    jne big_endian    # Se for diferente, pula para big_endian

    # Se for Little-endian, imprime mensagem
    movq $1, %rax     # syscall write (syscall number para write em 64 bits)
    movq $1, %rdi     # saída stdout
    movq $msg_little, %rsi
    movq $29, %rdx    # tamanho da string
    syscall           # chamada de sistema

    jmp exit          # Pula para o final

big_endian:
    # Se for Big-endian, imprime mensagem
    movq $1, %rax
    movq $1, %rdi
    movq $msg_big, %rsi
    movq $25, %rdx
    syscall

exit:
    movq $60, %rax    # syscall exit (syscall number para exit em 64 bits)
    xor %rdi, %rdi    # código de saída 0
    syscall
