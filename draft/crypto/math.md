# Fundamentos de Matemática para Criptografia

## Divisibilidade

Nós dizemos que $b$ divide $a$ se existe um inteiro $c$ tal que $a = bc$. Neste caso, dizemos que $b$ é um divisor de $a$ e que $a$ é um múltiplo de $b$.

<div style="
  max-width: 500px;
  margin: 2em auto;
  padding: 1.5em;
  border-left: 6px solid #007ACC;
  background: #f4faff;
  box-shadow: 0 2px 6px rgba(0,0,0,0.05);
  font-family: sans-serif;
  color: #222;
  border-radius: 8px;
  line-height: 1.6;
">
  <strong>Divisores:</strong><br>
  Os divisores positivos de <strong>24</strong> são: <em>1, 2, 3, 4, 6, 8, 12, 24</em>.<br><br>
  <strong>Ex.:</strong><br>
  <code>13|192; -5|30;&nbsp; 17|289; -3|33; 17|0</code>
</div>

As seguintes propriedades são válidas para todos os inteiros $a, b, c$:

$$
\begin{array}{l}
\text{Se } a \mid 1, \text{ então } a = \stackrel{+}{-}1; \\\\[0.5em]
\text{Se } a \mid b \text{ e } b \mid a, \text{ então } a = \stackrel{+}{-}b; \\\\[0.5em]
\text{Qualquer } b \neq 0 \text{ divide } 0; \\\\[0.5em]
\text{Se } a \mid b \text{ e } b \mid c, \text{ então } a \mid c; \\\\[0.5em]
\text{Se } b \mid g \land b \mid h, \text{ então } b \mid (mg + nh),\ \text{ para quaisquer } m, n.
\end{array}
$$



$\mathbb{F}_{2^8} = \frac{\mathbb{F}_2[x]}{(x^8 + x^4 + x^3 + x + 1)}$

$\mathbb{F}_{2^8} = \mathbb{F}_2[x] \big/ \left(x^8 + x^4 + x^3 + x + 1\right)$

<div style="
  max-width: 500px;
  margin: 2em auto;
  padding: 1.5em;
  border-left: 6px solid #007ACC;
  background: #f4faff;
  box-shadow: 0 2px 6px rgba(0,0,0,0.05);
  font-family: sans-serif;
  color: #222;
  border-radius: 8px;
  line-height: 1.6;
">
  <strong>Ex.:</strong><br>
  <code>
  11|66 e 66|198 &#x21D2; 11|198
  </code>
  
</div>


## Grupos

Um grupo é um conjunto $G$ com uma operação binária que associa cada par ordenado $(a, b)$ de elementos $\in G\ \text{a um elemento}\ (a \cdot b) \in G$, satisfazendo as seguintes propriedades:

$$
\begin{array}{llr}
\text{(P1)}\quad a \cdot b \in G,\quad \forall\, a, b \in G; & \text{Fechamento} \\
\text{(P2)}\quad (a \cdot b) \cdot c = a \cdot (b \cdot c),\quad \forall\, a, b, c \in G; & \text{Associatividade} \\
\text{(P3)}\quad \exists\, e \in G \mid a \cdot e = a = e \cdot a,\quad \forall\, a \in G; & \text{Elemento neutro} \\
\text{(P4)}\quad \forall\, a \in G,\ \exists\, a^{-1} \in G \mid a \cdot a^{-1} = a^{-1} \cdot a = e \wedge\ e \in G; & \text{Inverso}
\end{array}
$$









Definição 1.1.1

Um anel $(R, +, \cdot)$ é um conjunto $R$ com duas operações binárias, $+$ e $\cdot$, tal que:

$$
\begin{array}{l}

\text{(a)}\quad (R, +)\ \text{é um grupo abeliano};\\
\text{(b)}\quad a \cdot (b \cdot c) = (a \cdot b) \cdot c,\quad \forall\, a, b, c \in R;\\
\text{(c)}\quad a \cdot (b + c) = a \cdot b + a \cdot c,\quad (b + c) \cdot a = b \cdot a + c \cdot a,\quad \forall\, a, b, c \in R.
\end{array}
$$
