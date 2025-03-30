# Codificação de Dados: Representações Reversíveis para Armazenamento e Transmissão

Esta seção reúne os principais esquemas de codificação utilizados para converter dados binários em formatos textuais seguros e reversíveis. Embora essas codificações não tenham relação direta com criptografia, elas desempenham papel essencial em aplicações que exigem compatibilidade com sistemas baseados em texto, como e-mails (MIME), URLs, APIs web, arquivos JSON e representações de chaves ou hashes.

Começamos com a codificação Base64, amplamente utilizada para transmitir dados binários por meios que aceitam apenas texto. Em seguida, analisamos variações como Base64URL (usada em JWT), Base32 (utilizada em TOTP e sistemas de backup de chaves) e Base16 (hexadecimal), muito comum na representação de hashes e valores binários compactos.

Completamos a análise com codificações mais especializadas como Base85 (PostScript, Git), Base58 (endereços Bitcoin) e Percent-Encoding, usada para garantir segurança sintática em URLs.

> Todas as codificações aqui tratadas são **reversíveis** e **determinísticas**, ou seja, não envolvem chave, sigilo ou aleatoriedade. São ferramentas auxiliares para transmissão e representação segura de dados, e servirão de base para os capítulos futuros onde aparecerão como parte de algoritmos criptográficos reais.