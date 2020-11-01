chrome.runtime.onMessage.addListener(function (message, sender, sendResponse) {
    if (message.request === 'toContent') {
        const heading = document.getElementsByTagName("h1");
        const firstHeadingValue = heading[0]?.innerText;
        const url = window.location.href;
        sendResponse({firstHeadingValue: firstHeadingValue, url: url});
    }
    return true;
});
