# Manipulação de Dados

Esta seção estabelece os fundamentos práticos para entender como os dados binários são tratados internamente — ponto essencial para compreender cifras, hashes e protocolos.

Começaremos com uma explicação rápida sobre endianness em **Endianness – Organização de Bytes na Memória**, que aborda como diferentes arquiteturas armazenam inteiros e como isso afeta a interoperabilidade e o uso seguro da criptografia.

Em seguida, em **Operadores Bitwise**, exploraremos os operadores bit a bit, que são fundamentais para a construção de cifras e hashes. Mostraremos como AND, OR, XOR e deslocamentos são usados na manipulação de bits, no controle de flags e na implementação de PRNGs e mecanismos criptográficos.

Veremos depois, em **Operações em Blocos**, os modos de operação (ECB, CBC, etc.) como técnicas de controle de fluxo entre blocos de dados — uma estrutura essencial para cifras reais. Aqui a conexão com criptografia é mais acentuada, mas o foco continua sendo como manipular blocos binários de forma sistemática.

> Esses conceitos serão reaproveitados ao longo do livro e servem de base para o entendimento de algoritmos criptográficos reais. O objetivo é diminuir o esforço cognitivo ao apresentar tópicos mais avançados, dando a chance do leitor se familiarizar com tópicos mais gerais envolvendo manipulação de dados. 
