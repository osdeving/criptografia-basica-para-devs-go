# A Função $f$

A função $f$ é o núcleo da transformação não linear do DES e é aplicada 16 vezes, uma em cada rodada da cifra. Seu papel é operar sobre a metade direita do bloco atual ($R_{n-1}$) em conjunto com uma subchave de 48 bits ($K_n$), produzindo um bloco intermediário de 32 bits que será misturado com a outra metade do bloco por meio de uma operação XOR.

A expressão geral da rodada é:


$$
R_n = L_{n-1} \oplus f(R_{n-1}, K_n)
$$


A força do DES está concentrada na complexidade da função $f$, que combina expansão de bits, operações XOR, substituição não linear (S-boxes) e uma permutação fixa.

## Etapas internas da função $f$

A função $f$ recebe uma entrada de 32 bits ($R_{n-1}$) e uma subchave de 48 bits ($K_n$). A seguir, aplica as seguintes transformações:

### 1. Expansão (E)
A primeira etapa da função $f$ é a expansão de $R_{n-1}$ de 32 para 48 bits, realizada por uma tabela de permutação conhecida como **E-box**. Essa transformação utiliza uma tabela fixa de 48 posições que reordena e duplica determinados bits da entrada. A expansão é projetada para criar sobreposições entre blocos adjacentes de 4 bits, aumentando a difusão e tornando mais eficaz a mistura posterior com a subchave:

- **Entrada**: bloco de 32 bits (metade direita do bloco original)
- **Tabela**: vetor de 48 índices no intervalo [1, 32]
- **Saída**: bloco de 48 bits, com repetição de alguns bits

Essa expansão permite que bits nos limites de grupos de 4 bits se repitam em mais de um bloco de 6 bits, o que aumenta a interdependência entre as partes processadas pelas S-boxes.

Fazendo um exercício mais simples, suponha uma entrada de 6 bits representada por $[A, B, C, D, E, F]$ e uma tabela de expansão fictícia com 8 posições. Essa tabela indica quais bits da entrada devem compor a saída expandida, podendo haver repetições (no exemplo, as posições 4 e 5 são duplicadas):

$$
\begin{aligned}
\text{Entrada:} &\quad [A, B, C, D, E, F] \\
\text{Tabela:} &\quad [6, 1, 2, 3, 4, 5, 4, 5] \\
\text{Saída:}  &\quad [F, A, B, C, D, E, D, E]
\end{aligned}
$$

Note que os bits $D$ e $E$ aparecem duas vezes na saída, ilustrando o comportamento da expansão E no DES que consiste em: criar sobreposição entre blocos vizinhos, o que contribui para a difusão da informação ao longo do algoritmo.

Uma vez que alguns bits da entrada aparecem mais de uma vez na saída. Esse comportamento reflete o mesmo princípio da expansão real do DES e pode ser matematicamente representado como:


$$
E: \{0,1\}^{32} \rightarrow \{0,1\}^{48}
$$

A _Tabela 1_ a seguir é uma representação da tabela E do DES com seus valores reais de acordo com o FIPS PUB 46-3.

<p align="center"><em>Tabela 1 — Expansão E do DES (32 bits → 48 bits)</em></p>

$$
\begin{array}{cccccc}
32 & 1  & 2  & 3  & 4  & 5 \\
4  & 5  & 6  & 7  & 8  & 9 \\
8  & 9  & 10 & 11 & 12 & 13 \\
12 & 13 & 14 & 15 & 16 & 17 \\
16 & 17 & 18 & 19 & 20 & 21 \\
20 & 21 & 22 & 23 & 24 & 25 \\
24 & 25 & 26 & 27 & 28 & 29 \\
28 & 29 & 30 & 31 & 32 & 1 \\
\end{array}
$$


Essa tabela especifica quais bits da entrada de 32 bits devem ser selecionados, e em que ordem, para formar a saída expandida de 48 bits. Observe que os bits das bordas (como os bits 1, 4, 8, 12, etc.) são usados em mais de uma posição, garantindo a sobreposição entre os blocos de 6 bits que alimentam as S-boxes.

[TODO: considerar criar uma imagem para ilustrar a expansão E]


### 2. XOR com a subchave
A saída de 48 bits da expansão é combinada bit a bit com a subchave $K_n$:

$$
B = E(R_{n-1}) \oplus K_n
$$

O resultado $B$ é um vetor de 48 bits, dividido em oito blocos de 6 bits cada: $B_1, B_2, \dots, B_8$. A geração das subchaves (key schedule) será abordada em uma seção dedicada. Por ora, seguimos com a descrição da próxima etapa: as S-boxes.

### 3. Substituição (S-boxes)

Cada um dos oito blocos $B_i$ (com 6 bits) obtidos após o XOR com a subchave é processado por uma **S-box** ($S_i$), uma tabela de substituição não linear que gera uma saída de 4 bits. Essa etapa introduz não linearidade no DES e aumenta sua resistência a ataques de criptoanálise linear e diferencial.

Formalmente, cada $S_i$ implementa a função:

$$
S_i: \{0,1\}^6 \rightarrow \{0,1\}^4
$$

A estrutura das S-boxes é projetada de forma que os dois bits das extremidades do bloco de 6 bits definem a linha da tabela, enquanto os quatro bits centrais definem a coluna. O valor na interseção corresponde à saída de 4 bits.

Como temos 48 bits após a expansão, divididos em 8 blocos de 6 bits, e utilizamos 8 S-boxes, o resultado da substituição é um bloco de 32 bits, pronto para passar pela permutação $P$.

As tabelas completas das S-boxes, bem como exemplos de sua aplicação e a discussão sobre suas propriedades de segurança, são apresentadas na próxima seção.


### 4. Permutação (P)

Após a aplicação das S-boxes, obtemos um bloco de 32 bits formado pela concatenação das saídas $S_1(B_1), S_2(B_2), \dots, S_8(B_8)$. Esse bloco passa então por uma **permutação fixa** de 32 posições, conhecida como permutação $P$.

O objetivo dessa etapa é **redistribuir os bits de forma controlada**, promovendo maior difusão entre os blocos. Isso aumenta a complexidade de qualquer tentativa de prever como uma mudança em um bit da entrada pode afetar a saída — um dos princípios fundamentais da segurança de uma cifra de bloco.

Formalmente, a permutação $P$ é uma bijeção sobre $\{0,1\}^{32}$ definida por uma tabela de 32 posições que reorganiza os bits conforme índices pré-estabelecidos. A definição dessa permutação está especificada no **FIPS PUB 46-3**.

A seguir, apresentamos a tabela da permutação $P$ do DES:

<p align="center"><em>Tabela 10 — Permutação P do DES (32 bits → 32 bits)</em></p>

$$
\begin{array}{cccccccc}
16 & 7 & 20 & 21 & 29 & 12 & 28 & 17 \\
1  & 15 & 23 & 26 & 5  & 18 & 31 & 10 \\
2  & 8  & 24 & 14 & 32 & 27 & 3  & 9  \\
19 & 13 & 30 & 6  & 22 & 11 & 4  & 25
\end{array}
$$

Cada valor da tabela indica a posição original do bit na entrada que será colocado naquela posição da saída. Por exemplo, o primeiro valor da tabela é 16, o que significa que o **bit 16 da entrada** se tornará o **bit 1 da saída**, e assim por diante.

#### Exemplo: aplicação da Permutação P

Considere um vetor de 32 bits de entrada onde os bits estão numerados de forma incremental apenas para ilustrar o mapeamento da permutação (isto é, o bit na posição 1 tem valor 1, o da posição 2 tem valor 2, e assim por diante):

$$
\text{Entrada (posições): } [1, 2, 3, \dots, 32]
$$

A permutação $P$ do DES reorganiza esses bits de acordo com a seguinte _Tabela 10_.

Assim, a saída da permutação será o seguinte reordenamento dos bits da entrada (valores entre colchetes indicam o bit da posição original que foi movido):

$$
\text{Saída: } [16, 7, 20, 21, 29, 12, 28, 17,\ 1, 15, 23, 26, 5, 18, 31, 10,\ 2, 8, 24, 14, 32, 27, 3, 9,\ 19, 13, 30, 6, 22, 11, 4, 25]
$$

Ou seja, o **bit 16 da entrada original** agora aparece na **posição 1 da saída**, o **bit 7** aparece na posição 2, e assim por diante.

Esse exemplo mostra como a permutação P redistribui os bits da saída das S-boxes, embaralhando-os antes de seguir para a próxima etapa da cifra.


A etapa de permutação P finaliza a função $f$, produzindo o bloco de 32 bits que será combinado com $L_{n-1}$ por meio de uma operação XOR, completando uma rodada da cifra DES.

A Figura 2 ilustra cada etapa discutida acima para a função $f$

<br>
<br>


<p align="center">
    <em>Figura 2 — A função $f$ no algoritmo do DES</em>
    <img src="des-image.svg" alt="Arquitetura geral do algoritmo DES"/>
</p>
<br>
<br>

## Conslusão

A função $f$ é o elemento central de transformação dentro de cada rodada do DES. Sua estrutura combina expansão, mistura com a subchave, substituição não linear via S-boxes e uma permutação de alta difusão. Essa sequência de operações foi projetada para dificultar a análise matemática do algoritmo, mesmo com conhecimento completo de seu funcionamento.

Embora cada etapa da função $f$ seja relativamente simples por si só, a composição entre elas cria uma estrutura robusta que garante a propagação de pequenas diferenças nos blocos de entrada para múltiplos bits na saída. Essa propriedade — conhecida como avalanche — é fundamental para a segurança do DES.

Na próxima seção, aprofundaremos o estudo das **S-boxes**, apresentando suas tabelas completas, exemplos de uso e considerações sobre seu papel na segurança do DES. Em seguida, abordaremos o processo de **key schedule**, responsável por gerar as subchaves de 48 bits utilizadas em cada rodada da cifra.

