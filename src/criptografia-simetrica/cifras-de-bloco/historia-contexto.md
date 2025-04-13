# História, Personagens e Controvérsias do DES

Esta seção apresenta a trajetória do Data Encryption Standard (DES), desde suas origens nos laboratórios da IBM até as polêmicas envolvendo sua padronização e posterior quebra. O objetivo é fornecer um panorama histórico que permita entender por que o DES desempenhou papel central no desenvolvimento da criptografia moderna.

## Origens

O DES tem suas raízes no projeto **Lucifer**, desenvolvido pela IBM no início da década de 1970. Este algoritmo surgiu de uma série de pesquisas conduzidas por **Horst Feistel**, um cientista de origem alemã vinculado aos laboratórios Watson da IBM. Feistel propôs um modelo de cifra baseado em operações repetidas sobre metades de blocos de dados — estrutura hoje conhecida como rede de Feistel. Essa abordagem, já discutida em [A estrutura de Feistel](criptografia-simetrica/cifras-de-bloco/feistel.md), viria a influenciar diretamente o projeto do DES e de diversas cifras simétricas nas décadas seguintes [1].

A equipe de Feistel incluía nomes como **Don Coppersmith**, **Walter Tuchman** e **Alan Konheim**, que trabalharam na elaboração de variantes do Lucifer para uso prático. A versão original do algoritmo operava sobre blocos de 128 bits com chaves de igual comprimento. No entanto, considerações de desempenho e implementação motivaram a IBM a simplificar o esquema, reduzindo o tamanho dos blocos e da chave, o que resultaria em um protótipo mais adequado para padronização nacional.

## Adoção pelo NBS

Em 1973, o **National Bureau of Standards (NBS)**, agência federal dos Estados Unidos responsável por padrões técnicos, iniciou um processo para selecionar um algoritmo de cifra de bloco que pudesse ser adotado como padrão para comunicação segura no setor civil. Um dos requisitos centrais era que o algoritmo fosse **de conhecimento público**, permitindo auditoria independente por especialistas e livre implementação pela indústria.

A IBM apresentou uma versão adaptada do Lucifer, que rapidamente se destacou entre as propostas recebidas. O NBS, sem expertise criptográfica interna suficiente, recorreu à **National Security Agency (NSA)** para revisão técnica da proposta. A versão final, com alterações feitas sob orientação da NSA, foi aprovada em 1976 e publicada oficialmente em 1977 como o **Data Encryption Standard**. Foi publicado pelo **Federal Information Processing Standards** (FIPS), órgão responsável por definir padrões federais de processamento de informações nos Estados Unidos, sob a especificação FIPS PUB 46 [2].

## O papel da NSA

A participação da NSA no processo de revisão gerou incertezas e especulações. A agência recomendou modificações em dois aspectos cruciais do algoritmo. O primeiro foi a **redução do tamanho da chave** de 128 para **56 bits**. A justificativa oficial envolvia considerações de implementação, mas não foram apresentados argumentos técnicos transparentes. Os oito bits restantes, completando os 64 da chave, foram reservados para verificação de paridade — um mecanismo de controle de integridade simples, abordado em [Bit de Paridade](conceitos/checksum-crc/parity-bit.md) — e não participam do processo de cifragem propriamente dito [2].

A segunda alteração foi a substituição das **S-boxes** (cf. [S-Box](conceitos/manipulacao-de-dados/s-box.md)) originais por uma nova versão desenvolvida em colaboração com a NSA. Por décadas, não se soube se essas modificações comprometiam a segurança do DES. A ausência de transparência levantou suspeitas sobre a possibilidade de um **backdoor criptográfico** embutido.

Somente nos anos 1990 veio à tona que as S-boxes haviam sido ajustadas para resistir a uma técnica chamada **criptoanálise diferencial**, descoberta posteriormente por **Eli Biham** e **Adi Shamir**. Em declaração publicada em artigo da IBM, **Coppersmith** revelou que a equipe da IBM já conhecia essa vulnerabilidade, mas havia sido instruída a não revelar detalhes públicos [3].

## Polêmicas e críticas

O lançamento do DES causou reações imediatas por parte da comunidade acadêmica. **Whitfield Diffie** e **Martin Hellman** publicaram uma série de artigos apontando que uma chave de 56 bits não oferecia segurança adequada contra ataques de força bruta [4]. Outros pesquisadores, como **Ron Rivest**, ecoaram as críticas e solicitaram que os detalhes das alterações feitas pela NSA fossem divulgados.

As suspeitas foram alimentadas pela dificuldade de realizar uma análise independente das S-boxes. Sem saber sua motivação de projeto, não era possível avaliar sua robustez. A posterior descoberta da resistência à criptoanálise diferencial mudou essa perspectiva, mas não eliminou o desconforto quanto à atuação da NSA.

A crítica final ao DES veio em 1998, quando a **Electronic Frontier Foundation (EFF)** construiu o **Deep Crack**, um dispositivo dedicado capaz de quebrar o DES por força bruta em menos de 24 horas. O custo estimado do equipamento era inferior a US$ 250 mil, o que demonstrava que a cifra já não era adequada para proteger informações sensíveis [5].

## Impacto histórico

Apesar das limitações identificadas, o DES teve um papel central no amadurecimento da criptografia aplicada. Foi o primeiro algoritmo padronizado para uso civil, incentivou o surgimento de literatura técnica acessível e fomentou o estudo de criptoanálise como disciplina prática. Serviu também como base para o desenvolvimento do **Triple DES (3DES)** e, posteriormente, para a seleção do **Advanced Encryption Standard (AES)**, iniciado em 1997.

O DES permanece como um caso clássico de interação entre avanço técnico, regulação estatal e debate público sobre segurança da informação.

