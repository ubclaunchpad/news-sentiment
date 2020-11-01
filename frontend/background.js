console.log("background.js is executed when extension is installed or refreshed");

const listOfInjectableURLRegex = [ // change these to the list of allowed news sources
    "^https:\\/\\/www\\.google",
    "^https:\\/\\/www\\.youtube\\.com",
    "^https:\\/\\/www\\.stackoverflow\\.com",
    "^https:\\/\\/www\\.bbc\\.com\\/news\\/"
]

chrome.tabs.onActivated.addListener(tab => {
    console.log(tab); // will print out tabId and windowId as tab object in the background console
    chrome.tabs.get(tab.tabId, activeTabObject => {
        console.log(activeTabObject.url); // will print out active tab info in the background console
        if (testActiveUrlAgainstListOfInjectableURLRegex(activeTabObject.url)) { // inject script if on google.com (for testing, change this)
            try {
                chrome.tabs.executeScript(null, {file: 'frontend/foreground.js'}, // can also do executeCSS
                    () => console.log('background: injected foreground.js')
                );
            } catch (e) {
                console.log('Error when executing foreground.js from background', e);
            }
        }
    });

    chrome.runtime.onMessage.addListener(function (message, sender, sendResponse) {
        if (message.request === 'toBkg') {
            chrome.tabs.get(tab.tabId, activeTabObject => {
                sendResponse({currentUrl: activeTabObject?.url});
            });
        }
        return true;
    });
});

function testActiveUrlAgainstListOfInjectableURLRegex(activeTabUrl) {
    let allowedUrl = false;
    listOfInjectableURLRegex.forEach(r => {
        const re = new RegExp(r);
        if (re.test(activeTabUrl)) {
            allowedUrl = true;
        }
    });
    return allowedUrl;
}
