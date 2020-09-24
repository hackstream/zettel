window.addEventListener("load", function(evt) {
    var ws = new WebSocket("ws://" + window.location.host + "/ws");
    ws.onopen = function(evt) {
        console.log("OPEN");
    }
    ws.onclose = function(evt) {
        console.log("CLOSE");
    }
    ws.onmessage = function(evt) {
        console.log("got reload event");
        ws.close();
        window.location.reload();
    }
    ws.onerror = function(evt) {
        print("ERROR: " + evt.data);
    }
});