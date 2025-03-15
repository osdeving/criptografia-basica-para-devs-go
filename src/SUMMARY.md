# Sumário

[Introdução](README.md)

# Conceitos Fundamentais

- [Base64 e Codificação](conceitos/base64.md)
- [Hashing e SHA-256](conceitos/sha256.md)
- [HMAC - Assinaturas Seguras](conceitos/hmac.md)
- [Criptografia Simétrica vs Assimétrica](conceitos/criptografia.md)

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

# Segurança no Transporte

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


