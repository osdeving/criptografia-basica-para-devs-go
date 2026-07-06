# Introdução

Uma função hash transforma uma entrada de tamanho variável em uma saída de tamanho fixo.

```text
entrada qualquer -> função hash -> saída de tamanho fixo
```

Exemplo:

```text
"abc" -> SHA-256 -> ba7816bf8f01cfea414140de5dae2223...
```

Essa saída também recebe nomes como:

```text
digest
message digest
hash value
fingerprint
```

Também conhecidas como funções *one-way*, as funções hash criptográficas não são cifras. Uma cifra precisa permitir a volta, isto é, precisa ser reversível e retornar à mensagem original.

```text
encrypt(plaintext, key) -> gera um ciphertext
decrypt(ciphertext, key) -> gera novamente a mensagem original plaintext
```

Com funções hash, não é possível voltar à mensagem original:

```text
hash(message) -> digest
digest -> não é possível obter a mensagem original a partir do digest
```

## Propriedades

Funções hash criptográficas devem ter as seguintes propriedades:

```text
determinística:
  mesma entrada, mesma saída

saída fixa:
  digest sempre tem o mesmo tamanho

resistente à pré-imagem:
  dado um digest, deve ser inviável encontrar uma mensagem que gere esse digest

resistente à segunda pré-imagem:
  dada uma mensagem M1, deve ser inviável encontrar M2 diferente com o mesmo digest

resistente à colisão:
  deve ser inviável encontrar qualquer par M1 != M2 com o mesmo digest

efeito avalanche:
  uma pequena mudança na entrada deve causar uma mudança ampla na saída
```

Essas propriedades existem para aumentar a confiança no algoritmo.

Quando dizemos que uma função hash é **determinística**, estamos dizendo que a mesma mensagem precisa gerar exatamente o mesmo digest hoje, amanhã, no seu notebook, no servidor, em Go, Rust, C ou qualquer outra implementação correta. Isso é o que permite comparar hashes. Se você baixa uma ISO do Linux e compara o SHA-256 com o valor publicado no site, está confiando nisso: os mesmos bytes precisam produzir o mesmo resultado.

Nesse uso, o digest funciona como uma **impressão digital** do arquivo, também chamada de **fingerprint**: um identificador compacto derivado do conteúdo. Se a função tivesse qualquer aleatoriedade interna, o digest deixaria de ser essa impressão digital estável e não poderia ser usado para verificar se dois arquivos são exatamente iguais.

É importante separar alguns termos relacionados. O **digest** é o valor produzido pela função hash. A **fingerprint** é o uso desse digest como uma impressão digital do conteúdo. Um **checksum** também serve para detectar alteração, mas normalmente é pensado para erros acidentais, não para resistência criptográfica. Já uma **assinatura digital** é outra coisa: ela normalmente assina um hash usando uma chave privada, permitindo verificar autoria e integridade com uma chave pública. Hash pode ser parte de uma assinatura, mas hash sozinho não prova quem produziu o arquivo.

A **saída fixa** permite transformar uma entrada potencialmente enorme em algo manuseável, como no exemplo da imagem ISO do Linux mencionado acima. Assim, um contrato de duas páginas, um vídeo de 4 GB ou uma string vazia podem virar, por exemplo, apenas 32 bytes no SHA-256.

Essa compactação não significa que a função hash *guardou* a mensagem inteira dentro dela; isso é matematicamente impossível. O digest funciona mais como uma impressão digital matemática: pequeno o bastante para comparar, transmitir e armazenar, mas sensível o bastante para denunciar mudanças na entrada.

A **resistência à pré-imagem** é a propriedade que sustenta a ideia de mão única. A palavra vem da linguagem matemática das funções: se $H(mensagem) = digest$, então o digest é a imagem daquela mensagem, e a mensagem é uma pré-imagem daquele digest.

Em uma função hash, o domínio é o conjunto de todas as entradas possíveis: arquivos, textos ou sequências de bytes. O contradomínio é o conjunto de todos os digests possíveis, por exemplo, todos os valores de 256 bits no caso do SHA-256.

Portanto, se eu te dou apenas um digest, você não deveria conseguir encontrar uma mensagem que produza exatamente aquele mesmo valor. Em uma hash criptográfica boa, o digest não deve entregar pistas úteis para reconstruir a entrada original.

Em notação matemática, o problema de pré-imagem pode ser descrito assim:

$$
\text{dado } y, \text{ encontrar } x \text{ tal que } H(x) = y
$$

O armazenamento de senhas ilustra bem o problema, embora não represente necessariamente uma quebra da propriedade de pré-imagem. Se o banco armazena $H(\text{"senha123"})$, isto é, a hash da string `"senha123"`, então `"senha123"` é uma pré-imagem daquele digest.

Um atacante que tenha acesso ao banco pode tentar adivinhar entradas comuns e calcular seus hashes até encontrar um digest igual. O uso de *rainbow tables* acelera esse processo usando hashes pré-computados de senhas prováveis. Isso não significa que a função hash foi invertida matematicamente; significa que o atacante explorou o fato de senhas humanas virem de um espaço pequeno e previsível.

Como ainda veremos neste livro, para senhas, não se usa SHA-256 puro. Usa-se *salt* e funções próprias para senha, como Argon2, bcrypt ou scrypt, justamente para tornar ataques de tentativa, dicionário e *rainbow table* muito mais caros.

A **resistência à segunda pré-imagem** é parecida com a resistência à pré-imagem, mas o ponto de partida muda. Na pré-imagem, o atacante recebe apenas um digest e tenta encontrar alguma mensagem que produza aquele valor. Na segunda pré-imagem, ele já conhece uma mensagem legítima $M_1$ e tenta fabricar outra mensagem $M_2$, diferente, tal que:

$$
H(M_1) = H(M_2)
$$

com:

$$
M_1 \neq M_2
$$

Ou seja, o atacante não quer apenas encontrar qualquer entrada para um digest desconhecido; ele quer substituir uma mensagem específica por outra que preserve exatamente o mesmo resumo.

Isso seria perigoso em cenários como documentos, contratos, backups, ISOs e binários de software. Imagine que alguém publica o hash de um binário legítimo. Se um atacante consegue criar um binário malicioso com o mesmo digest, o hash continuaria batendo, mas o conteúdo já não seria o mesmo.

Portanto, não basta ser difícil voltar do digest para alguma mensagem. Também precisa ser difícil, dado um conteúdo específico, fabricar um substituto com o mesmo resumo.

A **resistência à colisão** é a propriedade que torna computacionalmente inviável encontrar quaisquer duas mensagens diferentes, $M_1$ e $M_2$, que produzam o mesmo digest:

$$
M_1 \neq M_2
$$

$$
H(M_1) = H(M_2)
$$

Há muito mais mensagens possíveis do que digests possíveis. No caso do SHA-256, por exemplo, qualquer entrada é comprimida para um valor de 256 bits. Pelo princípio da casa dos pombos, mensagens diferentes podem cair no mesmo digest. A segurança não depende de colisões serem impossíveis; depende de elas serem computacionalmente inviáveis de encontrar.

Diferente da segunda pré-imagem, na colisão o requisito é mais geral: a função deve resistir à tentativa de encontrar qualquer par de mensagens distintas que colidam.

A resistência à colisão é uma propriedade central em contextos como assinaturas digitais, certificados, integridade de pacotes e documentos. Se colisões se tornam práticas, passa a ser possível construir dois conteúdos diferentes com o mesmo digest, enfraquecendo a confiança no digest como identificador criptográfico daquele conteúdo.

Foi por esse caminho que algoritmos como MD5 e SHA-1 perderam credibilidade em aplicações sensíveis. Eles continuaram determinísticos, continuaram produzindo saída fixa e continuaram parecendo funções hash no formato. O problema é que se tornou viável encontrar colisões, e isso quebra a confiança necessária para usos criptográficos.

O **efeito avalanche** é o comportamento visualmente mais interessante. Você muda uma letra, um bit, uma vírgula, e a saída deveria mudar de forma ampla, sem preservar uma relação óbvia com o digest anterior.

Uma boa hash não deve funcionar como uma proporção que muda um pouquinho quando a entrada muda um pouquinho. Ela deve se comportar como um sistema de mistura: uma alteração local entra nas rodadas internas e se espalha pelo estado interno do algoritmo. Esse efeito não é apenas estética matemática; ele dificulta inferir a estrutura da entrada olhando diferenças entre digests.

Essas propriedades trabalham juntas. Uma função hash pode ser determinística e ter saída fixa, como uma soma de bytes, e ainda assim ser inútil para criptografia. O que torna a função criptográfica é o conjunto: digest fixo, difícil de inverter, difícil de substituir, difícil de colidir e com boa difusão. Quando uma dessas peças falha, o algoritmo não serve como base confiável para segurança.

## Colisão é inevitável

Uma hash recebe entradas de tamanho variável e gera saídas de tamanho fixo. Isso significa que colisões existem inevitavelmente.

Como já vimos, o que a criptografia exige não é:

```text
colisão impossível
```

O requisito é:

```text
colisão deve ser inviável de encontrar na prática
```

Se uma hash tem saída de $n$ bits, o custo ideal para achar uma colisão por busca genérica é da ordem de:

$$
2^{n/2}
$$

Isso vem do paradoxo do aniversário.

Alguns exemplos:

$$
MD5 \approx 2^{64}
$$

$$
SHA\text{-}1 \approx 2^{80}
$$

$$
SHA\text{-}256 \approx 2^{128}
$$

Isso parece muito, mas ataques criptanalíticos podem reduzir muito esse custo. Foi isso que aconteceu com MD4, MD5 e SHA-1.

## Funções Hash Não Criptográficas

Nem toda função hash é uma função hash criptográfica. Muitas funções hash existem para outros objetivos: distribuir dados em tabelas hash, identificar registros, particionar mensagens em sistemas distribuídos, detectar erros simples ou produzir identificadores rápidos.

Uma função hash não criptográfica pode ser perfeitamente útil nesses contextos e, ao mesmo tempo, ser completamente inadequada para segurança.

Considere uma função didática simples:

```rust
fn simple_hash(input: &[u8]) -> u8 {
    input.iter().fold(0u8, |acc, byte| acc.wrapping_add(*byte))
}
```

Ela soma os bytes da entrada e devolve apenas um byte como resultado. Como a saída é `u8`, existem apenas 256 valores possíveis.

Essa função tem algumas características de uma hash:

```text
é determinística:
  a mesma entrada sempre gera a mesma saída

tem saída fixa:
  qualquer entrada gera um resultado de 8 bits

é rápida:
  basta percorrer os bytes e somar
```

Mas ela viola praticamente tudo que esperamos de uma hash criptográfica.

Por exemplo:

```text
"abc" e "acb" colidem
```

Isso acontece porque a função só soma os bytes. A ordem dos caracteres não importa. Portanto, qualquer permutação dos mesmos bytes gera o mesmo resultado.

Também é fácil fabricar colisões:

```text
qualquer conjunto de bytes com a mesma soma colide
```

Além disso, como existem apenas 256 saídas possíveis, colisões são não apenas inevitáveis, mas extremamente fáceis de encontrar.

Essa função também não tem efeito avalanche real. Se você muda um byte da entrada de forma pequena, a saída tende a mudar de forma pequena e previsível. A alteração não se espalha por um estado interno complexo, não passa por várias rodadas de mistura e não destrói relações simples entre entrada e saída.

Podemos analisar essa função usando as propriedades que acabamos de estudar:

```text
determinismo:
  sim, a função é determinística

saída fixa:
  sim, a saída sempre tem 8 bits

resistência à pré-imagem:
  não, é fácil procurar entradas que gerem um valor específico

resistência à segunda pré-imagem:
  não, dada uma entrada, é fácil construir outra com a mesma soma

resistência à colisão:
  não, colisões são triviais

efeito avalanche:
  não, mudanças pequenas produzem efeitos previsíveis
```

Esse exemplo mostra a diferença entre uma `função hash não criptográfica` e uma `função hash criptograficamente segura`. Uma função pode transformar entrada variável em saída fixa e ainda assim não servir para criptografia.

Hashes criptográficas reais usam estruturas muito mais complexas. Em geral, elas seguem o seguinte modelo:

```text
1. aplicar padding
2. dividir a mensagem em blocos
3. manter um estado interno
4. processar cada bloco com uma função de compressão ou permutação
5. emitir o estado final como digest
```

Esse modelo aparece na família MD e em SHA-1/SHA-2. Já o SHA-3 usa uma construção diferente, baseada em esponja, mas a intenção geral continua parecida: absorver a entrada, misturar o estado interno e produzir uma saída que pareça imprevisível para quem tenta controlar a entrada.

## Confusão e difusão

Como já vimos no capítulo 1, Claude Shannon introduziu dois conceitos que ajudam a ler qualquer algoritmo criptográfico moderno:

```text
confusão:
  esconder relações simples entre entrada, chave e saída

difusão:
  espalhar a influência de cada bit da entrada por muitos bits da saída
```

Em hashes, a difusão aparece quando um bit alterado na mensagem mexe em muitas palavras, rodadas e bits do digest.

MD2 usa S-box e XOR. MD4, MD5, SHA-1 e SHA-2 usam funções booleanas, somas modulares e rotações. SHA-3 usa uma esponja baseada em permutação.

O objetivo é sempre parecido: misturar localmente várias vezes até que a saída pareça imprevisível para quem tenta controlar a entrada.

Nas próximas seções, veremos de perto os algoritmos da família MD e da família SHA.