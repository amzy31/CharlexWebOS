document.addEventListener('DOMContentLoaded',()=>{
 const name=document.getElementById('noteFileName'), body=document.getElementById('noteContent'), save=document.getElementById('saveDownloadBtn'), download=document.getElementById('downloadNoteBtn'), load=document.getElementById('loadBtn'), file=document.getElementById('loadEncryptedFile'), status=document.getElementById('noteStatus');
 if(!name||!body)return; const set=s=>{if(status)status.textContent=s};
 function filename(){let n=(name.value||'Untitled').trim()||'Untitled';return n.toLowerCase().endsWith('.txt')?n:n+'.txt'}
 save.onclick=async()=>{if(!body.value.trim())return set('Write something before saving.');if(!window.CryptoJS)return set('Encryption library unavailable.');const pw=prompt('Enter encryption password:');if(!pw)return;try{const n=filename(),enc=CryptoJS.AES.encrypt(body.value,pw).toString();await Charlex.FS.writeFile(n,enc);localStorage.setItem('charlex:lastNote',n);set(`Saved ${n} to WebDisk — staying on index.html`);}catch(e){console.error(e);set('Could not save note.')}};
 download.onclick=()=>{const n=filename();const blob=new Blob([body.value],{type:'text/plain'});const a=document.createElement('a');a.href=URL.createObjectURL(blob);a.download=n;a.click();set(`Downloaded ${n}`);setTimeout(()=>URL.revokeObjectURL(a.href),1000)};
 load.onclick=async()=>{const f=file.files?.[0];if(!f)return set('Choose an encrypted .txt file first.');const pw=prompt('Enter decryption password:');if(!pw)return;try{const enc=await f.text(),dec=CryptoJS.AES.decrypt(enc,pw).toString(CryptoJS.enc.Utf8);if(!dec)throw Error();name.value=f.name;body.value=dec;set(`Loaded ${f.name}`)}catch(e){set('Decryption failed — check the password or file.')}};
 name.value=localStorage.getItem('charlex:lastNote')||'';
});
