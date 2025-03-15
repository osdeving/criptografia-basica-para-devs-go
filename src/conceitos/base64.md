# Base64 e Codificação

<style>
  table {
    border-collapse: collapse;
    width: 50%;
  }
  th, td {
    border: 1px solid black;
    text-align: center;
    padding: 8px;
  }
</style>


Como funciona uma codificação? Imagine que você tem uma regra que diz o seguinte: caso a entrada for 'a', então o valor será 97, caso a entrada for 'b', então o valor será 98, caso for 'c', então o valor será 99, e assim em diante. Isso é um de-para e é também conhecido como tabela ASCII e envolve outros caracteres além do alfabeto. Descrevi um tipo de codificação. 

Base64 é uma regra de codificação (Encode) que aplicamos na entrada produzindo uma saída. Isso significa que podemos lêr a entrada em base64 e produzir uma saída decodificada ou ler uma entrada decodificada e codificar em base64, respectivamente conhecidos como Encode e Decode. 

A diferença em relação a tabela ASCII, é que a entrada pode ser qualquer coisa e não se limita apenas a caracteres individuais. No caso do Base64, a codificação funciona agrupando a entrada em blocos de 3 bytes (24 bits) e dividindo-os em 4 grupos de 6 bits. Cada um desses grupos de 6 bits é então convertido para um caractere correspondente em uma tabela de 64 símbolos. Essa table inclui as letras maiúsculas e minúsculas do alfabeto, seguido pelos números de 0 a 9 e então por dois caracteres especiais (+ e / no Base64 padrão e _ e - no Base64URL). Quando a entrada não é um múltiplo exato de 3 bytes, a codificação adiciona um caractere de preenchimento (=) para manter a consistência.

No Base64, a conversão ocorre em blocos e garante que qualquer sequência de bytes possam ser representadas apenas com caracteres seguros para transporte em protocolos como e-mails (MIME), URLs e JSON. Isso torna o Base64 útil para codificar binários, imagens, chaves criptográficas e outros dados que não são diretamente representáveis como texto legível.

O processo inverso, conhecido como decodificação (Decode), pega uma string codificada em Base64 e reconstrói os bytes originais, revertendo a conversão de 6 bits para 8 bits e removendo qualquer padding (=) que tenha sido adicionado na codificação. Assim como podemos converter a → 97 na tabela ASCII e depois reverter 97 → a, no Base64 podemos pegar "TWFu" e recuperar "Man" ou vice-versa, a partir de "Man" obter "TWFu".

No entanto, é importante lembrar que Base64 não é uma criptografia, pois qualquer pessoa pode decodificá-lo facilmente. Ele é apenas um método de representação de dados que facilita o transporte e armazenamento em sistemas que não suportam caracteres binários diretamente. Nota, o Base64 aumenta o tamanho dos dados em 33% em relação ao original.

## Como funciona na prática?

Vamos supor que a entrada sera a string "Rox". Cada caractere em um computador é representado por um número inteiro conforme a tabela ASCII. 

### Passo 1: Obter os valores ASCII

Primeiro, vamos converter os caracteres em seus valores numéricos:

* R → 82
* o → 111
* x → 120

Agora, precisamos representar esses números em binário (base 2), pois a codificação Base64 trabalha diretamente com bits:

### Passo 2: Converter os valores ASCII para binário (8 bits cada)

- R (82) → 01010010
- o (111) → 01101111
- x (120) → 01111000

Agora temos um total de 3 bytes (24 bits), que são agrupados assim:

01010010 01101111 01111000

### Passo 3: Dividir em blocos de 6 bits

O Base64 trabalha com grupos de 6 bits, então precisamos separar nossos 24 bits assim:

010100 100110 111101 111000

Agora, cada grupo de 6 bits será convertido em um decimal cujo qual representa um índice na tabela Base64.

### Passo 4: Converter os grupos de 6 bits para decimal

* 010100 → 20
* 100110 → 38
* 111101 → 61
* 111000 → 56

Agora, usamos a tabela Base64 para converter esses números em caracteres.

### Passo 5: Mapear os valores na tabela Base64

A tabela Base64 contém 64 caracteres indexados de 0 até 63. Essa tabela segue uma ordem específica:

1) Letras maiúsculas (A-Z) → índices 0 a 25
2) Letras minúsculas (a-z) → índices 26 a 51
3) Números (0-9) → índices 52 a 61
4) Caracteres especiais (+ e /) → índices 62 e 63

Representação em Go:

```go
const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
```
Agora, usamos essa tabela para mapear os valores 20, 38, 61 e 56:

* base64Table[20] → "U"
* base64Table[38] → "m"
* base64Table[61] → "9"
* base64Table[56] → "4"

Com isso, o texto "Rox" foi convertido para Base64 como "Um94"

### Resumo

1 - Convertemos os caracteres para seus valores ASCII (R = 82, o = 111, x = 120).

2 - Transformamos esses valores em binário (01010010 01101111 01111000).

3 - Dividimos em blocos de 6 bits (010100 100110 111101 111000).

4 - Convertemos cada grupo de 6 bits em um número decimal (20, 38, 61, 56).

5 - Mapeamos esses números na tabela Base64, resultando na string "Um94".

Bem tranquilo, certo?

## Implementação em Go

O maior desafio que teremos é agrupar os bits em grupos de 6 bits, pois normalmente estamos trabalhando com 8 bits (1 byte). Isso implica que, a cada 3 bytes (24 bits no total), teremos 4 grupos de 6 bits, pois 24 / 6 = 4. 


---
<summary>Extra: Entendendo o bitwise AND
<details>
<br>
<p>Se tivermos um valor binário qualquer e aplicarmos uma máscara bit a bit (bitwise AND) com 0b111111 (0x3F em hexadecimal), conseguimos extrair exatamente 6 bits da posição desejada.</p>

<p>Isso funciona porque o operador AND (&), mantém apenas os bits onde há 1 nos dois operadores, então podemos isolar porções específicas de um número maior.</p>

#### Exemplo prático:

```
BYTE qualquer:            01010010  (82 em decimal)
Máscara de 6 bits:        00111111  (0x3F em hexadecimal)
Resultado após AND:       00010010  (18 em decimal)
```

Note que AND tem o poder de desligar o bit ou mantê-lo sem modificação. 

<table border="1">
  <tr>
    <th>A</th>
    <th>AND</th>
    <th>B</th>
    <th>C</th>
  </tr>
  <tr>
    <td>1</td>
    <td>AND</td>
    <td>1</td>
    <td>1</td>
  </tr>
  <tr>
    <td>1</td>
    <td>AND</td>
    <td>0</td>
    <td>0</td>
  </tr>
  <tr>
    <td>0</td>
    <td>AND</td>
    <td>0</td>
    <td>0</td>
  </tr>
</table>

<br>

Ou seja:

* Se você aplicar a máscara b00000000, vai apagar tudo
* Se você aplicar a máscara b00000001 vai apagar tudo e, para o primeiro bit, vai depender se o outro valor tem 1 ou 0: se tiver 1, ele será mantido, se tiver 0, ele apaga.

Outro exemplo: pegar apenas os bits 3 e 4 de um byte qualquer

<pre>
BYTE qualquer:            <span style="color:#ffcc00;">01010010</span>  (82 em decimal)
Máscara de 6 bits:        <span style="color:#00ccff;">00111111</span>  (0x3F em hexadecimal)
Resultado após AND:       <span style="color:#ff6666;">00010010</span>  (18 em decimal)
</pre>

</details>
</summary>

---

Podemos trabalhar diretamente com 3 bytes em um inteiro e percorrer a mensagem de entrada de 3 em 3 bytes e para cada passagem jogar os 3 bytes no inteiro e extrair os valores usando máscara de bits.

```go
for i := 0; i < len(input); i += 6 {
    
    blk := int32(in[i    ]) << 16 | 
           int32(in[i + 1]) << 8  | 
           int32(in[i + 2])

    out = append(out,
            b64[(blk >> 18) & 0b00111111],
            b64[(blk >> 12) & 0b00111111],
            b64[(blk >>  6) & 0b00111111],
            b64[ blk        & 0b00111111],
		)

}
```

---
<summary>Extra: Explicação do código
<details>
<br>
<p>O que foi feito?</p>
<br>
<p>Criamos um inteiro de 32 bits (int32) e jogamos os 3 bytes do input em cima dele. Para o primeiro byte (na posição i) jogamos para a esquerda 16 bits, para o segundo byte (na posição i + 1) jogamos para a esquerda 8 bits e para o terceiro byte (na posição i + 2) não precisamos jogar nada para a esquerda.</p>
<p>
Imagine que são caixas que cabem 1 byte, o inteiro possui 4 dessas caixas, então precisamos jogar o primeiro byte para a esquerda 16 bits para que ele ocupe a terceira posição, o segundo byte para a esquerda 8 bits e o terceiro byte não precisamos jogar nada para a esquerda porque vai começar no bit 0. Graficamente temos isso:</p>
<pre>
Byte i + 0 = <span style="color:#ff6666;">01101111</span>  (111 em decimal)
Byte i + 1 = <span style="color:#00ccff;">01111000</span>  (120 em decimal)
Byte i + 2 = <span style="color:#ffcc00;">01010010</span>  (82 em decimal) <br>
blk = <span style="color:lightblue"> 000000000 00000000 00000000 00000000</span><br>
Primeiro byte entra começando no bit 16
blk = <span style="color:lightblue"> 000000000<span style="color:#ffcc00;">01010010</span>00000000 00000000</span><br>
Segundo byte entra começando no bit 8
blk = <span style="color:lightblue"> 00000000</span><span style="color:#ffcc00;">01010010</span><span style="color:#00ccff;">01111000</span>00000000</span><br>
Terceiro byte entra começando no bit 0
blk = <span style="color:lightblue"> 00000000</span><span style="color:#ffcc00;">01010010</span><span style="color:#00ccff;">01111000</span><span style="color:#ff6666;">01010010</span>
</pre>

<p>Vamos em câmera lenta. Suponha que</p>

<pre>
in[i    ] = b10000000
in[i + 1] = b00000001
in[i + 2] = b00010000
</pre>

```go
blk := int32(in[i    ]) << 16  // blk = 00000000_10000000_00000000_00000000
blk  |= int32(in[i + 1]) << 8  // blk = 00000000_10000000_00000001_00000000
blk |= int32(in[i + 2])        // blk = 00000000_10000000_00000001_00010000
```
<p>Ou seja, ligou o bit 32 vindo do primeiro byte (ele já estava na posição 8, deslocou 16), o bit 8 do segundo byte (estava na posição 0 e deslocou 8) e o bit 5 (estava na posição 5 do terceiro byte e não teve deslocamento).</p>

<p>Sei que vocẽ já entendeu, mas cabe lembrar que OR funciona da seguinte forma, você tem um valor qualquer com alguns bits ligados e outros não, quando você aplica o OR com outro valor, o que já existe no seu continua, o que não existe no seu, mas existe no outro cara, ele passa a existir no seu. P.ex.: o seu é b00010000 o outro cara é b00000001 agora o seu será b00010001. Veja a tabela do OR para refrescar a memória:</p>

<table border="1">
  <tr>
  <th>A</th>
  <th>OR</th>
  <th>B</th>
  <th>C</th>
  </tr>
  <tr>
  <td>1</td>
  <td>OR</td>
  <td>1</td>
  <td>1</td>
  </tr>
  <tr>
  <td>1</td>
  <td>OR</td>
  <td>0</td>
  <td>1</td>
  </tr>
  <tr>
  <td>0</td>
  <td>OR</td>
  <td>0</td>
  <td>0</td>
  </tr>
</table>
<br>

<p>Agora que temos os 3 bytes dentro de um único inteiro de 32 bits (blk), precisamos extrair 4 grupos de 6 bits, pois cada caractere Base64 é representado por exatamente 6 bits.</p>

<p>A extração é feita aplicando deslocamento de bits (>>) e uma máscara (& 0b00111111), que serve para zerar os bits irrelevantes e pegar exatamente os 6 bits desejados.</p>

<p>Graficamente, temos:</p>

<pre>
blk = <span style="color:lightblue">00000000</span><span style="color:#ffcc00;">01010010</span><span style="color:#00ccff;">01111000</span><span style="color:#ff6666;">01010010</span>
</pre>

<p>Agora, extraímos os grupos de 6 bits um por um:</p>

<ul>
  <li>Primeiros 6 bits: Para extrair os bits mais à esquerda, deslocamos 18 bits para a direita e aplicamos a máscara.</li>
</ul>

<pre>
b64_1 = (blk >> 18) & 0b00111111  <span style="color:darkgreen">// 00000000_00000000_00000000_00010010 → 000100</span>
b64_1 = <span style="color:#ffcc00;">000100</span>
</pre>

<ul>
  <li>Segundos 6 bits: Deslocamos 12 bits para a direita e aplicamos a máscara.</li>
</ul>

<pre>
b64_2 = (blk >> 12) & 0b00111111  <span style="color:darkgreen">// 00000000_00000000_00000010_01010010 → 010100</span>
b64_2 = <span style="color:#00ccff;">010100</span>
</pre>

<ul>
  <li>Terceiros 6 bits: Deslocamos 6 bits para a direita e aplicamos a máscara.</li>
</ul>

<pre>
b64_3 = (blk >> 6) & 0b00111111   <span style="color:darkgreen">// 00000000_00000000_01111000_01010010 → 011110
b64_3 = <span style="color:#ff6666;">011110</span>
</pre>

<ul>
  <li>Últimos 6 bits: Não há deslocamento, apenas aplicamos a máscara.</li>
</ul>

<pre>
b64_4 = blk & 0b00111111 <span style="color:darkgreen">// 00000000_00000000_01111000_01010010 → 100010
b64_4 = <span style="color:#ff9999;">100010</span>
</pre>

<p>Agora, temos os 4 índices da tabela Base64 prontos para serem mapeados!</p>
<p>Isso significa que podemos simplesmente usar o array <code>b64</code> para obter os caracteres correspondentes, indexando diretamente com <code>b64_1</code>, <code>b64_2</code>, <code>b64_3</code> e <code>b64_4</code>.</p>

<pre>
b64_1 = <span style="color:#ffcc00;">000100</span>  (4 em decimal)  →  Índice <code>b64[4]</code>  →  <span style="color:#ffcc00;">T</span>
b64_2 = <span style="color:#00ccff;">010100</span>  (20 em decimal) →  Índice <code>b64[20]</code> →  <span style="color:#00ccff;">U</span>
b64_3 = <span style="color:#ff6666;">011110</span>  (30 em decimal) →  Índice <code>b64[30]</code> →  <span style="color:#ff6666;">e</span>
b64_4 = <span style="color:#ff9999;">100010</span>  (34 em decimal) →  Índice <code>b64[34]</code> →  <span style="color:#ff9999;">Y</span>
</pre>

<p>Ou seja, os números binários extraídos representam <strong>índices na tabela Base64</strong>, e ao acessar <code>b64[4]</code>, <code>b64[20]</code>, etc., obtemos os caracteres finais da string codificada.</p>
</details>
</summary>

---


Agora só precisamos lidar com casos onde o input tem menos que 3 bytes ou o input não é múltiplo de 3 bytes

```go
if rem == 0 {
		return string(out)
	}
    
	var blk int64 = 0

 	if rem == 1 {
        blk = int64(in[len - rem]) << 16
        out = append(out,
            b64[(blk >> 18) & 0b00111111],
            b64[(blk >> 12) & 0b00111111],
            '=',
            '=',
        )
    } else if rem == 2 {
        blk = int64(in[len - rem]) << 16 | int64(in[len - rem + 1]) << 8
        out = append(out,
            b64[(blk >> 18) & 0b00111111],
            b64[(blk >> 12) & 0b00111111],
            b64[(blk >> 6)  & 0b00111111],
            '=',
        )
    }

	return string(out)

```

<p>
Quando a entrada nã é múltiplo de 3 significa que ao sairmos do laço principal restará 1 ou 2 bytes. A especificação do Base64 exige que tenhamos 4 grupos fechados de 6 bits, que resultam 4 caracteres da tabela. 
</p>
<p>
Quando sobrar apenas 1 byte, teremos que agrupar em blocos de 6 bits e ter um total de 4 blocos. Como obter um total de 4 blocos de 6 bits com apenas 8 bits? A especificação diz que temos que pegar os 6 bits e obter qual o valor da tabela, para os 2 bits que sobrarem colocamos zeros **à direita** para completar 6 bits e então obtemos o valor na tabela e para os dois grupos restantes simplesmente consideramos o valor '='.
</p>
<p>
Então temos que: O primeiro grupo de 6 bits do byte pega um valor normalmente da tabela, o outro valor é obtido completando os 2 bits que sobraram com quatro zeros à direita e os outros dois grupos de 6 bits consideramos o '=' para cada um.
</p>
<p>
Quando sobram 2 bytes a regra é a mesma. Isto é, dos 16 bits teremos 2 valores _cheios_ e os 4 bits que sobram recebem dois zeros à direita e, diferente do caso que sobra apenas 1 byte, nesse caso precisamos de apenas um '=' para completar o grupo de 4 valores uma vez que já temos 3 grupos prontos.
</p>
<p>
P.ex.:
</p>

<p>
Se tivermos 1 byte sobrando (A = 01000001 em binário = 65 em decimal), precismos transformar isso em 4 blocos de 6 bits. Então fazemos o seguinte:
</p>

<pre>
<span style="color:yellow">010000</span><span style="color:rebeccapurple">01</span><span style="color:red">0000</span><span style="color:lightblue">=</span><span style="color:lightgreen">=</span>
</pre>

<p>Para os dois primeiros (coincidentemente) teremos o mesmo valor na b64table ('Q') e para os outros dois teremos o '='. Assim, ficamos com</p>

<pre>
<strong>QQ==</strong>
</pre>


Com isso, c código completo fica assim:

```go
package main

import (
	"fmt"
	"encoding/base64"
)

const b64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func Encb64(in []byte) string {
	
    var out []byte
	len := len(in)
	rem := len % 3

    for i := 0; i < len - rem; i += 3 {
    
        blk := int64(in[i    ]) << 16 | 
               int64(in[i + 1]) << 8  | 
               int64(in[i + 2])

        out = append(out,
            b64[(blk >> 18) & 0b00111111],
            b64[(blk >> 12) & 0b00111111],
            b64[(blk >>  6) & 0b00111111],
            b64[ blk        & 0b00111111],
		)
	}

	if rem == 0 {
		return string(out)
	}
    
	var blk int64 = 0

 	if rem == 1 {
        blk = int64(in[len - rem]) << 16
        out = append(out,
            b64[(blk >> 18) & 0b00111111],
            b64[(blk >> 12) & 0b00111111],
            '=',
            '=',
        )
    } else if rem == 2 {
        blk = int64(in[len - rem]) << 16 | int64(in[len - rem + 1]) << 8
        out = append(out,
            b64[(blk >> 18) & 0b00111111],
            b64[(blk >> 12) & 0b00111111],
            b64[(blk >> 6)  & 0b00111111],
            '=',
        )
    }

	return string(out)
}

func main() {
    testCases := []string{"Manaed", "1234567890"}
	
    for _, test := range testCases {
		encoded := Encb64([]byte(test))
		encondedGo := base64.StdEncoding.EncodeToString([]byte(test))
		fmt.Println("Base64 de ", test, ": ", encoded)
		fmt.Println("Base64 Go : ", test, ": ", encondedGo)
		
	}
}
```

<p>Curiosidade: nossa versão fica alguns ms mais lenta que a versão do Go, sintam-se à vontade para estudar a versão em Go, agora poderão entender todo o código!</p>















