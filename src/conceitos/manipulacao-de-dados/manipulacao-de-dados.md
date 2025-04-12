# Manipulação de Dados

Esta seção estabelece os fundamentos práticos para entender como os dados binários são tratados internamente — essencial para compreender os algoritmos de cifra, hashes e protocolos.

Começaremos com uma explicação rápida sobre endianness em **Endianness – Organização de Bytes na Memória**, que aborda como diferentes arquiteturas armazenam inteiros e como isso afeta a interoperabilidade e o seu papel no
 com os variados algorítmos de criptografia.

Em seguida, em **Operadores Bitwise**, exploraremos os operadores bit a bit, que são fundamentais para a construção de cifras e hashes. Mostraremos como AND, OR, XOR e deslocamentos são usados na manipulação de bits, no controle de flags e na implementação de PRNGs e mecanismos criptográficos.

Veremos depois, em **Operações em Blocos**, os modos de operação (ECB, CBC, etc.) como técnicas de controle de fluxo entre blocos de dados — uma estrutura usada encadear operações de cifra. Aqui a conexão com criptografia é mais acentuada, mas o foco continua sendo como manipular blocos binários de forma sistemática.

Na sequência, estudaremos três conceitos fundamentais para o projeto de cifras modernas: **Confusão e Difusão**, **S-Box** e **Permutações e Trocas**. Embora estejam mais diretamente associados ao campo da criptografia, todos eles são, na essência, formas especializadas de manipulação de dados binários. Por esse motivo, optamos por incluí-los nesta seção.

> Esses conceitos serão reaproveitados ao longo do livro e servem de base para o entendimento de algoritmos criptográficos reais quando explorados de forma pŕatica estudando as suas implementações. O objetivo é diminuir o esforço cognitivo ao apresentar tópicos mais avançados, dando a chance do leitor se familiarizar com tópicos mais gerais envolvendo manipulação de dados.
