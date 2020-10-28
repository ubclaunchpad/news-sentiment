// alert('testing testing');
chrome.runtime.onMessage.addListener(function (request) {
    alert(request);
})
