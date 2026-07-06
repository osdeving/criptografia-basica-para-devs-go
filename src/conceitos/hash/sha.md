# Familia SHA - Secure Hash Algorithms

SHA significa `Secure Hash Algorithm`.

A familia SHA foi padronizada pelo NIST e tem varias geracoes:

```text
SHA-0
SHA-1
SHA-2
SHA-3
```

O nome parece uma sequencia simples, mas existe uma mudanca estrutural importante:

```text
SHA-0, SHA-1 e SHA-2:
  seguem o estilo Merkle-Damgard, com funcao de compressao por blocos

SHA-3:
  vem do Keccak e usa construcao sponge
```

## SHA-0

SHA-0 foi publicado em 1993 como o primeiro Secure Hash Standard.

Caracteristicas:

```text
saida: 160 bits
bloco: 512 bits
word: 32 bits
rodadas: 80
status: retirado/substituido
```

SHA-0 e parecido com SHA-1. A diferenca didatica mais importante esta no `message schedule`.

Em SHA-0:

```text
W[t] = W[t-3] XOR W[t-8] XOR W[t-14] XOR W[t-16]
```

Em SHA-1:

```text
W[t] = ROTL1(W[t-3] XOR W[t-8] XOR W[t-14] XOR W[t-16])
```

Ou seja: SHA-1 adiciona uma rotacao de 1 bit na expansao da mensagem.

Parece pouco. Mas em criptografia, uma rotacao colocada no lugar certo pode aumentar muito a difusao.

### Vetores SHA-0

```text
SHA0("")    = f96cea198ad1dd5617ac084a3d92c6107708c0ef
SHA0("abc") = 0164b8a914cd2a5e74c4f7ff082c4d97f1edf880
```

Teste:

```text
cargo test -p encryptor learning_hashes::sha::tests::sha0_vectors
```

## SHA-1

SHA-1 foi publicado em 1995 no FIPS 180-1.

Caracteristicas:

```text
saida: 160 bits
bloco: 512 bits
word: 32 bits
rodadas: 80
estado: 5 words de 32 bits
status moderno: nao usar para novas aplicacoes de seguranca
```

### Estado inicial

SHA-1 usa 5 words:

```text
H0 = 0x67452301
H1 = 0xefcdab89
H2 = 0x98badcfe
H3 = 0x10325476
H4 = 0xc3d2e1f0
```

### Padding

SHA-1 usa padding parecido com SHA-256:

```text
1. adiciona 0x80
2. adiciona zeros ate faltar 8 bytes no bloco
3. adiciona tamanho original em bits como u64 big-endian
```

Aqui ja aparece uma diferenca contra MD4/MD5:

```text
MD4/MD5: tamanho em little-endian
SHA-1:   tamanho em big-endian
```

### Rodadas

SHA-1 roda 80 passos divididos em 4 faixas:

```text
0..19:
  f = Ch(b,c,d)
  k = 0x5a827999

20..39:
  f = b XOR c XOR d
  k = 0x6ed9eba1

40..59:
  f = Maj(b,c,d)
  k = 0x8f1bbcdc

60..79:
  f = b XOR c XOR d
  k = 0xca62c1d6
```

Operacao central:

```text
temp = ROTL5(a) + f + e + k + W[t]
e = d
d = c
c = ROTL30(b)
b = a
a = temp
```

Tudo modulo `2^32`.

### O que quebrou no SHA-1

SHA-1 tem digest de 160 bits. A resistencia ideal contra colisao seria perto de `2^80`. Ataques diferenciais reduziram esse custo e culminaram em demonstracoes praticas.

Marcos importantes:

```text
2005:
  Wang, Yin e Yu publicam ataques importantes contra SHA-1

2017:
  SHAttered demonstra colisao publica para SHA-1

2020:
  SHA-1 is a Shambles demonstra ataque de prefixo escolhido mais pratico

2023:
  NIST anuncia decisao de revisar FIPS 180-4 para remover SHA-1
```

SHA-1 ainda pode aparecer em sistemas legados, mas nao deve ser usado para novas assinaturas, certificados ou integridade em ambiente adversarial.

### Vetores SHA-1

```text
SHA1("")    = da39a3ee5e6b4b0d3255bfef95601890afd80709
SHA1("abc") = a9993e364706816aba3e25717850c26c9cd0d89d
```

Teste:

```text
cargo test -p encryptor learning_hashes::sha::tests::sha1_vectors
```

## SHA-2

SHA-2 e uma familia, nao um unico algoritmo.

```text
SHA-224
SHA-256
SHA-384
SHA-512
SHA-512/224
SHA-512/256
```

As variantes se dividem em dois grupos:

```text
SHA-224 e SHA-256:
  bloco de 512 bits
  words de 32 bits
  64 rodadas

SHA-384, SHA-512, SHA-512/224, SHA-512/256:
  bloco de 1024 bits
  words de 64 bits
  80 rodadas
```

Tabela:

| Algoritmo | Bloco | Word | Rodadas | Saida |
|---|---:|---:|---:|---:|
| SHA-224 | 512 bits | 32 bits | 64 | 224 bits |
| SHA-256 | 512 bits | 32 bits | 64 | 256 bits |
| SHA-384 | 1024 bits | 64 bits | 80 | 384 bits |
| SHA-512 | 1024 bits | 64 bits | 80 | 512 bits |
| SHA-512/224 | 1024 bits | 64 bits | 80 | 224 bits |
| SHA-512/256 | 1024 bits | 64 bits | 80 | 256 bits |

### SHA-224 vs SHA-256

SHA-224 e basicamente SHA-256 com:

```text
estado inicial diferente
saida truncada para 224 bits
```

Isso e melhor do que simplesmente fazer:

```text
SHA256(mensagem)[0..28]
```

porque o estado inicial diferente faz separacao de dominio entre as variantes.

### SHA-384 vs SHA-512

SHA-384 segue a mesma ideia:

```text
usa o nucleo SHA-512
estado inicial diferente
saida truncada para 384 bits
```

### SHA-512/224 e SHA-512/256

Essas variantes usam o nucleo de 64 bits do SHA-512, mas produzem saidas menores.

Em maquinas 64-bit, SHA-512/256 pode ser interessante: produz 256 bits como SHA-256, mas usa operacoes de 64 bits.

### Vetores SHA-2

```text
SHA224("") =
d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f

SHA256("") =
e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855

SHA384("") =
38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b

SHA512("") =
cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e

SHA512/224("") =
6ed0dd02806fa89e25de060c19d3ac86cabb87d6a0ddd05c333b84f4

SHA512/256("") =
c672b8d1ef56ed28ab87c3622c5114069bdd3ad7b8f9737498d0c01ecef0967a
```

Testes:

```text
cargo test -p encryptor learning_hashes::sha::tests::sha224_vectors
cargo test -p encryptor learning_hashes::sha::tests::sha256_vectors
cargo test -p encryptor learning_hashes::sha::tests::sha384_vectors
cargo test -p encryptor learning_hashes::sha::tests::sha512_vectors
cargo test -p encryptor learning_hashes::sha::tests::sha512_224_vectors
cargo test -p encryptor learning_hashes::sha::tests::sha512_256_vectors
```

## SHA-3

SHA-3 vem do Keccak, escolhido pelo NIST em 2012 apos competicao publica.

Criadores:

```text
Guido Bertoni
Joan Daemen
Michael Peeters
Gilles Van Assche
```

SHA-3 foi padronizado no FIPS 202.

Variantes:

```text
SHA3-224
SHA3-256
SHA3-384
SHA3-512
SHAKE128
SHAKE256
```

SHAKE128 e SHAKE256 sao XOFs: `extendable-output functions`. Elas podem gerar saida de tamanho variavel.

### Sponge

SHA-3 nao usa Merkle-Damgard. Ele usa uma construcao de esponja:

```text
absorver entrada
aplicar permutacao
espremer saida
```

O estado e dividido conceitualmente em:

```text
rate:
  parte usada para absorver e espremer dados

capacity:
  parte reservada para seguranca
```

Essa diferenca estrutural e o grande motivo para estudar SHA-3 mesmo que SHA-256 continue seguro: diversidade de desenho.

## O que mudou de geracao para geracao

```text
MD4 -> SHA-0/SHA-1:
  digest maior, 160 bits
  mais rodadas
  message schedule expandido

SHA-0 -> SHA-1:
  adiciona ROTL1 no message schedule

SHA-1 -> SHA-2:
  mais variantes
  estado maior
  funcoes sigma diferentes
  mais constantes
  opcoes de 32 e 64 bits
  saidas de 224 a 512 bits

SHA-2 -> SHA-3:
  troca Merkle-Damgard por sponge
  vem de competicao publica
  inclui XOFs
```

Resumo pragmatico:

```text
nao use: MD2, MD4, MD5, SHA-0, SHA-1
use normalmente: SHA-256, SHA-512, SHA-3
use para senha: KDF propria para senha, como Argon2id, scrypt ou PBKDF2
```
