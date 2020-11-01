document.addEventListener('DOMContentLoaded', function() {
    document.getElementById('contentClick').addEventListener('click', contentOnClick, false);
    document.getElementById('backgroundClick').addEventListener('click', bkgOnClick, false);

    function contentOnClick() {
        chrome.tabs.query({currentWindow: true, active: true}, function (tabs) {
            chrome.tabs.sendMessage(tabs[0].id, {request: 'toContent'}, null, showHeading);
        });
    }
    function bkgOnClick() {
        chrome.tabs.query({currentWindow: true, active: true}, function (tabs) {
            chrome.tabs.sendMessage(tabs[0].id, {request: 'toBkg'}, null, showUrl);
        });
    }

    function showHeading(res) {
        if (res) {
            const headingDiv = document.createElement('div');
            headingDiv.textContent = res.firstHeadingValue;
            const br = document.createElement('br');
            const urlDiv = document.createElement('div');
            urlDiv.textContent = res.url;
            document.body.appendChild(headingDiv);
            document.body.appendChild(br);
            document.body.appendChild(urlDiv);
        }
    }

    function showUrl(res) {
        console.log("showUrl response: ", res);
        const div = document.createElement('div');
        div.textContent = res.firstHeadingValue;
        console.log("res.currentUrl in popup.js: ");
        document.body.appendChild(div);
    }

}, false);
