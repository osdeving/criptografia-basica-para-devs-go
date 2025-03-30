# Base64

Base64 é uma codificação reversível que converte qualquer sequência de bytes em uma representação textual composta por um conjunto de 64 caracteres, conforme especificado na RFC 4648 [1]. Ela é amplamente utilizada para permitir o transporte de dados binários por meios que só aceitam texto, como e-mails (MIME), URLs ou JSON.

A codificação Base64 funciona agrupando a entrada em blocos de 3 bytes (24 bits) e dividindo-os em 4 grupos de 6 bits. Cada grupo de 6 bits é então convertido em um caractere da tabela Base64, a qual é formada da seguinte forma:

- Letras maiúsculas (A–Z) → índices 0 a 25
- Letras minúsculas (a–z) → índices 26 a 51
- Números (0–9) → índices 52 a 61
- Símbolos especiais `+` e `/` → índices 62 e 63

Em Go, podemos representar essa tabela como uma string:

```go
const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
```

Quando a entrada não é múltiplo exato de 3 bytes, a codificação adiciona preenchimento (o caractere `=`) ao final para manter consistência com o número de blocos.

Cabe ressaltar que Base64 **não é criptografia**. Qualquer pessoa pode decodificar sua saída e seu uso é exclusivamente para **representação e compatibilidade**.

## Exemplo: Codificando "Rox"

### Passo 1: Conversão para ASCII
```
R → 82
o → 111
x → 120
```

### Passo 2: Conversão para binário (8 bits por caractere)
```
82  → 01010010
o  → 01101111
x  → 01111000
```

Agrupados: `01010010 01101111 01111000` (24 bits)

### Passo 3: Dividir em 4 blocos de 6 bits

```
010100 100110 111101 111000
```

### Passo 4: Converter cada bloco para decimal

```
20, 38, 61, 56
```

### Passo 5: Mapear na tabela Base64

```
base64Table[20] = "U"
base64Table[38] = "m"
base64Table[61] = "9"
base64Table[56] = "4"
```

Resultado: `Um94`


> Para extrair os grupos de 6 bits, utilizamos **bitwise AND** com uma máscara `0b00111111`. Para uma explicação mais detalhada sobre manipulação de bits, consulte a seção [Operadores Bitwise](../manipulacao-de-dados/operadores-bitwise.md).


## Implementação em Go

```go
const b64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func Encb64(in []byte) string {
    var out []byte
    len := len(in)
    rem := len % 3

    for i := 0; i < len - rem; i += 3 {
        blk := int64(in[i]) << 16 | int64(in[i+1]) << 8 | int64(in[i+2])
        out = append(out,
            b64[(blk>>18)&0b111111],
            b64[(blk>>12)&0b111111],
            b64[(blk>>6) &0b111111],
            b64[blk&0b111111],
        )
    }

    if rem == 1 {
        blk := int64(in[len - 1]) << 16
        out = append(out,
            b64[(blk>>18)&0b111111],
            b64[(blk>>12)&0b111111], '=', '=',
        )
    } else if rem == 2 {
        blk := int64(in[len - 2]) << 16 | int64(in[len - 1]) << 8
        out = append(out,
            b64[(blk>>18)&0b111111],
            b64[(blk>>12)&0b111111],
            b64[(blk>>6) &0b111111], '=',
        )
    }

    return string(out)
}
```

### Explicação do código

Cada grupo de 3 bytes (24 bits) é combinado em um único inteiro (`blk`) usando operações de shift e OR. Por exemplo:

```go
blk := int64(in[i]) << 16 | int64(in[i+1]) << 8 | int64(in[i+2])
```

Dessa forma, `blk` conseguimos compactar os 3 bytes dentro de um inteiro, onde:
- o primeiro byte ocupa os bits 16 a 23
- o segundo byte ocupa os bits 8 a 15
- o terceiro byte ocupa os bits 0 a 7

A extração dos grupos de 6 bits é feita da esquerda para a direita com deslocamento (`>>`) e máscara (`& 0b111111`), isolando cada conjunto de 6 bits:

```go
b64_1 = (blk >> 18) & 0b111111
b64_2 = (blk >> 12) & 0b111111
b64_3 = (blk >> 6)  & 0b111111
b64_4 =  blk        & 0b111111
```

Esses índices são então usados diretamente para acessar os caracteres na tabela Base64.

Quando sobram bytes (i.e. 1 ou 2 bytes) estes são tratados separadamente, incluindo o padding `=` conforme especificação.

## Variante: Base64URL

Uma variação comum do Base64 é o Base64URL, também definida na RFC 4648. Ela foi projetada para ser segura em contextos como URLs, nomes de arquivos e parâmetros HTTP, e, em especial, com JWT (Json Web Tokens). As modificações são as seguintes: 

1) Substituir os caracteres:

    `+` por `-` e `/` por `_`

2) O caractere de preenchimento `=` pode ser omitido, já que o comprimento da string pode indicar os bits ausentes na decodificação.

A lógica de codificação permanece idêntica ao Base64 padrão; apenas a tabela de caracteres muda:

```go
const b64Url = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
```


## Considerações finais

- Base64 é uma transformação **reversível**.
- Aumenta o tamanho da entrada em cerca de **33%**.
- Útil para compatibilidade textual, não para segurança.

Outras codificações com finalidades semelhantes incluem Base32, Base85 e Hex.



## Referências

[1] RFC 4648 – The Base16, Base32, and Base64 Data Encodings. Internet Engineering Task Force (IETF). Disponível em: https://datatracker.ietf.org/doc/html/rfc4648

