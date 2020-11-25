chrome.runtime.onMessage.addListener(function (message, sender, sendResponse) {
    if (message.request === 'toContent') {
      const heading = document.getElementsByTagName("h1");
      const firstHeadingValue = heading[0]?.innerText;
      const url = window.location.href;
      if (testActiveUrlAgainstListOfInjectableURLRegex(url)) {
        sendResponse({firstHeadingValue: firstHeadingValue, url: url});
      }
    } else if (message.request === 'getList') {
      getData('http://localhost:8090/articles').then(data => sendResponse({ listArticles: data }));
    } else if (message.request === 'getSources') {
      getData("http://localhost:8090/sources").then(sources =>  addToListOfInjectableUrlRegex(sources));
    }
    return true;
});

let listOfInjectableURLRegex = []


function getData(url) {
    return fetch(url)
      .then((res) => {
        return res.json();
      })
      .then((result) => {
        return result;
      })
}

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

function addToListOfInjectableUrlRegex(sources) {
  listOfInjectableURLRegex = [];
  sources.forEach(({ url }) => {
      if (url) {
          listOfInjectableURLRegex.push(new RegExp(url));
      }
  })
}
