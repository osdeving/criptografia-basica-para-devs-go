# Estrutura interna (SPN)

## Visão geral da estrutura

O **Advanced Encryption Standard (AES)** é uma cifra de bloco com bloco fixo de 128 bits e chaves de 128, 192 ou 256 bits. Seu projeto se baseia em uma rede de substituição-permutação (SPN), diferentemente do DES, que é estruturado como uma rede de Feistel.

Cada operação no AES atua sobre uma matriz de bytes chamada **State**, com 4 linhas e 4 colunas (para bloco de 128 bits). O número de rodadas depende do tamanho da chave:

| Tamanho da chave | Rodadas |
| ---------------- | ------- |
| 128 bits         | 10      |
| 192 bits         | 12      |
| 256 bits         | 14      |

## Representação do bloco

O bloco de 128 bits é representado como uma matriz 4×4 de bytes:

```plaintext
State =
[ b0  b4  b8  b12 ]
[ b1  b5  b9  b13 ]
[ b2  b6  b10 b14 ]
[ b3  b7  b11 b15 ]
```

## Fases da cifra

Cada rodada (exceto a inicial e a final) realiza quatro operações fundamentais:

1. **SubBytes:** substituição não-linear byte a byte usando uma S-box baseada em inverso multiplicativo sobre $\mathbb{F}_{2^8}$.
2. **ShiftRows:** permutação cíclica das linhas da matriz.
3. **MixColumns:** transformação linear das colunas da matriz com multiplicações no campo $\mathbb{F}_{2^8}$.
4. **AddRoundKey:** operação XOR entre o *state* e uma subchave derivada da chave principal.

A rodada inicial realiza apenas o AddRoundKey, e a rodada final omite o MixColumns.

### Exemplo de fluxo para AES-128

```plaintext
Input Block
↓
AddRoundKey (Chave 0)
↓
[Rodada 1 a 9]
   SubBytes → ShiftRows → MixColumns → AddRoundKey
↓
Rodada 10
   SubBytes → ShiftRows → AddRoundKey
↓
Output Ciphertext
```

## SubBytes

A S-box é construída combinando duas operações:

* Inverso multiplicativo no campo $\mathbb{F}_{2^8}$, com zero mapeado para zero.
* Transformação afim (operação linear sobre os bits).

Essa estrutura proporciona não-linearidade e resistência a criptoanálise diferencial e linear.

## ShiftRows

As linhas da matriz são rotacionadas:

* Linha 0: sem mudança.
* Linha 1: desloca 1 byte à esquerda.
* Linha 2: desloca 2 bytes.
* Linha 3: desloca 3 bytes.

Essa operação contribui para a difusão horizontal dos dados.

## MixColumns

Cada coluna da matriz é tratada como um vetor de 4 bytes e multiplicada por uma matriz fixa sobre $\mathbb{F}_{2^8}$:

```plaintext
[02 03 01 01]
[01 02 03 01]
[01 01 02 03]
[03 01 01 02]
```

A operação assegura a difusão entre os bytes de cada coluna.

## AddRoundKey

Cada byte do *state* é combinado com o byte correspondente da subchave da rodada por meio de XOR. Essa é a única operação que introduz segredo na cifra.

---

## Geração das subchaves (Key Schedule)

O algoritmo de derivação de subchaves expande a chave principal em um número fixo de palavras de 32 bits. O número total de subchaves depende do tamanho da chave original:

| Chave | Palavras derivadas |
| ----- | ------------------ |
| 128   | 44                 |
| 192   | 52                 |
| 256   | 60                 |

Cada nova palavra é construída a partir da anterior com operações como rotacionamento, substituição com S-box e XOR com constantes Rcon.

---

## Considerações finais

A estrutura interna do AES oferece forte segurança combinando **não-linearidade, difusão e chaveamento dependente da rodada**. Sua simetria matemática e clareza algorítmica tornam-no ideal tanto para análise teórica quanto para implementação prática em hardware e software.
