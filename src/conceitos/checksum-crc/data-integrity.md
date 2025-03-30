# Verificação de Integridade de Dados: Checksum, Bit de Paridade e CRC

Nesta seção, estudaremos mecanismos clássicos de verificação de integridade aplicados em transmissão e armazenamento de dados. Esses mecanismos são baseados em operações determinísticas que permitem identificar se houve erro ou corrupção durante o transporte ou gravação de informação.

Começamos com o **Bit de Paridade**, um dos métodos mais simples de detecção de erro, amplamente usado em memória RAM, protocolos seriais e dispositivos de baixo custo.

Em seguida, analisamos o **Checksum**, que agrupa e soma blocos de dados para gerar um valor de referência simples. Apesar de fraco do ponto de vista criptográfico, é eficiente para detectar corrupção acidental.

Por fim, estudamos o **CRC (Cyclic Redundancy Check)**, um algoritmo baseado em operações binárias e teoria de polinômios sobre GF(2), amplamente usado em redes, discos e protocolos robustos de comunicação como Ethernet e USB.

> Ao contrário das funções hash criptográficas, os algoritmos desta seção não são projetados para resistir a ataques intencionais, mas sim para detectar falhas acidentais. Servem como base histórica e técnica para compreender os limites da verificação simples de integridade antes de abordarmos as construções criptográficas propriamente ditas.