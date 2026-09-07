(function () {
  'use strict';

  const C = window.Charlex = window.Charlex || {};
  C.Browser = C.Browser || {};

  function normalizeURL(value) {
    let u = String(value || '').trim();
    if (!u) return '';
    if (!/^[a-z][a-z0-9+.-]*:\/\//i.test(u)) u = 'https://' + u;
    try { return new URL(u).href; } catch (_) { return ''; }
  }

  function setStatus(message, kind) {
    const status = document.getElementById('browserStatus');
    if (!status) return;
    status.textContent = message;
    status.dataset.kind = kind || '';
  }

  function navigate(url) {
    const normalized = normalizeURL(url);
    const input = document.getElementById('urlInput');
    const frame = document.getElementById('browserFrame');
    if (!normalized || !input || !frame) return false;

    input.value = normalized;
    frame.src = normalized;
    setStatus('Loading ' + normalized, 'loading');
    return true;
  }

  function init() {
    const input = document.getElementById('urlInput');
    const go = document.getElementById('goBtn');
    const frame = document.getElementById('browserFrame');
    if (!input || !go || !frame) return;

    const initial = 'https://example.com';
    input.value = initial;

    const nav = () => navigate(input.value);
    go.addEventListener('click', nav);
    input.addEventListener('keydown', e => { if (e.key === 'Enter') nav(); });

    frame.addEventListener('load', () => {
      setStatus('Loaded', 'ok');
    });
    frame.addEventListener('error', () => {
      setStatus('This site cannot be embedded. Use Open External.', 'error');
    });

    const external = document.getElementById('openExternalBtn');
    if (external) {
      external.addEventListener('click', () => {
        const url = normalizeURL(input.value);
        if (url) window.open(url, '_blank', 'noopener,noreferrer');
      });
    }

    navigate(initial);
  }

  C.Browser.navigate = navigate;
  C.Browser.init = init;
  window.openBrowserWindow = () => C.DOM && C.DOM.showWindow('browserWindow');

  if (document.readyState === 'loading') document.addEventListener('DOMContentLoaded', init, { once: true });
  else init();
})();
