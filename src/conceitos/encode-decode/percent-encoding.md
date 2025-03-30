# Percent-Encoding (URL Encoding)

O **Percent-Encoding**, também conhecido como **URL encoding**, é um mecanismo definido pela RFC 3986 [1] para representar caracteres arbitrários em URLs utilizando apenas um subconjunto seguro de ASCII. Esse método codifica bytes utilizando o símbolo `%` seguido de dois dígitos hexadecimais (0–9, A–F) que representam o valor do byte.

Essa codificação é comumente usada em:

- URLs e URIs
- Parâmetros de requisições HTTP
- JSON Web Tokens (JWT) no modo compact
- Formulários HTML (`application/x-www-form-urlencoded`)

---

## Exemplo

Se quisermos transmitir o valor `senha=meu segredo` como parte de uma URL:

```
senha=meu segredo
```

Espaços e caracteres especiais são codificados:

```
senha=meu%20segredo
```

Outro exemplo, com caracteres não-ASCII:

```
entrada: café
```

O caractere `é` (0xC3A9 em UTF-8) é convertido para:

```
caf%C3%A9
```

---

## Caracteres codificados obrigatoriamente

Qualquer caractere fora do conjunto seguro é codificado. O conjunto seguro inclui:

- Letras (`A`–`Z`, `a`–`z`)
- Dígitos (`0`–`9`)
- Caracteres especiais: `-`, `_`, `.`, `~`

Tudo mais é substituído por `%XX`, onde `XX` é o byte em hexadecimal.

---

## Implementação em Go

A biblioteca padrão oferece suporte completo via o pacote `net/url`:

```go
import (
	"fmt"
	"net/url"
)

func main() {
	entrada := "café & açúcar"
	encoded := url.QueryEscape(entrada)
	fmt.Println("Percent-encoded:", encoded)
}
```

Saída:

```
Percent-encoded: caf%C3%A9+%26+a%C3%A7%C3%BAcar
```

Obs: Espaços podem ser representados como `+` (formulário) ou `%20` (padrão URI).

## Como a codificação %XX é construída

A codificação Percent-Encoding é determinada diretamente pelo valor binário de cada byte. Se o byte não pertence ao conjunto de caracteres seguros, ele será convertido para sua representação hexadecimal de dois dígitos, prefixada por %. Por exemplo, o caractere @ (código decimal 64) é codificado como %40 porque 0x40 é a representação hexadecimal de 64.

Essa transformação segue a fórmula:

```
% + UPPERCASE(HEX(byte))

```

## Considerações

- Reversível: a decodificação é trivial.
- Não fornece sigilo ou autenticidade.
- Essencial para interoperabilidade com HTTP e navegadores.
- Aparece como pré-processamento em várias aplicações web.

---

## Referências

[1] RFC 3986 – Uniform Resource Identifier (URI): Generic Syntax. IETF. Disponível em: https://datatracker.ietf.org/doc/html/rfc3986

