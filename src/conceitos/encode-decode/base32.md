# Base32

Base32 é uma codificação definida na RFC 4648 [1] que transforma dados binários em uma representação textual usando 32 caracteres seguros para transporte e legíveis. Ela é comumente utilizada em sistemas que impõem restrições a caracteres especiais, como nomes de arquivos, QR Codes, e na geração de segredos em algoritmos de autenticação como o TOTP (Time-based One-Time Password), amplamente empregado por aplicativos como o Google Authenticator.

A codificação Base32 opera agrupando a entrada binária em blocos de 5 bytes (40 bits) e os dividindo em 8 grupos de 5 bits. Cada grupo de 5 bits é mapeado para um caractere da tabela Base32, composta por:

- Letras maiúsculas (A–Z) → índices 0 a 25
- Dígitos (2–7) → índices 26 a 31

**Observação**: Os dígitos 0 e 1 são evitados para reduzir ambiguidades com as letras O e I, o que torna o Base32 mais robusto para leitura manual e OCR.

---

## Exemplo: Codificando "Ma"

"Ma" tem dois caracteres, ou seja, 2 bytes:

```
M → 0x4D → 01001101
a → 0x61 → 01100001
```

Concatenando os bits: `0100110101100001` (16 bits)

Agora separamos os bits em blocos de 5:

```
01001 10101 10000 1[padding]
```

Completamos com zeros à direita até ter múltiplo de 5 bits:

```
01001 10101 10000 10000
```

Convertendo para decimal:

```
9, 21, 16, 16
```

Usando a tabela Base32:

```go
const base32Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
```

Temos:
```
base32Table[9]  = J
base32Table[21] = V
base32Table[16] = Q
base32Table[16] = Q
```

Resultado: `JVQQ` (com padding = se aplicável)

---

## Implementação em Go (básica)

```go
import (
	"fmt"
	"strings"
)

const base32Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

func EncodeBase32(data []byte) string {
	var builder strings.Builder
	var buffer uint64
	var bits uint

	for _, b := range data {
		buffer = (buffer << 8) | uint64(b)
		bits += 8
		for bits >= 5 {
			index := (buffer >> (bits - 5)) & 0b11111
			builder.WriteByte(base32Table[index])
			bits -= 5
		}
	}

	if bits > 0 {
		index := (buffer << (5 - bits)) & 0b11111
		builder.WriteByte(base32Table[index])
	}

	for builder.Len()%8 != 0 {
		builder.WriteByte('=')
	}

	return builder.String()
}

func main() {
	entrada := []byte("Ma")
	fmt.Println("Base32: ", EncodeBase32(entrada))
}
```

---

## Considerações

- É mais verboso que o Base64: a saída é cerca de 60% maior que a original.
- Melhor tolerância à digitação e OCR.
- Amplamente usado em aplicações como **TOTP** (RFC 6238) e **secrets URI** (RFC 3548 / 4648).

---

## Referências

[1] RFC 4648 – The Base16, Base32, and Base64 Data Encodings. IETF. Disponível em: https://datatracker.ietf.org/doc/html/rfc4648

