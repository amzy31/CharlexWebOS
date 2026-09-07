(function () {
  window.Charlex = window.Charlex || {};
  Charlex.toggleFullscreen = async function () {
    try {
      if (!document.fullscreenElement) await document.documentElement.requestFullscreen();
      else await document.exitFullscreen();
    } catch (_) {
      alert('Fullscreen is not available in this browser.');
    }
  };
  document.addEventListener('DOMContentLoaded', function () {
    if (Charlex.WindowManager && Charlex.WindowManager.initWindowManager) {
      Charlex.WindowManager.initWindowManager();
    }
    var dock = document.getElementById('dock');
    if (dock && window.innerWidth <= 480) dock.style.bottom = '10px';
  });
})();
