package app

import (
	"embed"
	"html/template"
	"io/fs"
	"mime"
	"net/http"
	"path"
	"strings"
)

//go:embed static/* static/css/* static/img/* static/js/*
var staticFiles embed.FS

var page = template.Must(template.New("index").Parse(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width,initial-scale=1,viewport-fit=cover">
<title>CharleX WebOS</title>
<link rel="icon" href="/img/logo.jpg">
<link rel="stylesheet" href="/css/style.css?v=20260908">
<script src="/js/crypto.js?v=20260908"></script>
</head>
<body>
<script>
window.CharlexWindowControls = window.CharlexWindowControls || {
  action: function(button, event) {
    if (!button) return;
    var now = Date.now();
    if (event && event.type === 'click' && button.__charlexLastPointerAction && now - button.__charlexLastPointerAction < 350) return;
    if (event && event.type === 'pointerdown') button.__charlexLastPointerAction = now;
    if (event) { event.preventDefault(); event.stopPropagation(); }
    var w = button.closest ? button.closest('.window') : null;
    if (!w) return;
    var restore = w.__charlexRestore || {};
    if (button.classList.contains('close')) {
      w.style.display = 'none';
      w.classList.remove('minimized','maximized','focused');
      return;
    }
    if (button.classList.contains('minimize')) {
      w.style.display = 'none';
      w.classList.add('minimized');
      w.classList.remove('focused');
      return;
    }
    if (button.classList.contains('maximize')) {
      if (!w.classList.contains('maximized')) {
        w.__charlexRestore = {position:w.style.position,left:w.style.left,top:w.style.top,width:w.style.width,height:w.style.height};
        w.classList.add('maximized');
        w.style.position = 'fixed';
        w.style.left = '0'; w.style.top = '0';
        w.style.width = '100vw'; w.style.height = '100vh';
      } else {
        w.classList.remove('maximized');
        w.style.position = restore.position || 'absolute';
        w.style.left = restore.left || ''; w.style.top = restore.top || '';
        w.style.width = restore.width || ''; w.style.height = restore.height || '';
      }
      w.style.display = 'flex';
      return;
    }
  }
};
</script>
<div id="desktop" style="background-image:url('/img/bg.jpg')">
  {{template "windows" .}}
</div>
<div id="dock" class="dock-container bg-dark container d-flex justify-content-center align-items-center" role="navigation" aria-label="Application dock">
  <button class="dock-icon btn btn-dark rounded-circle" type="button" title="Welcome" aria-label="Welcome" onclick="Charlex.DOM.showWindow('window1')"><div class="dock-icon-content"><img src="/img/logo.jpg" alt=""></div></button>
  <button class="dock-icon btn btn-dark rounded-circle" type="button" title="Notes" aria-label="Notes" onclick="Charlex.DOM.showWindow('noteWindow')"><div class="dock-icon-content"><img src="/img/accessories-text-editor.svg" alt=""></div></button>
  <button class="dock-icon btn btn-dark rounded-circle" type="button" title="System" aria-label="System" onclick="Charlex.DOM.showWindow('sysInfo')"><div class="dock-icon-content"><img src="/img/sysinfo.png" alt=""></div></button>
  <button class="dock-icon btn btn-dark rounded-circle" type="button" title="Terminal" aria-label="Terminal" onclick="Charlex.DOM.showWindow('shellWindow')"><div class="dock-icon-content"><img src="/img/terminal.svg" alt=""></div></button>
  <button class="dock-icon btn btn-dark rounded-circle" type="button" title="Browser" aria-label="Browser" onclick="Charlex.DOM.showWindow('browserWindow')"><div class="dock-icon-content"><img src="/img/browser.png" alt=""></div></button>
  <button class="dock-icon btn btn-dark rounded-circle" type="button" title="WebDisk" aria-label="WebDisk" onclick="Charlex.DOM.showWindow('webdiskWindow')"><div class="dock-icon-content"><img src="/img/file.png" alt=""></div></button>
  <button class="dock-icon btn btn-dark rounded-circle" type="button" title="Fullscreen" aria-label="Fullscreen" onclick="Charlex.toggleFullscreen()"><div class="dock-icon-content"><span class="dock-glyph">⛶</span></div></button>
  <button class="dock-icon btn btn-dark rounded-circle" type="button" title="Restart" aria-label="Restart" onclick="location.reload()"><div class="dock-icon-content"><span class="dock-glyph">↻</span></div></button>
</div>
<script src="/js/chrlex-dom.js?v=20260908"></script>
<script src="/js/window_manager.js?v=20260908"></script>
<script src="/js/dock.js?v=20260908"></script>
<script src="/js/browser.js?v=20260908"></script>
<script src="/js/webdisk.js?v=20260908"></script>
<script src="/js/note.js?v=20260908"></script>
<script src="/js/sysinfo.js?v=20260908"></script>
<script src="/js/shell.js?v=20260908"></script>
<script src="/js/go-runtime.js?v=20260908"></script>
</body>
</html>

{{define "windows"}}
<section class="window" id="window1" style="top:7%;left:9%;display:none">
<header class="window-header"><div class="window-controls"><button class="window-control-button close" type="button" aria-label="Close" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button><button class="window-control-button minimize" type="button" aria-label="Minimize" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button><button class="window-control-button maximize" type="button" aria-label="Maximize" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button></div><div class="window-title">Welcome</div><div class="window-header-spacer"></div></header>
<div class="window-content"><div class="welcome"><div class="welcome-icon">⌘</div><h1>CharleX WebOS</h1><p>A responsive browser desktop with local apps, WebDisk storage and encrypted notes.</p><button class="glass-button" onclick="Charlex.DOM.showWindow('noteWindow')">Open Notes</button></div></div></section>

<section class="window" id="noteWindow" style="top:12%;left:12%;display:none">
<header class="window-header"><div class="window-controls"><button class="window-control-button close" type="button" aria-label="Close" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button><button class="window-control-button minimize" type="button" aria-label="Minimize" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button><button class="window-control-button maximize" type="button" aria-label="Maximize" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button></div><div class="window-title">Notes</div><div class="window-header-spacer"></div></header>
<div class="window-content"><div class="note-app"><div class="note-top"><input id="noteFileName" placeholder="Note filename.txt"><span id="noteStatus">Local, private storage</span></div><textarea id="noteContent" placeholder="Start writing…"></textarea><div class="note-actions"><button id="saveDownloadBtn" class="glass-button primary" type="button">Save encrypted</button><button id="downloadNoteBtn" class="glass-button" type="button">Download</button><input id="loadEncryptedFile" type="file" accept=".txt"><button id="loadBtn" class="glass-button" type="button">Load encrypted</button></div></div></div></section>

<section class="window" id="sysInfo" style="top:15%;left:15%;display:none">
<header class="window-header"><div class="window-controls"><button class="window-control-button close" type="button" aria-label="Close" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button><button class="window-control-button minimize" type="button" aria-label="Minimize" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button><button class="window-control-button maximize" type="button" aria-label="Maximize" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button></div><div class="window-title">System Info</div><div class="window-header-spacer"></div></header>
<div class="window-content"><pre id="topOutput" class="sysinfo">Loading…</pre></div></section>

<section class="window" id="shellWindow" style="top:18%;left:18%;display:none">
<header class="window-header"><div class="window-controls"><button class="window-control-button close" type="button" aria-label="Close" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button><button class="window-control-button minimize" type="button" aria-label="Minimize" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button><button class="window-control-button maximize" type="button" aria-label="Maximize" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button></div><div class="window-title">Terminal</div><div class="window-header-spacer"></div></header>
<div class="window-content"><div id="shellContent" class="shell"><div>CharleX WebOS shell — local commands only</div><div class="shell-prompt"><span>$</span><input id="shellInput" autocomplete="off"></div></div></div></section>

<section class="window" id="browserWindow" style="top:8%;left:20%;display:none">
<header class="window-header"><div class="window-controls"><button class="window-control-button close" type="button" aria-label="Close" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button><button class="window-control-button minimize" type="button" aria-label="Minimize" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button><button class="window-control-button maximize" type="button" aria-label="Maximize" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button></div><div class="window-title">Browser</div><div class="window-header-spacer"></div></header>
<div class="window-content"><div class="browser-app"><div class="browser-bar"><input id="urlInput" placeholder="https://example.com"><button id="goBtn" class="glass-button primary" type="button">Go</button><button id="openExternalBtn" class="glass-button" type="button">Open External</button></div><div id="browserStatus" class="browser-status" role="status" aria-live="polite">Ready</div><iframe id="browserFrame" title="Browser"></iframe></div></div></section>

<section class="window" id="webdiskWindow" style="top:6%;left:6%;display:none">
<header class="window-header"><div class="window-controls"><button class="window-control-button close" type="button" aria-label="Close" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button><button class="window-control-button minimize" type="button" aria-label="Minimize" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button><button class="window-control-button maximize" type="button" aria-label="Maximize" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button></div><div class="window-title">WebDisk</div><div class="window-header-spacer"></div></header>
<div class="window-content"><div id="webdiskContent" class="webdisk-shell"><div class="webdisk-toolbar"><div class="webdisk-nav"><button id="backBtn" class="webdisk-tool">‹</button><button id="forwardBtn" class="webdisk-tool">›</button></div><div><strong>WebDisk</strong><small id="currentDir">/</small></div><div class="webdisk-actions"><button id="newFileBtn" class="webdisk-tool">＋ File</button><button id="newFolderBtn" class="webdisk-tool">＋ Folder</button><button id="renameBtn" class="webdisk-tool">Rename</button><button id="deleteBtn" class="webdisk-tool danger">Delete</button></div></div><div class="webdisk-body"><aside class="webdisk-sidebar"><b>Locations</b><button class="webdisk-location active" data-path="/">⌂ WebDisk</button></aside><main class="webdisk-main"><div class="webdisk-statusbar"><span id="webdiskStatus">Preparing…</span></div><div class="webdisk-head"><span>Name</span><span>Kind</span><span>Size</span><span>Modified</span></div><div id="webdiskList" class="webdisk-list"></div><div id="webdiskEmpty" class="webdisk-empty" hidden>📂<strong>This folder is empty</strong><span>Create a file or folder to get started.</span></div><div id="webdiskLoading" class="webdisk-loading"><span class="webdisk-spinner"></span><span>Loading files…</span></div></main></div></div></div></section>
{{end}}`))

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" || r.URL.Path == "/index.html" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_ = page.Execute(w, nil)
		return
	}
	serveStatic(w, r)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	p := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
	if p == "" || strings.Contains(p, "..") {
		http.NotFound(w, r)
		return
	}
	data, err := fs.ReadFile(staticFiles, "static/"+p)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if ct := mime.TypeByExtension(path.Ext(p)); ct != "" {
		w.Header().Set("Content-Type", ct)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}
