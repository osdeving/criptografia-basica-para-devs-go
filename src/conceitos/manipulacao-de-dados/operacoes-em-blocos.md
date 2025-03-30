# Operações em Blocos

## O que são modos de operação?

Muitos algoritmos lidam com blocos de dados de tamanho fixo, como 64 ou 128 bits. Quando a entrada é maior que um único bloco, é necessário dividir os dados e aplicar transformações **bloco a bloco**, seguindo um **modo de operação**.

Embora esses modos sejam usados em contextos criptográficos, o objetivo aqui é entender **como os dados são processados tecnicamente**, sem considerar a lógica de cifra. Vamos nos concentrar em:

- Como os blocos são encadeados
- Qual a dependência entre entradas e saídas
- Como o estado interno evolui
- Como entrada, saída e IV interagem

Para fins ilustrativos, faremos comentários como `// aplicar cifra aqui`, mas **não implementaremos cifra real nesta seção**.


## Divisão em blocos

Antes de aplicar qualquer modo de operação, os dados precisam ser divididos em blocos de tamanho fixo. Se os dados não são múltiplo do tamanho do bloco, algum tipo de preenchimento (padding) será necessário. Veremos mais sobre padding em uma seção posterior, em especial quando estivermos falando sobre Base64.

Exemplo com blocos de 8 bytes:

```
Entrada (24 bytes): "ExemploDeTextoCriptografado"

Blocos:
B1 = "ExemploD"
B2 = "eTextoCr"
B3 = "iptograf"
B4 = "ado\x00\x00\x00" // padding se aplicado
```

A seguir, vamos examinar os principais modos de operação: ECB, CBC, CFB, OFB, CTR e outros opcionais.


## ECB — Electronic Codebook

O modo ECB é o mais simples de todos: cada bloco de entrada é processado **de forma independente** dos demais. Não há encadeamento, nem dependência de bloco anterior. Por isso, é possível aplicar a transformação a todos os blocos **em paralelo**.

Uma analogia útil é pensar em ECB como um carimbo: para cada pedaço do texto, aplica-se sempre o mesmo molde. Se dois blocos forem iguais, o resultado será o mesmo.

Apesar de sua simplicidade, esse comportamento pode expor padrões quando blocos repetidos resultam em saídas repetidas — por isso o ECB é desaconselhado em aplicações criptográficas. Mas como modelo de operação, é útil para entender a base dos modos.

### Funcionamento

```
Entrada:     B1   B2   B3   B4
             |    |    |    |
             v    v    v    v
Saída:      C1   C2   C3   C4

// Cada Ci = cifra(Bi)
```

### Implementação simulada em Go

```go
package main

import "fmt"

func ecb(blocks []string) []string {
	out := make([]string, len(blocks))
	for i, b := range blocks {
		// Simulando uma "cifra" com reverso de string
		out[i] = reverse(b) // aplicar cifra real aqui
	}
	return out
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	blocos := []string{"ExemploD", "eTextoCr", "iptograf", "ado\x00\x00\x00"}
	saida := ecb(blocos)
	for i, b := range saida {
		fmt.Printf("C%d: %q\n", i+1, b)
	}
}
```

Esse código demonstra a ideia de que cada bloco é transformado isoladamente. A função `reverse` simula uma transformação sem depender de contexto.

Nos próximos modos, veremos como introduzir **estado** e **encadeamento** entre os blocos.


## CBC — Cipher Block Chaining

O modo CBC introduz encadeamento entre os blocos de dados. Aqui, cada bloco de entrada é combinado (via `XOR`) com o bloco de saída anterior **antes** da transformação ser aplicada. Isso garante que blocos idênticos de entrada não gerem saídas idênticas, desde que a inicialização seja diferente.

O primeiro bloco é combinado com um **vetor de inicialização**  (_IV - Initialization Vector_ ), que deve ser único e imprevisível para cada mensagem.

Pense em CBC como uma corrente: cada elo (bloco) depende do anterior. Se você mudar um elo, o restante da corrente será alterado. Esse encadeamento dá mais segurança estrutural ao fluxo de blocos.

### Funcionamento

```
Entrada:     B1   B2   B3   B4
              |    |    |    |
IV ---> XOR   |    |    |    |
        v     v    v    v    v
      cifra  XOR  XOR  XOR  ...
        |     |    |    |    |
       C1 -> C2 -> C3 -> C4

// C1 = cifra(B1 ⊕ IV)
// C2 = cifra(B2 ⊕ C1)
// C3 = cifra(B3 ⊕ C2)
// C4 = cifra(B4 ⊕ C3)
```

### Implementação simulada em Go

```go
package main

import "fmt"

func xor(a, b string) string {
	out := make([]byte, len(a))
	for i := range a {
		out[i] = a[i] ^ b[i]
	}
	return string(out)
}

func cbc(blocks []string, iv string) []string {
	out := make([]string, len(blocks))
	prev := iv
	for i, b := range blocks {
		mix := xor(b, prev)
		out[i] = reverse(mix) // aplicar cifra real aqui
		prev = out[i]         // saída vira entrada para o próximo XOR
	}
	return out
}

func main() {
	blocos := []string{"ExemploD", "eTextoCr", "iptograf", "ado\x00\x00\x00"}
	iv := "12345678" // mesmo tamanho de bloco
	saida := cbc(blocos, iv)
	for i, b := range saida {
		fmt.Printf("C%d: %q\n", i+1, b)
	}
}
```

Aqui, usamos uma função `xor` simples para combinar blocos. A função `reverse`, como antes, simula a aplicação de uma cifra. O encadeamento é visível no uso da variável `prev`, que propaga o resultado entre os blocos.

No próximo modo, veremos como esse encadeamento pode ser reorganizado para permitir recuperação de dados em tempo real.

# CFB — Cipher Feedback

O modo CFB transforma uma cifra de bloco em um sistema que se comporta como uma cifra de fluxo. Diferente do ECB e CBC, o CFB permite operar em unidades menores que o bloco (como bytes ou bits), o que é útil para aplicações onde os dados não chegam em blocos completos.

Apesar de utilizar uma cifra de bloco internamente, o que realmente é cifrado é o **vetor de realimentação**. Esse vetor é processado com a cifra, e o resultado é combinado (via XOR) com o bloco ou byte de entrada para gerar a saída.

### Funcionamento (versão em blocos inteiros)

```
Entrada:     P1     P2     P3     P4
              |      |      |      |
IV -----> cifra     |      |      |
           |         |      |      |
           v         v      v      v
          XOR       XOR   XOR   XOR
           |         |      |      |
          C1  --->  C2 ->  C3 ->  C4

// C1 = P1 ⊕ cifra(IV)
// C2 = P2 ⊕ cifra(C1)
// C3 = P3 ⊕ cifra(C2)
// C4 = P4 ⊕ cifra(C3)
```

Diferente do CBC, a cifra é aplicada sobre a saída anterior, não sobre a entrada. Isso permite que a cifra seja usada como um gerador de fluxo pseudoaleatório que é então combinado com os dados reais.

### Implementação simulada em Go

```go
package main

import "fmt"

func xor(a, b string) string {
	out := make([]byte, len(a))
	for i := range a {
		out[i] = a[i] ^ b[i]
	}
	return string(out)
}

func cfb(entrada []string, iv string) []string {
	saida := make([]string, len(entrada))
	estado := iv
	for i, p := range entrada {
		keystream := reverse(estado) // aplicar cifra real aqui
		saida[i] = xor(p, keystream)
		estado = saida[i] // realimentado com saída
	}
	return saida
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	blocos := []string{"ExemploD", "eTextoCr", "iptograf", "ado\x00\x00\x00"}
	iv := "12345678"
	saida := cfb(blocos, iv)
	for i, b := range saida {
		fmt.Printf("C%d: %q\n", i+1, b)
	}
}
```

Neste exemplo:
- `reverse` representa a cifra aplicada sobre o estado interno
- `estado` é atualizado com a saída anterior, simulando o feedback
- `xor` combina a saída da cifra com o bloco de entrada

Nos modos seguintes, veremos como separar o fluxo gerado da entrada original (OFB) e como introduzir contadores (CTR).

# OFB — Output Feedback

O modo OFB (Output Feedback) é semelhante ao CFB, mas com uma diferença importante: o **vetor de realimentação é atualizado independentemente da entrada**. Isso significa que o fluxo pseudoaleatório gerado pela cifra é completamente dissociado do texto de entrada, permitindo paralelismo na geração e reutilização do keystream.

Na prática, o OFB transforma uma cifra de bloco em uma cifra de fluxo pura, onde o texto de entrada é combinado com uma sequência pseudoaleatória de bits gerados iterativamente a partir do IV.

### Funcionamento

```
Entrada:     P1     P2     P3     P4
              |      |      |      |
IV -----> cifra ---> cifra ---> cifra --->
              |       |       |      |
              v       v       v      v
            XOR     XOR     XOR    XOR
              |       |       |      |
             C1      C2      C3     C4

// E1 = cifra(IV)
// E2 = cifra(E1)
// E3 = cifra(E2)
// E4 = cifra(E3)
// Ci = Pi ⊕ Ei
```

Como o OFB não depende da entrada anterior, ele é ideal para ambientes em que erros de transmissão não podem propagar corrompendo blocos seguintes.

### Implementação simulada em Go

```go
package main

import "fmt"

func xor(a, b string) string {
	out := make([]byte, len(a))
	for i := range a {
		out[i] = a[i] ^ b[i]
	}
	return string(out)
}

func ofb(entrada []string, iv string) []string {
	saida := make([]string, len(entrada))
	estado := iv
	for i, p := range entrada {
		keystream := reverse(estado) // aplicar cifra real aqui
		saida[i] = xor(p, keystream)
		estado = keystream // fluxo é autônomo
	}
	return saida
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	blocos := []string{"ExemploD", "eTextoCr", "iptograf", "ado\x00\x00\x00"}
	iv := "12345678"
	saida := ofb(blocos, iv)
	for i, b := range saida {
		fmt.Printf("C%d: %q\n", i+1, b)
	}
}
```

Neste exemplo:
- O `estado` é atualizado **somente pela cifra anterior**, sem depender da entrada
- O fluxo gerado é previsível e determinístico para o mesmo IV

O modo OFB é adequado para aplicações onde a integridade parcial dos dados deve ser preservada mesmo em caso de erros de transmissão.

No próximo modo, introduziremos um contador para gerar blocos independentes — o modo CTR.

# CTR — Counter Mode

O modo CTR (Counter) transforma uma cifra de bloco em uma cifra de fluxo por meio da geração de um **keystream determinístico**, baseado em um contador. Ao invés de realimentar saídas anteriores, ele cifra valores sequenciais derivados de um IV (Initialization Vector).

Cada bloco de texto é combinado (via XOR) com a cifra do contador correspondente. Isso permite que cada bloco seja processado **de forma independente**, facilitando a paralelização.

### Funcionamento

```
Entrada:     P1     P2     P3     P4
              |      |      |      |
IV+0 --> cifra        
         |           |      |      |
         v           v      v      v
        XOR         XOR   XOR    XOR
         |           |      |      |
        C1          C2     C3     C4

// C1 = P1 ⊕ cifra(IV + 0)
// C2 = P2 ⊕ cifra(IV + 1)
// C3 = P3 ⊕ cifra(IV + 2)
// C4 = P4 ⊕ cifra(IV + 3)
```

Esse modo é muito utilizado por sua eficiência e por permitir operações em paralelo. Como o contador é previsível, os blocos de keystream podem ser gerados antes mesmo da chegada da mensagem.

### Implementação simulada em Go

```go
package main

import (
	"fmt"
	"strconv"
)

func xor(a, b string) string {
	out := make([]byte, len(a))
	for i := range a {
		out[i] = a[i] ^ b[i]
	}
	return string(out)
}

func padCounter(base string, count int) string {
	suffix := fmt.Sprintf("%08d", count)
	return base[:len(base)-len(suffix)] + suffix
}

func ctr(entrada []string, iv string) []string {
	saida := make([]string, len(entrada))
	for i, p := range entrada {
		counter := padCounter(iv, i)
		keystream := reverse(counter) // aplicar cifra real aqui
		saida[i] = xor(p, keystream)
	}
	return saida
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	blocos := []string{"ExemploD", "eTextoCr", "iptograf", "ado\x00\x00\x00"}
	iv := "0000000000000000" // deve ter tamanho fixo
	saida := ctr(blocos, iv)
	for i, b := range saida {
		fmt.Printf("C%d: %q\n", i+1, b)
	}
}
```

Nesta simulação:
- O contador é representado por um sufixo decimal no `IV`
- A função `reverse` simula a cifra aplicada ao contador
- Cada bloco é processado com base em um valor de contador diferente

O CTR é amplamente utilizado em prática (por exemplo, em AES-CTR), sendo considerado seguro e eficiente quando o `IV` não é reutilizado com a mesma chave.


# XTS — XEX-based Tweaked Codebook Mode with Ciphertext Stealing

O modo XTS foi projetado especificamente para cifragem de **unidades de armazenamento em blocos**, como discos e SSDs. Ele combina uma cifra de bloco (como AES) com uma estrutura especial chamada **tweak**, que personaliza a operação para cada setor ou bloco físico, evitando repetições previsíveis.

XTS é um modo baseado em **XEX (XOR-Encrypt-XOR)** com **roubo de texto cifrado (ciphertext stealing)**, permitindo lidar com entradas que não são múltiplas do tamanho do bloco sem recorrer a padding.

> XTS é o modo de operação recomendado por padrões como IEEE P1619 para cifragem de dados em discos.

### Estrutura geral

- Usa **duas chaves**: uma para a cifra principal (`K1`) e outra para o tweak (`K2`)
- Gera um tweak a partir do número de setor ou posição lógica do bloco
- Aplica o tweak com XOR antes e depois da cifra (XOR-Encrypt-XOR)

### Funcionamento simplificado

```
Para cada bloco i:

Tweak_i = cifra(K2, bloco_logico) << i // ajustado por multiplicador

M_i' = P_i ⊕ Tweak_i
C_i' = cifra(K1, M_i')
C_i  = C_i' ⊕ Tweak_i
```

Essa construção garante que dois blocos com mesmo conteúdo, mas localização diferente no disco, produzam saídas diferentes. Além disso, XTS suporta entradas que não completam um bloco inteiro usando roubo de bits do bloco seguinte (ciphertext stealing).

### Simulação conceitual

Para simplificar, podemos demonstrar apenas a ideia do uso de tweak e da dupla XOR:

```go
package main

import "fmt"

func xor(a, b string) string {
	out := make([]byte, len(a))
	for i := range a {
		out[i] = a[i] ^ b[i]
	}
	return string(out)
}

func xts(blocks []string, tweakBase string) []string {
	saida := make([]string, len(blocks))
	for i, p := range blocks {
		tweak := genTweak(tweakBase, i)
		m1 := xor(p, tweak)
		cifra := reverse(m1) // cifra principal simulada
		saida[i] = xor(cifra, tweak)
	}
	return saida
}

func genTweak(base string, i int) string {
	b := []byte(base)
	for j := range b {
		b[j] = b[j] ^ byte(i) // mistura simples com índice
	}
	return string(b)
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	blocos := []string{"ExemploD", "eTextoCr", "iptograf", "ado\x00\x00\x00"}
	tweak := "SETORESP"
	saida := xts(blocos, tweak)
	for i, b := range saida {
		fmt.Printf("C%d: %q\n", i+1, b)
	}
}
```

### Características

- Suporta operação paralela
- Seguro contra repetições de padrões em setores
- Recomendado para cifragem de disco e armazenamento

Embora mais complexo que os modos anteriores, o XTS resolve limitações práticas importantes na cifragem de dispositivos reais.

Com isso, encerramos os principais modos de operação. A seguir, faremos uma comparação entre eles, destacando características técnicas e aplicações ideais.

# Comparação entre Modos de Operação

A análise técnica dos modos de operação não deve ser feita apenas com base em características isoladas, mas também considerando seu contexto histórico e a motivação por trás de cada proposta. A seguir, apresentamos uma visão evolutiva e comparativa entre os modos ECB, CBC, CFB, OFB, CTR e XTS.


## Evolução histórica e motivação

O primeiro modo amplamente utilizado foi o **ECB (Electronic Codebook)**. Seu funcionamento direto e paralelizável o tornava atraente, mas rapidamente ficou claro que ele era **inadequado para dados estruturados**: blocos iguais produziam saídas idênticas, expondo padrões — o que compromete completamente a segurança.

Para contornar esse problema, surgiu o **CBC (Cipher Block Chaining)**. Ele introduziu **encadeamento entre os blocos**: cada bloco depende do anterior. Isso impede que blocos idênticos de entrada produzam saídas idênticas, mesmo com a mesma chave. Contudo, essa segurança adicional vem ao custo da **impossibilidade de paralelismo na cifragem** — já que é preciso aguardar o resultado do bloco anterior.

O modo **CFB (Cipher Feedback)** foi criado como alternativa ao CBC para contextos de **transmissão de dados em tempo real**, permitindo cifragem de **tamanhos menores que o bloco** (como bytes ou bits). Seu encadeamento ocorre via saída anterior, o que o torna seguro, mas ainda dependente do processamento sequencial.

Já o **OFB (Output Feedback)** nasce como uma variação do CFB, buscando resolver a propagação de erros: ao invés de realimentar com a saída cifrada, o OFB **gera um fluxo independente da entrada**, tornando o sistema mais tolerante a falhas e adequado para ambientes ruidosos (como transmissões digitais).

O modo **CTR (Counter Mode)** representa uma mudança significativa: elimina o encadeamento e adota um **contador como fonte de entropia**. Cada bloco da entrada é combinado com a cifra do contador correspondente. Com isso, torna-se **totalmente paralelizável**, ideal para ambientes de alto desempenho, redes e sistemas distribuídos. A cifragem pode inclusive ser precomputada, desde que o contador e a chave sejam fixos.

Por fim, o modo **XTS** surge para atender **um problema prático não resolvido por nenhum dos anteriores**: a cifragem segura de blocos físicos em discos. Ele se baseia em XEX (XOR-Encrypt-XOR), utiliza dois conjuntos de chaves, gera um tweak exclusivo por posição e ainda lida com blocos incompletos sem padding (via ciphertext stealing). XTS é, portanto, uma construção especializada, recomendada por padrões como o IEEE P1619.


## Relações e compensações

Cada modo surgiu como resposta a uma limitação do anterior:
- **CBC** corrige a repetição de padrões do **ECB**.
- **CFB** adapta o CBC para trabalhar em unidades menores.
- **OFB** remove a propagação de erros do CFB.
- **CTR** substitui todos os anteriores quando o foco é paralelismo e desempenho.
- **XTS** trata o caso específico de cifragem de setores físicos.

É importante compreender essas relações para fazer escolhas adequadas em projetos reais. Modos de operação não são equivalentes: cada um atende a um cenário específico de uso e impõe restrições distintas.

Na próxima seção, faremos considerações finais sobre segurança, desempenho e práticas recomendadas.

# Considerações Finais sobre Modos de Operação

Os modos de operação analisados neste capítulo evidenciam como decisões de projeto em criptografia envolvem trocas entre **segurança**, **eficiência**, **flexibilidade** e **resiliência a falhas**. Nenhum modo é universalmente superior: cada um resolve um subconjunto de problemas e impõe restrições específicas.

O objetivo deste capítulo foi **isolar os modos de operação**, apresentando suas características técnicas e funcionais de forma independente de algoritmos específicos. A relação com cifras reais (como AES ou ChaCha20), bem como as **implicações práticas de segurança** na escolha do modo, serão retomadas e aprofundadas em seções posteriores do livro.

## Sobre segurança

- **ECB** deve ser evitado em qualquer situação onde haja estrutura ou repetição nos dados. Seu uso é restrito a testes ou casos didáticos.
- **CBC**, **CFB** e **OFB** oferecem segurança razoável quando combinados com IVs aleatórios e não reutilizados.
- **CTR** e **XTS** são considerados seguros e modernos, desde que o **contador ou tweak não se repita** com a mesma chave.

## Sobre desempenho e paralelismo

- **CTR** e **XTS** são os únicos modos que permitem paralelismo total na cifra e decifra.
- **ECB** também é paralelizável, mas inseguro.
- **CBC**, **CFB** e **OFB** são sequenciais na cifra e, portanto, menos eficientes para grandes volumes de dados ou ambientes concorrentes.

## Sobre robustez

- **OFB** e **CTR** não propagam erros em cadeia, tornando-os ideais para canais com ruído.
- **CBC** e **CFB** propagam erros de forma localizada, e podem corromper blocos subsequentes na decifra.

## Escolha prática

A escolha do modo de operação deve considerar:
- A natureza dos dados (texto, fluxo, armazenamento físico)
- A sensibilidade a erros de transmissão
- A necessidade de desempenho e paralelismo
- As garantias de segurança esperadas

| Caso de uso                 | Modo recomendado |
|----------------------------|------------------|
| Criptografia de disco      | XTS              |
| Comunicação em tempo real  | CFB ou OFB       |
| Alta performance em redes  | CTR              |
| Arquivos e backups         | CBC              |

Em todos os casos, o uso de **vetores de inicialização (IVs)** únicos e imprevisíveis é essencial para garantir a segurança dos modos baseados em encadeamento ou contadores.

Com isso, concluímos a análise dos modos de operação. As próximas seções tratarão da implementação de cifras reais e da integração com essas estruturas de controle de fluxo.

