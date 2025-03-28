# Construindo Key Streams a partir de PRNGs

A maioria dos PRNGs possuem boas propriedades estatísticas. Se aplicarmos testes estatísticos em um PRNG é provável que consigamos resultados parecidos com os resultados de uma sequência aleatória gerada a partir de um "cara ou coroa" jogando uma moeda pra cima. É tentador, portanto, usar um PRNG para gerar uma sequência aleatória para ser usada como chave para um algoritmo de criptografia simétrica. No entanto, essa percepção de aleatoriedade é enganosa e não é suficiente para garantir a segurança de um sistema criptográfico. Eve é esperta o suficiente para calcular a key stream de Alice e Bob se ela puder deduzir parte da mensagem cifrada. Considere o seguinte exemplo de ataque.

Vamos assumir um PRNG baseado em uma sequência linear.

```plaintext
S0 = seed
Si + 1 = ASi + B mod m, com i = 0, 1, 2, ...
```

Onde m tem 100 bits e Si, A e B pertencem ao conjunto {0, 1, ..., m - 1}. O módulo m é parte do esquema de encriptação e é publicamente conhecido. O segredo A e B, cada qual com 100 bits, são mantidos em segredo. Isso nos dá uma chave de 200 bits, que é mais que suficiente para proteger contra ataques de força bruta. Uma vez que isso é uma stream cipher, Alice pode encriptar com:

```plaintext
yi = xi + Si mod m

ou

yi = xi XOR Si
```

Onde Si são os bits da representação binária da saída do PRNG Sj. Vamos assumir que Eve conhece os primeiros 300 bits do texto cifrado. Ela pode então calcular os primeiros 300 bits da key stream da seguinte forma:

```plaintext
si = yi + xi mod m, i = 0, 1, 2, ..., 300
```

Esses 300 bits imediatamente fornecem os primeiros 300 bits da saída do PRNG: S1 = (s1, ..., s100), S2 = (s101, ..., s200), S3 = (s201, ..., s300). Isso permite Eve gerar duas equações lineares para A e B:

```plaintext
S2 = AS1 + B mod m
S3 = AS2 + B mod m
```

Isso é um sistema de equações lineares com duas incógnitas A e B. Eve pode resolver esse sistema de equações para A e B. Uma vez que ela conhece A e B, ela pode calcular a key stream inteira e decifrar a mensagem.

```plaintext
A = (S2 - S3) / (S1 - S2) mod m
B = S2 - S1(S2 - S3) / (S1 - S2) mod m
```

No caso do mdc(S1 - S2, m) diferente de 1, teremos múltiplas soluções já que isso é uma equação em Zm. Entretanto, com uma quarta informação conhecida a partir do plaintext, Eve pode isolar a solução correta. Alternativamente, Eve pode simplesmente tentar todas as possíveis soluções para A e B e testar qual delas resulta em uma key stream que decifra a mensagem. Em resumo, se Even conhecer parte do texto cifrado, ela pode deduzir a key stream e decifrar a mensagem.


## Gerador Congruencial Linear (LCG)

O gerador congruencial linear (LCG) é um PRNG que gera uma sequência de números inteiros usando a seguinte fórmula:

```plaintext
$S_i+1 = A \dot S_i + B \mod m$
```

Onde:
- Xn é o número gerado na iteração n
- a é o multiplicador
- c é o incremento
- m é o módulo

O gerador congruencial linear é um PRNG determinístico, o que significa que a sequência de números gerados é completamente determinado pelos valores iniciais (X0, a, c, m).

O gerador congruencial linear é amplamente utilizado em simulações e em criptografia, mas também é vulnerável a ataques de força bruta e análise estatística. Se Eve conhecer parte da sequência gerada pelo LCG, ela pode deduzir os parâmetros do LCG e prever a sequência futura. Isso é particularmente perigoso em criptografia, onde a previsibilidade da sequência de números gerados pode levar a vulnerabilidades de segurança.



# Construindo Key Streams a partir de PRNGs

A maioria dos PRNGs possuem boas propriedades estatísticas. Se aplicarmos testes estatísticos em um PRNG é provável que consigamos resultados parecidos com os resultados de uma sequência aleatória gerada a partir de um "cara ou coroa" jogando uma moeda pra cima. É tentador, portanto, usar um PRNG para gerar uma sequência aleatória para ser usada como chave para um algoritmo de criptografia simétrica. No entanto, essa percepção de aleatoriedade é enganosa e não é suficiente para garantir a segurança de um sistema criptográfico. Eve é esperta o suficiente para calcular a key stream de Alice e Bob se ela puder deduzir parte da mensagem cifrada. Considere o seguinte exemplo de ataque.

Vamos assumir um PRNG baseado em uma sequência linear.

```plaintext
S0 = seed
Si + 1 = ASi + B mod m, com i = 0, 1, 2, ...
```

Onde m tem 100 bits e Si, A e B pertencem ao conjunto {0, 1, ..., m - 1}. O módulo m é parte do esquema de encriptação e é publicamente conhecido. O segredo A e B, cada qual com 100 bits, são mantidos em segredo. Isso nos dá uma chave de 200 bits, que é mais que suficiente para proteger contra ataques de força bruta. Uma vez que isso é uma stream cipher, Alice pode encriptar com:

```plaintext
yi = xi + Si mod m

ou

yi = xi XOR Si
```

Onde Si são os bits da representação binária da saída do PRNG Sj. Vamos assumir que Eve conhece os primeiros 300 bits do texto cifrado. Ela pode então calcular os primeiros 300 bits da key stream da seguinte forma:

```plaintext
si = yi + xi mod m, i = 0, 1, 2, ..., 300
```

Esses 300 bits imediatamente fornecem os primeiros 300 bits da saída do PRNG: S1 = (s1, ..., s100), S2 = (s101, ..., s200), S3 = (s201, ..., s300). Isso permite Eve gerar duas equações lineares para A e B:

```plaintext
S2 = AS1 + B mod m
S3 = AS2 + B mod m
```

Isso é um sistema de equações lineares com duas incógnitas A e B. Eve pode resolver esse sistema de equações para A e B. Uma vez que ela conhece A e B, ela pode calcular a key stream inteira e decifrar a mensagem.

```plaintext
A = (S2 - S3) / (S1 - S2) mod m
B = S2 - S1(S2 - S3) / (S1 - S2) mod m
```

No caso do mdc(S1 - S2, m) diferente de 1, teremos múltiplas soluções já que isso é uma equação em Zm. Entretanto, com uma quarta informação conhecida a partir do plaintext, Eve pode isolar a solução correta. Alternativamente, Eve pode simplesmente tentar todas as possíveis soluções para A e B e testar qual delas resulta em uma key stream que decifra a mensagem. Em resumo, se Even conhecer parte do texto cifrado, ela pode deduzir a key stream e decifrar a mensagem.



