# Família MD — Message Digest

A família **MD** (*Message Digest*) reúne algoritmos de hash criados ou influenciados pelo trabalho de Ronald Rivest. Esses algoritmos foram importantes historicamente porque formaram a base para modelos modernos de funções hash criptográficas.

A sigla **MD** significa *Message Digest*. A ideia central é exatamente a que vimos na introdução: transformar uma mensagem arbitrária em um digest compacto.


```text
mensagem -> algoritmo MD -> digest
```

Atualmente a família MD não é considerada segura. Pelo contrário: MD2, MD4 e MD5 são hoje algoritmos obsoletos para aplicações criptográficas. O valor deles neste livro é histórico e didático. Eles mostram vários conceitos que aparecem em algoritmos posteriores: padding, blocos, estado interno, funções de mistura, S-boxes, operações booleanas, somas modulares, rotações, confusão e difusão.

De forma simplificada:

```text
MD2:
  projetado para máquinas de 8 bits
  usa blocos de 16 bytes
  usa checksum interno
  usa S-box e XOR

MD4:
  projetado para máquinas de 32 bits
  usa blocos de 512 bits
  usa registradores de 32 bits
  usa funções booleanas, soma modular e rotações

MD5:
  evolução mais conservadora do MD4
  mantém saída de 128 bits
  aumenta o número de rodadas
  também acabou quebrado para usos criptográficos
```

A família MD é uma boa forma de consolidar os temas apresentados nos capítulos anteriores e uma boa preparação para enteder os algoritmos SHA que veremos depois. MD4 e MD5, em especial, influenciaram o desenho de SHA-1 e ajudam a entender por que tantos algoritmos de hash trabalham com blocos, words de 32 bits, constantes, rotações e rodadas.

> Importante: os algoritmos deste livro, em especial os deste capítulo, devem ser tratados como material de estudo. Não use MD2, MD4 ou MD5 para segurança moderna.

## MD2

O primeiro algoritmo publicado da família foi o **MD2**, especificado na RFC 1319. Ele produz um digest de 128 bits, isto é, 16 bytes.

O MD2 foi projetado para computadores de 8 bits, o que explica várias decisões internas do algoritmo. Diferente de MD4, MD5 e SHA-1, ele não trabalha naturalmente com words de 32 bits nem usa rotações como operação principal. Em vez disso, o MD2 usa blocos de 16 bytes, uma S-box de 256 posições, operações XOR e um checksum interno anexado à mensagem antes da transformação final.

Suas características principais são:

```text
saída:
  128 bits, ou 16 bytes

entrada:
  tamanho variável

bloco interno:
  16 bytes

operações principais:
  XOR e substituição por S-box

arquitetura-alvo:
  máquinas de 8 bits

status atual:
  obsoleto para segurança moderna
```

O MD2 possui três etapas principais:

```text
1. padding
2. checksum
3. transformação principal
```

Veremos cada uma delas a seguir.

## 1. Padding no MD2

Como o MD2 processa a mensagem em blocos de 16 bytes, ele precisa garantir que a entrada tenha tamanho múltiplo de 16.

A regra de padding é simples:

```text
se faltam N bytes para completar o bloco,
adicionamos N bytes, cada um com o valor N
```

Se a mensagem já tiver tamanho múltiplo de 16, adicionamos um bloco inteiro de padding:

```text
16 bytes, cada um com o valor 0x10
```

Exemplo com a entrada `"ABC"`:

```text
"A" = 0x41
"B" = 0x42
"C" = 0x43
```

A mensagem tem 3 bytes. Para completar um bloco de 16 bytes, faltam 13 bytes.

Portanto:

```text
padding necessário = 16 - 3 = 13
valor do padding   = 0x0D
```
A tabela a seguir sumariza o processo:


| Etapa | Dados |
|---|---|
| Entrada original | `41 42 43`, isto é, `"ABC"` |
| Tamanho original | `3 bytes` |
| Padding necessário | `16 - 3 = 13 bytes` |
| Byte de padding | `0x0D` |
| Resultado final | `41 42 43 0D 0D 0D 0D 0D 0D 0D 0D 0D 0D 0D 0D 0D` |

Se a entrada tivesse 15 bytes, o padding seria:

```text
01
```

Se a entrada tivesse exatamente 16 bytes, o padding seria:

```text
10 10 10 10 10 10 10 10 10 10 10 10 10 10 10 10
```

Esse padrão de padding é  usado em alguns modos de cifragem por bloco: sempre se adiciona pelo menos um byte, e o valor do byte indica quantos bytes de padding foram adicionados.

## 2. Checksum no MD2

Depois do padding, o MD2 calcula um **checksum de 16 bytes** sobre a mensagem preenchida. Esse checksum é anexado ao final da mensagem antes da transformação principal.

Aqui é importante não confundir esse checksum com aquele visto no capítulo anterior. No MD2, o checksum faz parte da estrutura criptográfica do algoritmo. Ele adiciona dependência extra entre os blocos e dificulta a construção de colisões simples.

O cálculo usa uma tabela de substituição chamada **S-box**. A S-box do MD2 possui 256 valores e é definida na especificação do algoritmo.

Uma **S-box** (*Substitution Box*), conforme vimos no capítulo 1, é uma tabela de substituição. Ela recebe um valor de entrada e devolve outro valor. Em criptografia, S-boxes são usadas para introduzir não linearidade e aumentar a **confusão** tornando mais difícil estabelecer relações simples entre entrada e saída.

O checksum do MD2 é atualizado bloco por bloco. Para cada bloco de 16 bytes, percorremos os bytes e atualizamos o vetor `C`.

A fórmula de atualização do checksum é:

$$
C[j] \leftarrow C[j] \oplus S[M[16i + j] \oplus L]
$$

Em seguida, atualizamos o acumulador:

$$
L \leftarrow C[j]
$$

Onde:

$$
\begin{aligned}
C[j] &:= \text{byte } j \text{ do checksum} \\
M[16i + j] &:= \text{byte } j \text{ do bloco } i \text{ da mensagem} \\
S[x] &:= \text{valor da S-box na posição } x \\
L &:= \text{último byte atualizado do checksum} \\
\oplus &:= \text{operação XOR}
\end{aligned}
$$

O valor inicial de `L` é zero.

Exemplo simplificado:

$$
C[3] = \texttt{0xA5}
$$

$$
M[3] = \texttt{0x7F}
$$

$$
L = \texttt{0x2C}
$$

Primeiro calculamos o índice usado na S-box:

$$
M[3] \oplus L = \texttt{0x7F} \oplus \texttt{0x2C} = \texttt{0x53}
$$

Suponha que a S-box tenha o seguinte valor nessa posição:

$$
S[\texttt{0x53}] = \texttt{0xD4}
$$

Então atualizamos o byte do checksum:

$$
\begin{aligned}
C[3]
&\leftarrow C[3] \oplus S[\texttt{0x53}] \\
&= \texttt{0xA5} \oplus \texttt{0xD4} \\
&= \texttt{0x71}
\end{aligned}
$$

Depois, atualizamos o acumulador `L`:

$$
L \leftarrow C[3] = \texttt{0x71}
$$

Ao final do processo, os 16 bytes do checksum são anexados à mensagem.

```text
mensagem com padding -> calcula checksum -> mensagem com padding + checksum
```

Essa etapa reforça a ideia que vimos na introdução: uma hash criptográfica não pode ser apenas determinística e ter saída fixa. Ela precisa misturar a entrada de forma que alterações locais se espalhem e dificultem a construção de substitutos ou colisões.

### Implementação do padding e checksum do MD2 em Rust

Abaixo está uma implementação didática do padding e do checksum do MD2 em Rust.

```rust
const S: [u8; 256] = [
    41, 46, 67, 201, 162, 216, 124, 1, 61, 54, 84, 161, 236, 240, 6, 19,
    98, 167, 5, 243, 192, 199, 115, 140, 152, 147, 43, 217, 188, 76, 130, 202,
    30, 155, 87, 60, 253, 212, 224, 22, 103, 66, 111, 24, 138, 23, 229, 18,
    190, 78, 196, 214, 218, 158, 222, 73, 160, 251, 245, 142, 187, 47, 238, 122,
    169, 104, 121, 145, 21, 178, 7, 63, 148, 194, 16, 137, 11, 34, 95, 33,
    128, 127, 93, 154, 90, 144, 50, 39, 53, 62, 204, 231, 191, 247, 151, 3,
    255, 25, 48, 179, 72, 165, 181, 209, 215, 94, 146, 42, 172, 86, 170, 198,
    79, 184, 56, 210, 150, 164, 125, 182, 118, 252, 107, 226, 156, 116, 4, 241,
    69, 157, 112, 89, 100, 113, 135, 32, 134, 91, 207, 101, 230, 45, 168, 2,
    27, 96, 37, 173, 174, 176, 185, 246, 28, 70, 97, 105, 52, 64, 126, 15,
    85, 71, 163, 35, 221, 81, 175, 58, 195, 92, 249, 206, 186, 197, 234, 38,
    44, 83, 13, 110, 133, 40, 132, 9, 211, 223, 205, 244, 65, 129, 77, 82,
    106, 220, 55, 200, 108, 193, 171, 250, 36, 225, 123, 8, 12, 189, 177, 74,
    120, 136, 149, 139, 227, 99, 232, 109, 233, 203, 213, 254, 59, 0, 29, 57,
    242, 239, 183, 14, 102, 88, 208, 228, 166, 119, 114, 248, 235, 117, 75, 10,
    49, 68, 80, 180, 143, 237, 31, 26, 219, 153, 141, 51, 159, 17, 131, 20,
];

fn md2_padding(input: &[u8]) -> Vec<u8> {
    let padding_len = 16 - (input.len() % 16);

    let mut message = input.to_vec();

    for _ in 0..padding_len {
        message.push(padding_len as u8);
    }

    message
}

fn md2_checksum(message: &[u8]) -> [u8; 16] {
    let mut checksum = [0u8; 16];
    let mut l = 0u8;

    for block in message.chunks_exact(16) {
        for j in 0..16 {
            let index = (block[j] ^ l) as usize;
            checksum[j] ^= S[index];
            l = checksum[j];
        }
    }

    checksum
}
```

## 3. Transformação principal do MD2

Depois do padding e do checksum, o MD2 processa a mensagem final usando um buffer intermediário de 48 bytes chamado `X`.

Esse buffer é dividido em três partes de 16 bytes:

```text
X[0..16]: estado atual do digest

X[16..32]: bloco atual da mensagem

X[32..48]: mistura entre estado atual e bloco atual
```

A estrutura é:

```text
X = [
  H(0) H(1) ... H(15)
  M(0) M(1) ... M(15)
  H(0) XOR M(0), ..., H(15) XOR M(15)
]
```

Ou, de forma compacta:

$$
X[32+j] = X[j] \oplus X[16+j]
$$

A cada bloco de 16 bytes da mensagem, o MD2 faz:

```text
1. copia o bloco para X[16..32]
2. preenche X[32..48] com X[0..16] XOR X[16..32]
3. executa 18 rodadas de mistura
4. mantém X[0..16] como estado atualizado
```

As 18 rodadas aplicam XOR com valores da S-box. Em cada rodada, percorremos os 48 bytes do buffer e atualizamos um acumulador `t`.

A operação central é:

$$
X[j] = X[j] \oplus S[t]
$$

Depois:

$$
t = X[j]
$$

Ao final de cada rodada:

$$
t = t + round \pmod{256}
$$

Essa transformação mostra um padrão importante: a hash mantém um estado interno e o atualiza repetidamente. A segurança não vem de uma única operação mágica, mas da repetição de operações simples que visando produzir confusão e difusão.

### Implementação completa do MD2 em Rust

```rust
const S: [u8; 256] = [
    41, 46, 67, 201, 162, 216, 124, 1, 61, 54, 84, 161, 236, 240, 6, 19,
    98, 167, 5, 243, 192, 199, 115, 140, 152, 147, 43, 217, 188, 76, 130, 202,
    30, 155, 87, 60, 253, 212, 224, 22, 103, 66, 111, 24, 138, 23, 229, 18,
    190, 78, 196, 214, 218, 158, 222, 73, 160, 251, 245, 142, 187, 47, 238, 122,
    169, 104, 121, 145, 21, 178, 7, 63, 148, 194, 16, 137, 11, 34, 95, 33,
    128, 127, 93, 154, 90, 144, 50, 39, 53, 62, 204, 231, 191, 247, 151, 3,
    255, 25, 48, 179, 72, 165, 181, 209, 215, 94, 146, 42, 172, 86, 170, 198,
    79, 184, 56, 210, 150, 164, 125, 182, 118, 252, 107, 226, 156, 116, 4, 241,
    69, 157, 112, 89, 100, 113, 135, 32, 134, 91, 207, 101, 230, 45, 168, 2,
    27, 96, 37, 173, 174, 176, 185, 246, 28, 70, 97, 105, 52, 64, 126, 15,
    85, 71, 163, 35, 221, 81, 175, 58, 195, 92, 249, 206, 186, 197, 234, 38,
    44, 83, 13, 110, 133, 40, 132, 9, 211, 223, 205, 244, 65, 129, 77, 82,
    106, 220, 55, 200, 108, 193, 171, 250, 36, 225, 123, 8, 12, 189, 177, 74,
    120, 136, 149, 139, 227, 99, 232, 109, 233, 203, 213, 254, 59, 0, 29, 57,
    242, 239, 183, 14, 102, 88, 208, 228, 166, 119, 114, 248, 235, 117, 75, 10,
    49, 68, 80, 180, 143, 237, 31, 26, 219, 153, 141, 51, 159, 17, 131, 20,
];

fn md2(input: &[u8]) -> [u8; 16] {
    let mut message = apply_md2_padding(input);

    let checksum = compute_md2_checksum(&message);
    message.extend_from_slice(&checksum);

    process_md2_blocks(&message)
}

fn apply_md2_padding(input: &[u8]) -> Vec<u8> {
    let padding_len = 16 - (input.len() % 16);

    let mut message = input.to_vec();

    for _ in 0..padding_len {
        message.push(padding_len as u8);
    }

    message
}

fn compute_md2_checksum(message: &[u8]) -> [u8; 16] {
    let mut checksum = [0u8; 16];
    let mut l = 0u8;

    for block in message.chunks_exact(16) {
        for j in 0..16 {
            let index = (block[j] ^ l) as usize;
            checksum[j] ^= S[index];
            l = checksum[j];
        }
    }

    checksum
}

fn process_md2_blocks(message: &[u8]) -> [u8; 16] {
    let mut x = [0u8; 48];

    for block in message.chunks_exact(16) {
        for j in 0..16 {
            x[16 + j] = block[j];
            x[32 + j] = x[j] ^ x[16 + j];
        }

        let mut t = 0u8;

        for round in 0..18 {
            for j in 0..48 {
                x[j] ^= S[t as usize];
                t = x[j];
            }

            t = t.wrapping_add(round as u8);
        }
    }

    let mut digest = [0u8; 16];
    digest.copy_from_slice(&x[0..16]);
    digest
}

fn to_hex(bytes: &[u8]) -> String {
    let mut output = String::new();

    for byte in bytes {
        output.push_str(&format!("{byte:02x}"));
    }

    output
}

fn main() {
    let digest = md2(b"abc");

    println!("{}", to_hex(&digest));
    println!("esperado: da853b0d3f88d99b30283a69e6ded6bb");
}
```

Para a entrada `"abc"`, o MD2 deve produzir:

```text
da853b0d3f88d99b30283a69e6ded6bb
```

## Propriedades aplicadas no MD2

Comparado ao `simple_hash` visto na introdução, o MD2 já incorpora mecanismos típicos de uma função hash criptográfica utilizando os seguintes conceitos:

```text
determinismo:
  a mesma mensagem sempre gera o mesmo digest

saída fixa:
  o digest sempre possui 128 bits

padding:
  a mensagem é ajustada para o tamanho de bloco do algoritmo

estado interno:
  o buffer X mantém o estado intermediário

confusão:
  a S-box dificulta relações lineares simples entre entrada e saída

difusão:
  as rodadas propagam alterações locais pelo estado interno

avalanche:
  pequenas mudanças na entrada tendem a alterar muitos bits do digest final
```

Contudo, MD2 não é considerado seguro para uso moderno. Ele é útil aqui porque mostra a evolução histórica dos algoritmos MD e a partir dele poderemos enxergar a transição para MD4, que troca a lógica voltada a 8 bits por um algoritmo voltado para arquiteturas de 32 bits.

## MD4: A evolução do Message Digest

O **MD4** foi desenvolvido por Ronald Rivest em 1990 como uma evolução mais rápida da família MD. Enquanto o MD2 foi pensado para máquinas de 8 bits, o MD4 foi projetado para processadores de 32 bits.

Essa mudança altera a aparência do algoritmo e suas características são as seguintes:

```text
blocos de 512 bits
words de 32 bits
quatro registradores principais
soma modular
funções booleanas
rotações à esquerda
três rodadas principais
```

O digest final continua tendo 128 bits, mas a estrutura interna é muito diferente da do MD2.

O MD4 é importante historicamente porque influenciou diretamente o MD5 e também a faz parte da linguagem que leva ao SHA-1. Porém, a exemplo do MD2, ele é criptograficamente fraco e não deve ser usado para segurança.

## Estrutura geral do MD4

O MD4 segue quatro etapas principais:

```text
1. padding da mensagem
2. inicialização do estado
3. processamento em blocos de 512 bits
4. concatenação do estado final
```

## 1. Padding no MD4

O MD4 processa blocos de 512 bits, isto é, 64 bytes.

A mensagem é preenchida da seguinte forma:

```text
1. adiciona-se um bit 1
2. adicionam-se bits 0 até que faltem 64 bits para completar o bloco
3. os últimos 64 bits armazenam o tamanho original da mensagem em bits
```

Na prática, em bytes, isso começa com:

```text
0x80
```

porque `0x80` em binário é:

```text
10000000
```

Depois vêm bytes `0x00` até que o tamanho da mensagem seja congruente a 56 bytes módulo 64. Os últimos 8 bytes guardam o tamanho original da mensagem em bits, usando little-endian.

Em notação:

$$
\text{len}(mensagem\_preenchida) \equiv 56 \pmod{64}
$$

antes de adicionar os 8 bytes finais de tamanho.

Depois dos 8 bytes finais:

$$
\text{len}(mensagem\_final) \equiv 0 \pmod{64}
$$

Exemplo conceitual com `"abc"`:

```text
"abc" = 3 bytes = 24 bits

mensagem:
61 62 63

adiciona 0x80:
61 62 63 80

adiciona zeros até chegar a 56 bytes:
61 62 63 80 00 00 00 ... 00

adiciona o tamanho original em bits, little-endian:
18 00 00 00 00 00 00 00
```

O valor `0x18` é 24 em decimal.

## 2. Inicialização do estado

O MD4 usa quatro registradores de 32 bits:

```text
A
B
C
D
```

Eles começam com valores fixos:

```text
A = 0x67452301
B = 0xefcdab89
C = 0x98badcfe
D = 0x10325476
```

Esses quatro registradores formam o estado interno do algoritmo.

```text
estado = A || B || C || D
```

Como cada registrador tem 32 bits:

$$
4 \cdot 32 = 128
$$

Portanto, ao final do processamento, a concatenação de `A`, `B`, `C` e `D` forma o digest de 128 bits.

## 3. Funções booleanas do MD4

O MD4 usa três funções booleanas principais: `F`, `G` e `H`.

### Função F

$$
F(X, Y, Z) = (X \land Y) \lor (\lnot X \land Z)
$$

Em Rust:

```rust
fn f(x: u32, y: u32, z: u32) -> u32 {
    (x & y) | (!x & z)
}
```

Essa função escolhe bits de `Y` ou `Z` dependendo dos bits de `X`.

Se um bit de `X` é `1`, o resultado tende a pegar o bit correspondente de `Y`.

Se um bit de `X` é `0`, o resultado tende a pegar o bit correspondente de `Z`.

### Função G

$$
G(X, Y, Z) = (X \land Y) \lor (X \land Z) \lor (Y \land Z)
$$

Em Rust:

```rust
fn g(x: u32, y: u32, z: u32) -> u32 {
    (x & y) | (x & z) | (y & z)
}
```

Essa função se comporta como uma votação bit a bit. Para cada posição, se pelo menos dois dos três bits forem `1`, o resultado será `1`.

### Função H

$$
H(X, Y, Z) = X \oplus Y \oplus Z
$$

Em Rust:

```rust
fn h(x: u32, y: u32, z: u32) -> u32 {
    x ^ y ^ z
}
```

Essa função mistura os bits usando XOR.

## 4. Rodadas do MD4

Cada bloco de 512 bits é dividido em 16 words de 32 bits:

```text
X[0], X[1], ..., X[15]
```

Essas words são lidas em little-endian.

Depois, o bloco passa por três rodadas:

```text
Rodada 1:
  usa F
  não usa constante adicional

Rodada 2:
  usa G
  adiciona a constante 0x5a827999

Rodada 3:
  usa H
  adiciona a constante 0x6ed9eba1
```

A operação básica de uma rodada é:

$$
A = (A + F(B,C,D) + X[k]) \lll s
$$

O símbolo:

$$
\lll
$$

representa rotação à esquerda.

Em Rust, usamos:

```rust
value.rotate_left(s)
```

A soma é modular em 32 bits. Em Rust, usamos:

```rust
wrapping_add
```

porque, em algoritmos criptográficos desse tipo, o overflow não é erro: ele faz parte da operação matemática.

## Implementação do MD4 em Rust

A implementação abaixo mostra a estrutura do MD4 sem depender de bibliotecas externas. Ideal para entender na prática os mecanismos internos do algorítmo.

```rust
fn f(x: u32, y: u32, z: u32) -> u32 {
    (x & y) | (!x & z)
}

fn g(x: u32, y: u32, z: u32) -> u32 {
    (x & y) | (x & z) | (y & z)
}

fn h(x: u32, y: u32, z: u32) -> u32 {
    x ^ y ^ z
}

fn round1(a: &mut u32, b: u32, c: u32, d: u32, x: u32, s: u32) {
    *a = a
        .wrapping_add(f(b, c, d))
        .wrapping_add(x)
        .rotate_left(s);
}

fn round2(a: &mut u32, b: u32, c: u32, d: u32, x: u32, s: u32) {
    *a = a
        .wrapping_add(g(b, c, d))
        .wrapping_add(x)
        .wrapping_add(0x5a82_7999)
        .rotate_left(s);
}

fn round3(a: &mut u32, b: u32, c: u32, d: u32, x: u32, s: u32) {
    *a = a
        .wrapping_add(h(b, c, d))
        .wrapping_add(x)
        .wrapping_add(0x6ed9_eba1)
        .rotate_left(s);
}

fn md4(input: &[u8]) -> [u8; 16] {
    let mut message = input.to_vec();

    let bit_len = (message.len() as u64) * 8;

    message.push(0x80);

    while message.len() % 64 != 56 {
        message.push(0x00);
    }

    message.extend_from_slice(&bit_len.to_le_bytes());

    let mut a0: u32 = 0x6745_2301;
    let mut b0: u32 = 0xefcd_ab89;
    let mut c0: u32 = 0x98ba_dcfe;
    let mut d0: u32 = 0x1032_5476;

    for block in message.chunks_exact(64) {
        let mut x = [0u32; 16];

        for i in 0..16 {
            let start = i * 4;

            x[i] = u32::from_le_bytes([
                block[start],
                block[start + 1],
                block[start + 2],
                block[start + 3],
            ]);
        }

        let mut a = a0;
        let mut b = b0;
        let mut c = c0;
        let mut d = d0;

        round1(&mut a, b, c, d, x[0], 3);
        round1(&mut d, a, b, c, x[1], 7);
        round1(&mut c, d, a, b, x[2], 11);
        round1(&mut b, c, d, a, x[3], 19);

        round1(&mut a, b, c, d, x[4], 3);
        round1(&mut d, a, b, c, x[5], 7);
        round1(&mut c, d, a, b, x[6], 11);
        round1(&mut b, c, d, a, x[7], 19);

        round1(&mut a, b, c, d, x[8], 3);
        round1(&mut d, a, b, c, x[9], 7);
        round1(&mut c, d, a, b, x[10], 11);
        round1(&mut b, c, d, a, x[11], 19);

        round1(&mut a, b, c, d, x[12], 3);
        round1(&mut d, a, b, c, x[13], 7);
        round1(&mut c, d, a, b, x[14], 11);
        round1(&mut b, c, d, a, x[15], 19);

        round2(&mut a, b, c, d, x[0], 3);
        round2(&mut d, a, b, c, x[4], 5);
        round2(&mut c, d, a, b, x[8], 9);
        round2(&mut b, c, d, a, x[12], 13);

        round2(&mut a, b, c, d, x[1], 3);
        round2(&mut d, a, b, c, x[5], 5);
        round2(&mut c, d, a, b, x[9], 9);
        round2(&mut b, c, d, a, x[13], 13);

        round2(&mut a, b, c, d, x[2], 3);
        round2(&mut d, a, b, c, x[6], 5);
        round2(&mut c, d, a, b, x[10], 9);
        round2(&mut b, c, d, a, x[14], 13);

        round2(&mut a, b, c, d, x[3], 3);
        round2(&mut d, a, b, c, x[7], 5);
        round2(&mut c, d, a, b, x[11], 9);
        round2(&mut b, c, d, a, x[15], 13);

        round3(&mut a, b, c, d, x[0], 3);
        round3(&mut d, a, b, c, x[8], 9);
        round3(&mut c, d, a, b, x[4], 11);
        round3(&mut b, c, d, a, x[12], 15);

        round3(&mut a, b, c, d, x[2], 3);
        round3(&mut d, a, b, c, x[10], 9);
        round3(&mut c, d, a, b, x[6], 11);
        round3(&mut b, c, d, a, x[14], 15);

        round3(&mut a, b, c, d, x[1], 3);
        round3(&mut d, a, b, c, x[9], 9);
        round3(&mut c, d, a, b, x[5], 11);
        round3(&mut b, c, d, a, x[13], 15);

        round3(&mut a, b, c, d, x[3], 3);
        round3(&mut d, a, b, c, x[11], 9);
        round3(&mut c, d, a, b, x[7], 11);
        round3(&mut b, c, d, a, x[15], 15);

        a0 = a0.wrapping_add(a);
        b0 = b0.wrapping_add(b);
        c0 = c0.wrapping_add(c);
        d0 = d0.wrapping_add(d);
    }

    let mut digest = [0u8; 16];

    digest[0..4].copy_from_slice(&a0.to_le_bytes());
    digest[4..8].copy_from_slice(&b0.to_le_bytes());
    digest[8..12].copy_from_slice(&c0.to_le_bytes());
    digest[12..16].copy_from_slice(&d0.to_le_bytes());

    digest
}

fn to_hex(bytes: &[u8]) -> String {
    let mut output = String::new();

    for byte in bytes {
        output.push_str(&format!("{byte:02x}"));
    }

    output
}

fn main() {
    let digest = md4(b"abc");

    println!("{}", to_hex(&digest));
    println!("esperado: a448017aaf21d8525fc10ae87aa6729d");
}
```

Para a entrada `"abc"`, o MD4 deve produzir:

```text
a448017aaf21d8525fc10ae87aa6729d
```

## O que o MD4 ensina

O MD4 é muito mais próximo da família SHA-1/SHA-2 do que o MD2.

Ele introduz um padrão que veremos de novo:

```text
padding com tamanho da mensagem
processamento em blocos de 512 bits
estado interno de words de 32 bits
funções booleanas
soma modular
rotações
várias rodadas de mistura
digest final como concatenação do estado
```

Esse modelo mostra bem o papel de **confusão** e **difusão**.

As funções booleanas `F`, `G` e `H` ajudam a esconder relações simples entre os bits. As rotações e a ordem das words espalham a influência dos bits da mensagem pelo estado interno.

Mas o MD4 também mostra uma lição importante: ser rápido demais pode ser perigoso quando a margem de segurança é pequena. O MD4 foi desenhado para ser extremamente eficiente, mas suas três rodadas não foram suficientes para resistir a ataques criptanalíticos posteriores.

## Ataques contra MD4

O MD4 foi um avanço importante na época, mas acabou se mostrando frágil. Ataques diferenciais exploraram a estrutura interna do algoritmo e tornaram possível encontrar colisões de maneira prática.

Conforme vimos na introdução, uma colisão significa encontrar duas mensagens diferentes, $M_1$ e $M_2$, tais que:

$$
M_1 \neq M_2
$$

e:

$$
H(M_1) = H(M_2)
$$

Isso não é a mesma coisa que uma segunda pré-imagem.

Na **segunda pré-imagem**, o atacante recebe uma mensagem específica $M_1$ e tenta encontrar outra mensagem $M_2$ com o mesmo digest.

Na **colisão**, o atacante pode construir ou escolher os dois lados do par. O requisito da função hash é que até esse caso mais geral seja computacionalmente inviável.

Quando colisões em MD4 se tornaram práticas, ele deixou de ser adequado para usos como:

```text
assinaturas digitais
certificados
integridade de pacotes
identificação criptográfica de documentos
```

## MD4 e senhas NTLM

O MD4 aparece historicamente no contexto de autenticação do Windows. O chamado **NT hash**, usado em NTLM, é essencialmente o MD4 da senha codificada em UTF-16LE.

De forma simplificada:

```text
senha -> UTF-16LE -> MD4 -> NT hash
```

O problema principal, nesse contexto, não é apenas a existência de colisões. Para senhas, o problema também é que MD4 é extremamente rápido e não foi projetado como função de derivação de senha. Isso facilita ataques de força bruta e dicionário, especialmente quando as senhas são curtas, previsíveis ou reutilizadas.

Para senhas, é necessário usar salt e funções próprias para armazenamento de senhas, como Argon2, bcrypt ou scrypt. O salt dificulta ataques com rainbow tables, enquanto essas funções tornam cada tentativa mais cara, desencorajando computacionalmente ataques de força bruta e dicionário.

## MD5: A próxima evolução

O **MD5** surgiu como uma evolução mais conservadora do MD4. Ele manteve várias ideias da mesma linhagem:

```text
digest de 128 bits
blocos de 512 bits
estado de quatro registradores de 32 bits
soma modular
funções booleanas
rotações
padding com tamanho da mensagem
```

Mas o MD5 adicionou mais rodadas e modificou as funções e constantes internas para tentar aumentar a margem de segurança em relação ao MD4.

Por muito tempo, MD5 foi amplamente usado para verificar arquivos, identificar conteúdos e construir sistemas que precisavam de um digest curto e rápido.

Hoje, porém, MD5 também é considerado inseguro para aplicações criptográficas. Colisões práticas tornaram o algoritmo inadequado para assinaturas digitais, certificados e qualquer cenário em que um atacante possa tentar construir dois conteúdos diferentes com o mesmo digest.

Ainda assim, estudar MD5 é útil porque ele mostra claramente a transição entre MD4 e SHA-1. Muitos elementos que aparecem em SHA-1 ficam mais fáceis de entender depois de observar a família MD.

Nas próximas seções, veremos o MD5 com mais detalhe e depois avançaremos para a família SHA.
