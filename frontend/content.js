chrome.runtime.onMessage.addListener(function (message, sender, sendResponse) {
    if (message.request === 'toContent') {
        const heading = document.getElementsByTagName("h1");
        const firstHeadingValue = heading[0]?.innerText;
        const url = window.location.href;
        sendResponse({firstHeadingValue: firstHeadingValue, url: url});
    } else if (message.request === 'getList') {
        getData('http://localhost:8090/articles').then(data => sendResponse({ listArticles: data }));
    }
    return true;
});

function getData(url) {
    return fetch(url)
      .then((res) => {
        return res.json();
      })
      .then((result) => {
        return result;
      })
  }
  
  