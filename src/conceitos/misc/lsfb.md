# LSFB - Linear Shift Feedback

O LFSR (Linear Feedback Shift Register) é um tipo de registrador de deslocamento com feedback linear. Ele é usado em várias aplicações de criptografia, incluindo a geração de números pseudoaleatórios (PRNGs) e a criptografia de fluxo (stream cipher).

## Como Funciona

Você tem um vetor de bits inicial (a semente, ou estado inicial). A cada ciclo, o registrador é deslocado para a direita e um novo bit é calculado como XOR de alguns bits do estado atual, esse bit é então adicionado ao bit mais à esquerda.

## Exemplo com 4 bits

Estado inicial (semente): 1001

Taps: posição 1 e 4 (ou seja, os bits 1 e 4 são usados para calcular o novo bit)

Iterações:

1. Estado atual: 1001
   Feedback: 1 XOR 1 = 0
   Novo estado: 0100

2. Estado atual: 0100
   Feedback: 0 XOR 0 = 0
   Novo estado: 0010

3. Estado atual: 0010
   Feedback: 0 XOR 0 = 0
   Novo estado: 0001

4. Estado atual: 0001
   Feedback: 0 XOR 1 = 1
   Novo estado: 1000

E asssim por diante...


## Fórmula Geral










