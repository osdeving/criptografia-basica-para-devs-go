usa blocos de 64 bits
usa chaves de 56 bits
usa 16 rodadas de cifragem para cada bloco
cada rodada usa uma chave diferente derivada da chave principal
usa uma estrutura de Feistel (usada em vários algoritmos, mas não todos, p.ex.: AES não é uma cifra de Feistel)
vantagem do feistel: processo de cifrar e decifrar é praticamente o mesmo, mudando a ordem das subchaves
primeiro temo uma permutação inicial sobre o bloco de 64 bits dividido em duas partes: L0 e R0
L0 e R0 são a entrada da rede de Feistel que consiste em 16 rodadas
Ri é passado pela função f e o resultado é feito XOR com a outra metade esquerda Li
Então as metades são trocadas

```
Lᵢ = Rᵢ₋₁
Rᵢ = Lᵢ₋₁ ⊕ f(Rᵢ₋₁, kᵢ)

para i = 1, ..., 16.
```

```go
L[0] = L0
R[0] = R0

for i := 1; i <= 16; i++ {
    L[i] = R[i-1]
    R[i] = xor(L[i-1], f(R[i-1], K[i-1]))
}
```

cada rodada a chave ki é derivada da chave principal k de  56 bits via key schedule
a metade direita não é cifrada, ela é usada para cifrar a metade esquerda e se torna a metade esquerda da rodada seguinte

(a função f pode ser visto como um PRNG que recebe Ri e ki, se retorno serve para embaralhar Li via XOR. A saída de f precisa ser imprevisível)


A estrutura Feistel como um todo é uma função bijetora nos 64 bits de entrada:
Isso significa que cada entrada de 64 bits gera exatamente uma saída distinta de 64 bits, e toda saída pode ser revertida à entrada original (a função tem inversa)
A função f não é injetora, i.e. ela pode gerar a mesma saída para entradas Li e ki diferentes.



