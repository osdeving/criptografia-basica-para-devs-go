# Parity Bit

Parity bit é um código de detecção de erros em que um bit extra é adicionado ao final da mensagem para garantir que o número total de bits "1" seja par (paridade par) ou ímpar (paridade ímpar). Se um único bit for invertido durante a transmissão, o receptor perceberá a inconsistência e poderá solicitar o reenvio.

Exemplo de paridade par:

```
Mensagem original: 10101110 (8 bits)
Número de bits "1": 5
Bit de paridade: 1 (para tornar o total par)
Mensagem transmitida: 101011101
```

Se um erro ocorrer e inverter um bit:

```
Recebido: 101001101 (erro no terceiro bit)
Número de bits "1": 4 → Inconsistência detectada.
```
Embora simples, a paridade tem limitações: não detecta erros em um número par de bits invertidos, tornando-a inadequada para muitas aplicações.

Como alternativa, poderíamos enviar a mensagem duas vezes e comparar ambas as cópias. No entanto, isso dobra o consumo de banda e não garante detecção total de erros se ambas as cópias forem corrompidas de forma idêntica.

O ideal é encontrar um código de verificação que minimize o custo de transmissão e maximize a capacidade de detecção de erros. Os dois métodos mais utilizados são:

- Checksums – Somam os valores dos bytes da mensagem e armazenam o resultado junto com os dados.
- Cyclic Redundancy Check (CRC) – Usa operações matemáticas sobre polinômios para gerar um código de verificação mais robusto.

Nos próximos tópicos, exploraremos esses métodos em detalhes, mostrando suas implementações e aplicações.