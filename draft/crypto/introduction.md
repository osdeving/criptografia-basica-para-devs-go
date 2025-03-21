       criptologia
criptografia - criptoanálise


criptografia: ciência de proteger a comunicação de pessoas mal intencionadas, historicamente o foco estava em esconder o conteúdo da mensagem, hoje o foco é também em autenticar a origem da mensagem e garantir a integridade da mensagem.

criptoanálise: ciência de quebrar a criptografia

a criptografia pode ser feita usando algoritmos simétricos e assimétricos e usa protocolos para garantir a segurança da comunicação.

cifra de substituição é protegida contra brute force attack, mas não contra análise de frequência. 26! aproximadamente 2^88 seria o key space de uma cifra de substituição com 88 caracteres inviável de ser quebrada por brute force. Mas a cifra de substituição é vulnerável a análise de frequência porque a as propriedades estatísticas da língua portuguesa são conhecidas e elas são mantidas na cifra de substituição.


isso significa que o keyspace de uma cifra não garante a segurança da cifra, a segurança da cifra depende de uma cifra ser segura contra análise de frequência o que implica em garantir que as propriedades estatísticas do texto plano não sejam mantidas no texto cifrado.

quase todos os algoritmos de criptografia trabalham com um conjunto finito. Os conjuntos que estamos acostumados, como os números reais e os números inteiros, são infinitos.
Pense que vc continua adicionando 1h para um relógio de 24 horas, vc vai chegar ao mesmo ponto de onde vc começou, e repetirá esse ciclo infinitamente. Logo, apesar do ser finito (24h) podemos adicionar 1h infinitamente.


