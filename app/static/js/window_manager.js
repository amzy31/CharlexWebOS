(function () {
  'use strict';

  const C = window.Charlex = window.Charlex || {};
  C.WindowManager = C.WindowManager || {};
  let z = 100;
  let drag = null;

  const $ = (selector, root = document) => root.querySelector(selector);
  const $$ = (selector, root = document) => Array.from(root.querySelectorAll(selector));

  function focus(el) {
    if (!el) return;
    z = Math.max(z + 1, 100);
    el.style.zIndex = String(z);
    $$('.window').forEach(w => w.classList.remove('focused'));
    el.classList.add('focused');
  }

  function openWindow(id) {
    const w = document.getElementById(id);
    if (!w) return false;
    w.style.display = 'flex';
    w.classList.remove('minimized');
    focus(w);
    return true;
  }

  function closeWindow(id) {
    const w = document.getElementById(id);
    if (!w) return false;
    w.style.display = 'none';
    w.classList.remove('minimized', 'maximized', 'focused');
    return true;
  }

  function minimizeWindow(id) {
    const w = document.getElementById(id);
    if (!w) return false;
    w.style.display = 'none';
    w.classList.add('minimized');
    w.classList.remove('focused');
    return true;
  }

  function maximizeWindow(id) {
    const w = document.getElementById(id);
    if (!w) return false;
    const fallback = window.CharlexWindowControls;
    if (fallback) {
      fallback.action(w.querySelector('.maximize'));
      focus(w);
      return true;
    }
    return false;
  }

  function startDrag(e, id) {
    const w = document.getElementById(id);
    if (!w || w.classList.contains('maximized') || e.target.closest('.window-controls')) return;
    const p = e.touches?.[0] || e;
    const rect = w.getBoundingClientRect();
    drag = { el: w, ox: p.clientX - rect.left, oy: p.clientY - rect.top };
    focus(w);
  }

  function move(e) {
    if (!drag) return;
    const p = e.touches?.[0] || e;
    const w = drag.el;
    const maxX = Math.max(0, window.innerWidth - w.offsetWidth);
    const maxY = Math.max(0, window.innerHeight - w.offsetHeight);
    w.style.left = Math.max(0, Math.min(p.clientX - drag.ox, maxX)) + 'px';
    w.style.top = Math.max(0, Math.min(p.clientY - drag.oy, maxY)) + 'px';
  }

  function endDrag() { drag = null; }

  function bindWindows(root = document) {
    $$('.window', root).forEach(w => {
      if (w.dataset.wmBound === '1') return;
      w.dataset.wmBound = '1';
      const header = $('.window-header', w);
      if (header) {
        header.addEventListener('pointerdown', e => { if (e.button === 0) startDrag(e, w.id); });
        header.addEventListener('dblclick', e => {
          if (e.target.closest('.window-controls')) return;
          maximizeWindow(w.id);
        });
      }
      w.addEventListener('pointerdown', () => focus(w));
    });
  }

  C.WindowManager.openWindow = openWindow;
  C.WindowManager.closeWindow = closeWindow;
  C.WindowManager.minimizeWindow = minimizeWindow;
  C.WindowManager.maximizeWindow = maximizeWindow;
  C.WindowManager.startDrag = startDrag;
  C.WindowManager.initWindowManager = bindWindows;

  document.addEventListener('pointermove', move);
  document.addEventListener('pointerup', endDrag);
  document.addEventListener('pointercancel', endDrag);
  document.addEventListener('DOMContentLoaded', () => bindWindows(), { once: true });
  if (document.readyState !== 'loading') bindWindows();
})();
