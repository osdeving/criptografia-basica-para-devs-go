# Divisibilidade

## Algoritmo de Euclides

O algoritmo de Euclides é um método que permite encontrar o máximo divisor comum (MDC) de dois números inteiros. O MDC é o maior número que divide ambos os números sem deixar resto.

### Exemplos

---
a) Vamos encontrar o MDC de 12 e 18.

1. Divida o maior número pelo menor:
    $$
        18 \div 12 = 1 \quad \text{(quociente)}
    $$
2. Calcule o resto:
    $$
        18 - (12 \cdot 1) = 6 \quad \text{(resto)}
    $$

3. Repita o processo com o divisor e o resto:
    $$
        12 \div 6 = 2 \quad \text{(quociente)}
    $$

4. Calcule o novo resto:
    $$
        12 - (6 \cdot 2) = 0 \quad \text{(resto)}
    $$
5. Como o resto é zero, o MDC é o último divisor, que é 6.
6. Portanto, o MDC de 12 e 18 é 6.

---
b) Vamos encontrar o MDC de 48 e 18.
1. Divida o maior número pelo menor:
    $$
        48 \div 18 = 2 \quad \text{(quociente)}
    $$
2. Calcule o resto:
    $$
        48 - (18 \cdot 2) = 12 \quad \text{(resto)}
    $$
3. Repita o processo com o divisor e o resto:
    $$
        18 \div 12 = 1 \quad \text{(quociente)}
    $$
4. Calcule o novo resto:
    $$
        18 - (12 \cdot 1) = 6 \quad \text{(resto)}
    $$
5. Repita o processo com o divisor e o resto:
    $$
        12 \div 6 = 2 \quad \text{(quociente)}
    $$
6. Calcule o novo resto:
    $$
        12 - (6 \cdot 2) = 0 \quad \text{(resto)}
    $$
7. Como o resto é zero, o MDC é o último divisor, que é 6.
8. Portanto, o MDC de 48 e 18 é 6.

---
c) Vamos encontrar o MDC de coprimos, como 15 e 28.
1. Divida o maior número pelo menor:
    $$
        28 \div 15 = 1 \quad \text{(quociente)}
    $$
2. Calcule o resto:
    $$
        28 - (15 \cdot 1) = 13 \quad \text{(resto)}
    $$
3. Repita o processo com o divisor e o resto:
    $$
        15 \div 13 = 1 \quad \text{(quociente)}
    $$
4. Calcule o novo resto:
    $$
        15 - (13 \cdot 1) = 2 \quad \text{(resto)}
    $$
5. Repita o processo com o divisor e o resto:
    $$
        13 \div 2 = 6 \quad \text{(quociente)}
    $$
6. Calcule o novo resto:
    $$
        13 - (2 \cdot 6) = 1 \quad \text{(resto)}
    $$
7. Repita o processo com o divisor e o resto:
    $$
        2 \div 1 = 2 \quad \text{(quociente)}
    $$
8. Calcule o novo resto:
    $$
        2 - (1 \cdot 2) = 0 \quad \text{(resto)}
    $$
9. Como o resto é zero, o MDC é o último divisor, que é 1.
10. Portanto, o MDC de 15 e 28 é 1, indicando que são primos entre si.

---
d) Vamos encontrar o MDC de dois números primos, como 7 e 13.
1. Divida o maior número pelo menor:
    $$
        13 \div 7 = 1 \quad \text{(quociente)}
    $$
2. Calcule o resto:
    $$
        13 - (7 \cdot 1) = 6 \quad \text{(resto)}
    $$
3. Repita o processo com o divisor e o resto:
    $$
        7 \div 6 = 1 \quad \text{(quociente)}
    $$
4. Calcule o novo resto:
    $$
        7 - (6 \cdot 1) = 1 \quad \text{(resto)}
    $$
5. Repita o processo com o divisor e o resto:
    $$
        6 \div 1 = 6 \quad \text{(quociente)}
    $$
6. Calcule o novo resto:
    $$
        6 - (1 \cdot 6) = 0 \quad \text{(resto)}
    $$
7. Como o resto é zero, o MDC é o último divisor, que é 1.
8. Portanto, o MDC de 7 e 13 é 1, indicando que são primos entre si, note que 7 e 13 são primos, visto que um número primo é aquele que só é divisível por 1 e por ele mesmo, dois primos distintos sempre terão MDC igual a 1.

### Implementação em Go

Uma implementação simples em Go, sem usar recursão para maior clareza, é a seguinte:

```go
func mdc(a, b int) int {
    for {
        // a > b?
        if a > b {
            // Divide a por b, e obtém o resto (% faz isso)
            r := a % b

            // r > 0?
            if r > 0 {
                // Troca a por b e b por r e repete
                a, b = b, r
                continue
            } else {
                // Se r == 0, então b é o MDC
                return b
            }
        } else {
            // Se b > a, então troca a por b e b por a e repete
            a, b = b, a
        }
    }
}
```

## Algoritmo de Euclides estendido

O algoritmo de Euclides estendido é uma extensão do algoritmo de Euclides que, além de encontrar o MDC de dois números, também fornece os coeficientes da combinação linear que resulta no MDC.

### Exemplo
Vamos encontrar o MDC de 12 e 18 usando o algoritmo de Euclides estendido.
1. Inicialize os valores:
    $$
        a = 12, b = 18
    $$
2. Aplique o algoritmo de Euclides:
    $$
        18 = 1 \cdot 12 + 6
    $$
3. Continue aplicando o algoritmo:
    $$
        12 = 2 \cdot 6 + 0
    $$
4. O MDC é 6.
5. Agora, vamos encontrar os coeficientes $x$ e $y$ tais que:
    $$
        12x + 18y = 6
    $$
6. A partir do passo 2, podemos reescrever a equação:
    $$
        6 = 18 - 1 \cdot 12
    $$
7. Agora, substituímos $12$ por $18 - 1 \cdot 12$:
    $$
        6 = 18 - 1 \cdot (18 - 1 \cdot 12)
    $$
8. Simplificando, obtemos:
    $$
        6 = 2 \cdot 18 - 1 \cdot 12
    $$
9. Portanto, os coeficientes são $x = -1$ e $y = 2$.
10. Assim, temos:
    $$
        12 \cdot (-1) + 18 \cdot 2 = 6
    $$
11. O que confirma que o MDC de 12 e 18 é 6.
12. Portanto, o MDC de 12 e 18 é 6, com os coeficientes $x = -1$ e $y = 2$.
13. Isso significa que $12 \cdot (-1) + 18 \cdot 2 = 6$.
14. O algoritmo de Euclides estendido é útil em várias aplicações, como na criptografia, onde precisamos encontrar inversos multiplicativos em corpos finitos.


