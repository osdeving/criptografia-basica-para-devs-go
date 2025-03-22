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
        this.innerHTML = '<ol class="chapter"><li class="chapter-item expanded affix "><a href="index.html">Introdução</a></li><li class="chapter-item expanded affix "><li class="part-title">Conceitos Fundamentais</li><li class="chapter-item expanded "><a href="conceitos/manipulacao-de-dados/manipulacao-de-dados.html"><strong aria-hidden="true">1.</strong> Manipulação de Dados</a></li><li><ol class="section"><li class="chapter-item expanded "><a href="conceitos/manipulacao-de-dados/introduction.html"><strong aria-hidden="true">1.1.</strong> Introdução</a></li><li class="chapter-item expanded "><a href="conceitos/manipulacao-de-dados/endianness.html"><strong aria-hidden="true">1.2.</strong> Endianness</a></li><li class="chapter-item expanded "><a href="conceitos/manipulacao-de-dados/operadores-bitwise.html"><strong aria-hidden="true">1.3.</strong> Operaadores Bitwise</a></li><li class="chapter-item expanded "><a href="conceitos/manipulacao-de-dados/operacoes-em-blocos.html"><strong aria-hidden="true">1.4.</strong> Operações em Blocos</a></li></ol></li><li class="chapter-item expanded "><a href="conceitos/encode-decode/introduction.html"><strong aria-hidden="true">2.</strong> Base64 e Codificação</a></li><li><ol class="section"><li class="chapter-item expanded "><a href="conceitos/encode-decode/introduction.html"><strong aria-hidden="true">2.1.</strong> Introdução</a></li><li class="chapter-item expanded "><a href="conceitos/encode-decode/base64.html"><strong aria-hidden="true">2.2.</strong> Base64</a></li></ol></li><li class="chapter-item expanded "><a href="conceitos/checksum-crc/checksum-crc.html"><strong aria-hidden="true">3.</strong> Checksum e CRC - Verificação de Integridade</a></li><li><ol class="section"><li class="chapter-item expanded "><a href="conceitos/checksum-crc/introduction.html"><strong aria-hidden="true">3.1.</strong> Introdução</a></li><li class="chapter-item expanded "><a href="conceitos/checksum-crc/parity-bit.html"><strong aria-hidden="true">3.2.</strong> Parity Bit</a></li><li class="chapter-item expanded "><a href="conceitos/checksum-crc/checksum.html"><strong aria-hidden="true">3.3.</strong> Checksum</a></li><li class="chapter-item expanded "><a href="conceitos/checksum-crc/crc.html"><strong aria-hidden="true">3.4.</strong> CRC - Cyclic Redundandy Check</a></li></ol></li><li class="chapter-item expanded "><a href="conceitos/hash/sha256.html"><strong aria-hidden="true">4.</strong> Funções Hash Criptográficas</a></li><li><ol class="section"><li class="chapter-item expanded "><a href="conceitos/hash/sha256.html"><strong aria-hidden="true">4.1.</strong> Funções Hash</a></li><li class="chapter-item expanded "><a href="conceitos/hash/message-digest.html"><strong aria-hidden="true">4.2.</strong> Família de Algoritmos MD - Message Digest</a></li><li class="chapter-item expanded "><a href="conceitos/hash/des.html"><strong aria-hidden="true">4.3.</strong> Família de Algoritmos DES - Data Encryption Standard</a></li><li class="chapter-item expanded "><a href="conceitos/hash/sha.html"><strong aria-hidden="true">4.4.</strong> Família de Algoritmos SHA - Secure Hash Algorithms</a></li><li class="chapter-item expanded "><a href="conceitos/hash/pbkdf2-argon2.html"><strong aria-hidden="true">4.5.</strong> PBKDF2, Argon2 e Scrypt - Hash para Senhas</a></li></ol></li><li class="chapter-item expanded "><a href="conceitos/crypto/crypto-algorithms.html"><strong aria-hidden="true">5.</strong> Algorítmos de Criptografia</a></li><li><ol class="section"><li class="chapter-item expanded "><a href="criptografia/crypto/des.html"><strong aria-hidden="true">5.1.</strong> DES - Data Encryption Standard</a></li><li class="chapter-item expanded "><a href="criptografia/crypto/3des.html"><strong aria-hidden="true">5.2.</strong> 3DES - Triple DES</a></li><li class="chapter-item expanded "><a href="criptografia/crypto/aes.html"><strong aria-hidden="true">5.3.</strong> AES - Advanced Encryption Standard</a></li><li class="chapter-item expanded "><a href="criptografia/crypto/rsa.html"><strong aria-hidden="true">5.4.</strong> RSA - Rivest, Shamir e Adleman</a></li><li class="chapter-item expanded "><a href="criptografia/crypto/ecdsa.html"><strong aria-hidden="true">5.5.</strong> ECDSA - Elliptic Curve Digital Signature Algorithm</a></li><li class="chapter-item expanded "><a href="conceitos/crypto/hmac.html"><strong aria-hidden="true">5.6.</strong> HMAC - Assinaturas Seguras</a></li><li class="chapter-item expanded "><a href="conceitos/crypto/criptografia.html"><strong aria-hidden="true">5.7.</strong> Criptografia Simétrica vs Assimétrica</a></li></ol></li><li class="chapter-item expanded "><li class="part-title">Referências</li><li class="chapter-item expanded "><a href="referencias/bibliotecas.html"><strong aria-hidden="true">6.</strong> Bibliotecas Recomendadas</a></li><li class="chapter-item expanded "><a href="referencias/leituras.html"><strong aria-hidden="true">7.</strong> Links e Leituras Adicionais</a></li></ol>';
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
