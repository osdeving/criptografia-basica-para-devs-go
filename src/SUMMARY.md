# Sumário

[Introdução](README.md)

# Conceitos Fundamentais

- [Base64 e Codificação](conceitos/base64.md)
- [Checksum e CRC - Verificação de Integridade](conceitos/checksum-crc.md)
- [Funções Hash Criptográficas](conceitos/sha256.md)
    - [Funções Hash](conceitos/sha256.md)
    - [Família de Algoritmos MD - Message Digest](conceitos/message-digest.md)
    - [Família de Algoritmos DES - Data Encryption Standard](conceitos/des.md)
    - [Família de Algoritmos SHA - Secure Hash Algorithms](conceitos/sha.md)
    - [PBKDF2, Argon2 e Scrypt - Hash para Senhas](conceitos/pbkdf2-argon2.md)
- [Algorítmos de Criptografia](conceitos/crypto-algorithms.md)
    - [DES - Data Encryption Standard](criptografia/des.md)
    - [3DES - Triple DES](criptografia/3des.md)
    - [AES - Advanced Encryption Standard](criptografia/aes.md)
    - [RSA - Rivest, Shamir e Adleman](criptografia/rsa.md)
    - [ECDSA - Elliptic Curve Digital Signature Algorithm](criptografia/ecdsa.md)
    - [HMAC - Assinaturas Seguras](conceitos/hmac.md)
    - [Criptografia Simétrica vs Assimétrica](conceitos/criptografia.md)
# Aplicações Práticas em Segurança

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
- [Implementando um Webhook Seguro com HMAC](implementacao/hmac-webhook.md)

# Referências

- [Bibliotecas Recomendadas](referencias/bibliotecas.md)
- [Links e Leituras Adicionais](referencias/leituras.md)

[Contribuidores](misc/contributors.md)