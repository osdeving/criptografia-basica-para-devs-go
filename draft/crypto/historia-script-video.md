[Bloco 1]

A criptografia existe desde que alguém quis esconder uma mensagem de outra pessoa.
Mas a criptografia moderna, do tipo que usamos hoje em sistemas, redes e aplicações, só surgiu no século vinte.
E a parte dela que vamos estudar agora — a criptografia simétrica — evoluiu principalmente por causa de três coisas: guerra, espionagem e computação.

[Bloco 2]

Antes dos computadores, criptografar era basicamente uma brincadeira de reorganizar símbolos.
O exemplo mais famoso é a cifra de César, que simplesmente desloca letras no alfabeto.
Também existiam técnicas de embaralhar letras, substituir símbolos, ou usar ferramentas físicas, como rodas codificadoras ou tiras enroladas em bastões.

Esses métodos eram simples, mas já traziam dois conceitos que ainda usamos hoje: confusão e difusão.
Eles não tinham teoria, mas plantaram a base do que viria depois.

[Bloco 3]

A Primeira e, principalmente, a Segunda Guerra Mundial mudaram tudo.
A criptografia passou a ser uma questão de sobrevivência.

A máquina mais conhecida dessa época foi a Enigma, usada pela Alemanha nazista.
Ela parecia segura, mas tinha falhas sérias — não no mecanismo em si, mas no jeito como era usada.

Foi quebrada por matemáticos poloneses e depois por britânicos, liderados por Alan Turing em Bletchley Park.
Isso ajudou diretamente a encurtar a guerra.

[Bloco 4]

Do outro lado, os Estados Unidos quebraram a cifra Purple, do Japão.
E o mais curioso: fizeram isso sem nunca ter colocado as mãos na máquina original.
Eles reconstruíram tudo por engenharia reversa e análise de padrões.

Esses casos mostraram que segurança real depende de prever o que o inimigo pode fazer — mesmo que ele saiba tudo sobre o sistema.

[Bloco 5]

Durante a Guerra Fria, surgiram máquinas ainda mais sofisticadas, como a SIGABA, nos EUA, e a FIALKA, na União Soviética.
Eram complexas, mas tinham um limite claro: eram mecânicas.
E o mundo estava se tornando digital.

A necessidade de algo mais eficiente levou ao desenvolvimento de algoritmos eletrônicos.
Foi aí que surgiu o Lucifer, da IBM — o primeiro passo para o DES.

[Bloco 6]

O DES, ou Data Encryption Standard, foi adotado em 1977 como padrão oficial.
Era o primeiro algoritmo simétrico civil, aberto ao público.
E isso foi histórico — a criptografia saiu dos quartéis e foi parar em bancos, governos e empresas.

Mas o DES também gerou desconfiança.
A NSA participou da sua criação e reduziu o tamanho da chave para 56 bits.
Hoje sabemos que isso foi um erro — e o DES acabou sendo quebrado por força bruta.

[Bloco 7]

Pra contornar o problema, criaram o 3DES, que basicamente aplica o DES três vezes.
Funcionava, mas era lento.

Então, em 1997, o NIST abriu uma chamada pública para escolher um novo padrão.
Esse concurso foi um marco: transparente, com propostas do mundo todo e critérios bem definidos.

O vencedor foi o algoritmo Rijndael, criado por Joan Daemen e Vincent Rijmen.
Ele se tornou o AES, adotado oficialmente em 2001 — e continua até hoje como o principal padrão de criptografia simétrica.

[Bloco 8]

Hoje, algoritmos simétricos como o AES ou o ChaCha20 estão por toda parte.
São usados em conexões seguras na web, em VPNs, em sistemas de arquivos e até em apps de mensagens.

Mesmo quando usamos criptografia assimétrica — como em certificados e chaves públicas — quem protege os dados de verdade são os algoritmos simétricos.

[Bloco 9]

E tem uma coisa importante aqui:
Toda essa evolução aconteceu porque, ao longo do tempo, os códigos foram quebrados.

A Enigma caiu por falhas no uso.
A Purple caiu por padrões previsíveis.
Até máquinas supercomplexas foram vencidas por inteligência e insistência.

E é justamente por causa dessas falhas que hoje seguimos dois princípios fundamentais:

Primeiro: a segurança deve estar no algoritmo, não no segredo do sistema.
Segundo: um algoritmo precisa resistir mesmo se o adversário souber tudo sobre ele.

[Bloco 10]

Ou seja, a criptografia moderna não é feita pra esconder como funciona.
Pelo contrário: ela é feita pra resistir mesmo sendo pública.

Então, quando você começa a estudar um algoritmo como o AES — ou implementa uma cifra como o RC4 ou o 3DES — você não está só programando.

Você está aplicando décadas de erros, acertos e aprendizados.
E ajudando a não repetir os mesmos erros que, no passado, custaram muito mais do que dados.