# Run on Chrome
Go to
```
chrome://extensions/
```

in your chrome browser.

Turn on developer mode through the Developer mode toggle on the upper right-hand side.

Click "Load unpacked" in the upper left-hand side.

Choose the news-sentiment root folder.

Click on the refresh icon if you make a change to the file.

You should be able to use the extension through the browser now.

------------------
As of Oct 31:
We can get the user's url using the content.js script or the background.js script. I've implemented both for now.

To use:
- pull the changes and reload the extension in chrome://extensions.
- there will be two buttons in the extension popup. Click the first button to get the h1 text and the current url 
(from content.js) or click the second button to get the url (from background.js). They should be displayed as lines of 
text in the popup.

Also, we can restrict when we want our extension to only do things when 
on certain sites by checking the url with regex (there is an example in the background.js file) or we can restrict 
where the content.js file should work in the manifest.json using the "matches" key under "content_scripts". 


