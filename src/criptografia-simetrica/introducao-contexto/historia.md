## O que é criptografia simétrica

A criptografia simétrica é um tipo de criptografia em que a mesma chave é usada tanto para cifrar quanto para decifrar dados. Isso significa que quem envia e quem recebe precisam ter acesso à mesma chave secreta — o que implica em um desafio logístico importante: a troca segura dessa chave.

Esse modelo contrasta com a criptografia assimétrica, onde as chaves de cifragem e decifragem são diferentes. A criptografia simétrica é, em geral, mais rápida e simples de implementar. É usada em grande escala em sistemas que exigem desempenho, como conexões seguras, discos criptografados, redes privadas virtuais (VPNs) e proteção de arquivos.

A segurança dos algoritmos simétricos depende totalmente da chave: se ela for descoberta ou mal protegida, todo o sistema fica vulnerável. Isso torna a gestão de chaves um ponto crítico.

Existem dois grandes grupos de cifras simétricas: **cifras de fluxo**, que processam dados em sequência contínua, e **cifras de bloco**, que operam sobre blocos fixos de bits. Essa distinção será explorada na próxima seção.

---

## Breve história da criptografia moderna

A criptografia existe desde a antiguidade, mas a criptografia moderna começa a tomar forma no século XX. Nesse período, a guerra, a espionagem e o surgimento da computação moldaram os fundamentos do campo. Foi um momento em que o sigilo deixou de ser apenas uma ferramenta diplomática ou militar isolada, e passou a influenciar diretamente estratégias geopolíticas e o curso da história.

Antes dos computadores, cifras como a **de César**, as transposições e artefatos físicos como as escítalas espartanas eram comuns. Elas já aplicavam ideias de confusão e difusão, mesmo sem formalização matemática.

Durante a Segunda Guerra Mundial, a criptografia ganhou escala e importância estratégica. Máquinas como a **Enigma**, da Alemanha, e a **Purple**, do Japão, exigiram o uso de técnicas de criptoanálise sistemática — o que levou à criação dos primeiros computadores.

A quebra da Enigma pelos aliados, liderada pelos trabalhos em **Bletchley Park**, não foi apenas um feito técnico: estima-se que tenha encurtado a guerra em pelo menos dois anos e salvado milhões de vidas. O esforço envolveu matemáticos como **Alan Turing** e **Gordon Welchman**, mas também operadores, linguistas e especialistas em interceptação de sinais. O impacto foi tão profundo que grande parte do projeto permaneceu classificado por décadas.

No caso da Purple, os norte-americanos, liderados por William Friedman e sua equipe, conseguiram reconstruir completamente a lógica da máquina japonesa sem nunca terem visto um exemplar. A partir da análise de padrões e repetições em mensagens interceptadas, conseguiram acessar comunicações diplomáticas e militares de alto nível.

No pós-guerra, os EUA padronizaram internamente o uso de sistemas como a **SIGABA**, que incorporava mecanismos de troca aleatória de rotores para evitar padrões repetitivos, dificultando a criptoanálise. Já a União Soviética operava com a **FIALKA**, uma evolução da Enigma com recursos adicionais como verificação redundante e alfabetos em cirílico.

Essas máquinas mostraram limites fundamentais da abordagem mecânica. Elas eram grandes, caras, lentas, e exigiam sincronização precisa entre remetente e receptor. Além disso, vulnerabilidades práticas — como reutilização de configurações ou erros de operador — tornavam o sistema quebrável mesmo sem falhas matemáticas.

Com o avanço da eletrônica e da computação, tornou-se inevitável o desenvolvimento de algoritmos puramente digitais. Na década de 1970, a IBM desenvolveu o **Lucifer**, que, com a colaboração da NSA, se transformaria no **DES** — o primeiro algoritmo de cifra simétrica padronizado para uso civil. Isso marcou uma mudança de paradigma: a criptografia saiu do domínio exclusivamente governamental e passou a fazer parte da infraestrutura de empresas, bancos e, eventualmente, da internet.

Concursos como o do **AES**, e mais tarde o **eSTREAM**, marcaram a transição para uma era de padronizações abertas, revisão por pares e competição internacional entre algoritmos candidatos. Esses processos mostraram que segurança não é apenas uma questão técnica, mas também política, institucional e social.

A criptografia moderna nasceu da necessidade urgente de vencer guerras — mas sobreviveu, amadureceu e se democratizou como uma das bases fundamentais da era digital.

A criptografia existe desde a antiguidade, mas a criptografia moderna começa a tomar forma no século XX. Nesse período, a guerra, a espionagem e o surgimento da computação moldaram os fundamentos do campo.

Antes dos computadores, cifras como a **de César**, as transposições e artefatos físicos como as escítalas espartanas eram comuns. Elas já aplicavam ideias de confusão e difusão, mesmo sem formalização matemática.

Durante a Segunda Guerra Mundial, a criptografia ganhou escala e importância estratégica. Máquinas como a **Enigma**, da Alemanha, e a **Purple**, do Japão, exigiram o uso de técnicas de criptoanálise sistemática — o que levou à criação dos primeiros computadores.

No pós-guerra, os EUA padronizaram internamente o uso de sistemas como a **SIGABA**, enquanto a União Soviética adotava sistemas como a **FIALKA**. A complexidade dessas máquinas mecânicas logo se mostrou insuficiente para o novo ritmo das comunicações digitais, e foi necessário o desenvolvimento de algoritmos eletrônicos.

Na década de 1970, a IBM desenvolveu o **Lucifer**, que, com a colaboração da NSA, se transformaria no **DES** — o primeiro algoritmo de cifra simétrica padronizado para uso civil.

Concursos como o do **AES**, e mais tarde o **eSTREAM**, marcaram a transição para uma era de padronizações abertas, revisão por pares e competição internacional entre algoritmos candidatos. Esses processos mostraram que segurança não é apenas uma questão técnica, mas também política e institucional.

---

## Códigos militares, espionagem e cifras quebradas

Boa parte da evolução da criptografia vem das quebras — das falhas que revelaram as limitações de métodos anteriores.

A **máquina Enigma**, por exemplo, foi considerada segura por oficiais alemães, mas foi derrotada por falhas operacionais, reutilização de padrões e interceptação massiva. O trabalho de Alan Turing, Gordon Welchman e equipes em **Bletchley Park** foi crucial para quebrá-la. A cifra **Purple**, do Japão, foi reconstruída por criptoanalistas americanos com base em análise estatística e engenharia reversa.

Décadas depois, já no contexto digital, algoritmos amplamente utilizados também foram alvo de ataques eficazes. Um dos exemplos mais importantes é o **RC4**, uma cifra de fluxo criada por Ron Rivest em 1987. Apesar de sua simplicidade e velocidade, ela apresentava uma fraqueza crítica: os primeiros bytes gerados pelo seu PRGA (Pseudo-Random Generation Algorithm) tinham viés estatístico. Esse viés pôde ser explorado em ataques práticos, especialmente em protocolos como WEP e TLS, que usavam RC4 sem descartar os bytes iniciais. Em 2013, ataques de recuperação de texto claro em tempo real contra RC4 no TLS levaram à sua obsolescência definitiva.

No caso do **MD5**, uma função de hash originalmente considerada resistente a colisões, os primeiros ataques teóricos datam do final da década de 1990. Em 2004, Xiaoyun Wang e sua equipe demonstraram um ataque prático de colisão, com custo computacional muito abaixo do esperado por força bruta. Com o tempo, ataques mais sofisticados, como a criação de certificados SSL fraudulentos com colisões prefixadas, mostraram que o MD5 não podia mais ser usado em contextos de segurança. Hoje, ele é considerado inseguro para qualquer aplicação que dependa de unicidade ou integridade.

O avanço dos ataques levou ao desenvolvimento de novas abordagens para cifras de fluxo e PRNGs (geradores pseudoaleatórios criptograficamente seguros). O **projeto eSTREAM**, parte do consórcio europeu ECRYPT, buscou identificar cifras de fluxo seguras para aplicações leves e de alto desempenho. Dali surgiram algoritmos como **Salsa20**, **Grain**, **HC-128** e **Rabbit**, que oferecem segurança com desempenho competitivo, mesmo em dispositivos embarcados.

Esses episódios mostram que a criptografia é um campo em constante movimento: algoritmos são propostos, testados, criticados, melhorados — ou aposentados. E cada quebra bem-sucedida representa, ao mesmo tempo, uma vulnerabilidade corrigida e um avanço para o campo.

Boa parte da evolução da criptografia vem das quebras — das falhas que revelaram as limitações de métodos anteriores.

---

## O nascimento do DES e a influência da IBM

O algoritmo **DES (Data Encryption Standard)** nasceu a partir do projeto **Lucifer**, desenvolvido por **Horst Feistel** na IBM. O Lucifer introduziu a **rede de Feistel**, estrutura que permite que o mesmo algoritmo seja usado para cifrar e decifrar apenas invertendo a ordem das rodadas.

A proposta foi adaptada com a colaboração da **NSA**, que sugeriu alterações nas S-boxes e uma redução no tamanho da chave. Em 1977, o **NIST** oficializou o DES como padrão federal dos EUA. Foi a primeira cifra amplamente usada fora de ambientes militares.

Apesar de sua importância histórica, o DES logo mostrou fragilidades: o tamanho da chave (56 bits) tornou-se pequeno diante da capacidade computacional crescente. Em 1998, a **EFF** demonstrou um ataque de força bruta que quebrou o DES em poucos dias usando hardware dedicado.

A resposta temporária foi o **3DES**, que aplicava o DES três vezes com chaves diferentes, aumentando a segurança mas também o custo computacional.

---

## NIST, NSA e as controvérsias na padronização

A padronização do DES expôs tensões entre agências governamentais, instituições civis e a comunidade científica.

A **NSA** teve participação ativa no processo, influenciando detalhes técnicos do algoritmo — o que gerou suspeitas. O design das **S-boxes**, inicialmente mantido em sigilo, foi posteriormente revelado como otimizado contra criptoanálise diferencial — uma técnica ainda não publicada à época, mas conhecida pela NSA.

Essa falta de transparência gerou críticas, especialmente entre acadêmicos e defensores da criptografia civil. Como resposta, o NIST passou a adotar processos mais abertos em concursos posteriores, como o do **AES** e do **SHA-3**, com submissões públicas e avaliações internacionais.

A relação entre agências como a NSA e padrões civis permanece delicada, e episódios como o vazamento do algoritmo de geração de números pseudoaleatórios **Dual\_EC\_DRBG**, suspeito de conter backdoors, reforçaram a necessidade de auditoria pública e desconfiança saudável.

---

## Personagens e contribuições

A criptografia moderna foi moldada por contribuições de matemáticos, engenheiros e cientistas da computação. A seguir, alguns nomes centrais:

> **Claude Shannon (1916–2001)**  
> Nasceu em Michigan, EUA. Doutorou-se no MIT, onde também trabalhou como pesquisador.  
> É considerado o "pai da teoria da informação". Publicou "Communication Theory of Secrecy Systems" [1], onde formalizou a segurança perfeita e os conceitos de confusão e difusão. Seu trabalho definiu a base teórica da criptografia moderna.

> **Horst Feistel (1915–1990)**  
> Alemão naturalizado americano, trabalhou na IBM. Desenvolveu o algoritmo Lucifer, baseado na estrutura de rede que leva seu nome: a **rede de Feistel**. Essa estrutura foi usada no DES e inspirou algoritmos como Blowfish, CAST e AES. 

> **Whitfield Diffie (n. 1944) & Martin Hellman (n. 1945)**  
> Norte-americanos, ambos com formação em engenharia elétrica e ciência da computação. Trabalharam juntos na Stanford University.  
> Em 1976, publicaram "New Directions in Cryptography" [2], introduzindo o conceito de criptografia assimétrica e troca de chaves pública, revolucionando o campo.

> **Ron Rivest (n. 1947)**  
> Professor do MIT. Coautor do algoritmo RSA, ao lado de Adleman e Shamir. Criador das cifras RC (Rivest Cipher), como RC4, RC5 e RC6. Contribuiu para o desenvolvimento de funções de hash e protocolos de autenticação. 

> **Joan Daemen (n. 1965) & Vincent Rijmen (n. 1970)**  
> Belgas, ambos formados na KU Leuven. Criaram o algoritmo Rijndael, vencedor do concurso AES. Seu design baseado em substituições e permutações é eficiente em hardware e software.  
> Publicaram diversos artigos sobre criptografia leve, estruturas SPN e segurança de blocos simétricos [3].

> **Eli Biham (n. 1960)**  
> Israelense, professor no Technion (Israel Institute of Technology). Coautor da criptoanálise diferencial com Adi Shamir — técnica que influenciou o design de S-boxes no DES e no AES.  
> Autor de livros e artigos sobre segurança simétrica, criptoanálise linear e criptografia leve.

> **Daniel J. Bernstein (n. 1971)**  
> Professor na University of Illinois at Chicago. Ativista da criptografia de código aberto. Criador do ChaCha20, Salsa20 e Poly1305. Também desenvolveu o conjunto de ferramentas NaCl (Networking and Cryptography library).  
> Publicou trabalhos influentes sobre segurança prática e criptografia pós-quântica [4].

> **Adversários históricos**  
> - **Marian Rejewski**: matemático polonês que, com Zygalski e Rozycki, quebrou a Enigma antes da Segunda Guerra.  
> - **Alan Turing**: matemático britânico, atuou em Bletchley Park e contribuiu para a quebra automatizada da Enigma.  
> - **Bletchley Park Team**: equipe interdisciplinar de criptoanalistas, engenheiros e operadores.

> - **NSA**: embora não creditada publicamente, contribuiu para a resistência do DES à criptoanálise diferencial, antecipando ataques publicados anos depois.

Esses nomes aparecem ligados a algoritmos, conceitos ou estruturas usados até hoje — e ajudam a contextualizar o motivo pelo qual os sistemas atuais têm a forma que têm.


Referências:  
[1] C. Shannon, "Communication Theory of Secrecy Systems," Bell System Technical Journal, vol. 28, pp. 656–715, 1949.  
[2] W. Diffie and M. E. Hellman, "New Directions in Cryptography," IEEE Transactions on Information Theory, vol. 22, no. 6, pp. 644–654, 1976.  
[3] J. Daemen and V. Rijmen, "The Design of Rijndael: AES – The Advanced Encryption Standard," Springer, 2002.  
[4] D. J. Bernstein, "The Salsa20 family of stream ciphers," in New Stream Cipher Designs, Springer, 2008, pp. 84–97.

---

## Linha do tempo da criptografia simétrica moderna

|  Ano | Evento |
|------|--------|
| ~1900 a.C.  | Uso de formas primitivas de substituição em mensagens militares egípcias. |
| ~58 a.C.  | Cifra de César é usada pelo Império Romano. |
| ~850 d.C.  | Al-Kindi descreve a análise de frequência como técnica de criptoanálise. |
| 1917  | A cifra ADFGVX é usada pela Alemanha na Primeira Guerra Mundial. |
| 1920–1945  | Enigma (Alemanha), Purple (Japão), SIGABA (EUA), FIALKA (URSS) são desenvolvidas. |
| 1943–1945  | Quebra da Enigma em Bletchley Park; criação dos primeiros computadores especializados. |
| 1949  | Claude Shannon publica "Communication Theory of Secrecy Systems". |
| 1970  | Lucifer é desenvolvido pela IBM (Horst Feistel). |
| 1977  | DES é padronizado pelo NIST com apoio da NSA. |
| 1987  | Ron Rivest propõe o algoritmo RC4. |
| 1990  | MD5 é publicado como função de hash. |
| 1997  | Início do concurso AES. |
| 2000  | Rijndael é selecionado como AES. |
| 2004  | Ataques práticos contra MD5 são divulgados. |
| 2008  | eSTREAM publica finalistas para cifras de fluxo. |
| 2013  | RC4 é considerado quebrado para uso em TLS. |
| 2015  | ChaCha20-Poly1305 é adotado no TLS como alternativa segura ao RC4. |

---

## Criptografia simétrica no mundo real hoje

Hoje, a criptografia simétrica está no coração de praticamente toda comunicação segura.

- **TLS/HTTPS**: protege a comunicação web. Após a troca de chaves, os dados trafegam com AES ou ChaCha20.
- **SSH**: conexão segura com servidores.
- **VPNs (IPsec, WireGuard)**: usam criptografia simétrica para encapsular e proteger pacotes.
- **Sistemas de arquivos (LUKS, BitLocker)**: usam AES para proteger dados em disco.
- **Aplicativos de mensagem (Signal, WhatsApp)**: usam criptografia ponta a ponta baseada em chaves simétricas derivadas.

Além disso, algoritmos simétricos estão presentes em sistemas embarcados, criptomoedas, roteadores, bancos e firmware. A combinação entre eficiência, simplicidade e maturidade torna a criptografia simétrica uma escolha sólida para o sigilo de dados — desde que bem implementada e atualizada conforme novas vulnerabilidades forem descobertas.

