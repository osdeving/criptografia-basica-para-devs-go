# Matemática para Criptografia

## 1. Introdução
A criptografia moderna depende fortemente de estruturas matemáticas abstratas como grupos, anéis, corpos, polinômios e aritmética modular. Estas ferramentas fornecem a base teórica para algoritmos amplamente utilizados, como o DES, AES e curvas elípticas (ECC). Este documento oferece uma abordagem gradual e didática da matemática por trás dessas construções, com definições, propriedades e aplicações práticas.

## 1. Aritmética Modular

### 1.1 Definição

A **aritmética modular** é um sistema de aritmética para inteiros, onde os números "voltam ao início" após atingir um certo valor, chamado de **módulo**. A notação $a \equiv b \mod m$ significa que $a$ e $b$ têm o mesmo resto quando divididos por $m$. É também chamada de **aritmética do relógio**, pois o relógio volta ao início após cada volta completa. Por exemplo, o número 13 em um relógio de 12 horas é representado como 1, pois $13 \mod 12 = 1$.

### 1.2 Propriedades

### 1.3 Inverso Multiplicativo

Definição

$$
a \cdot a^{-1} \equiv 1 \mod m
$$

onde $a^{-1}$ é o inverso multiplicativo de $a$ módulo $m$. O inverso existe se e somente se $a$




### 1.3 Aritmética Modular em Computadores


## 2. Fundamentos Algébricos

### 2.1 Conjuntos e operações binárias
Um **conjunto** é uma coleção de elementos distintos. Uma **operação binária** sobre um conjunto $S$ é uma função que associa a cada par ordenado de elementos de $S$ um único elemento de $S$.

Exemplo: adição de inteiros, multiplicação de polinômios, XOR entre bits.

### 2.2 Grupos

Um **grupo** é um conjunto $G$, munido de uma operação binária $\cdot$, que satisfaz:
- **Fechamento**: $\forall a, b \in G, \; a \cdot b \in G$
- **Associatividade**: $a \cdot (b \cdot c) = (a \cdot b) \cdot c$
- **Elemento neutro**: $\exists e \in G$ tal que $a \cdot e = e \cdot a = a$
- **Elemento inverso**: $\forall a \in G, \; \exists a^{-1} \in G$ tal que $a \cdot a^{-1} = e$

Se, além disso, $a \cdot b = b \cdot a$, o grupo é chamado **abeliano**.

Exemplos:
- $(\mathbb{Z}, +)$ é um grupo abeliano
- $(\mathbb{Z}_n^*, \cdot)$, com $n$ primo, também é abeliano

### 2.3 Anéis (rings)
Um **anel** é um conjunto $R$ com duas operações:
- $(R, +)$ é um grupo abeliano
- Multiplicação é associativa: $a(bc) = (ab)c$
- Distributividade: $a(b + c) = ab + ac$

Se a multiplicação é comutativa, diz-se que o anel é comutativo. Se existe unidade $1 \in R$ tal que $a \cdot 1 = a$, diz-se que o anel tem identidade.

Exemplo: $(\mathbb{Z}, +, \cdot)$ é um anel comutativo com identidade

### 2.4 Corpos (campos/fields)
Um **corpo** (ou field) é um anel comutativo com unidade onde todos os elementos não nulos possuem inverso multiplicativo.

Exemplos:
- $\mathbb{Q}, \mathbb{R}, \mathbb{C}$
- $\mathbb{F}_2 = \{0,1\}$, com soma e multiplicação mod 2
- $\mathbb{F}_{2^m}$, corpos finitos construídos a partir de polinômios

## 3. Aritmética Modular
(... a expandir ...)

## 4. Polinômios
(... a expandir ...)

## 5. Estruturas em Criptografia
(... a expandir ...)

## 6. Aplicações Práticas
(... a expandir ...)

## 7. Tabelas, Definições e Teoremas
(... a expandir ...)

## 8. Conclusão
(... a expandir ...)

## Definições

<div style="border: 1px solid #ccc; padding: 16px; background-color:RGB(0,0,0); border-radius: 6px;font-size: 18px; color: white;">
<b>Definição: Multiplicação em corpos de extensão</b><br><br>

Sejam $A(x), B(x) \in \mathbb{F}_{2^m}$, e seja

$$
P(x) \equiv \sum_{i=0}^{m} p_i x^i, \quad p_i \in \mathbb{F}_2
$$

um polinômio irredutível de grau $m$. A multiplicação entre os dois elementos $A(x)$ e $B(x)$ é realizada da seguinte forma:

$$
C(x) \equiv A(x) \cdot B(x) \mod P(x)
$$
</div>

<div style="border: 1px solid #ccc; padding: 16px; background-color:RGB(0,0,0); border-radius: 6px;font-size: 18px; color: white;">

$$
P(x) = 1x^8 + 0x^7 + 0x^6 + 0x^5 + 1x^4 + 1x^3 + 0x^2 + 1x^1 + 1
$$

### 📚 Exemplo: Multiplicação em $\mathbb{F}_{2^4}$ com Redução Modular

Seja o corpo finito $\mathbb{F}_{2^4} = \mathbb{F}_2[x] / (P(x))$, onde:

$$
P(x) = x^4 + x + 1
$$

Queremos multiplicar os polinômios:
- $A(x) = x^3 + x^2 + 1 \Rightarrow (1101)_2$
- $B(x) = x^2 + x \Rightarrow (0110)_2$
</div>

### 🧠 Etapa 1: Multiplicação polinomial (sem módulo)

$$
C'(x) = A(x) \cdot B(x) = (x^3 + x^2 + 1)(x^2 + x)
$$

Distribuindo os termos:
- $x^3 \cdot x^2 = x^5$
- $x^3 \cdot x = x^4$
- $x^2 \cdot x^2 = x^4$
- $x^2 \cdot x = x^3$
- $1 \cdot x^2 = x^2$
- $1 \cdot x = x$

Somando em $\mathbb{F}_2$ (usando XOR):
$$
C'(x) = x^5 + x^4 + x^4 + x^3 + x^2 + x = x^5 + x^3 + x^2 + x
$$

---

### 🔁 Etapa 2: Redução módulo $P(x) = x^4 + x + 1$

Sabemos que:
- $x^4 \equiv x + 1 \mod P(x)$
- Multiplicando por $x$: $x^5 \equiv x^2 + x \mod P(x)$

Substituímos no polinômio:

$$
C(x) = x^5 + x^3 + x^2 + x \equiv (x^2 + x) + x^3 + x^2 + x
$$

Agrupando os termos (com XOR):
$$
C(x) \equiv x^3 + (x^2 \oplus x^2) + (x \oplus x) = x^3
$$

---

### ✅ Resultado final:

$$
A(x) \cdot B(x) \mod P(x) = x^3 \Rightarrow (1000)_2
$$

---

### ⚠️ Importante: isso **não** é multiplicação de inteiros!

Se fizermos:
- $(1101)_2 = 13$
- $(0110)_2 = 6$
- $13 \cdot 6 = 78 \Rightarrow (1001110)_2$

➡️ Esse **não é** o mesmo resultado que $(1000)_2 = 8$ da multiplicação em $\mathbb{F}_{2^4}$.

---

### 💡 Conclusão

Mesmo representando elementos do corpo como inteiros binários, as operações devem seguir a **aritmética de polinômios em $\mathbb{F}_2[x]$ com redução módulo um polinômio irredutível**.



