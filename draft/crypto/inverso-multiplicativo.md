A expressão:

$$
a \cdot a^{-1} \equiv 1 \mod p
$$

Significa:

> Quando você multiplica $a$ por seu inverso multiplicativo $a^{-1}$, e divide o resultado por $p$, o resto é 1.

Exemplo:

Vamos usar $a = 3$ e $p = 7$.

Sabemos que $3^{-1} \equiv 5 \mod 7$.\
Porque: $3 \cdot 5 = 15$, e $15 \mod 7 = 1$.

Ou seja: $15 \div 7 = 2$ com resto 1 $\Rightarrow 3 \cdot 3^{-1} \mod 7$.

Interpretação:

> O inverso multiplicativo módulo $p$ é o número que, ao multiplicar por $a$, dá 1 como resto quando dividido por $p$.

Tabela com os inversos multiplicativos no corpo $\mathbb{Z}_7$:

| $a$ | $a^{-1}$ | Verificação: $a \cdot a^{-1} \mod 7$ |
| --- | ------- | ---------------------------------|
| 1   | 1       | 1 \* 1 = 1                       |
| 2   | 4       | 2 \* 4 = 8, 8 mod 7 = 1         |
| 3   | 5       | 3 \* 5 = 15, 15 mod 7 = 1       |
| 4   | 2       | 4 \* 2 = 8, 8 mod 7 = 1         |
| 5   | 3       | 5 \* 3 = 15, 15 mod 7 = 1       |
| 6   | 6       | 6 \* 6 = 36, 36 mod 7 = 1       |
| 0   | -       | Não existe (0 não tem inverso)   |

Alguns números são auto inversos, como:
- $1^{-1} = 1$
- $6^{-1} = 6$

Os pares $(a, a^{-1})$ são simétricos

- $2 \Leftrightarrow  4$
- $3 \Leftrightarrow  5$

Observação:

No conjunto dos números reais:

$$
a^{-1} = \frac{1}{a}
$$

Exemplo:
$$
3^{-1} = \frac{1}{3},\quad \text{pois}\ 3 \cdot \frac{1}{3} = \frac{3}{3} = 1
$$

Mas no módulo $p$ (ex: $\mathbb{Z}_7$), não existe fração como $\frac{1}{3}$.

Em aritmética modular dizemos:

O inverso multiplicativo de $a \mod p$ é o número $x$ tal que:
$a \cdot x \equiv 1 \mod p$.

Esse $x$ é análogo de $\frac{1}{a}$ no corpo $\mathbb{Z}_p$.

Exemplo com $a = 3$ e $p = 7$.

Queremos:

$$
3 \cdot x \equiv 1 \mod 7
$$

Testando:
$$
\begin{aligned}
3 \cdot 1 &= 3 \equiv 3\\
3 \cdot 2 &= 6 \equiv 6\\
3 \cdot 3 &= 9 \equiv 2\\
3 \cdot 4 &= 12 \equiv 5\\
3 \cdot 5 &= 15 \equiv 1\\
\end{aligned}
$$

Achamos: $x = 5 \Rightarrow 3^{-1} \equiv 5 \mod 7$.

No mundo dos reais:

$$
3^{-1} = \frac{1}{3}
$$

No mundo dos módulos:
$$
3^{-1} \mod 7 = 5,  \quad \text{pois}\ 3 \cdot 5 \equiv 1 \mod 7
$$

* o papel de $\frac{1}{3}$ está sendo desempenhado por $5$ no módulo $7$.
* a "fração" se expressa como outro inteiro que cumpre a mesma função.

Conclusão:

Em $\mathbb{Z}_p$, o inverso multiplicativo de $a$ não é $\frac{1}{a}$ como número real, mas sim o número $x \in \mathbb{Z}_p$ que:
$$
a \cdot x \equiv 1 \mod p
$$

Esse $x$ é o substituto de $\frac{1}{a}$ no corpo $\mathbb{Z}_p$.
















