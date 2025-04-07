# Sumário

[capa](capa.md)
[Prefácio](prefacio.md)
[Introdução](README.md)


# Conceitos Fundamentais
- [Manipulação de Dados](conceitos/manipulacao-de-dados/manipulacao-de-dados.md)
    - [Introdução](conceitos/manipulacao-de-dados/introduction.md)
    - [Endianness](conceitos/manipulacao-de-dados/endianness.md)
    - [Operadores Bitwise](conceitos/manipulacao-de-dados/operadores-bitwise.md)
    - [Operações em Blocos](conceitos/manipulacao-de-dados/operacoes-em-blocos.md)
    - [Confusão e Difusão](conceitos/manipulacao-de-dados/confusao-difusao.md)
    - [S-Box](conceitos/manipulacao-de-dados/s-box.md)
    - [Permutações e Trocas](conceitos/manipulacao-de-dados/permutacoes.md)
- [Codificação de Dados: Representações Reversíveis para Armazenamento e Transmissão](conceitos/encode-decode/codification.md)
    - [Introdução](conceitos/encode-decode/introduction.md)
    - [Base64](conceitos/encode-decode/base64-2.md)
    - [Base32](conceitos/encode-decode/base32.md)
    - [Base16](conceitos/encode-decode/base16.md)
    - [Base85](conceitos/encode-decode/base85.md)
    - [Base58](conceitos/encode-decode/base58.md)
    - [Percent-Encoding](conceitos/encode-decode/percent-encoding.md)
- [Veirificação de Integridade de Dados: Checksum, Bit de Paridade e CRC](conceitos/checksum-crc/data-integrity.md)
    - [Introdução](conceitos/checksum-crc/introduction.md)
    - [Bit de Paridade](conceitos/checksum-crc/parity-bit.md)
    - [Checksum](conceitos/checksum-crc/checksum.md)
    - [CRC - Cyclic Redundancy Check](conceitos/checksum-crc/crc.md)
- [Funções Hash Criptográficas](conceitos/hash/sha256.md)
    - [Funções Hash](conceitos/hash/sha256.md)
    - [Família de Algoritmos MD - Message Digest](conceitos/hash/message-digest.md)
    - [Família de Algoritmos DES - Data Encryption Standard](conceitos/hash/des.md)
    - [Família de Algoritmos SHA - Secure Hash Algorithms](conceitos/hash/sha.md)
    - [PBKDF2, Argon2 e Scrypt - Hash para Senhas](conceitos/hash/pbkdf2-argon2.md)
- [Criptografia Simétrica](criptografia-simetrica/index.md)
  
  - [Introdução e Contexto](criptografia-simetrica/introducao-contexto/index.md)
    - [Breve história da criptografia moderna](criptografia-simetrica/introducao-contexto/historia.md)
    - [Fluxo ou bloco: modos distintos de cifrar](criptografia-simetrica/introducao-contexto/fluxo-vs-bloco.md)
    
  - [Cifras de Fluxo](criptografia-simetrica/cifras-de-fluxo/index.md)
    - [Conceito e aplicações](criptografia-simetrica/cifras-de-fluxo/conceito.md)
    - [Geradores de fluxo pseudoaleatório](criptografia-simetrica/cifras-de-fluxo/prng.md)
    - [XOR como primitiva central](criptografia-simetrica/cifras-de-fluxo/xor.md)
    - [RC4: funcionamento, popularidade e quedas](criptografia-simetrica/cifras-de-fluxo/rc4.md)
    - [Cifras modernas: ChaCha e Salsa](criptografia-simetrica/cifras-de-fluxo/chacha-salsa.md)
    - [RC4 em Go](criptografia-simetrica/cifras-de-fluxo/rc4-go.md)
    - [Segurança em cifras de fluxo](criptografia-simetrica/cifras-de-fluxo/seguranca.md)
  
  - [Cifras de Bloco](criptografia-simetrica/cifras-de-bloco/index.md)
    - [O que é uma cifra de bloco](criptografia-simetrica/cifras-de-bloco/o-que-e.md)
    - [A estrutura de Feistel](criptografia-simetrica/cifras-de-bloco/feistel.md)
    - [DES: arquitetura, S-boxes e permutações](criptografia-simetrica/cifras-de-bloco/des.md)
    - [Limitações e abandono do DES](criptografia-simetrica/cifras-de-bloco/queda-des.md)
    - [3DES: extensão temporária](criptografia-simetrica/cifras-de-bloco/3des.md)
    - [Modos de operação](criptografia-simetrica/cifras-de-bloco/modos-operacao.md)
    - [RC5 e RC6](criptografia-simetrica/cifras-de-bloco/rc5-rc6.md)
    - [AES: concurso e escolha](criptografia-simetrica/cifras-de-bloco/aes-concurso.md)
    - [AES: funcionamento interno](criptografia-simetrica/cifras-de-bloco/aes-interno.md)
    - [AES em modos de operação](criptografia-simetrica/cifras-de-bloco/aes-modos.md)
    - [AES em Go](criptografia-simetrica/cifras-de-bloco/aes-go.md)
    - [Armadilhas e boas práticas](criptografia-simetrica/cifras-de-bloco/boas-praticas.md)
  
  - [Conclusão](criptografia-simetrica/conclusao.md)

- [Algorítmos de Criptografia](conceitos/crypto/crypto-algorithms.md)
    - [DES - Data Encryption Standard](criptografia/crypto/des.md)
    - [3DES - Triple DES](criptografia/crypto/3des.md)
    - [AES - Advanced Encryption Standard](criptografia/crypto/aes.md)
    - [RSA - Rivest, Shamir e Adleman](criptografia/crypto/rsa.md)
    - [ECDSA - Elliptic Curve Digital Signature Algorithm](criptografia/crypto/ecdsa.md)
    - [HMAC - Assinaturas Seguras](conceitos/crypto/hmac.md)
    - [Criptografia Simétrica vs Assimétrica](conceitos/crypto/criptografia.md)
- [Miscelânea](conceitos/misc/misc.md)
    - [Confusão e Difusão](conceitos/misc/confusao-difusao.md)
    - [S-Box](conceitos/misc/s-box.md)
    - [Permutações e Trocas](conceitos/misc/permutacoes.md)
    - [Geradores de Números Pseudoaleatórios (PRNG)](conceitos/misc/prng.md)
    - [LFSR - Linear Feedback Shift Register](conceitos/misc/lfsr.md)
    - [LCG - Linear Congruential Generator](conceitos/misc/lcg.md)
    - [S-DES - Simplified DES](conceitos/misc/s-des.md)
    - [S-AES - Simplified AES](conceitos/misc/s-aes.md)
    - [MiniDES - Mini Data Encryption Standard](conceitos/misc/mini-des.md)
    - [MiniAES - Mini Advanced Encryption Standard](conceitos/misc/mini-aes.md)
    - [MiniRSA - Mini Rivest, Shamir e Adleman](conceitos/misc/mini-rsa.md)
    - [MiniECDSA - Mini Elliptic Curve Digital Signature Algorithm](conceitos/misc/mini-ecdsa.md)
    - [MiniHMAC - Mini Assinaturas Seguras](conceitos/misc/mini-hmac.md)

<!-- # Aplicações Práticas em Segurança

- [TLS e Certificados SSL](seguranca/tls.md)
    - [Como Funciona o Handshake TLS](seguranca/tls/handshake.md)
    - [Certificados e Autoridades Certificadoras](seguranca/tls/certificados.md)
    - [Cipher Suites e Forward Secrecy](seguranca/tls/ciphers.md)
    - [Exemplo Prático com Go](seguranca/tls/exemplo.md)
- [Segurança em APIs REST](seguranca/apis.md)
    - [HTTPS Obrigatório](seguranca/apis/https.md)
    - [Protegendo Endpoints com JWT](seguranca/apis/jwt-seguranca.md)
    - [Rate Limiting e Proteção contra DoS](seguranca/apis/rate-limit.md)
    - [Assinaturas HMAC para Webhooks](seguranca/apis/hmac-webhooks.md)

# Autenticação e Assinaturas Digitais

- [JSON Web Token (JWT)](autenticacao/jwt.md)
    - [Estrutura de um JWT](autenticacao/jwt/estrutura.md)
    - [Assinatura e Verificação](autenticacao/jwt/assinatura.md)
    - [JWT Expiração e Refresh Token](autenticacao/jwt/expiracao-refresh.md)
    - [Exemplo Prático com Go](autenticacao/jwt/exemplo.md)
- [OAuth 2.0 e OpenID Connect](autenticacao/oauth2.md)
    - [Fluxo Authorization Code](autenticacao/oauth2/authorization-code.md)
    - [Client Credentials e Machine-to-Machine](autenticacao/oauth2/client-credentials.md)
    - [OAuth Scopes e Permissões](autenticacao/oauth2/scopes.md)
    - [Exemplo Prático com Go](autenticacao/oauth2/exemplo.md)

# Autenticação e Assinaturas Digitais

- [JSON Web Token (JWT)](autenticacao/jwt.md)
    - [Estrutura de um JWT](autenticacao/jwt/estrutura.md)
    - [Assinatura e Verificação](autenticacao/jwt/assinatura.md)
    - [JWT Expiração e Refresh Token](autenticacao/jwt/expiracao-refresh.md)
    - [Exemplo Prático com Go](autenticacao/jwt/exemplo.md)
- [OAuth 2.0 e OpenID Connect](autenticacao/oauth2.md)
    - [Fluxo Authorization Code](autenticacao/oauth2/authorization-code.md)
    - [Client Credentials e Machine-to-Machine](autenticacao/oauth2/client-credentials.md)
    - [OAuth Scopes e Permissões](autenticacao/oauth2/scopes.md)
    - [Exemplo Prático com Go](autenticacao/oauth2/exemplo.md)

# Ataques e Vulnerabilidades

- [Colisão de Hash e Birthday Attack](ataques/hash-colisao.md)
- [Ataques a TLS - BEAST, POODLE, MITM](ataques/tls.md)
- [Quebra de Senhas - Ataques de Força Bruta e Rainbow Tables](ataques/senhas.md)

# Práticas Recomendadas

- [Armazenamento Seguro de Senhas](praticas/senhas.md)
    - [Bcrypt, Argon2 e PBKDF2](praticas/senhas/bcrypt-argon2.md)
    - [Erros Comuns no Hashing de Senhas](praticas/senhas/erros-comuns.md)
- [Gerenciamento de Tokens Seguros](praticas/tokens.md)
    - [Armazenando Tokens com Segurança](praticas/tokens/armazenamento.md)
    - [Expiração e Rotação de Tokens](praticas/tokens/expiracao-rotacao.md)
- [Checklist de Segurança para APIs](praticas/checklist-apis.md)

# Implementação Prática

- [Criando um Servidor OAuth2 com Go](implementacao/oauth2-server.md)
- [Protegendo um Backend com JWT e TLS](implementacao/backend-seguro.md)
- [Implementando um Webhook Seguro com HMAC](implementacao/hmac-webhook.md) -->

# Referências

- [Bibliotecas Recomendadas](referencias/bibliotecas.md)
- [Links e Leituras Adicionais](referencias/leituras.md)
