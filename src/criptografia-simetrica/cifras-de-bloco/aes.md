# Advanced Encryption Standard (AES)

O **Advanced Encryption Standard (AES)** é o algoritmo simétrico de cifra de bloco mais utilizado atualmente, adotado mundialmente como padrão de segurança para dados em repouso e em trânsito. Substituto do DES, o AES foi escolhido em um concurso público conduzido pelo NIST no final dos anos 1990, com critérios transparentes que marcaram uma nova era no desenvolvimento de algoritmos criptográficos.

Começaremos com o contexto do processo seletivo em [AES: concurso e escolha](aes-concurso.md), onde será discutida a motivação para a substituição do DES, os critérios definidos pelo NIST, os candidatos finalistas e os motivos técnicos que levaram à vitória do algoritmo Rijndael, proposto por Vincent Rijmen e Joan Daemen.

Em [AES: funcionamento interno](aes-interno.md), analisaremos a estrutura matemática da cifra — baseada em uma rede de substituição-permutação (SPN), em contraste com o modelo de Feistel —, detalhando os passos de SubBytes, ShiftRows, MixColumns e AddRoundKey. A seção também abordará a geração de subchaves por meio do processo de key schedule.

Na sequência, em [Modos de operação com AES](aes-modos.md), exploraremos como o AES é aplicado na prática a blocos múltiplos de dados, incluindo modos como CBC, CTR e GCM, com ênfase nos impactos de cada modo sobre segurança e desempenho.

Em [Armadilhas e boas práticas](boas-praticas.md), apresentaremos erros comuns de implementação e uso, como reutilização de IVs e escolha inadequada de modos, além de diretrizes para o uso seguro do algoritmo em sistemas reais.

Por fim, em [Implementação de AES em Go](aes-go.md), demonstramos uma aplicação prática da cifra em código, consolidando os conceitos aprendidos.

> O AES é a base da criptografia simétrica moderna: sua compreensão é essencial para qualquer profissional que lide com segurança da informação, engenharia de software ou sistemas criptográficos.
