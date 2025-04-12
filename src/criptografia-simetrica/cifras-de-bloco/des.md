# Data Encryption Standard (DES)

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

![DES](des-image.svg)

