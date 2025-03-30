# Base16 (Hexadecimal)

Base16, ou codificação **hexadecimal**, é uma forma de representar dados binários em texto usando 16 símbolos: os dígitos de `0` a `9` e as letras `A` a `F`. Cada byte (8 bits) é representado por dois caracteres hexadecimais — ou seja, 4 bits por caractere.

Essa codificação é amplamente utilizada em sistemas criptográficos para exibir valores binários de forma legível e padronizada, como hashes (MD5, SHA-1, SHA-256), chaves, IVs, e blocos de dados.

Está formalmente definida na RFC 4648 [1] como uma das representações canônicas para dados binários.

---

## Tabela de correspondência

| Binário | Hex |
|--------|-----|
| 0000   | 0   |
| 0001   | 1   |
| 0010   | 2   |
| 0011   | 3   |
| 0100   | 4   |
| 0101   | 5   |
| 0110   | 6   |
| 0111   | 7   |
| 1000   | 8   |
| 1001   | 9   |
| 1010   | A   |
| 1011   | B   |
| 1100   | C   |
| 1101   | D   |
| 1110   | E   |
| 1111   | F   |

---

## Exemplo: "Hi"

A string "Hi" possui dois bytes:

```
H → 0x48 → 01001000
i → 0x69 → 01101001
```

Convertendo byte a byte para hexadecimal:

```
0x48 → 4 e 8 → "48"
0x69 → 6 e 9 → "69"
```

Resultado final: `4869`

---

## Implementação em Go

```go
import (
	"encoding/hex"
	"fmt"
)

func main() {
	entrada := []byte("Hi")
	saida := hex.EncodeToString(entrada)
	fmt.Println("Hex: ", saida)
}
```

Essa é a forma canônica implementada na biblioteca padrão de Go, e pode ser usada para representar qualquer dado binário de forma segura, compacta e legível.

---

## Considerações

- Cada byte vira 2 caracteres → aumento de 100% no tamanho.
- Alta legibilidade e padronização.
- Forma mais comum de visualizar hashes e chaves.
- Diferente do Base64, **não precisa de padding**.

---

## Referências

[1] RFC 4648 – The Base16, Base32, and Base64 Data Encodings. IETF. Disponível em: https://datatracker.ietf.org/doc/html/rfc4648

