(function(){
  const C=window.Charlex=window.Charlex||{}; C.DOM=C.DOM||{};
  C.DOM.createWindow=function(id,title,html,top='10%',left='10%',display='none'){
    const old=document.getElementById(id); if(old)old.remove();
    const w=document.createElement('section'); w.className='window';w.id=id;Object.assign(w.style,{top,left,display});
    w.innerHTML=`<header class="window-header"><div class="window-controls"><button class="window-control-button close" type="button" aria-label="Close" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button><button class="window-control-button minimize" type="button" aria-label="Minimize" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button><button class="window-control-button maximize" type="button" aria-label="Maximize" onpointerdown="window.CharlexWindowControls.action(this,event)" onclick="window.CharlexWindowControls.action(this,event)"></button></div><div class="window-title">${title}</div><div class="window-header-spacer"></div></header><div class="window-content">${html}</div>`;
    document.getElementById('desktop').appendChild(w);
    if (C.WindowManager && C.WindowManager.initWindowManager) C.WindowManager.initWindowManager(document);
    return w;
  };
  C.DOM.showWindow=id=>C.WindowManager.openWindow(id); C.DOM.hideWindow=id=>C.WindowManager.closeWindow(id);
})();
