# Arquitetura Geral

O Data Encryption Standard (DES) segue a arquitetura de uma **rede de Feistel**, composta por 16 rodadas idênticas de transformação e duas permutações fixas nas extremidades: a **permutação inicial (IP)** e a **permutação final (IP⁻¹)**. Cada rodada aplica uma função complexa à metade direita do bloco de dados e a combina com a metade esquerda, promovendo confusão e difusão ao longo das etapas.

## Estrutura geral do processo de cifragem

A cifragem em DES inicia com a permutação inicial (IP), que reordena os 64 bits do bloco de entrada segundo uma tabela fixa. Em seguida, o bloco é dividido em duas metades de 32 bits: **L₀** (esquerda) e **R₀** (direita).

Cada uma das 16 rodadas realiza as seguintes operações:

$$
\begin{aligned}
L_i &= R_{i-1} \\
R_i &= L_{i-1} \oplus f(R_{i-1}, K_i)
\end{aligned}
$$

A função $f$ opera sobre 32 bits de entrada e uma subchave de 48 bits $K_i$, gerada a partir da chave principal. O resultado passa por expansões, substituições via S-boxes e uma permutação interna (mais detalhes nas próximas seções).

Após a 16ª rodada, os blocos $L_{16}$ e $R_{16}$ são recombinados na ordem inversa (isto é, $R_{16} \parallel L_{16} $), e então passa pela permutação final (IP⁻¹), produzindo o bloco cifrado.

<br>
<br>


<p align="center">
    <em>Figura 1 — Arquitetura geral do algoritmo DES</em>
    <img src="des-arquitetura-geral.svg" alt="Arquitetura geral do algoritmo DES"/>
</p>
<br>
<br>

A Figura 1 ilustra a arquitetura geral do DES. O processo de cifragem de um bloco de 64 bits pode ser descrito nas seguintes etapas:

<br>

> $$
\begin{aligned}
&\textbf{1. Permutação Inicial (IP)} \\
&\quad X = IP(M), \quad \text{onde } M \in \{0,1\}^{64} \\
&\quad X \rightarrow L_0 \parallel R_0, \quad L_0, R_0 \in \{0,1\}^{32} \\[1em]

&\textbf{2. Rodadas de Feistel (i = 1,\dots,16)} \\
&\quad L_i = R_{i-1} \\
&\quad R_i = L_{i-1} \oplus f(R_{i-1}, k_i) \\[1em]

&\textbf{3. Troca final (Swap)} \\
&\quad X' = R_{16} \parallel L_{16} \\[1em]

&\textbf{4. Permutação Final (IP}^{-1}\textbf{)} \\
&\quad C = IP^{-1}(X') \quad \text{(bloco cifrado)}
\end{aligned}
$$

<br>


As permutações inicial e final não interferem na segurança e têm papel meramente operacional: visam otimizar a manipulação dos dados em implementações físicas, especialmente em circuitos integrados da década de 1970.

A inversão da ordem dos blocos ao fim das rodadas — ou seja, o uso de $R_{16} \parallel L_{16}$ em vez de $L_{16} \parallel R_{16}$ — decorre diretamente da estrutura de Feistel e permite que o processo de decifragem seja simetricamente equivalente ao de cifragem, trocando-se apenas a ordem das subchaves.

Por fim, a função $f$, que introduz não linearidade ao processo, constitui o núcleo da segurança do DES. Nas próximas seções, exploraremos em detalhe sua estrutura interna, as S-boxes envolvidas e o processo de derivação das subchaves a partir da chave principal.


