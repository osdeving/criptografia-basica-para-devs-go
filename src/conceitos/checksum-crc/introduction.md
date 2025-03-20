# Introdução

Na seção anterior, discutimos a codificação e o transporte de dados. Agora, avançamos para um aspecto essencial da transmissão e armazenamento de informações: a verificação de integridade. Esse processo garante que os dados recebidos são os mesmos que foram enviados, sem terem sido alterados ou corrompidos durante a transmissão ou armazenamento.

A integridade dos dados pode ser comprometida por diversos fatores, como erros humanos, falhas sistêmicas (por exemplo, setores defeituosos em discos rígidos) ou perdas de pacotes em redes devido a interferências ou congestionamento.

Para detectar essas falhas, foram desenvolvidos mecanismos como os checksums e os CRC (Cyclic Redundancy Check), que permitem identificar se os dados foram alterados de forma acidental.

É importante diferenciar esses métodos das funções hash criptográficas, como MD5 e SHA-256, que abordaremos mais adiante neste livro. Enquanto checksums e CRCs são projetados para detectar erros acidentais, funções hash criptográficas garantem a integridade contra modificações intencionais, oferecendo resistência a ataques e manipulação maliciosa.

Nesta seção, exploraremos como checksums simples podem ser utilizados para verificar integridade e, em seguida, como os CRC oferecem um método mais robusto para detectar erros em sistemas computacionais.