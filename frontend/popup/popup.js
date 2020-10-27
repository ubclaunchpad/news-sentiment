document.addEventListener('DOMContentLoaded', function() {

    document.querySelector('button').addEventListener('click', onclick, false);

    function onclick() { // currently connected to the button in popup.html with ^ listener
        chrome.tabs.query({currentWindow: true, active: true},
            function (tabs) {
                chrome.tabs.sendMessage(null,'hi');
                alert('hellooooo');
            });
    }
}, false)
