# Key Schedule [TODO: criar imagem para a key schedule, revisar texto e didática]

O processo de **key schedule** do DES é responsável por gerar as 16 subchaves de 48 bits utilizadas nas rodadas da cifra. A partir de uma chave principal de 64 bits, uma série de transformações é aplicada para derivar, de forma controlada, uma subchave distinta para cada rodada.

## 1. Remoção de bits de paridade (PC-1)

A chave fornecida ao DES possui 64 bits, mas apenas 56 são efetivamente utilizados para a cifra. Os outros 8 bits (um em cada byte) são reservados para verificação de paridade e são descartados na primeira etapa.

Essa remoção é feita através de uma permutação conhecida como **PC-1 (Permuted Choice 1)**, que reorganiza os bits e elimina os de paridade. O resultado é um bloco de 56 bits:

$$
K^+ = PC1(K) \in \{0,1\}^{56}
$$

## 2. Divisão em blocos C e D

O bloco de 56 bits é então dividido em dois blocos de 28 bits:

- $C_0$: metade esquerda
- $D_0$: metade direita

Esses blocos são manipulados ao longo das 16 rodadas de geração de subchaves.

## 3. Rotações controladas

Em cada rodada $i$, os blocos $C_{i-1}$ e $D_{i-1}$ são rotacionados circularmente para a esquerda. O número de bits a rotacionar depende da rodada e é definido por uma tabela fixa chamada **Rotations**:

$$
\text{Rotations} = [1, 1, 2, 2, 2, 2, 2, 2,\ 1, 2, 2, 2, 2, 2, 2, 1]
$$

Por exemplo:

- Na rodada 1: $C_1 = \text{ROL}(C_0, 1)$, $D_1 = \text{ROL}(D_0, 1)$
- Na rodada 2: $C_2 = \text{ROL}(C_1, 1)$, $D_2 = \text{ROL}(D_1, 1)$
- Da rodada 3 à 8: 2 rotações por rodada
- E assim por diante, conforme a tabela acima.

Essas rotações visam diversificar a disposição dos bits ao longo das rodadas, aumentando a complexidade da relação entre chave e saída da cifra.

#### Exemplo ilustrativo de rotação

Suponha um bloco de 6 bits para fins didáticos:

```text
Entrada: [A, B, C, D, E, F]
Rotacionado 1 à esquerda: [B, C, D, E, F, A]
Rotacionado 2 à esquerda: [C, D, E, F, A, B]
```

## 4. Geração da subchave (PC-2)

Após a rotação dos blocos $C_i$ e $D_i$, eles são concatenados para formar um novo bloco de 56 bits:

$$
CD_i = C_i \parallel D_i
$$

Esse bloco é então submetido a uma segunda permutação, chamada **PC-2 (Permuted Choice 2)**, que seleciona e reorganiza 48 dos 56 bits para formar a subchave da rodada:

$$
K_i = PC2(C_i \parallel D_i) \in \{0,1\}^{48}
$$

Esse procedimento é repetido 16 vezes, gerando as subchaves $K_1, K_2, \dots, K_{16}$, cada uma usada em sua respectiva rodada na função $f$.

A tabela PC-2 será apresentada na próxima seção, juntamente com exemplos de geração de subchaves e observações sobre sua segurança.

