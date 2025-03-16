# Checksum e CRC - Verificação de Integridade


## Introdução
Na seção anterior, discutimos a codificação e o transporte de dados. Agora, avançamos para um aspecto essencial da transmissão e armazenamento de informações: a verificação de integridade. Esse processo garante que os dados recebidos são os mesmos que foram enviados, sem terem sido alterados ou corrompidos durante a transmissão ou armazenamento.

A integridade dos dados pode ser comprometida por diversos fatores, como erros humanos, falhas sistêmicas (por exemplo, setores defeituosos em discos rígidos) ou perdas de pacotes em redes devido a interferências ou congestionamento.

Para detectar essas falhas, foram desenvolvidos mecanismos como os checksums e os CRC (Cyclic Redundancy Check), que permitem identificar se os dados foram alterados de forma acidental.

É importante diferenciar esses métodos das funções hash criptográficas, como MD5 e SHA-256, que abordaremos mais adiante neste livro. Enquanto checksums e CRCs são projetados para detectar erros acidentais, funções hash criptográficas garantem a integridade contra modificações intencionais, oferecendo resistência a ataques e manipulação maliciosa.

Nesta seção, exploraremos como checksums simples podem ser utilizados para verificar integridade e, em seguida, como os CRC oferecem um método mais robusto para detectar erros em sistemas computacionais.

## Bit de Paridade (Parity Bit)

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


## Checksum

O checksum é a forma mais simples de verificação de integridade, sendo calculado somando os valores dos bytes ou palavras de um dado e armazenando essa soma junto com os dados. Ao receber ou ler os dados, a soma é recalculada e comparada com o checksum armazenado. Se os valores diferirem, isso indica um erro na transmissão ou no armazenamento.

### Exemplo de Checksum Simples (Soma de Bytes)

Vamos criar uma implementação básica de checksum em Go, somando os valores dos bytes da mensagem e calculando o módulo 256 para manter um valor fixo de 8 bits.

```go
package main

import (
	"fmt"
)

// Função de checksum simples (soma dos bytes módulo 256)
func SimpleChecksum(data []byte) byte {
	var checksum byte = 0
	for _, b := range data {
		checksum += b
	}
	return checksum
}

func main() {
	data := []byte("HELLO")
	fmt.Printf("Checksum de 'HELLO': 0x%X\n", SimpleChecksum(data))
}
```

Um algorítmo simples como o demonstrado acima possui algumas limitações importantes:

- Não detecta trocas na ordem dos bytes (ABC e CBA geram o mesmo checksum)
- Colisões serão frequentes (diferentes mensagens podem gerar o mesmo valor de checksum)
- Não detecta alguns tiopos de erro bit a bit

Para resolver esses problemas, foram criadas versões mais sofisticadas, como o CRC.

## Checksum no UDP/IP

O protocolo UDP (User Datagram Protocol) e o cabeçalho IP (Internet Protocol) utilizam um checksum baseado em complemento de um, que é uma técnica mais avançada do que a simples soma de bytes. Esse método melhora a detecção de erros, sendo amplamente usado na transmissão de pacotes de rede.

Como o checksum do UDP/IP funciona?

- A mensagem (ou cabeçalho IP) é dividida em blocos de 16 bits.
- Os blocos são somados usando aritmética de complemento de um.
- O complemento de um do resultado final é armazenado como o checksum.
- No destino, o receptor soma todos os blocos novamente, incluindo o checksum recebido.
- Se o resultado for 0xFFFF, os dados são considerados íntegros; caso contrário, há erro na transmissão.

Isso permite detectar a maioria dos erros comuns em transmissões de rede.

### Implementação do Checksum UDP/IP em Go

```go
package main

import (
	"encoding/binary"
	"fmt"
)

// Função para calcular o checksum do UDP/IP usando complemento de um
func UDPChecksum(data []byte) uint16 {
	var sum uint32

	// Processa os dados em blocos de 16 bits
	for i := 0; i < len(data)-1; i += 2 {
		word := binary.BigEndian.Uint16(data[i : i+2]) // Converte dois bytes em um uint16
		sum += uint32(word)
	}

	// Se o número de bytes for ímpar, adiciona o último byte com padding
	if len(data)%2 != 0 {
		sum += uint32(data[len(data)-1]) << 8
	}

	// Adiciona os "overflows" da soma (carry bits)
	for (sum >> 16) > 0 {
		sum = (sum & 0xFFFF) + (sum >> 16)
	}

	// Retorna o complemento de um do resultado final
	return uint16(^sum)
}

func main() {
	data := []byte("HELLOUDP")
	checksum := UDPChecksum(data)
	fmt.Printf("Checksum UDP de 'HELLOUDP': 0x%X\n", checksum)
}

```

Essa implementação reflete como os protocolos UDP e IP realizam a verificação de integridade, garantindo que os pacotes de rede cheguem corretamente ao destino.


## CRC - Cyclic Redundancy Check


O CRC (Cyclic Redundancy Check) é um método de detecção de erros baseado em operações matemáticas sobre polinômios binários. Em vez de apenas somar valores, o CRC trata os dados como um número binário e divide esse número por um polinômio pré-determinado, registrando o resto da divisão como código de verificação.

O CRC é amplamente usado em:

- Transmissões de rede (Ethernet, Wi-Fi)
- Armazenamento de arquivos (ZIP, RAR)
- Sistemas embarcados (CD-ROM, comunicações seriais)

### Implementação do CRC-8 em Go

O CRC-8 é uma das versões mais simples, usando um polinômio de 8 bits.


```go
package main

import (
	"fmt"
)

// Tabela CRC-8 para polinômio 0x07 (x^8 + x^2 + x + 1)
var crc8Table [256]byte

// Inicializa a tabela CRC-8
func init() {
	const poly = byte(0x07)
	for i := 0; i < 256; i++ {
		crc := byte(i)
		for j := 0; j < 8; j++ {
			if (crc & 0x80) != 0 {
				crc = (crc << 1) ^ poly
			} else {
				crc <<= 1
			}
		}
		crc8Table[i] = crc
	}
}

// Calcula CRC-8 para um conjunto de dados
func CRC8(data []byte) byte {
	crc := byte(0)
	for _, b := range data {
		crc = crc8Table[crc^b]
	}
	return crc
}

func main() {
	data := []byte("HELLO")
	fmt.Printf("CRC-8 de 'HELLO': 0x%X\n", CRC8(data))
}
```

O CRC, diferente do checksum, detecta trocas na ordem dos bytes, detecta erros simples de mútiplos bits errados e possui menor taxa de colisões. Mas o CRC-8 não é reistente a ataques maliciosos (não deve ser usado para segurança) e pode falhar na detecção de alguns padrões específicos de erro se o polinômio não for bem escolhido.


## Variações do CRC

CRC-16
CRC-32
CRC-64
Adler-32
Fletcher-16/32

Cada algoritmo tem um nível diferente de eficiência e resistência a erros. CRC-32, por exemplo, é usado no ZIP e no protocolo de rede Ethernet devido à sua confiabilidade e baixo custo computacional.

## Diferença entre Checksum, CRC e Funções Hash


### Conclusão


O checksum é uma técnica básica e eficiente para detecção de erros simples, mas tem falhas para detecção de padrões mais complexos. O CRC usa divisão polinomial para criar uma verificação mais robusta contra corrupção de dados, sendo amplamente utilizado em redes e armazenamento. Nenhum desses métodos deve ser confundido com funções hash criptográficas, que garantem segurança contra alterações intencionais nos dados.

Se queremos segurança contra modificações maliciosas, devemos usar HMAC, SHA-256 ou outras funções hash seguras.

Agora que entendemos essas formas de verificação de integridade fraca, podemos avançar para a família MD (Message Digest), que evolui de checksums simples para funções hash mais sofisticadas.
