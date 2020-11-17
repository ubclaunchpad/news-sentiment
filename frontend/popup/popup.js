document.addEventListener('DOMContentLoaded', function() {
    const startingView = `<div id="outerContainer">
                                        <div>STARTING PAGE</div>
                                        <button id="toMain">click to set user rating to center and enter main</button>
                                </div>`;
    const mainView = `<div id="childContainer">
                                    <button id="contentClick">click to get h1 and url from content.js</button>
                                    <button id="backgroundClick">click for url from background.js</button>
                                    <button id="dummyURLs">show list of dummy articles</button>
                                    <script src="popup.js" charset="UTF-8"></script>
                               </div>`;
    const outerContainer = document.getElementById('outerContainer');

    chrome.storage.sync.get('userRating', function(result) {
        if (!result?.['userRating']) {
            outerContainer.innerHTML = startingView;
            document.getElementById('toMain')?.addEventListener('click', setUserRating, false);
        }
        console.log('Value currently is ' + result?.['userRating']);
    });

    document.getElementById('contentClick')?.addEventListener('click', contentOnClick, false);
    document.getElementById('backgroundClick')?.addEventListener('click', bkgOnClick, false);
    document.getElementById('dummyURLs')?.addEventListener('click', dummyOnClick, false);

    function setUserRating() {
        console.log("in set user rating");
        // add a slider and allow user to set their value
        chrome.storage.sync.set({'userRating': 'center'}, function() {
            console.log('Value is set');
            chrome.storage.sync.get('userRating', function(result) {
                console.log('Value is ' + result?.['userRating']);
            });
            outerContainer.innerHTML = mainView;
        });
    }

    function contentOnClick() {
        chrome.tabs.query({currentWindow: true, active: true}, function (tabs) {
            chrome.tabs.sendMessage(tabs[0].id, {request: 'toContent'}, null, fromContent);
        });
    }
    function bkgOnClick() {
        chrome.tabs.query({currentWindow: true, active: true}, () => {
            chrome.runtime.sendMessage({request: 'toBkg'}, fromBkg);
        });
    }

    function dummyOnClick() {
        chrome.tabs.query({currentWindow: true, active: true}, function (tabs) {
            chrome.tabs.sendMessage(tabs[0].id,{request: 'getDummy'}, null, showDummy);
        });
    }

    function fromContent(res) {
        if (res?.firstHeadingValue) {
            const headingDiv = document.createElement('div');
            headingDiv.textContent = res.firstHeadingValue;
            const br = document.createElement('br');
            document.body.appendChild(headingDiv);
            document.body.appendChild(br);
        }
        if (res?.url) {
            const urlDiv = document.createElement('div');
            urlDiv.textContent = res.url;
            const br = document.createElement('br');
            document.body.appendChild(urlDiv);
            document.body.appendChild(br);
        }
    }

    function fromBkg(res) {
        if (res?.currentUrl) {
            const div = document.createElement('div');
            div.textContent = res.currentUrl;
            const br = document.createElement('br');
            document.body.appendChild(div);
            document.body.appendChild(br);
        }
    }

    function showDummy(res) {

        if (res?.listArticles) {
            const listArticles = res.listArticles;
            listArticles.forEach(a => {
                const anchorEl = document.createElement('a');
                anchorEl.textContent = a.TITLE;
                anchorEl.href = a.URL;
                anchorEl.onclick = () => openUrlInNewTab(a.URL);
                document.body.appendChild(anchorEl);
            });
        }
    }

    function openUrlInNewTab(url) {
        chrome.tabs.create({ url: url });
    }

}, false);
