// Populate the sidebar
//
// This is a script, and not included directly in the page, to control the total size of the book.
// The TOC contains an entry for each page, so if each page includes a copy of the TOC,
// the total size of the page becomes O(n**2).
class MDBookSidebarScrollbox extends HTMLElement {
    constructor() {
        super();
    }
    connectedCallback() {
        this.innerHTML = '<ol class="chapter"><li class="chapter-item expanded affix "><a href="index.html">Introdução</a></li><li class="chapter-item expanded affix "><li class="part-title">Conceitos Fundamentais</li><li class="chapter-item expanded "><a href="conceitos/base64.html"><strong aria-hidden="true">1.</strong> Base64 e Codificação</a></li><li class="chapter-item expanded "><a href="conceitos/sha256.html"><strong aria-hidden="true">2.</strong> Hashing e SHA-256</a></li><li class="chapter-item expanded "><a href="conceitos/hmac.html"><strong aria-hidden="true">3.</strong> HMAC - Assinaturas Seguras</a></li><li class="chapter-item expanded "><a href="conceitos/criptografia.html"><strong aria-hidden="true">4.</strong> Criptografia Simétrica vs Assimétrica</a></li><li class="chapter-item expanded affix "><li class="part-title">Autenticação e Assinaturas Digitais</li><li class="chapter-item expanded "><a href="autenticacao/jwt.html"><strong aria-hidden="true">5.</strong> JSON Web Token (JWT)</a></li><li><ol class="section"><li class="chapter-item expanded "><a href="autenticacao/jwt/estrutura.html"><strong aria-hidden="true">5.1.</strong> Estrutura de um JWT</a></li><li class="chapter-item expanded "><a href="autenticacao/jwt/assinatura.html"><strong aria-hidden="true">5.2.</strong> Assinatura e Verificação</a></li><li class="chapter-item expanded "><a href="autenticacao/jwt/expiracao-refresh.html"><strong aria-hidden="true">5.3.</strong> JWT Expiração e Refresh Token</a></li><li class="chapter-item expanded "><a href="autenticacao/jwt/exemplo.html"><strong aria-hidden="true">5.4.</strong> Exemplo Prático com Go</a></li></ol></li><li class="chapter-item expanded "><a href="autenticacao/oauth2.html"><strong aria-hidden="true">6.</strong> OAuth 2.0 e OpenID Connect</a></li><li><ol class="section"><li class="chapter-item expanded "><a href="autenticacao/oauth2/authorization-code.html"><strong aria-hidden="true">6.1.</strong> Fluxo Authorization Code</a></li><li class="chapter-item expanded "><a href="autenticacao/oauth2/client-credentials.html"><strong aria-hidden="true">6.2.</strong> Client Credentials e Machine-to-Machine</a></li><li class="chapter-item expanded "><a href="autenticacao/oauth2/scopes.html"><strong aria-hidden="true">6.3.</strong> OAuth Scopes e Permissões</a></li><li class="chapter-item expanded "><a href="autenticacao/oauth2/exemplo.html"><strong aria-hidden="true">6.4.</strong> Exemplo Prático com Go</a></li></ol></li><li class="chapter-item expanded "><li class="part-title">Segurança no Transporte</li><li class="chapter-item expanded "><a href="seguranca/tls.html"><strong aria-hidden="true">7.</strong> TLS e Certificados SSL</a></li><li><ol class="section"><li class="chapter-item expanded "><a href="seguranca/tls/handshake.html"><strong aria-hidden="true">7.1.</strong> Como Funciona o Handshake TLS</a></li><li class="chapter-item expanded "><a href="seguranca/tls/certificados.html"><strong aria-hidden="true">7.2.</strong> Certificados e Autoridades Certificadoras</a></li><li class="chapter-item expanded "><a href="seguranca/tls/ciphers.html"><strong aria-hidden="true">7.3.</strong> Cipher Suites e Forward Secrecy</a></li><li class="chapter-item expanded "><a href="seguranca/tls/exemplo.html"><strong aria-hidden="true">7.4.</strong> Exemplo Prático com Go</a></li></ol></li><li class="chapter-item expanded "><a href="seguranca/apis.html"><strong aria-hidden="true">8.</strong> Segurança em APIs REST</a></li><li><ol class="section"><li class="chapter-item expanded "><a href="seguranca/apis/https.html"><strong aria-hidden="true">8.1.</strong> HTTPS Obrigatório</a></li><li class="chapter-item expanded "><a href="seguranca/apis/jwt-seguranca.html"><strong aria-hidden="true">8.2.</strong> Protegendo Endpoints com JWT</a></li><li class="chapter-item expanded "><a href="seguranca/apis/rate-limit.html"><strong aria-hidden="true">8.3.</strong> Rate Limiting e Proteção contra DoS</a></li><li class="chapter-item expanded "><a href="seguranca/apis/hmac-webhooks.html"><strong aria-hidden="true">8.4.</strong> Assinaturas HMAC para Webhooks</a></li></ol></li><li class="chapter-item expanded "><li class="part-title">Práticas Recomendadas</li><li class="chapter-item expanded "><a href="praticas/senhas.html"><strong aria-hidden="true">9.</strong> Armazenamento Seguro de Senhas</a></li><li><ol class="section"><li class="chapter-item expanded "><a href="praticas/senhas/bcrypt-argon2.html"><strong aria-hidden="true">9.1.</strong> Bcrypt, Argon2 e PBKDF2</a></li><li class="chapter-item expanded "><a href="praticas/senhas/erros-comuns.html"><strong aria-hidden="true">9.2.</strong> Erros Comuns no Hashing de Senhas</a></li></ol></li><li class="chapter-item expanded "><a href="praticas/tokens.html"><strong aria-hidden="true">10.</strong> Gerenciamento de Tokens Seguros</a></li><li><ol class="section"><li class="chapter-item expanded "><a href="praticas/tokens/armazenamento.html"><strong aria-hidden="true">10.1.</strong> Armazenando Tokens com Segurança</a></li><li class="chapter-item expanded "><a href="praticas/tokens/expiracao-rotacao.html"><strong aria-hidden="true">10.2.</strong> Expiração e Rotação de Tokens</a></li></ol></li><li class="chapter-item expanded "><a href="praticas/checklist-apis.html"><strong aria-hidden="true">11.</strong> Checklist de Segurança para APIs</a></li><li class="chapter-item expanded affix "><li class="part-title">Implementação Prática</li><li class="chapter-item expanded "><a href="implementacao/oauth2-server.html"><strong aria-hidden="true">12.</strong> Criando um Servidor OAuth2 com Go</a></li><li class="chapter-item expanded "><a href="implementacao/backend-seguro.html"><strong aria-hidden="true">13.</strong> Protegendo um Backend com JWT e TLS</a></li><li class="chapter-item expanded "><a href="implementacao/hmac-webhook.html"><strong aria-hidden="true">14.</strong> Implementando um Webhook Seguro com HMAC</a></li><li class="chapter-item expanded affix "><li class="part-title">Referências</li><li class="chapter-item expanded "><a href="referencias/bibliotecas.html"><strong aria-hidden="true">15.</strong> Bibliotecas Recomendadas</a></li><li class="chapter-item expanded "><a href="referencias/leituras.html"><strong aria-hidden="true">16.</strong> Links e Leituras Adicionais</a></li><li class="chapter-item expanded affix "><a href="misc/contributors.html">Contribuidores</a></li></ol>';
        // Set the current, active page, and reveal it if it's hidden
        let current_page = document.location.href.toString().split("#")[0];
        if (current_page.endsWith("/")) {
            current_page += "index.html";
        }
        var links = Array.prototype.slice.call(this.querySelectorAll("a"));
        var l = links.length;
        for (var i = 0; i < l; ++i) {
            var link = links[i];
            var href = link.getAttribute("href");
            if (href && !href.startsWith("#") && !/^(?:[a-z+]+:)?\/\//.test(href)) {
                link.href = path_to_root + href;
            }
            // The "index" page is supposed to alias the first chapter in the book.
            if (link.href === current_page || (i === 0 && path_to_root === "" && current_page.endsWith("/index.html"))) {
                link.classList.add("active");
                var parent = link.parentElement;
                if (parent && parent.classList.contains("chapter-item")) {
                    parent.classList.add("expanded");
                }
                while (parent) {
                    if (parent.tagName === "LI" && parent.previousElementSibling) {
                        if (parent.previousElementSibling.classList.contains("chapter-item")) {
                            parent.previousElementSibling.classList.add("expanded");
                        }
                    }
                    parent = parent.parentElement;
                }
            }
        }
        // Track and set sidebar scroll position
        this.addEventListener('click', function(e) {
            if (e.target.tagName === 'A') {
                sessionStorage.setItem('sidebar-scroll', this.scrollTop);
            }
        }, { passive: true });
        var sidebarScrollTop = sessionStorage.getItem('sidebar-scroll');
        sessionStorage.removeItem('sidebar-scroll');
        if (sidebarScrollTop) {
            // preserve sidebar scroll position when navigating via links within sidebar
            this.scrollTop = sidebarScrollTop;
        } else {
            // scroll sidebar to current active section when navigating via "next/previous chapter" buttons
            var activeSection = document.querySelector('#sidebar .active');
            if (activeSection) {
                activeSection.scrollIntoView({ block: 'center' });
            }
        }
        // Toggle buttons
        var sidebarAnchorToggles = document.querySelectorAll('#sidebar a.toggle');
        function toggleSection(ev) {
            ev.currentTarget.parentElement.classList.toggle('expanded');
        }
        Array.from(sidebarAnchorToggles).forEach(function (el) {
            el.addEventListener('click', toggleSection);
        });
    }
}
window.customElements.define("mdbook-sidebar-scrollbox", MDBookSidebarScrollbox);
