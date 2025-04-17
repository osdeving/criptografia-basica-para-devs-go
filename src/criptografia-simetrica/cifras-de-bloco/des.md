# DES: arquitetura, S-boxes e permutações

Esta seção apresenta em profundidade o funcionamento interno do **Data Encryption Standard (DES)**, uma das cifras simétricas mais influentes da criptografia moderna. Embora atualmente considerada obsoleta para aplicações seguras, sua estrutura serviu de base para o entendimento e desenvolvimento de diversos algoritmos posteriores — e continua sendo amplamente utilizada como ferramenta didática.

Começaremos com uma visão geral em [O que é o DES?](o-que-e-des.md), contextualizando sua origem como padrão federal nos Estados Unidos e suas principais características técnicas. Em seguida, em [História, Personagens e Controvérsias do DES](historia-contexto.md), exploraremos o processo de padronização conduzido pelo NBS, o envolvimento da IBM e da NSA, e as polêmicas que surgiram em torno da transparência e segurança do algoritmo.

Na parte técnica, analisaremos a estrutura da cifra em [Arquitetura Geral](arquitetura-geral.md), com foco nas 16 rodadas de transformação baseadas na **rede de Feistel**. Em [A Função f](funcao-f.md), estudaremos em detalhe a lógica interna da principal operação do DES, incluindo a expansão, substituição e permutação. Em seguida, nas seções [S-boxes](s-boxes.md) e [Key Schedule](key-schedule.md), veremos como são construídas as tabelas de substituição e como a chave principal de 64 bits dá origem às 16 subchaves utilizadas ao longo das rodadas.

Apresentaremos também uma [Implementação em Go](des-go.md), demonstrando como os conceitos são aplicados na prática, e encerraremos com uma análise sobre [Segurança e Ataques](seguranca.md), além de [Extensões e Substitutos](extensoes-substitutos.md) como o 3DES, Twofish e AES.

> A arquitetura do DES continua sendo uma das mais estudadas da história da criptografia. Entendê-la em profundidade é compreender o ponto de transição entre a criptografia clássica e a criptografia computacional moderna. >
