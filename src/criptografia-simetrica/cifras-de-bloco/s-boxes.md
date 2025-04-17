# Substituição (S-boxes)

Cada um dos oito blocos $B_i$ (com 6 bits) obtidos após o XOR com a subchave é passado por uma **S-box** ($S_i$), uma tabela de substituição não linear que gera uma saída de 4 bits. Essas caixas, conhecidas como **S-boxes** (*Substitution Boxes*), introduzem não linearidade no processo de cifragem, dificultando a aplicação de técnicas como criptoanálise linear e diferencial.

Formalmente, cada $S_i$ implementa uma função:

$$
S_i: \{0,1\}^6 \rightarrow \{0,1\}^4
$$

A entrada da função consiste em 6 bits: os **bits das extremidades** (1º e 6º) determinam a **linha** da tabela ($0 \leq \text{linha} < 4$), enquanto os **quatro bits centrais** determinam a **coluna** ($0 \leq \text{coluna} < 16$). O valor na interseção da linha e coluna define a saída de 4 bits.

Como a entrada expandida para a função $f$ possui 48 bits, ela é dividida em 8 blocos de 6 bits. Cada bloco é processado por uma S-box distinta ($S_1$ a $S_8$), produzindo ao final um total de 32 bits de saída.

As S-boxes foram projetadas com critérios de segurança cuidadosamente escolhidos, embora esses detalhes só tenham sido tornados públicos anos depois da padronização do DES. Posteriormente, revelou-se que as tabelas foram ajustadas pela NSA com o objetivo de aumentar a resistência à **criptoanálise diferencial**, uma técnica que ainda não havia sido publicada na época, o que levou a questionamentos sobre a transparência do processo.

A seguir, apresentamos as oito S-boxes do DES ($S_1$ a $S_8$), conforme especificado no **FIPS PUB 46-3**. Cada tabela possui 4 linhas e 16 colunas, cobrindo todas as possíveis combinações de entrada de 6 bits.

<p align="center"><em>Tabela 2 — S-box 1 do DES (S₁: 6 bits → 4 bits)</em></p>

$$
\begin{array}{cccccccccccccccc}
   & 0 & 1 & 2 & 3 & 4 & 5 & 6 & 7 & 8 & 9 & 10 & 11 & 12 & 13 & 14 & 15 \\
0: & 14 & 4 & 13 & 1 & 2 & 15 & 11 & 8 & 3 & 10 & 6 & 12 & 5 & 9 & 0 & 7 \\
1: & 0 & 15 & 7 & 4 & 14 & 2 & 13 & 1 & 10 & 6 & 12 & 11 & 9 & 5 & 3 & 8 \\
2: & 4 & 1 & 14 & 8 & 13 & 6 & 2 & 11 & 15 & 12 & 9 & 7 & 3 & 10 & 5 & 0 \\
3: & 15 & 12 & 8 & 2 & 4 & 9 & 1 & 7 & 5 & 11 & 3 & 14 & 10 & 0 & 6 & 13
\end{array}
$$

<p align="center"><em>Tabela 3 — S-box 2 do DES (S₂: 6 bits → 4 bits)</em></p>

$$
\begin{array}{cccccccccccccccc}
   & 0 & 1 & 2 & 3 & 4 & 5 & 6 & 7 & 8 & 9 & 10 & 11 & 12 & 13 & 14 & 15 \\
0: & 15 & 1 & 8 & 14 & 6 & 11 & 3 & 4 & 9 & 7 & 2 & 13 & 12 & 0 & 5 & 10 \\
1: & 3 & 13 & 4 & 7 & 15 & 2 & 8 & 14 & 12 & 0 & 1 & 10 & 6 & 9 & 11 & 5 \\
2: & 0 & 14 & 7 & 11 & 10 & 4 & 13 & 1 & 5 & 8 & 12 & 6 & 9 & 3 & 2 & 15 \\
3: & 13 & 8 & 10 & 1 & 3 & 15 & 4 & 2 & 11 & 6 & 7 & 12 & 0 & 5 & 14 & 9
\end{array}
$$

<p align="center"><em>Tabela 4 — S-box 3 do DES (S₃: 6 bits → 4 bits)</em></p>

$$
\begin{array}{cccccccccccccccc}
   & 0 & 1 & 2 & 3 & 4 & 5 & 6 & 7 & 8 & 9 & 10 & 11 & 12 & 13 & 14 & 15 \\
0: & 10 & 0 & 9 & 14 & 6 & 3 & 15 & 5 & 1 & 13 & 12 & 7 & 11 & 4 & 2 & 8 \\
1: & 13 & 7 & 0 & 9 & 3 & 4 & 6 & 10 & 2 & 8 & 5 & 14 & 12 & 11 & 15 & 1 \\
2: & 13 & 6 & 4 & 9 & 8 & 15 & 3 & 0 & 11 & 1 & 2 & 12 & 5 & 10 & 14 & 7 \\
3: & 1 & 10 & 13 & 0 & 6 & 9 & 8 & 7 & 4 & 15 & 14 & 3 & 11 & 5 & 2 & 12
\end{array}
$$

<p align="center"><em>Tabela 5 — S-box 4 do DES (S₄: 6 bits → 4 bits)</em></p>

$$
\begin{array}{cccccccccccccccc}
   & 0 & 1 & 2 & 3 & 4 & 5 & 6 & 7 & 8 & 9 & 10 & 11 & 12 & 13 & 14 & 15 \\
0: & 7 & 13 & 14 & 3 & 0 & 6 & 9 & 10 & 1 & 2 & 8 & 5 & 11 & 12 & 4 & 15 \\
1: & 13 & 8 & 11 & 5 & 6 & 15 & 0 & 3 & 4 & 7 & 2 & 12 & 1 & 10 & 14 & 9 \\
2: & 10 & 6 & 9 & 0 & 12 & 11 & 7 & 13 & 15 & 1 & 3 & 14 & 5 & 2 & 8 & 4 \\
3: & 3 & 15 & 0 & 6 & 10 & 1 & 13 & 8 & 9 & 4 & 5 & 11 & 12 & 7 & 2 & 14
\end{array}
$$

<p align="center"><em>Tabela 6 — S-box 5 do DES (S₅: 6 bits → 4 bits)</em></p>

$$
\begin{array}{cccccccccccccccc}
   & 0 & 1 & 2 & 3 & 4 & 5 & 6 & 7 & 8 & 9 & 10 & 11 & 12 & 13 & 14 & 15 \\
0: & 2 & 12 & 4 & 1 & 7 & 10 & 11 & 6 & 8 & 5 & 3 & 15 & 13 & 0 & 14 & 9 \\
1: & 14 & 11 & 2 & 12 & 4 & 7 & 13 & 1 & 5 & 0 & 15 & 10 & 3 & 9 & 8 & 6 \\
2: & 4 & 2 & 1 & 11 & 10 & 13 & 7 & 8 & 15 & 9 & 12 & 5 & 6 & 3 & 0 & 14 \\
3: & 11 & 8 & 12 & 7 & 1 & 14 & 2 & 13 & 6 & 15 & 0 & 9 & 10 & 4 & 5 & 3
\end{array}
$$

<p align="center"><em>Tabela 7 — S-box 6 do DES (S₆: 6 bits → 4 bits)</em></p>

$$
\begin{array}{cccccccccccccccc}
   & 0 & 1 & 2 & 3 & 4 & 5 & 6 & 7 & 8 & 9 & 10 & 11 & 12 & 13 & 14 & 15 \\
0: & 12 & 1 & 10 & 15 & 9 & 2 & 6 & 8 & 0 & 13 & 3 & 4 & 14 & 7 & 5 & 11 \\
1: & 10 & 15 & 4 & 2 & 7 & 12 & 9 & 5 & 6 & 1 & 13 & 14 & 0 & 11 & 3 & 8 \\
2: & 9 & 14 & 15 & 5 & 2 & 8 & 12 & 3 & 7 & 0 & 4 & 10 & 1 & 13 & 11 & 6 \\
3: & 4 & 3 & 2 & 12 & 9 & 5 & 15 & 10 & 11 & 14 & 1 & 7 & 6 & 0 & 8 & 13
\end{array}
$$

<p align="center"><em>Tabela 8 — S-box 7 do DES (S₇: 6 bits → 4 bits)</em></p>

$$
\begin{array}{cccccccccccccccc}
   & 0 & 1 & 2 & 3 & 4 & 5 & 6 & 7 & 8 & 9 & 10 & 11 & 12 & 13 & 14 & 15 \\
0: & 4 & 11 & 2 & 14 & 15 & 0 & 8 & 13 & 3 & 12 & 9 & 7 & 5 & 10 & 6 & 1 \\
1: & 13 & 0 & 11 & 7 & 4 & 9 & 1 & 10 & 14 & 3 & 5 & 12 & 2 & 15 & 8 & 6 \\
2: & 1 & 4 & 11 & 13 & 12 & 3 & 7 & 14 & 10 & 15 & 6 & 8 & 0 & 5 & 9 & 2 \\
3: & 6 & 11 & 13 & 8 & 1 & 4 & 10 & 7 & 9 & 5 & 0 & 15 & 14 & 2 & 3 & 12
\end{array}
$$

## Exemplo: aplicação da S-box $S_1$

Considere o bloco de 6 bits abaixo, extraído da entrada da função $f$:

$$
B_1 = 101110
$$

Dividimos o bloco conforme a regra da S-box:

- **Bits 1 e 6** (primeiro e último): determinam a linha:
  $$
  \text{Linha} = (b_1, b_6) = (1, 0) = 10_2 = 2
  $$

- **Bits 2 a 5**: determinam a coluna:
  $$
  \text{Coluna} = (b_2 b_3 b_4 b_5) = (0111) = 7
  $$

Consultando a tabela da $S_1$, na linha 2 e coluna 7:

$$
S_1[2][7] = 11_{10} = 1011
$$

Logo, a saída da S-box $S_1$ para esse bloco é:

$$
\text{Saída de } B_1 = 1011
$$

Esse procedimento é aplicado a cada um dos oito blocos $B_i$ de 6 bits, com cada um sendo processado por sua respectiva S-box $S_i$. A concatenação das oito saídas gera o bloco final de 32 bits que será permutado na próxima etapa da função $f$.

## Conclusão

Até aqui, vimos as quatro etapas que compõem a função $f$: expansão, XOR com a subchave, substituição pelas S-boxes e a permutação final. Juntas, essas etapas garantem a combinação de confusão e difusão necessárias à segurança do DES.

Na próxima seção, analisaremos o processo de **key schedule**, responsável por derivar as 16 subchaves de 48 bits a partir da chave principal de 64 bits. Cada rodada do DES utiliza uma dessas subchaves, o que introduz variabilidade no processamento e impede ataques por repetição de padrões.

