# Concurso e escolha do AES

## Motivação para um novo padrão

O **Data Encryption Standard (DES)** foi, por décadas, o principal algoritmo de cifra de bloco utilizado globalmente. Contudo, a limitação da chave de 56 bits tornou-o inseguro frente aos avanços computacionais. Em 1998, a Electronic Frontier Foundation (EFF) demonstrou que um ataque de força bruta contra o DES era factível com hardware dedicado.

Como resposta, o **National Institute of Standards and Technology (NIST)** iniciou, em 1997, o processo público para a seleção de um novo padrão criptográfico — o **Advanced Encryption Standard (AES)**.

---

## Requisitos estabelecidos pelo NIST

O NIST definiu requisitos técnicos e políticos rigorosos:

* Bloco de dados fixo de **128 bits**.
* Suporte a chaves de **128, 192 e 256 bits**.
* Algoritmo seguro contra criptoanálise diferencial, linear e por chave relacionada.
* Eficiência em hardware e software, em diferentes arquiteturas.
* Clareza na descrição e ausência de propriedade intelectual.
* **Transparência no processo de avaliação**.

> “Our intent is to evaluate these algorithms in an open and public process.”
> — *FIPS 197 Introduction*

---

## Linha do tempo do concurso AES

```plaintext
1997 ───────────── 1998 ───────────── 1999 ───────────── 2000 ───────────── 2001
 ↓                  ↓                  ↓                  ↓                  ↓
Chamada pública     15 candidatos      5 finalistas       Rijndael escolhido FIPS 197 publicado
```

---

## Etapas do processo

### 1. Submissão inicial

15 algoritmos foram submetidos, incluindo candidatos com estruturas variadas como SPN, Feistel, redes híbridas e operações em campos finitos.

### 2. Seleção dos finalistas

Com base em critérios técnicos, o NIST selecionou os seguintes algoritmos:

| Algoritmo    | Proponentes              | Estrutura    | Origem |
| ------------ | ------------------------ | ------------ | ------ |
| **MARS**     | IBM                      | Híbrida      | EUA    |
| **RC6**      | RSA Laboratories         | Feistel mod. | EUA    |
| **Serpent**  | Anderson, Biham, Knudsen | SPN          | Europa |
| **Twofish**  | Schneier et al.          | Feistel      | EUA    |
| **Rijndael** | Daemen & Rijmen          | SPN          | Europa |

---

## Comparação técnica dos finalistas

| Algoritmo    | Segurança  | Desempenho (SW) | Simplicidade | Flexibilidade |
| ------------ | ---------- | --------------- | ------------ | ------------- |
| MARS         | Alta       | Baixo           | Média        | Alta          |
| RC6          | Média      | Alta            | Média        | Média         |
| Serpent      | Muito alta | Média           | Baixa        | Baixa         |
| Twofish      | Alta       | Alta            | Média        | Alta          |
| **Rijndael** | Alta       | **Muito alta**  | **Alta**     | Alta          |

> “Rijndael offers an elegant design with strong security and excellent performance on a wide range of platforms.”
> — *NIST AES Report, 2000*

---

## Razões para a escolha do Rijndael

* Base matemática sobre $\mathbb{F}_{2^8}$, com operações bem definidas e eficientes.
* Baixa complexidade computacional: nenhuma operação condicional nem tabelas grandes.
* Boa escalabilidade em chaves e blocos.
* Alta paralelizabilidade: operações em bytes independentes.
* Performance superior em CPUs de 8, 16, 32 e 64 bits.
* Estrutura modular que facilita implementação segura.

---

## Formalização

O algoritmo Rijndael foi oficializado como **AES** em outubro de 2000, sendo padronizado como:

> **FIPS PUB 197 – Specification for the Advanced Encryption Standard (AES)**
> Publicado em 26 de novembro de 2001. Disponível em:
> [https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.197.pdf](https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.197.pdf)

---

## Impacto e legado

O concurso AES estabeleceu um novo modelo para a seleção de algoritmos criptográficos:

* Processo aberto e auditável, com participação da comunidade.
* Documentação formal e ampla revisão científica.
* Iniciativa replicada posteriormente em concursos como o SHA-3 e post-quantum (NIST PQC).

O AES é hoje a cifra padrão em protocolos como:

* TLS (HTTPS)
* IPsec
* SSH
* WPA2/WPA3
* OpenPGP

---

> O processo de seleção do AES consolidou um novo paradigma na criptografia: algoritmos públicos, seguros por design, e legitimados pela revisão científica.
