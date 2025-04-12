# Introdução

Antes de explorarmos algoritmos criptográficos como funções hash, cifras de bloco ou mecanismos de assinatura, é importante compreender **como os dados são representados e manipulados em baixo nível**. A criptografia opera sobre sequências binárias, e muitas de suas construções dependem de operações como **máscaras de bits, deslocamentos, permutações, concatenação e particionamento de blocos**.

Esta seção reúne os fundamentos práticos que formam a base para a implementação e compreensão de algoritmos criptográficos. Embora não envolvam segurança por si só, esses conceitos estruturam a **transformação e o fluxo de dados**, e são usados repetidamente em construções como:

* Expansão e compressão de blocos (como em DES)

* Rotação de palavras (como em SHA e AES)

* Aplicação de máscaras e XORs (presentes em todas as cifras)

* Organização de blocos e endianness (em PBKDF2, HMAC, etc.)

Algoritmos clássicos, como MD2, baseiam parte de sua segurança e funcionamento em **operações simples sobre bytes e bits**. Em todos os algorítmos teremos de uma forma ou outra operações em nível de bits. Por isso, essa preparação técnica é indispensável para leitura crítica e compreensão da implementação de sistemas criptográficos.

Nos próximos tópicos, cobriremos:

* A ordem dos bytes na memória (**endianness**)

* Os operadores de bit a bit (**bitwise**)

* Como manipular dados em blocos, rotacionar e reorganizar estruturas

Esses conceitos serão reaproveitados ao longo de todo o conteúdo técnico do livro e facilitarão a compreensão tanto matemática quanto prática dos algoritmos, reduzindo o esforço cognitivo necessário para entender as diversas manipulações de dados.

