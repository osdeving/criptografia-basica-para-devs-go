# O que é o DES?

O **Data Encryption Standard (DES)** é uma cifra de bloco simétrica padronizada em 1977 pelo governo dos Estados Unidos, por meio do então **National Bureau of Standards (NBS)**, hoje **NIST**. Ele opera sobre blocos de 64 bits de dados, utilizando chaves de 56 bits para cifragem e decifragem de informações.

O DES é baseado em uma **estrutura de Feistel** (cf. [A estrutura de Feistel](criptografia-simetrica/cifras-de-bloco/feistel.md)), que divide cada bloco de entrada em duas metades e aplica uma função iterativa sobre elas ao longo de 16 rodadas. Em cada rodada, uma subchave derivada da chave principal é usada em uma operação de mistura que envolve expansão, substituição e permutação de bits (cf. [S-Box](conceitos/manipulacao-de-dados/s-box.md)
    e [Permutações e Trocas](conceitos/manipulacao-de-dados/permutacoes.md)). Essa estrutura garante que, mesmo usando operações relativamente simples, o resultado final seja uma transformação complexa e difícil de inverter sem a chave correta.

O objetivo principal do DES era fornecer um padrão confiável de criptografia para uso não militar, especialmente no setor financeiro e em comunicações empresariais. Sua publicação representou uma mudança significativa, ao tornar disponível ao público um algoritmo antes restrito a aplicações governamentais ou militares. Isso impulsionou o desenvolvimento da criptografia como disciplina prática fora do âmbito estatal.

Apesar de sua influência e adoção generalizada, o DES foi gradualmente substituído a partir dos anos 2000 por algoritmos mais seguros, como o **Advanced Encryption Standard (AES)**. Isso se deve, principalmente, à redução da sua resistência a ataques de força bruta, uma vez que seu tamanho de chave (56 bits) passou a ser considerado pequeno frente à capacidade computacional moderna.

Ao longo das próximas seções, exploraremos com mais detalhes a história do DES, sua arquitetura interna, sua função de mistura, o agendamento de chaves e os principais ataques que motivaram sua substituição como padrão federal de criptografia.

